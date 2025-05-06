package checker

import (
	"context"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/command"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type DatabaseWordWriterChecker struct {
	dbChecker      *DatabaseWordChecker
	tdkWordChecker *TdkWordChecker
	vc             *command.WordCommand
}

func NewDatabaseWordWriterChecker(dbChecker *DatabaseWordChecker, tdkWordChecker *TdkWordChecker, vc *command.WordCommand) *DatabaseWordWriterChecker {
	return &DatabaseWordWriterChecker{
		dbChecker:      dbChecker,
		tdkWordChecker: tdkWordChecker,
		vc:             vc,
	}
}

func (d DatabaseWordWriterChecker) Check(ctx context.Context, word domain.Word) (bool, error) {
	existsInDb, err := d.dbChecker.Check(ctx, word)

	if err != nil {
		return false, err
	}

	if existsInDb {
		return true, nil
	}

	existsInTdk, err := d.tdkWordChecker.Check(ctx, word)

	if err != nil || !existsInTdk {
		return false, err
	}

	_ = d.vc.AddWord(ctx, word) // TODO:: handle error

	return true, nil
}
