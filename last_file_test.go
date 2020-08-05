// Package lastfile has 1 exported function

package lastfile

import "testing"

func TestWithPrefix(t *testing.T) {
	type args struct {
		dir    string
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{{
		name:    "1",
		args:    args{dir: ".", prefix: "last"},
		want:    "last_file_test.go",
		wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WithPrefix(tt.args.dir, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithPrefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
