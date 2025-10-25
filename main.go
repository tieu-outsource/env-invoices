package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/gif"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"env-invoices/activation"
	"env-invoices/client"
	"env-invoices/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

type user struct {
	username        string
	password        string
	status          string
	log             []string
	invoiceCount    int
	downloadedCount int
	err             error
	processing      bool
}

var (
	users                   []*user
	table                   *widget.Table
	mainW                   fyne.Window
	totalInvoicesToDownload int
	totalInvoicesDownloaded int
	progressMutex           sync.Mutex
	totalProgressLabel      *widget.Label
	kyEntry                 *widget.Entry
	thangEntry              *widget.Entry
	namEntry                *widget.Entry
	downloadDir             string
	finalDownloadDir        string
	apiKeyEntry             *widget.Entry
)

func main() {
	a := app.New()
	mainW = a.NewWindow("Bulk Invoice Downloader")

	apiKeyEntry = widget.NewPasswordEntry()
	apiKeyEntry.SetPlaceHolder("Enter Captcha Solver API Key")

	if savedKey, err := loadCaptchaKey(); err == nil {
		apiKeyEntry.SetText(savedKey)
	}

	now := time.Now()
	kyEntry = widget.NewEntry()
	kyEntry.SetText("1")
	thangEntry = widget.NewEntry()
	thangEntry.SetText(fmt.Sprintf("%d", now.Month()))
	namEntry = widget.NewEntry()
	namEntry.SetText(fmt.Sprintf("%d", now.Year()))

	filterForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Kỳ", Widget: kyEntry},
			{Text: "Tháng", Widget: thangEntry},
			{Text: "Năm", Widget: namEntry},
		},
	}

	homeDir, _ := os.UserHomeDir()
	if homeDir != "" {
		downloadDir = filepath.Join(homeDir, "Downloads")
	}
	downloadDirLabel := widget.NewLabel(downloadDir)
	selectFolderButton := widget.NewButton("Select Download Folder", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, mainW)
				return
			}
			if uri != nil {
				downloadDir = uri.Path()
				downloadDirLabel.SetText(downloadDir)
			}
		}, mainW)
	})

	table = widget.NewTable(
		func() (int, int) { return len(users), 5 },
		func() fyne.CanvasObject { return container.NewMax(widget.NewLabel("wide content")) },
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row >= len(users) {
				return
			}
			u := users[i.Row]
			max := o.(*fyne.Container)
			switch i.Col {
			case 0:
				max.Objects = []fyne.CanvasObject{widget.NewLabel(u.username)}
			case 1:
				if u.processing {
					max.Objects = []fyne.CanvasObject{widget.NewProgressBarInfinite()}
				} else {
					max.Objects = []fyne.CanvasObject{widget.NewLabel(u.status)}
				}
			case 2:
				max.Objects = []fyne.CanvasObject{widget.NewLabel(fmt.Sprintf("%d/%d", u.downloadedCount, u.invoiceCount))}
			case 3:
				if u.err != nil {
					viewLogBtn := widget.NewButton("View Log", func() {
						logContent := strings.Join(u.log, "\n")

						// Create a multi-line text widget (read-only)
						textEntry := widget.NewMultiLineEntry()
						textEntry.SetText(logContent)

						// Wrap inside a scrollable container
						scroll := container.NewVScroll(textEntry)
						scroll.SetMinSize(fyne.NewSize(600, 400)) // set dialog size

						// Show in dialog
						dialog.ShowCustom("User Log", "Close", scroll, mainW)
					})

					max.Objects = []fyne.CanvasObject{viewLogBtn}
				} else {
					max.Objects = []fyne.CanvasObject{widget.NewLabel("")}
				}
			case 4:
				if u.err != nil {
					button := widget.NewButton("Retry", func() {
						ky, _ := strconv.Atoi(kyEntry.Text)
						thang, _ := strconv.Atoi(thangEntry.Text)
						nam, _ := strconv.Atoi(namEntry.Text)
						go downloadForUser(u, apiKeyEntry.Text, ky, thang, nam, finalDownloadDir)
					})
					max.Objects = []fyne.CanvasObject{button}
				} else {
					max.Objects = []fyne.CanvasObject{widget.NewLabel("")}
				}
			}
			max.Refresh()
		},
	)
	table.SetColumnWidth(0, 150)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 100)
	table.SetColumnWidth(4, 100)

	selectFileButton := widget.NewButton("Select Excel File", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, mainW)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()
			f, err := excelize.OpenReader(reader)
			if err != nil {
				dialog.ShowError(err, mainW)
				return
			}
			rows, err := f.GetRows(f.GetSheetName(0))
			if err != nil {
				dialog.ShowError(err, mainW)
				return
			}
			users = []*user{}
			for _, row := range rows[1:] {
				if len(row) >= 2 {
					users = append(users, &user{username: row[0], password: row[1], status: "Pending"})
				}
			}
			table.Refresh()
		}, mainW)
		fileDialog.Show()
	})

	var startButton *widget.Button
	startButton = widget.NewButton("Start Download", func() {
		// Check activation before allowing download
		if !activation.IsActivated() {
			dialog.ShowInformation("Activation Required",
				"This application is not activated. Please activate it first.", mainW)
			return
		}

		apiKey := apiKeyEntry.Text
		if apiKey == "" {
			dialog.ShowError(fmt.Errorf("please enter your captcha solver API key"), mainW)
			return
		}

		key := strings.TrimSpace(apiKeyEntry.Text)
		if key != "" {
			_ = saveCaptchaKey(key)
		}

		ky, err := strconv.Atoi(kyEntry.Text)
		if err != nil {
			dialog.ShowError(fmt.Errorf("invalid value for Kỳ"), mainW)
			return
		}
		thang, err := strconv.Atoi(thangEntry.Text)
		if err != nil {
			dialog.ShowError(fmt.Errorf("invalid value for Tháng"), mainW)
			return
		}
		nam, err := strconv.Atoi(namEntry.Text)
		if err != nil {
			dialog.ShowError(fmt.Errorf("invalid value for Năm"), mainW)
			return
		}

		baseDirName := fmt.Sprintf("EVN_%d_%d_%d", ky, thang, nam)
		fullPath := filepath.Join(downloadDir, baseDirName)
		counter := 1
		for {
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				finalDownloadDir = fullPath
				break
			}
			fullPath = filepath.Join(downloadDir, fmt.Sprintf("%s_%d", baseDirName, counter))
			counter++
		}

		if err := os.MkdirAll(finalDownloadDir, 0755); err != nil {
			dialog.ShowError(err, mainW)
			return
		}

		startButton.Disable()
		startButton.SetText("Downloading...")

		progressMutex.Lock()
		totalInvoicesToDownload = 0
		totalInvoicesDownloaded = 0
		progressMutex.Unlock()
		updateTotalProgressLabel()

		go func() {
			var wg sync.WaitGroup
			for _, u := range users {
				wg.Add(1)
				go func(u *user) {
					defer wg.Done()
					downloadForUser(u, apiKey, ky, thang, nam, finalDownloadDir)
				}(u)
			}
			wg.Wait()
			fyne.Do(func() {
				startButton.Enable()
				startButton.SetText("Start Download")
			})
		}()
	})

	totalProgressLabel = widget.NewLabel("Total Progress: 0/0")

	// Add activation button
	var activationButton *widget.Button // declare first
	activationButton = widget.NewButton("Activation", func() {
		showActivationDialog(activationButton)
	})

	if activation.IsActivated() {
		activationButton.Hide()
	}

	topContent := container.NewVBox(apiKeyEntry, filterForm, selectFolderButton, downloadDirLabel, selectFileButton, activationButton)
	bottomContent := container.NewVBox(totalProgressLabel, startButton)

	content := container.NewBorder(topContent, bottomContent, nil, nil, table)

	mainW.SetContent(content)
	mainW.Resize(fyne.NewSize(800, 600))
	mainW.ShowAndRun()
}

