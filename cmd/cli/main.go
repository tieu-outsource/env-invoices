package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"env-invoices/client"
	"env-invoices/utils"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "download":
		download()
	case "cc":
		convertCaptcha()
	default:
		fmt.Println("Usage: cli download")
	}
}

func download() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: cli login <username> <password>")
	}

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

	username := os.Args[2]
	password := os.Args[3]

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
	zipFile, err := c.DownloadInvoice(invoice)
	if err != nil {
		log.Fatalf("Failed to download invoice: %v", err)
	}

	// save zip file to invoice.zip
	err = os.WriteFile("invoice.zip", zipFile, 0644)
	if err != nil {
		log.Fatalf("Failed to save zip file: %v", err)
	}
}

func convertCaptcha() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: cli cc <imgpath> <key>")
	}

	log.Printf("Reading image from %s", os.Args[2])
	img, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	imgBase64 := base64.StdEncoding.EncodeToString(img)

	key := os.Args[3]
	out, err := utils.ResolveCaptcha(imgBase64, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
