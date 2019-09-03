package controllers

import (
	"reflect"
	"testing"
)

func TestGetURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "一般測試",
			args: args{
				url: "https://ipfs.globalupload.io/QmWqePaG57UJhcZyt1cKMPh8ZfqdnoEueaR7JpS7r6vbLy",
			},
			want: []byte("{\n    \"status\": true,\n    \"subject\": \"test\",\n    \"content\": \"this is a REST API test\"\n}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetURL(tt.args.url)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostURL(t *testing.T) {
	type args struct {
		url    string
		params []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "一般測試",
			args: args{
				url:    "https://ipfs.globalupload.io/QmWqePaG57UJhcZyt1cKMPh8ZfqdnoEueaR7JpS7r6vbLy",
				params: []byte("{\"status\": true, \"subject\": \"test\", \"content\": \"this is a REST API test\"}"),
			},
			want: []byte("Method POST not allowed: read only access"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostURL(tt.args.url, tt.args.params)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
