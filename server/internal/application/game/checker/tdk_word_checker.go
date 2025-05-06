package checker

import (
	"context"
	"fmt"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

const tdkUrlFormat = "https://sozluk.gov.tr/gts?ara=%s"

type TdkResponse interface {
	IsWordAcceptable() bool
}

type httpClient interface {
	Get(url string) (TdkResponse, error)
}

type TdkWordChecker struct {
	client httpClient
}

func NewTdkWordChecker(client httpClient) *TdkWordChecker {
	return &TdkWordChecker{
		client: client,
	}
}

func (twc *TdkWordChecker) Check(ctx context.Context, word domain.Word) (bool, error) {
	resp, err := twc.client.Get(fmt.Sprintf(tdkUrlFormat, word))

	if err != nil {
		return false, err
	}

	return resp.IsWordAcceptable(), nil
}
