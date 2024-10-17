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
	DefaultSettingChargeStart      = "64"
	DefaultSettingChargeEnd        = "65"
	DefaultSettingChargeEnabled    = "66"
	DefaultSettingChargeLimit      = "77"
	DefaultSettingDischargeEnabled = "56"
	DefaultSettingDischargeStart   = "53"
	DefaultSettingDischargeEnd     = "54"
	DefaultSettingEcoModeEnabled   = "24"
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

type ReadSettingChargeStartResponse struct {
	Data struct {
		Value string `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingChargeStart(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargeStartResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargeStartResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingChargeStartArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                string  `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargeStartResponse struct {
	Data struct {
		Value   string `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargeStart(
	ctx context.Context,
	args *WriteSettingChargeStartArgs,
) (*WriteSettingChargeStartResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargeStartResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingChargeEndResponse struct {
	Data struct {
		Value string `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingChargeEnd(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargeEndResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargeEndResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingChargeEndArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                string  `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargeEndResponse struct {
	Data struct {
		Value   string `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargeEnd(
	ctx context.Context,
	args *WriteSettingChargeEndArgs,
) (*WriteSettingChargeEndResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargeEndResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingChargeEnabledResponse struct {
	Data struct {
		Value bool `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingChargeEnabled(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargeEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargeEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingChargeEnabledArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                bool    `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargeEnabledResponse struct {
	Data struct {
		Value   bool   `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargeEnabled(
	ctx context.Context,
	args *WriteSettingChargeEnabledArgs,
) (*WriteSettingChargeEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargeEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingChargeLimitResponse struct {
	Data struct {
		Value int `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingChargeLimit(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingChargeLimitResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingChargeLimitResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingChargeLimitArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                int     `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingChargeLimitResponse struct {
	Data struct {
		Value   int    `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingChargeLimit(
	ctx context.Context,
	args *WriteSettingChargeLimitArgs,
) (*WriteSettingChargeLimitResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingChargeLimitResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingDischargeEnabledResponse struct {
	Data struct {
		Value bool `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingDischargeEnabled(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingDischargeEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingDischargeEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingDischargeEnabledArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                bool    `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingDischargeEnabledResponse struct {
	Data struct {
		Value   bool   `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingDischargeEnabled(
	ctx context.Context,
	args *WriteSettingDischargeEnabledArgs,
) (*WriteSettingDischargeEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingDischargeEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingDischargeStartResponse struct {
	Data struct {
		Value string `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingDischargeStart(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingDischargeStartResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingDischargeStartResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingDischargeStartArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                string  `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingDischargeStartResponse struct {
	Data struct {
		Value   string `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingDischargeStart(
	ctx context.Context,
	args *WriteSettingChargeStartArgs,
) (*WriteSettingDischargeStartResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingDischargeStartResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingDischargeEndResponse struct {
	Data struct {
		Value string `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingDischargeEnd(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingDischargeEndResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingDischargeEndResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingDischargeEndArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                string  `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingDischargeEndResponse struct {
	Data struct {
		Value   string `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingDischargeEnd(
	ctx context.Context,
	args *WriteSettingDischargeEndArgs,
) (*WriteSettingDischargeEndResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingDischargeEndResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ReadSettingEcoModeEnabledResponse struct {
	Data struct {
		Value bool `json:"value"`
	} `json:"data"`
}

func (c *Client) ReadSettingEcoModeEnabled(
	ctx context.Context,
	args *ReadSettingArgs,
) (*ReadSettingEcoModeEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingRead, c.baseURL, args.InverterSerialNumber, args.SettingID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	res := new(ReadSettingEcoModeEnabledResponse)
	if err := c.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

type WriteSettingEcoModeEnabledArgs struct {
	InverterSerialNumber string  `json:"-"`
	SettingID            string  `json:"-"`
	Value                bool    `json:"value"`
	Context              *string `json:"context,omitempty"`
}

type WriteSettingEcoModeEnabledResponse struct {
	Data struct {
		Value   bool   `json:"value"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) WriteSettingEcoModeEnabled(
	ctx context.Context,
	args *WriteSettingEcoModeEnabledArgs,
) (*WriteSettingEcoModeEnabledResponse, error) {
	u := fmt.Sprintf(fmtSettingWrite, c.baseURL, args.InverterSerialNumber, args.SettingID)

	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := new(WriteSettingEcoModeEnabledResponse)
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
