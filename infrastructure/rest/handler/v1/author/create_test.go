package author

import (
	"testing"

	"github.com/iwanjunaid/basesvc/domain/model"
	uuid "github.com/satori/go.uuid"
)

func Test_isRequestValid(t *testing.T) {
	id, _ := uuid.FromString("dc9fb74b-6ffb-4d49-8879-da9627afdd14")

	type args struct {
		a model.Author
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"gmail", args{
			model.Author{
				ID:        id,
				Name:      "admin",
				Email:     "admin@gmail.com",
				CreatedAt: 1604980687,
				UpdatedAt: 1604980687,
			},
		},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isRequestValid(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("isRequestValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isRequestValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
