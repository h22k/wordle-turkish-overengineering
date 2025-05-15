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

func (v VocableRepository) IsWordExists(ctx context.Context, word domain.Word) (bool, error) {
	return v.query.IsWordExists(ctx, word.String())
}

func (v VocableRepository) FindRandom(ctx context.Context) (domain.Word, int32, error) {
	wordPool, err := v.query.GetRandomWord(ctx)

	if err != nil {
		return "", 0, err
	}

	return domain.Word(wordPool.Word), wordPool.ID, nil
}

func (v VocableRepository) Update(ctx context.Context, vocable domain.Vocable) error {
	//TODO implement me
	panic("implement me")
}

func (v VocableRepository) Save(ctx context.Context, vocable domain.Vocable) error {
	_, err := v.query.AddWordToPool(ctx, vocable.Word.String())

	return err
}

func (v VocableRepository) FindByWord(ctx context.Context, word domain.Word) (domain.Vocable, error) {
	wp, err := v.query.FindWord(ctx, word.String())

	if err != nil {
		return domain.Vocable{}, err
	}

	return domain.Vocable{
		Word: domain.Word(wp.Word),
	}, nil
}