func downloadForUser(u *user, apiKey string, ky, thang, nam int, downloadDir string) {
	updateUser(u, func(u *user) {
		u.processing = true
		u.status = "In Progress"
		u.err = nil
		u.log = []string{"Starting download..."}
		u.downloadedCount = 0
	})

	c, err := client.New()
	if err != nil {
		updateUser(u, func(u *user) {
			u.processing = false
			u.status = "Error"
			u.err = err
			u.log = append(u.log, err.Error())
		})
		return
	}

	var loginSuccessful bool
	var details *client.LoginDetails

	for i := 0; i < 3; i++ {
		details, err = c.GetLoginDetails()
		if err != nil {
			updateUser(u, func(u *user) {
				u.log = append(u.log, fmt.Sprintf("Attempt %d: Failed to get login details: %v", i+1, err))
			})
			continue
		}
		imgData, err := c.DownloadLoginCaptcha(details.CaptchaImgURL)
		if err != nil {
			updateUser(u, func(u *user) {
				u.log = append(u.log, fmt.Sprintf("Attempt %d: Failed to download captcha: %v", i+1, err))
			})
			continue
		}
		imgBase64 := base64.StdEncoding.EncodeToString(imgData)
		captcha, err := utils.ResolveCaptcha(imgBase64, apiKey)
		if err != nil {
			updateUser(u, func(u *user) {
				u.log = append(u.log, fmt.Sprintf("Attempt %d: Failed to resolve captcha: %v", i+1, err))
			})
			continue
		}
		err = c.Login(details, u.username, u.password, captcha)
		if err == nil {
			loginSuccessful = true
			break
		} else if err == client.ErrWrongCaptcha {
			updateUser(u, func(u *user) { u.log = append(u.log, fmt.Sprintf("Attempt %d: Wrong captcha", i+1)) })
		} else {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = err
				u.log = append(u.log, fmt.Sprintf("Login failed: %v", err))
			})
			return
		}
	}

	if !loginSuccessful {
		updateUser(u, func(u *user) { u.log = append(u.log, "Automatic captcha solving failed. Asking for manual input.") })
		details, err = c.GetLoginDetails()
		if err != nil {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = err
				u.log = append(u.log, "Failed to get login details for manual captcha: "+err.Error())
			})
			return
		}
		imgData, err := c.DownloadLoginCaptcha(details.CaptchaImgURL)
		if err != nil {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = err
				u.log = append(u.log, "Failed to download captcha for manual input: "+err.Error())
			})
			return
		}
		captchaChan := make(chan string)
		askForManualCaptcha(imgData, captchaChan)
		manualCaptcha := <-captchaChan
		if manualCaptcha == "" {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = fmt.Errorf("manual captcha input cancelled")
				u.log = append(u.log, "Manual captcha input cancelled.")
			})
			return
		}
		err = c.Login(details, u.username, u.password, manualCaptcha)
		if err != nil {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = err
				u.log = append(u.log, fmt.Sprintf("Manual captcha login failed: %v", err))
			})
			return
		}
	}

	log.Printf("Login successful for %s", u.username)
	updateUser(u, func(u *user) { u.log = append(u.log, "Login successful") })

	invoices, err := c.SearchInvoices(ky, thang, nam)
	if err != nil {
		updateUser(u, func(u *user) {
			u.processing = false
			u.status = "Error"
			u.err = err
			u.log = append(u.log, "Failed to search invoices: "+err.Error())
		})
		return
	}
	progressMutex.Lock()
	totalInvoicesToDownload += len(invoices)
	progressMutex.Unlock()
	updateUser(u, func(u *user) { u.invoiceCount = len(invoices) })

	if len(invoices) == 0 {
		updateUser(u, func(u *user) {
			u.processing = false
			u.status = "No invoices"
		})
		return
	}

	for _, invoice := range invoices {
		zipFile, err := c.DownloadInvoice(invoice)
		if err != nil {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = err
				u.log = append(u.log, fmt.Sprintf("Failed to download invoice %s: %v", invoice.IDHoaDon, err))
			})
			return
		}

		// fileName := filepath.Join(downloadDir, fmt.Sprintf("%s_invoice_%s.zip", u.username, invoice.IDHoaDon))
		// err = os.WriteFile(fileName, zipFile, 0644)

		extractDir := filepath.Join(downloadDir, fmt.Sprintf("%s_%s", u.username, invoice.IDHoaDon))
		if err := os.MkdirAll(extractDir, 0755); err != nil {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = err
				u.log = append(u.log, fmt.Sprintf("Failed to create folder %s: %v", extractDir, err))
			})
			return
		}

		// Extract the ZIP content into the folder
		if err := extractZipBytes(zipFile, extractDir); err != nil {
			updateUser(u, func(u *user) {
				u.processing = false
				u.status = "Error"
				u.err = err
				u.log = append(u.log, fmt.Sprintf("Failed to extract invoice %s: %v", invoice.IDHoaDon, err))
			})
			return
		}

		progressMutex.Lock()
		totalInvoicesDownloaded++
		progressMutex.Unlock()
		updateUser(u, func(u *user) {
			u.downloadedCount++
			u.log = append(u.log, fmt.Sprintf("Saved invoice %s", invoice.IDHoaDon))
		})
	}

	updateUser(u, func(u *user) {
		u.processing = false
		u.status = "Completed"
	})
}

