package activation

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// This is the secret key for signing JWTs - you should keep this secret
	// For production, you might want to use environment variables or a config file
	jwtSecret = "your-secret-key-change-this-in-production-2024"
	
	activationFile = "activation.key"
)

// Claims represents the JWT claims for activation
type Claims struct {
	DeviceID string `json:"device_id"`
	jwt.RegisteredClaims
}

// GetDeviceID generates a unique device identifier based on hardware information
func GetDeviceID() (string, error) {
	var identifiers []string

	// Get hostname
	hostname, err := os.Hostname()
	if err == nil {
		identifiers = append(identifiers, hostname)
	}

	// Get MAC addresses
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			// Skip loopback and down interfaces
			if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
				mac := iface.HardwareAddr.String()
				if mac != "" {
					identifiers = append(identifiers, mac)
				}
			}
		}
	}

	if len(identifiers) == 0 {
		return "", fmt.Errorf("could not generate device ID")
	}

	// Create a hash of all identifiers
	combined := strings.Join(identifiers, "|")
	hash := sha256.Sum256([]byte(combined))
	deviceID := hex.EncodeToString(hash[:])

	// Return first 32 characters for readability
	return deviceID[:32], nil
}

// GenerateActivationKey creates a JWT token for a specific device ID
// This function should only be used by the admin
func GenerateActivationKey(deviceID string) (string, error) {
	claims := Claims{
		DeviceID: deviceID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)), // 1 year expiration
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "env-invoices-admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// ValidateActivationKey checks if the activation key is valid for the current device
func ValidateActivationKey(activationKey string) error {
	deviceID, err := GetDeviceID()
	if err != nil {
		return fmt.Errorf("failed to get device ID: %w", err)
	}

	token, err := jwt.ParseWithClaims(activationKey, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return fmt.Errorf("invalid activation key: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid activation key claims")
	}

	// Check if the device ID matches
	if claims.DeviceID != deviceID {
		return fmt.Errorf("activation key is not valid for this device")
	}

	// Check if token is expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return fmt.Errorf("activation key has expired")
	}

	return nil
}

// SaveActivationKey saves the activation key to a file
func SaveActivationKey(activationKey string) error {
	// First validate the key
	if err := ValidateActivationKey(activationKey); err != nil {
		return err
	}

	return os.WriteFile(activationFile, []byte(activationKey), 0600)
}

// LoadActivationKey loads and validates the saved activation key
func LoadActivationKey() (string, error) {
	data, err := os.ReadFile(activationFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("no activation key found")
		}
		return "", fmt.Errorf("failed to read activation key: %w", err)
	}

	activationKey := strings.TrimSpace(string(data))
	
	// Validate the loaded key
	if err := ValidateActivationKey(activationKey); err != nil {
		return "", err
	}

	return activationKey, nil
}

// IsActivated checks if the application is activated
func IsActivated() bool {
	_, err := LoadActivationKey()
	return err == nil
}
