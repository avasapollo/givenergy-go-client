package inverter_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/givenergy-go-client/v1/inverter"
)

const (
	baseURL   = "https://api.givenergy.cloud/v1"
	testToken = "toke"
)

// MockRoundTripper is a custom RoundTripper for mocking HTTP responses
type MockRoundTripper struct {
	ExpectedURL string
	// Map URLs to responses
	Response *http.Response
}

// RoundTrip implements the RoundTripper interface
func (mrt *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() != mrt.ExpectedURL {
		return nil, fmt.Errorf("expected URL: %s, got: %s", mrt.ExpectedURL, req.URL.String())
	}

	if mrt.Response != nil {
		return mrt.Response, nil
	}

	// Default response if URL not mocked
	return &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader("Not Found")),
		Header:     make(http.Header),
	}, nil
}

func newMockClient(t *testing.T, path string, statusCode int, testURL string) *http.Client {
	t.Helper()

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	// Create a mock response
	mockResponse := &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewBuffer(b)),
		Header:     make(http.Header),
	}
	mockResponse.Header.Set("Content-Type", "application/json")

	// Initialize the MockRoundTripper with desired responses
	mockTransport := &MockRoundTripper{
		ExpectedURL: testURL,
		Response:    mockResponse,
	}

	// Create an http.Client with the mock transport
	return &http.Client{
		Transport: mockTransport,
	}
}

func TestClient_ListSettings(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ListSettingsArgs{
			InverterSerialNumber: "inverter-1",
		}

		testURL := fmt.Sprintf("%s/inverter/%s/settings", baseURL, args.InverterSerialNumber)
		mockHTTPClient := newMockClient(
			t,
			"testdata/list_settings_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(testToken, inverter.WithHTTPClient(mockHTTPClient))

		data, err := cl.ListSettings(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ListSettingsResponse{
			Data: []*inverter.Settings{
				{
					ID:         266,
					Name:       "DC Discharge 3 End Time",
					Validation: "Value format should be HH:mm. Use correct time range for hour and minutes",
					ValidationRules: []string{
						"date_format:H:i",
					},
				},
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_ReadSettingChargerStart(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerStart,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charger_start_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargerStart(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargerStartResponse{
			Data: struct {
				Value string `json:"value"`
			}{
				Value: "01:00",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargerStart(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargerStartArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerEnd,
			Value:                "16:00",
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/write",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/write_charger_start_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargerStart(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargerStartResponse{
			Data: struct {
				Value   string `json:"value"`
				Success bool   `json:"success"`
				Message string `json:"message"`
			}{
				Value:   "16:00",
				Success: true,
				Message: "Written Successfully",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_ReadSettingChargerEnd(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerEnd,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charger_end_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargerEnd(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargerEndResponse{
			Data: struct {
				Value string `json:"value"`
			}{
				Value: "01:00",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargerEnd(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargerEndArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerEnd,
			Value:                "16:00",
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/write",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/write_charger_end_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargerEnd(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargerEndResponse{
			Data: struct {
				Value   string `json:"value"`
				Success bool   `json:"success"`
				Message string `json:"message"`
			}{
				Value:   "16:00",
				Success: true,
				Message: "Written Successfully",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_ReadSettingChargerEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerEnabled,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charger_enabled_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargerEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargerEnabledResponse{
			Data: struct {
				Value bool `json:"value"`
			}{
				Value: true,
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargerEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargerEnabledArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerEnabled,
			Value:                true,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/write",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/write_charger_enabled_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargerEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargerEnabledResponse{
			Data: struct {
				Value   bool   `json:"value"`
				Success bool   `json:"success"`
				Message string `json:"message"`
			}{
				Value:   true,
				Success: true,
				Message: "Written Successfully",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_ReadSettingChargerLimit(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerLimit,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charger_limit_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargerLimit(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargerLimitResponse{
			Data: struct {
				Value int `json:"value"`
			}{
				Value: 100,
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargerLimit(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargerLimitArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargerLimit,
			Value:                100,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/write",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/write_charger_limit_200.json",
			http.StatusOK,
			testURL,
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargerLimit(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargerLimitResponse{
			Data: struct {
				Value   int    `json:"value"`
				Success bool   `json:"success"`
				Message string `json:"message"`
			}{
				Value:   100,
				Success: true,
				Message: "Written Successfully",
			},
		}
		require.Equal(t, expected, data)
	})
}
