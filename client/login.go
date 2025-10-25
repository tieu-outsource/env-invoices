package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// LoginDetails holds the information extracted from the login page
// that is required to perform a login.
type LoginDetails struct {
	Token         string
	CaptchaDeText string
	CaptchaImgURL string
}

// GetLoginDetails fetches the login page and extracts the necessary details for logging in.
func (c *Client) GetLoginDetails() (*LoginDetails, error) {
	// GET the AccountNPC page to get the login form, captcha, and token.
	getResp, err := c.client.Get(BaseURL + "/home/AccountNPC")
	if err != nil {
		return nil, fmt.Errorf("failed to get account page: %w", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get account page: status code %d", getResp.StatusCode)
	}

	token, captchaDeText, captchaImgURL, err := _extractLoginDetails(getResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to extract login details from /home/AccountNPC: %w", err)
	}

	return &LoginDetails{
		Token:         token,
		CaptchaDeText: captchaDeText,
		CaptchaImgURL: BaseURL + captchaImgURL, // Make it an absolute URL
	}, nil
}

// Get login captcha image
func (c *Client) DownloadLoginCaptcha(fileUrl string) ([]byte, error) {
	u, _ := url.Parse(fileUrl)
	cookies := c.client.Jar.Cookies(u)

	for _, cookie := range cookies {
		fmt.Printf(" ___ %s = %s\n", cookie.Name, cookie.Value)
	}

	downloadResq, err := http.NewRequest("GET", fileUrl, nil)
	downloadResq.Header.Set("Referer", "https://cskh.npc.com.vn/home/AccountNPC")

	resp, err := c.client.Do(downloadResq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}


// Login performs the login action.
func (c *Client) Login(details *LoginDetails, username, password, captcha string) error {
	// POST to login
	form := url.Values{}
	form.Set("__RequestVerificationToken", details.Token)
	form.Set("Username", username)
	form.Set("Password", password)
	form.Set("CaptchaDeText", details.CaptchaDeText)
	form.Set("CaptchaInputText", captcha)
	form.Set("previousLink", "")

	postReq, err := http.NewRequest("POST", BaseURL+"/Account/Login", strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create login request: %w", err)
	}

	postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	postReq.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0")
	postReq.Header.Set("Referer", BaseURL+"/home/AccountNPC")

	postResp, err := c.client.Do(postReq)
	if err != nil {
		return fmt.Errorf("login request failed: %w", err)
	}
	defer postResp.Body.Close()

	if postResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(postResp.Body)
		return fmt.Errorf("login failed: status code %d, body: %s", postResp.StatusCode, string(body))
	}

	bodyBytes, err := io.ReadAll(postResp.Body)
	if err != nil {
		return fmt.Errorf("failed to read login response body: %w", err)
	}

	if strings.Contains(string(bodyBytes), "login-form") {
		return fmt.Errorf("login failed, response contains login form")
	}

	return nil
}

func _extractLoginDetails(body io.Reader) (string, string, string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return "", "", "", err
	}

	var token, captchaDeText, captchaImgURL string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			var value string
			isToken := false
			isCaptcha := false
			for _, a := range n.Attr {
				if a.Key == "name" && a.Val == "__RequestVerificationToken" {
					isToken = true
				}
				if a.Key == "name" && a.Val == "CaptchaDeText" {
					isCaptcha = true
				}
				if a.Key == "value" {
					value = a.Val
				}
			}
			if isToken {
				token = value
			}
			if isCaptcha {
				captchaDeText = value
			}
		}
		if n.Type == html.ElementNode && n.Data == "img" {
			isCaptchaImg := false
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "CaptchaImage" {
					isCaptchaImg = true
				}
			}
			if isCaptchaImg {
				for _, a := range n.Attr {
					if a.Key == "src" {
						captchaImgURL = a.Val
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if token == "" {
		return "", "", "", fmt.Errorf("__RequestVerificationToken not found")
	}
	if captchaDeText == "" {
		return "", "", "", fmt.Errorf("CaptchaDeText not found")
	}
	if captchaImgURL == "" {
		return "", "", "", fmt.Errorf("CaptchaImage not found")
	}

	return token, captchaDeText, captchaImgURL, nil
}
