package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// MockHTTPClient is a mock implementation of http.Client for testing.
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestSearchInvoices(t *testing.T) {
	sampleHTML := `
<table class="uk-table uk-table-small uk-table-middle uk-table-divider table2">
    <thead>
        <tr>
            <th>STT</th>
            <th>M&#227; kh&#225;ch h&#224;ng</th>
            <th>ID h&#243;a đơn</th>
            <th>Seri-K&#253; hiệu</th>
            <th>Tổng Tiền</th>
            <th>Loại h&#243;a đơn</th>
            <th>Xem thể hiện</th>
        </tr>
    </thead>
    <tbody>
                    <tr>
                        <td>1</td>
                        <td class="uk-text-truncate">PA04GT7017040</td>
                        <td>1671732781</td>
                        <td class="uk-text-truncate">1046124</td>
                        <td>
28.386.839
                        </td>
                            <td>Tiền Điện</td>
                        <td>
                                <button 
                                        id="TraCuuHD_XemHoaDon"
                                        onclick="ChonThaoTac('XIq6RC1EGQud6DC___PGjwiQ==-PA04GT7017040-1-9-2025');"
                                        class="uk-button uk-button-primary btn-send-contact"
                                        style="min-width: inherit; padding: 0 10px; line-height: 32px;">
                                    <span uk-icon="search"></span>
                                </button>
                                <button class="uk-button uk-button-primary btn-send-contact"
                                        style="min-width: inherit; padding: 0 10px; line-height: 32px;"
                                        onclick="TraCuuHD_onClickTaiHoaDonNPC('XIq6RC1EGQud6DC___PGjwiQ==-PA04GT7017040');">
                                    <span uk-icon="download"></span>
                                </button>

                        </td>

                    </tr>

    </tbody>
</table>
`

	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(sampleHTML)),
			}, nil
		},
	}

	c := &Client{
		client: &http.Client{
			Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
				return mockClient.Do(req)
			}),
		},
	}

	invoices, err := c.SearchInvoices(1, 9, 2025)
	if err != nil {
		t.Fatalf("SearchInvoices failed: %v", err)
	}

	expected := []Invoice{
		{IDHoaDon: "1671732781", SeriKyHieu: "1046124", MaKH: "PA04GT7017040", StrIDHoaDon: "XIq6RC1EGQud6DC___PGjwiQ==", Ky: 1, Thang: 9, Nam: 2025},
	}

	if len(invoices) != len(expected) {
		t.Fatalf("Expected %d invoices, got %d", len(expected), len(invoices))
	}

	for i := range invoices {
		got, want := invoices[i], expected[i]
		if got != want {
			t.Errorf("Invoice %d mismatch:\n got  %+v\n want %+v", i, got, want)
		}
	}
}

func TestSearchInvoices_NoInvoices(t *testing.T) {
	emptyHTML := `
<table class="uk-table uk-table-small uk-table-middle uk-table-divider table2">
    <thead><tr><th>STT</th></tr></thead>
    <tbody></tbody>
</table>
`

	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(emptyHTML)),
			}, nil
		},
	}

	c := &Client{
		client: &http.Client{
			Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
				return mockClient.Do(req)
			}),
		},
	}

	_, err := c.SearchInvoices(1, 9, 2025)
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	expectedError := "Không có tìm thấy hóa đơn nào"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
	}
}

// --- Helper to mock Transport ---
type roundTripperFunc func(req *http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
