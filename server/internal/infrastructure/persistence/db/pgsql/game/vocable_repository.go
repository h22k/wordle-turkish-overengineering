package game

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/persistence/db/pgsql/game/query"
)

type VocableRepository struct {
	query *query.Queries
}

func NewVocableRepository(q *query.Queries) *VocableRepository {
	return &VocableRepository{
		query: q,
	}
}

func (v VocableRepository) FindRandom(ctx context.Context) (domain.Word, error) {
	stringWord, err := v.query.GetRandomSecretWord(ctx)

	if err != nil {
		return "", err
	}

	return domain.Word(stringWord), nil
}

func (v VocableRepository) Update(ctx context.Context, vocable domain.Vocable) error {
	//TODO implement me
	panic("implement me")
}

func (v VocableRepository) Save(ctx context.Context, vocable domain.Vocable) error {
	_, err := v.query.AddWordToPool(ctx, query.AddWordToPoolParams{
		Word:     vocable.Word.String(),
		IsValid:  true,
		IsAnswer: false,
	})

	return err
}

func (v VocableRepository) FindByWord(ctx context.Context, word domain.Word) (domain.Vocable, error) {
	//TODO implement me
	panic("implement me")
}
