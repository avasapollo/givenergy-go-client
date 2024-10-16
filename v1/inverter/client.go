package inverter

import (
	"bytes"
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
	fmtSettingRead  = "%s/inverter/%s/settings/%s/read"
	fmtSettingWrite = "%s/inverter/%s/settings/%s/write"
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

type ReadSettingChargerStartResponse struct {
	Data struct {
		Value string `json:"value"`
	} `json:"data"`
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

type WriteSettingChargerStartArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                string  `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargerStartResponse struct {
	Data struct {
		Value   string `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargerStart(
	ctx context.Context,
	args *WriteSettingChargerStartArgs,
) (*WriteSettingChargerStartResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargerStartResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingChargerEndResponse struct {
	Data struct {
		Value string `json:"value"`
	} `json:"data"`
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

type WriteSettingChargerEndArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                string  `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargerEndResponse struct {
	Data struct {
		Value   string `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargerEnd(
	ctx context.Context,
	args *WriteSettingChargerEndArgs,
) (*WriteSettingChargerEndResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargerEndResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingChargerEnabledResponse struct {
	Data struct {
		Value bool `json:"value"`
	} `json:"data"`
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

type WriteSettingChargerEnabledArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                bool    `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargerEnabledResponse struct {
	Data struct {
		Value   bool   `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargerEnabled(
	ctx context.Context,
	args *WriteSettingChargerEnabledArgs,
) (*WriteSettingChargerEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargerEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingChargerLimitResponse struct {
	Data struct {
		Value int `json:"value"`
	} `json:"data"`
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

type WriteSettingChargerLimitArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                int     `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargerLimitResponse struct {
	Data struct {
		Value   int    `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargerLimit(
	ctx context.Context,
	args *WriteSettingChargerLimitArgs,
) (*WriteSettingChargerLimitResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargerLimitResponse)
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
