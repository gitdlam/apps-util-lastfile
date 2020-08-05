// Package lastfile has 1 exported function
package lastfile

import (
	"io/ioutil"
	"strings"
)


// WithPrefix returns last file (by name) with the given prefix
func WithPrefix(dir string, prefix string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	lastFile := ""
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if !strings.HasPrefix(strings.ToLower(f.Name()), strings.ToLower(prefix)) {
			continue
		}

		if lastFile < f.Name() {
			lastFile = f.Name()
		}

	}
	return dir + "\\" + lastFile, nil

}
