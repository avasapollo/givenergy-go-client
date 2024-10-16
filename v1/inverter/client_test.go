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

	"github.com/givenergy-client/v1/inverter"
)

const (
	baseURL   = "https://api.givenergy.cloud/v1"
	testToken = "toke"
)

// MockRoundTripper is a custom RoundTripper for mocking HTTP responses
type MockRoundTripper struct {
	// Map URLs to responses
	Responses map[string]*http.Response
}

// RoundTrip implements the RoundTripper interface
func (mrt *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if resp, exists := mrt.Responses[req.URL.String()]; exists {
		return resp, nil
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
		Responses: map[string]*http.Response{
			testURL: mockResponse,
		},
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
			Data: &inverter.ChargerStartValue{
				Value: "01:00",
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
			Data: &inverter.ChargerEndValue{
				Value: "01:00",
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
			Data: &inverter.ChargerEnabledValue{
				Value: true,
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
			Data: &inverter.ChargerLimitValue{
				Value: 100,
			},
		}
		require.Equal(t, expected, data)
	})
}
