package main

import (
	"reflect"
	"testing"
)

func Test_extractYear(t *testing.T) {
	type args struct {
		name string
	}
	var want []string
	want = append(want,"1979")
	var want2 []string
	want2 = append(want2,"1981")
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{ "Alien.1979.BluRay-1080P.mkv",args{"Alien.1979.BluRay-1080P.mkv"},want ,false},
		{ "For.Your.Eyes.Only.1981.BluRay-1080P.mkv",args{"For.Your.Eyes.Only.1981.BluRay-1080P.mkv"},want2 ,false},

	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractYear(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractYear() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractName(t *testing.T) {
	type args struct {
		year  string
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Amélie.2001.BluRay-1080P.mkv",args{"2001","Amélie.2001.BluRay-1080P.mkv"},"Amélie",false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractName(tt.args.year, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
