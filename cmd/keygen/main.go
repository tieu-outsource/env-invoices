package main

import (
	"bufio"
	"env-invoices/activation"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== EVN Invoice Downloader - Activation Key Generator ===")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the Device ID from the user:")
	fmt.Print("> ")

	deviceID, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	deviceID = strings.TrimSpace(deviceID)

	if deviceID == "" {
		fmt.Println("Device ID cannot be empty")
		os.Exit(1)
	}

	fmt.Println("Total days:")
	days, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nGenerating activation key for device: %s\n", deviceID)

	dayNum, err := strconv.Atoi(strings.TrimSpace(days))

	activationKey, err := activation.GenerateActivationKey(deviceID, dayNum)
	if err != nil {
		fmt.Printf("Error generating activation key: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n=== ACTIVATION KEY (send this to the user) ===")
	fmt.Println(activationKey)
	fmt.Println("==============================================")
	fmt.Println("\nThe user should paste this key into the activation dialog.")
}
