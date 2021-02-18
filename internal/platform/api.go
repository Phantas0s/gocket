package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const URL = "https://getpocket.com"

// Client represents a Pocket client that grants OAuth access to your application
type Client struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

// NewClient creates a new Pocket client.
func NewClient(consumerKey, accessToken string) *Client {
	return &Client{
		ConsumerKey: consumerKey,
		AccessToken: accessToken,
	}
}

func Post(action string, data, res interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", URL+action, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(res)
}
