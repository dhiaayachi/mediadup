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
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{ "Alien.1979.BluRay-1080P.mkv",args{"Alien.1979.BluRay-1080P.mkv"},want ,false},
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
