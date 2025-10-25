package client

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Invoice struct {
	IDHoaDon    string
	SeriKyHieu  string
	MaKH        string // TraCuuHD_MA_KH
	StrIDHoaDon string // strid_hdon
	Nam         int
	Thang       int
	Ky          int
}

// SearchInvoices fetches invoice table and parses it using HTML parser
func (c *Client) SearchInvoices(ky, thang, nam int) ([]Invoice, error) {
	url := fmt.Sprintf("https://cskh.npc.com.vn/HoaDon/TraCuuHDSPC?ky=%d&thang=%d&nam=%d&_=%d", ky, thang, nam, time.Now().UnixMilli())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "text/html, */*; q=0.01")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Referer", "https://cskh.npc.com.vn/DichVuTTCSKH/IndexNPC?index=2")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to search invoices: %s", resp.Status)
	}

	return parseInvoiceTable(resp.Body)
}

// parseInvoiceTable extracts invoice data (IDHoaDon, SeriKyHieu) from HTML table
func parseInvoiceTable(r io.Reader) ([]Invoice, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var invoices []Invoice

	// Traverse to find <table class="table2">
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "table2") {
					rows := extractRows(n)
					invoices = append(invoices, rows...)
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if len(invoices) == 0 {
		return nil, fmt.Errorf("no invoice rows found in table")
	}
	return invoices, nil
}

// extractRows finds all <tr> in <tbody> and extracts IDHoaDon and SeriKyHieu
func extractRows(table *html.Node) []Invoice {
	var invoices []Invoice

	var tbody *html.Node
	for c := table.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "tbody" {
			tbody = c
			break
		}
	}
	if tbody == nil {
		return invoices
	}

	for tr := tbody.FirstChild; tr != nil; tr = tr.NextSibling {
		if tr.Type != html.ElementNode || tr.Data != "tr" {
			continue
		}

		tdTexts := []string{}
		for td := tr.FirstChild; td != nil; td = td.NextSibling {
			if td.Type == html.ElementNode && td.Data == "td" {
				text := extractText(td)
				tdTexts = append(tdTexts, text)
			}
		}

		if len(tdTexts) >= 4 {
			idHoaDon := strings.TrimSpace(tdTexts[2])
			seriKyHieu := strings.TrimSpace(tdTexts[3])

			invoice := Invoice{
				IDHoaDon:   idHoaDon,
				SeriKyHieu: seriKyHieu,
			}

			// Search for the button within the current tr
			var f func(*html.Node)
			f = func(n *html.Node) {
				if n.Type == html.ElementNode && n.Data == "button" {
					for _, attr := range n.Attr {
						if attr.Key == "id" && attr.Val == "TraCuuHD_XemHoaDon" {
							for _, onclickAttr := range n.Attr {
								if onclickAttr.Key == "onclick" {
									maKH, strIDHoaDon, ky, thang, nam, err := parseOnclick(onclickAttr.Val)
									if err == nil {
										invoice.MaKH = maKH
										invoice.StrIDHoaDon = strIDHoaDon
										invoice.Ky = ky
										invoice.Thang = thang
										invoice.Nam = nam
									}
									break
								}
							}
						}
					}
				}
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
			f(tr) // Start searching from the current tr

			invoices = append(invoices, invoice)
		}
	}

	return invoices
}
func parseOnclick(onclick string) (maKH, strIDHoaDon string, ky, thang, nam int, err error) {
	// Expected format: ChonThaoTac('strid_hdon-ma_kh-ky-thang-nam');

	start := strings.Index(onclick, "'")
	end := strings.LastIndex(onclick, "'")
	if start == -1 || end == -1 || start == end {
		return "", "", 0, 0, 0, fmt.Errorf("invalid onclick format: %s", onclick)
	}

	data := onclick[start+1 : end]
	parts := strings.Split(data, "-")

	if len(parts) != 5 {
		return "", "", 0, 0, 0, fmt.Errorf("invalid data parts in onclick: %s", data)
	}

	strIDHoaDon = parts[0]
	maKH = parts[1]

	ky, err = strconv.Atoi(parts[2])
	if err != nil {
		return "", "", 0, 0, 0, fmt.Errorf("invalid ky in onclick: %w", err)
	}
	thang, err = strconv.Atoi(parts[3])
	if err != nil {
		return "", "", 0, 0, 0, fmt.Errorf("invalid thang in onclick: %w", err)
	}
	nam, err = strconv.Atoi(parts[4])
	if err != nil {
		return "", "", 0, 0, 0, fmt.Errorf("invalid nam in onclick: %w", err)
	}

	return maKH, strIDHoaDon, ky, thang, nam, nil
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(extractText(c))
	}
	return sb.String()
}
