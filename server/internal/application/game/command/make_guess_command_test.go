package command

import (
	"context"
	"reflect"
	"testing"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

func TestMakeGuessCommand_Execute(t *testing.T) {
	type fields struct {
		GuessRepository domain.GuessRepository
	}
	type args struct {
		ctx   context.Context
		input MakeGuessInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MakeGuessResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgc := MakeGuessCommand{
				guessRepository: tt.fields.GuessRepository,
			}
			got, err := mgc.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}
