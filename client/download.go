package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) DownloadInvoice(invoice Invoice) ([]byte, error) {
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
		return nil, err
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
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request body: %s", form.Encode())
		return nil, fmt.Errorf("failed to download invoice: status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return bodyBytes, nil
}
