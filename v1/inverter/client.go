package inverter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	DefaultSettingChargerStart   = "64"
	DefaultSettingChargerEnd     = "65"
	DefaultSettingChargerEnabled = "66"
	DefaultSettingChargerLimit   = "77"
)

const (
	fmtSettingRead = "%s/inverter/%s/settings/%s/read"
)

type Client struct {
	baseURL string
	token   string
	httpCl  *http.Client
}

func NewClient(token string, opts ...Option) *Client {
	conf := defaultOptions()
	for _, opt := range opts {
		opt(conf)
	}

	return &Client{
		token:   token,
		baseURL: conf.baseURL,
		httpCl:  conf.httpClient,
	}
}

type ListSettingsArgs struct {
	InverterSerialNumber string
}

type Settings struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Validation      string   `json:"validation"`
	ValidationRules []string `json:"validation_rules"`
}

type ListSettingsResponse struct {
	Data []*Settings `json:"data"`
}

func (c *Client) ListSettings(ctx context.Context, args *ListSettingsArgs) (*ListSettingsResponse, error) {
	u := fmt.Sprintf("%s/inverter/%s/settings", c.baseURL, args.InverterSerialNumber)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ListSettingsResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingArgs struct {
	InverterSerialNumber string
	SettingID            string
}

func NewReadSettingArgs(serialNumber string, settingID string) *ReadSettingArgs {
	return &ReadSettingArgs{
		InverterSerialNumber: serialNumber,
		SettingID:            settingID,
	}
}

type ChargerStartValue struct {
	Value string `json:"value"`
}

type ReadSettingChargerStartResponse struct {
	Data *ChargerStartValue `json:"data"`
}

func (c *Client) ReadSettingChargerStart(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargerStartResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargerStartResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ChargerEndValue struct {
	Value string `json:"value"`
}

type ReadSettingChargerEndResponse struct {
	Data *ChargerEndValue `json:"data"`
}

func (c *Client) ReadSettingChargerEnd(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargerEndResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargerEndResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ChargerEnabledValue struct {
	Value bool `json:"value"`
}

type ReadSettingChargerEnabledResponse struct {
	Data *ChargerEnabledValue `json:"data"`
}

func (c *Client) ReadSettingChargerEnabled(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargerEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargerEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ChargerLimitValue struct {
	Value int `json:"value"`
}

type ReadSettingChargerLimitResponse struct {
	Data *ChargerLimitValue `json:"data"`
}

func (c *Client) ReadSettingChargerLimit(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargerLimitResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargerLimitResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) do(req *http.Request, res any) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.httpCl.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= 300 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d body: %s", resp.StatusCode, msg)
	}

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return err
	}
	return nil
}
