package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/checker"
)

var TdkClientErr = errors.New("tdk client error")

type httpGetFunc func(url string) (*http.Response, error)

type TdkResponse struct {
	isExists bool
}

func (t TdkResponse) IsWordAcceptable() bool {
	return t.isExists
}

type TdkClient struct {
	client httpGetFunc
}

func NewTdkClient(client httpGetFunc) *TdkClient {
	return &TdkClient{
		client: client,
	}
}

func (t TdkClient) Get(url string) (checker.TdkResponse, error) {
	resp, err := t.client(url)
	if err != nil {
		return TdkResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return TdkResponse{}, TdkClientErr
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TdkResponse{}, err
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)

	var ute *json.UnmarshalTypeError
	if err != nil && !errors.As(err, &ute) {
		return TdkResponse{}, err
	}

	_, isErrorOccurred := responseMap["error"]

	return TdkResponse{isExists: !isErrorOccurred}, nil
}
