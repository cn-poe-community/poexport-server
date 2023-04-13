package translate

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Server string
}

func (c *Client) CreateBuilding(buildingJson string) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/pob/create", c.Server), bytes.NewBuffer([]byte(buildingJson)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		return body, nil
	}
	return nil, fmt.Errorf("bad gateway")
}
