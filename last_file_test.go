// Package lastfile has 1 exported function

package lastfile

import (
	"testing"

	appname "github.com/gitdlam/apps-appname"
)

func TestWithPrefix(t *testing.T) {
	type args struct {
		dir    string
		prefix string
	}
	folder, _ := appname.Folder()
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{{
		name:    "1",
		args:    args{dir: folder, prefix: "x"},
		want:    "x",
		wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := WithPrefix(tt.args.dir, tt.args.prefix)

			t.Log("**** Please manual check:", got, folder)

		})
	}
}
