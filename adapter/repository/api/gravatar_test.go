package api

import (
	"reflect"
	"testing"

	"github.com/iwanjunaid/basesvc/domain/model"
)

func TestGravatar_GetProfile(t *testing.T) {
	type fields struct {
		email        string
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes *model.GravatarProfiles
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test user", fields{
			email:        "ilmi.mris@gmail.com",
			hash:         "cd601941419730dbc79bbc41180ab703",
			defaultURL:   "",
			defaultValue: "",
			size:         0,
			forceDefault: false,
			rating:       "",
		}, &model.GravatarProfiles{
			Entry: []model.Profile{
				{
					"103714164",
					"cd601941419730dbc79bbc41180ab703",
					"cd601941419730dbc79bbc41180ab703",
					"http://gravatar.com/mrisilmi",
					"https://secure.gravatar.com/avatar/cd601941419730dbc79bbc41180ab703",
					[]model.Photo{{"https://secure.gravatar.com/avatar/cd601941419730dbc79bbc41180ab703", "thumbnail"}},
					[]string{},
					"mrisilmi",
					[]string{},
				},
			},
		},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &AuthorGravatarRepositoryImpl{
				email:        tt.fields.email,
				hash:         tt.fields.hash,
				defaultURL:   tt.fields.defaultURL,
				defaultValue: tt.fields.defaultValue,
				size:         tt.fields.size,
				forceDefault: tt.fields.forceDefault,
				rating:       tt.fields.rating,
			}
			gotRes, err := g.GetProfile()
			if (err != nil) != tt.wantErr {
				// t.Errorf("result is %v", gotRes)
				t.Errorf("Gravatar.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Gravatar.GetProfile() = %v, want %v", gotRes, tt.wantRes)
			}

		})
	}
}
