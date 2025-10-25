package main

import (
	"env-invoices/activation"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Testing Activation System ===\n")

	// Test 1: Get Device ID
	fmt.Println("1. Getting Device ID...")
	deviceID, err := activation.GetDeviceID()
	if err != nil {
		fmt.Printf("   ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("   ✓ Device ID: %s\n\n", deviceID)

	// Test 2: Generate activation key
	fmt.Println("2. Generating activation key...")
	key, err := activation.GenerateActivationKey(deviceID)
	if err != nil {
		fmt.Printf("   ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("   ✓ Key generated: %s...\n\n", key[:50])

	// Test 3: Validate the key
	fmt.Println("3. Validating activation key...")
	err = activation.ValidateActivationKey(key)
	if err != nil {
		fmt.Printf("   ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("   ✓ Key is valid!\n")

	// Test 4: Save and load key
	fmt.Println("4. Testing save and load...")
	err = activation.SaveActivationKey(key)
	if err != nil {
		fmt.Printf("   ERROR saving: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("   ✓ Key saved to activation.key")

	loadedKey, err := activation.LoadActivationKey()
	if err != nil {
		fmt.Printf("   ERROR loading: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("   ✓ Key loaded successfully\n\n")

	// Test 5: Check activation status
	fmt.Println("5. Checking activation status...")
	if activation.IsActivated() {
		fmt.Println("   ✓ Application is activated!\n")
	} else {
		fmt.Println("   ERROR: Application is not activated\n")
		os.Exit(1)
	}

	// Test 6: Try to validate a wrong device ID key
	fmt.Println("6. Testing with wrong device ID...")
	wrongKey, _ := activation.GenerateActivationKey("wrong-device-id-12345678")
	err = activation.ValidateActivationKey(wrongKey)
	if err != nil {
		fmt.Printf("   ✓ Correctly rejected invalid key: %v\n\n", err)
	} else {
		fmt.Println("   ERROR: Should have rejected key for wrong device!\n")
		os.Exit(1)
	}

	fmt.Println("=== All Tests Passed! ===")
	fmt.Printf("\nYour device ID for this machine is:\n%s\n", deviceID)
	fmt.Printf("\nYour activation key is:\n%s\n", loadedKey)
	
	// Clean up test file
	os.Remove("activation.key")
	fmt.Println("\nTest activation.key file removed.")
}
