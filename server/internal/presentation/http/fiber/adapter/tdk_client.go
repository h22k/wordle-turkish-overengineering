package adapter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/checker"
)

var TdkClientErr = errors.New("tdk client error")

type fiberGetFunc func(url string) *fiber.Agent

type TdkResponse struct {
	isExists bool
}

func (t TdkResponse) IsWordAcceptable() bool {
	return t.isExists
}

type TdkClient struct {
	client fiberGetFunc
}

func NewTdkClient(client fiberGetFunc) *TdkClient {
	return &TdkClient{
		client: client,
	}
}

func (t TdkClient) Get(url string) (checker.TdkResponse, error) {
	agent := t.client(url)
	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return TdkResponse{}, errs[0]
	}

	if statusCode >= http.StatusBadRequest {
		return TdkResponse{}, TdkClientErr
	}

	var fiberMap fiber.Map
	err := json.Unmarshal(body, &fiberMap)

	var ute *json.UnmarshalTypeError
	if err != nil && !errors.As(err, &ute) {
		return TdkResponse{}, err
	}

	_, isErrorOccurred := fiberMap["error"]

	return TdkResponse{isExists: !isErrorOccurred}, nil
}