// imgData: raw bytes from DownloadLoginCaptcha
func askForManualCaptcha(imgData []byte, captchaChan chan string) {
	fyne.Do(func() {
		// Detect if image is GIF
		var displayBytes []byte
		if _, err := gif.Decode(bytes.NewReader(imgData)); err == nil {
			// It's a GIF -> convert first frame to PNG
			gifImg, err := gif.Decode(bytes.NewReader(imgData))
			if err != nil {
				log.Println("Failed to decode GIF captcha:", err)
				captchaChan <- ""
				close(captchaChan)
				return
			}
			buf := new(bytes.Buffer)
			if err := png.Encode(buf, gifImg); err != nil {
				log.Println("Failed to encode PNG captcha:", err)
				captchaChan <- ""
				close(captchaChan)
				return
			}
			displayBytes = buf.Bytes()
		} else {
			// Not GIF, use raw bytes
			displayBytes = imgData
		}

		imgReader := bytes.NewReader(displayBytes)
		img := canvas.NewImageFromReader(imgReader, "captcha.png")
		img.FillMode = canvas.ImageFillOriginal
		img.SetMinSize(fyne.NewSize(200, 80)) // adjust as needed

		entry := widget.NewEntry()
		entry.SetPlaceHolder("Enter captcha here")

		content := container.NewVBox(img, entry)

		dialog.ShowCustomConfirm("Enter Captcha", "Submit", "Cancel", content, func(ok bool) {
			if ok {
				captchaChan <- entry.Text
			} else {
				captchaChan <- ""
			}
			close(captchaChan)
		}, mainW)
	})
}

