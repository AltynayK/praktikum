package service

import (
	"testing"
)

func TestWriteURLByID(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple test #1",
			args: args{"https://practicum.yandex.ru/learn/go-developer/"},
			want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WriteURLByID(tt.args.url); got != tt.want {
				t.Errorf("WriteURLByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetURLFromID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "simple test #1",
			args: args{1},
			want: "https://practicum.yandex.ru/learn/go-developer/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetURLFromID(tt.args.id); got != tt.want {
				t.Errorf("GetURLFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}
