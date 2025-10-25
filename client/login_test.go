package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetLoginDetails(t *testing.T) {
	// Sample HTML from the login page
	const sampleHTML = `
		<html>
			<body>
				<form>
					<input name="__RequestVerificationToken" type="hidden" value="test_token" />
					<input name="CaptchaDeText" type="hidden" value="test_captcha_de_text" />
					<img id="CaptchaImage" src="/captcha/123.jpg" />
				</form>
			</body>
		</html>
	`

	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/home/AccountNPC" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, sampleHTML)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Create a client with the mock server's URL
	c, err := New()
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	// Override the BaseURL to point to our mock server
	originalBaseURL := BaseURL
	BaseURL = server.URL
	defer func() { BaseURL = originalBaseURL }()

	// --- Test successful case ---
	t.Run("successful extraction", func(t *testing.T) {
		details, err := c.GetLoginDetails()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if details.Token != "test_token" {
			t.Errorf("expected token 'test_token', got '%s'", details.Token)
		}
		if details.CaptchaDeText != "test_captcha_de_text" {
			t.Errorf("expected captchaDeText 'test_captcha_de_text', got '%s'", details.CaptchaDeText)
		}
		expectedCaptchaURL := server.URL + "/captcha/123.jpg"
		if details.CaptchaImgURL != expectedCaptchaURL {
			t.Errorf("expected captchaImgURL '%s', got '%s'", expectedCaptchaURL, details.CaptchaImgURL)
		}
	})

	// --- Test server error case ---
	t.Run("server error", func(t *testing.T) {
		// A bit of a hack to simulate a server error: create a new server, close it immediately,
		// and point the client to it.
		errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		BaseURL = errorServer.URL
		errorServer.Close()

		_, err := c.GetLoginDetails()
		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
	})
}

func Test_extractLoginDetails(t *testing.T) {
	t.Run("successful extraction", func(t *testing.T) {
		const html = `
			<html>
				<body>
					<form>
						<input name="__RequestVerificationToken" type="hidden" value="test_token" />
						<input name="CaptchaDeText" type="hidden" value="test_captcha_de_text" />
						<img id="CaptchaImage" src="/captcha/123.jpg" />
					</form>
				</body>
			</html>
		`
		reader := strings.NewReader(html)
		token, captchaDeText, captchaImgURL, err := _extractLoginDetails(reader)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if token != "test_token" {
			t.Errorf("expected token 'test_token', got '%s'", token)
		}
		if captchaDeText != "test_captcha_de_text" {
			t.Errorf("expected captchaDeText 'test_captcha_de_text', got '%s'", captchaDeText)
		}
		if captchaImgURL != "/captcha/123.jpg" {
			t.Errorf("expected captchaImgURL '/captcha/123.jpg', got '%s'", captchaImgURL)
		}
	})

	t.Run("token not found", func(t *testing.T) {
		const html = `<html><body><input name="CaptchaDeText" value="v" /><img id="CaptchaImage" src="s" /></body></html>`
		reader := strings.NewReader(html)
		_, _, _, err := _extractLoginDetails(reader)
		if err == nil || !strings.Contains(err.Error(), "__RequestVerificationToken not found") {
			t.Errorf("expected error '__RequestVerificationToken not found', got '%v'", err)
		}
	})

	t.Run("captcha de text not found", func(t *testing.T) {
		const html = `<html><body><input name="__RequestVerificationToken" value="v" /><img id="CaptchaImage" src="s" /></body></html>`
		reader := strings.NewReader(html)
		_, _, _, err := _extractLoginDetails(reader)
		if err == nil || !strings.Contains(err.Error(), "CaptchaDeText not found") {
			t.Errorf("expected error 'CaptchaDeText not found', got '%v'", err)
		}
	})

	t.Run("captcha image not found", func(t *testing.T) {
		const html = `<html><body><input name="__RequestVerificationToken" value="v" /><input name="CaptchaDeText" value="v" /></body></html>`
		reader := strings.NewReader(html)
		_, _, _, err := _extractLoginDetails(reader)
		if err == nil || !strings.Contains(err.Error(), "CaptchaImage not found") {
			t.Errorf("expected error 'CaptchaImage not found', got '%v'", err)
		}
	})
}
