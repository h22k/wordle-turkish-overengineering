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

	var result interface{} // tdk's response is not consistent, so we use an empty interface. WTF bro
	if err := json.Unmarshal(body, &result); err != nil {
		return TdkResponse{}, TdkClientErr
	}

	switch result.(type) {
	case []interface{}: // if it's an array, we can assume that the word exists in tdk
		return TdkResponse{isExists: true}, nil
	case map[string]string: // if it's a map, we can assume that the word doesn't exist in tdk, due to tdk's response
		_, isErrorOccurred := result.(map[string]string)["error"]
		return TdkResponse{isExists: !isErrorOccurred}, nil
	default:
		return TdkResponse{isExists: false}, nil
	}
}