func updateUser(u *user, updateFunc func(*user)) {
	fyne.Do(func() {
		updateFunc(u)
		table.Refresh()
		updateTotalProgressLabel()
	})
}

func updateTotalProgressLabel() {
	progressMutex.Lock()
	defer progressMutex.Unlock()
	totalProgressLabel.SetText(fmt.Sprintf("Total Progress: %d/%d", totalInvoicesDownloaded, totalInvoicesToDownload))
}

// SaveCaptchaKey writes the API key to a file
func saveCaptchaKey(key string) error {
	return os.WriteFile("captcha_key.txt", []byte(strings.TrimSpace(key)), 0600)
}

// LoadCaptchaKey reads the API key from file (if exists)
func loadCaptchaKey() (string, error) {
	data, err := os.ReadFile("captcha_key.txt")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func extractZipBytes(data []byte, dest string) error {
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		fpath := filepath.Join(dest, f.Name)

		// Prevent zip-slip attack
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// showActivationDialog displays the activation dialog to the user
func showActivationDialog(activationButton *widget.Button) {
	deviceID, err := activation.GetDeviceID()
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to get device ID: %w", err), mainW)
		return
	}

	deviceIDLabel := widget.NewLabel("Your Device ID:")
	deviceIDEntry := widget.NewEntry()
	deviceIDEntry.SetText(deviceID)
	deviceIDEntry.Disable()

	copyButton := widget.NewButton("Copy Device ID", func() {
		mainW.Clipboard().SetContent(deviceID)
		dialog.ShowInformation("Copied", "Device ID copied to clipboard!", mainW)
	})

	instructionLabel := widget.NewLabel(
		"To activate this application:\n" +
			"1. Copy your Device ID using the button above\n" +
			"2. Send it to the administrator\n" +
			"3. Paste the activation key you receive below")
	instructionLabel.Wrapping = fyne.TextWrapWord

	activationKeyEntry := widget.NewMultiLineEntry()
	activationKeyEntry.SetPlaceHolder("Paste your activation key here...")
	activationKeyEntry.SetMinRowsVisible(3)

	var activationDialog dialog.Dialog

	activateButton := widget.NewButton("Activate", func() {
		activationKey := strings.TrimSpace(activationKeyEntry.Text)
		if activationKey == "" {
			dialog.ShowError(fmt.Errorf("please enter an activation key"), mainW)
			return
		}

		err := activation.SaveActivationKey(activationKey)
		if err != nil {
			dialog.ShowError(fmt.Errorf("activation failed: %w", err), mainW)
			return
		}

		// Hide activation button after success
		activationButton.Hide()

		dialog.ShowInformation("Success", "Application activated successfully!", mainW)
		activationDialog.Hide()
	})

	closeButton := widget.NewButton("Close", func() {
		activationDialog.Hide()
	})

	content := container.NewVBox(
		widget.NewLabel("Application Activation Required"),
		widget.NewSeparator(),
		deviceIDLabel,
		deviceIDEntry,
		copyButton,
		widget.NewSeparator(),
		instructionLabel,
		widget.NewLabel("Activation Key:"),
		activationKeyEntry,
		widget.NewSeparator(),
		container.NewHBox(activateButton, closeButton),
	)

	activationDialog = dialog.NewCustom("Activation", "Close", content, mainW)
	activationDialog.Resize(fyne.NewSize(500, 450))
	activationDialog.Show()
}
