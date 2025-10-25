package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"env-invoices/client"
)

func main() {
	c, err := client.New()
	if err != nil {
		log.Fatal(err)
	}

	// Get login details
	details, err := c.GetLoginDetails()
	if err != nil {
		log.Fatalf("Failed to get login details: %v", err)
	}
	fmt.Printf("Login details: %+v\n", details)

	// Download captcha image
	imgData, err := c.DownloadLoginCaptcha(details.CaptchaImgURL)
	if err != nil {
		log.Fatalf("Failed to download captcha image: %v", err)
	}

	// Save captcha image to a temporary file
	tmpfile, err := os.CreateTemp("", "captcha-*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(imgData); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// Open the captcha image for the user
	cmd := exec.Command("xdg-open", tmpfile.Name())
	err = cmd.Start()
	if err != nil {
		log.Printf("Failed to open captcha image: %v", err)
		fmt.Printf("Please open the captcha image manually: %s\n", tmpfile.Name())
	}

	// Prompt for captcha
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter CAPTCHA: ")
	captcha, _ := reader.ReadString('\n')
	captcha = strings.TrimSpace(captcha)

	// Hardcode credentials for now
	username := "PA04GT7017040"
	password := "123456"

	// Login
	err = c.Login(details, username, password, captcha)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	fmt.Println("Login successful")

	// Search invoices
	invoices, err := c.SearchInvoices(1, 9, 2025)
	if err != nil {
		log.Fatalf("Failed to search invoices: %v", err)
	}

	for _, invoice := range invoices {
		fmt.Printf("Invoice: %+v\n", invoice)
	}

	invoice := invoices[0]
	base64ZipFile, err := c.DownloadInvoice(invoice)
	if err != nil {
		log.Fatalf("Failed to download invoice: %v", err)
	}

	fmt.Printf("Downloaded invoice: %s\n", base64ZipFile)
}
