package inverter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type ReadChargerStartValue struct {
	Value string `json:"value"`
}

type ReadChargerStartResponse struct {
	Data *ReadChargerEndValue `json:"data"`
}

func (c *Client) ReadChargerStart(ctx context.Context, args *ReadSettingArgs) (*ReadChargerStartResponse, error) {
	u := fmt.Sprintf("%s/inverter/%s/settings/%s", c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadChargerStartResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadChargerEndValue struct {
	Value string `json:"value"`
}

type ReadChargerEndResponse struct {
	Data *ReadChargerEndValue `json:"data"`
}

func (c *Client) ReadChargerEnd(ctx context.Context, args *ReadSettingArgs) (*ReadChargerEndResponse, error) {
	u := fmt.Sprintf("%s/inverter/%s/settings/%s", c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadChargerEndResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadChargerEnabledValue struct {
	Value bool `json:"value"`
}

type ReadChargerEnabledResponse struct {
	Data *ReadChargerEnabledValue `json:"data"`
}

func (c *Client) ReadChargerEnabled(ctx context.Context, args *ReadSettingArgs) (*ReadChargerEnabledResponse, error) {
	u := fmt.Sprintf("%s/inverter/%s/settings/%s", c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadChargerEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadChargerLimitValue struct {
	Value bool `json:"value"`
}

type ReadChargerLimitResponse struct {
	Data *ReadChargerLimitValue `json:"data"`
}

func (c *Client) ReadChargerLimit(ctx context.Context, args *ReadSettingArgs) (*ReadChargerLimitResponse, error) {
	u := fmt.Sprintf("%s/inverter/%s/settings/%s", c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadChargerLimitResponse)
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
