package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/checker"
)

var TdkClientErr = errors.New("tdk client error")

type TdkResponse struct {
	isExists bool
}

func (t TdkResponse) IsWordAcceptable() bool {
	return t.isExists
}

type TdkClient struct {
	client *http.Client
}

func NewTdkClient(client *http.Client) *TdkClient {
	return &TdkClient{
		client: client,
	}
}

func (t TdkClient) Get(url string) (checker.TdkResponse, error) {
	resp, err := t.client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return TdkResponse{}, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return TdkResponse{}, TdkClientErr
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return TdkResponse{}, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return TdkResponse{}, err
	}

	_, isErrorOccurred := result["error"]

	return TdkResponse{isExists: !isErrorOccurred}, nil
}
