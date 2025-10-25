package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//	await fetch("https://cskh.npc.com.vn/HoaDon/DownloadHD1", {
//	    "credentials": "include",
//	    "headers": {
//	        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0",
//	        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
//	        "Accept-Language": "en-US,en;q=0.5",
//	        "Content-Type": "application/x-www-form-urlencoded",
//	        "Upgrade-Insecure-Requests": "1",
//	        "Sec-Fetch-Dest": "document",
//	        "Sec-Fetch-Mode": "navigate",
//	        "Sec-Fetch-Site": "same-origin",
//	        "Sec-Fetch-User": "?1",
//	        "Idempotency-Key": "\"1200986034772628023\"",
//	        "Priority": "u=0, i",
//	        "Pragma": "no-cache",
//	        "Cache-Control": "no-cache"
//	    },
//	    "referrer": "https://cskh.npc.com.vn/DichVuTTCSKH/IndexNPC?index=2",
//	    "body": "TraCuuHD_MA_KH=PA04PB9042895&TraCuuHD_IDHoaDon_MauMoi=&strid_hdon=CRcd___4feOQGDRvXpr1Ajjw%3D%3D&nam=2025&thang=9&ky=1&TenThuMuc=",
//	    "method": "POST",
//	    "mode": "cors"
//	});
func (c *Client) DownloadInvoice(invoice Invoice) (string, error) {
	form := url.Values{}
	form.Set("TraCuuHD_MA_KH", invoice.MaKH)
	form.Set("TraCuuHD_IDHoaDon_MauMoi", "")
	form.Set("strid_hdon", invoice.StrIDHoaDon)
	form.Set("nam", strconv.Itoa(int(invoice.Nam)))
	form.Set("thang", strconv.Itoa(int(invoice.Thang)))
	form.Set("ky", strconv.Itoa(int(invoice.Ky)))
	form.Set("TenThuMuc", "")

	req, err := http.NewRequest(
		"POST",
		BaseURL+"/HoaDon/DownloadHD1",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Referer", "https://cskh.npc.com.vn/DichVuTTCSKH/IndexNPC?index=2")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request body: %s", form.Encode())
		return "", fmt.Errorf("failed to download invoice: status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Define output file name
	filePath := "invoice.zip"

	// Write the bytes to the file
	err = os.WriteFile(filePath, bodyBytes, 0644)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
