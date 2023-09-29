package run

import (
	"testing"
)

func TestEslint(t *testing.T) {
	type args struct {
		command []string
		files   string
		flg     Flg
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eslint(tt.args.command, tt.args.files, tt.args.flg)
		})
	}
}

func TestGetFiles(t *testing.T) {
	type args struct {
		flg Flg
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFiles(tt.args.flg); got != tt.want {
				t.Errorf("GetFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitFetch(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GitFetch()
		})
	}
}
