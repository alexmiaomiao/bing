package offline

import (
	"os"
	"path"
)

var binPath string

func init() {
	if gp := os.Getenv("GOPATH"); gp != "" {
		binPath = path.Join(os.Getenv("GOPATH"), "bin")
	} else {
		binPath = os.Getenv("HOME")
	}
	
}

func wordToPath(w string) string {
	p := "offline"
	for _, r := range w {
		if ' ' == r {
			p = path.Join(p, "sp")
		} else {
			p = path.Join(p, string(r))
		}
	}
	return p
}

func Open(w string, r string) (*os.File, error) {
	dir, _ := os.Getwd()
	os.Chdir(binPath)
	defer os.Chdir(dir)
	return os.Open(path.Join(wordToPath(w), r))
}

func Create(w string, r string) (*os.File, error) {
	dir, _ := os.Getwd()
	os.Chdir(binPath)
	defer os.Chdir(dir)
	if err := os.MkdirAll(wordToPath(w), 0755); err != nil {
		return nil, err
	}
	return os.Create(path.Join(wordToPath(w), r))
}
