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
	"time"

	"github.com/stretchr/testify/require"

	"github.com/givenergy-go-client/v1/inverter"
)

const (
	baseURL   = "https://api.givenergy.cloud/v1"
	testToken = "toke"
)

// MockRoundTripper is a custom RoundTripper for mocking HTTP responses
type MockRoundTripper struct {
	t                *testing.T
	ExpectedBodyJSON string
	ExpectedURL      string
	// Map URLs to responses
	Response *http.Response
}

// RoundTrip implements the RoundTripper interface
func (mrt *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() != mrt.ExpectedURL {
		return nil, fmt.Errorf("expected URL: %s, got: %s", mrt.ExpectedURL, req.URL.String())
	}

	if req.Header.Get("Authorization") != "Bearer "+testToken {
		return nil, fmt.Errorf("expected Authorization: %s, got: %s", testToken, req.Header.Get("Authorization"))
	}

	if req.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("expected Content-Type: application/json, got: %s", req.Header.Get("Content-Type"))
	}

	if mrt.ExpectedBodyJSON != "" {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		require.JSONEq(mrt.t, mrt.ExpectedBodyJSON, string(b))
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

func newMockClient(
	t *testing.T,
	path string,
	statusCode int,
	testURL string,
	expectedBodyJson string,
) *http.Client {
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
		t:                t,
		ExpectedBodyJSON: expectedBodyJson,
		ExpectedURL:      testURL,
		Response:         mockResponse,
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
			"",
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

func TestClient_ReadSettingChargeStart(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeStart,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charge_start_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargeStart(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargeStartResponse{
			Data: struct {
				Value string `json:"value"`
			}{
				Value: "01:00",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargeStart(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargeStartArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeEnd,
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
			"testdata/write_charge_start_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\":\"16:00\" }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargeStart(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargeStartResponse{
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

func TestClient_ReadSettingChargeEnd(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeEnd,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charge_end_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargeEnd(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargeEndResponse{
			Data: struct {
				Value string `json:"value"`
			}{
				Value: "01:00",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargeEnd(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargeEndArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeEnd,
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
			"testdata/write_charge_end_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\":\"16:00\" }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargeEnd(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargeEndResponse{
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

func TestClient_ReadSettingChargeEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeEnabled,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charge_enabled_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargeEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargeEnabledResponse{
			Data: struct {
				Value bool `json:"value"`
			}{
				Value: true,
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargeEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargeEnabledArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeEnabled,
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
			"testdata/write_charge_enabled_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\": true }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargeEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargeEnabledResponse{
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

func TestClient_ReadSettingChargeLimit(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeLimit,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_charge_limit_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingChargeLimit(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingChargeLimitResponse{
			Data: struct {
				Value int `json:"value"`
			}{
				Value: 100,
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingChargeLimit(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingChargeLimitArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeLimit,
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
			"testdata/write_charge_limit_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\": 100 }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingChargeLimit(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingChargeLimitResponse{
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

func TestClient_ReadSettingDischargeStart(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingDischargeStart,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_discharge_start_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingDischargeStart(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingDischargeStartResponse{
			Data: struct {
				Value string `json:"value"`
			}{
				Value: "01:00",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingDischargeStart(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingDischargeStartArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingDischargeStart,
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
			"testdata/write_discharge_start_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\":\"16:00\" }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingDischargeStart(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingDischargeStartResponse{
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

func TestClient_ReadSettingDischargeEnd(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingChargeEnd,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_discharge_end_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingDischargeEnd(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingDischargeEndResponse{
			Data: struct {
				Value string `json:"value"`
			}{
				Value: "01:00",
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingDischargeEnd(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingDischargeEndArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingDischargeEnd,
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
			"testdata/write_discharge_end_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\":\"16:00\" }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingDischargeEnd(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingDischargeEndResponse{
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

func TestClient_ReadSettingDischargeEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingDischargeEnabled,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_discharge_enabled_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingDischargeEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingDischargeEnabledResponse{
			Data: struct {
				Value bool `json:"value"`
			}{
				Value: true,
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingDischargeEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingDischargeEnabledArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingDischargeEnabled,
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
			"testdata/write_discharge_enabled_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\": true }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingDischargeEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingDischargeEnabledResponse{
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

func TestClient_ReadSettingEcoModeEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.ReadSettingArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingEcoModeEnabled,
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/settings/%s/read",
			baseURL,
			args.InverterSerialNumber,
			args.SettingID,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/read_eco_mode_enabled_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.ReadSettingEcoModeEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.ReadSettingEcoModeEnabledResponse{
			Data: struct {
				Value bool `json:"value"`
			}{
				Value: true,
			},
		}
		require.Equal(t, expected, data)
	})
}

func TestClient_WriteSettingEcoModeEnabled(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.WriteSettingEcoModeEnabledArgs{
			InverterSerialNumber: "inverter-1",
			SettingID:            inverter.DefaultSettingEcoModeEnabled,
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
			"testdata/write_eco_mode_enabled_200.json",
			http.StatusOK,
			testURL,
			"{ \"value\": true }",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.WriteSettingEcoModeEnabled(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.WriteSettingEcoModeEnabledResponse{
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

func TestClient_SystemDataLatest(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		args := &inverter.SystemDataLatestArgs{
			InverterSerialNumber: "inverter-1",
		}

		testURL := fmt.Sprintf(
			"%s/inverter/%s/system-data/latest",
			baseURL,
			args.InverterSerialNumber,
		)

		mockHTTPClient := newMockClient(
			t,
			"testdata/system_data_latest_200.json",
			http.StatusOK,
			testURL,
			"",
		)

		cl := inverter.NewClient(
			testToken,
			inverter.WithHTTPClient(mockHTTPClient),
		)

		data, err := cl.SystemDataLatest(context.Background(), args)
		require.NoError(t, err)
		expected := &inverter.SystemDataLatestResponse{
			Data: &inverter.SystemData{
				Time:   time.Date(2024, 10, 17, 15, 22, 3, 0, time.UTC),
				Status: "Normal",
				Solar: &inverter.SystemDataSolar{
					Power: 2684,
					Arrays: []*inverter.DataSolar{
						{
							Array:   1,
							Voltage: 244.3,
							Current: 10.9,
							Power:   2684,
						},
						{
							Array:   2,
							Voltage: 0,
							Current: 0,
							Power:   0,
						},
					},
				},
				Grid: &inverter.SystemDataGrid{
					Voltage:   245.5,
					Current:   5.2,
					Power:     1235,
					Frequency: 50,
				},
				Battery: &inverter.SystemDataBattery{
					Percent:     96,
					Power:       -1251,
					Temperature: 21,
				},
				Inverter: &inverter.SystemDataInverter{
					Temperature:     31.2,
					Power:           -1265,
					OutputVoltage:   244.3,
					OutputFrequency: 50.04,
					EpsPower:        0,
				},
				Consumption: 183,
			},
		}
		require.Equal(t, expected, data)
	})
}
