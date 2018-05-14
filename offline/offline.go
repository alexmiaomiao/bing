package offline

import (
	"os"
	"path"
	"fmt"
)

func wordToPath(w string) string {
	p := "offline"
	for _, r := range w {
		if ' ' == r {
			p = path.Join(p, "sp")
		} else {
			p = path.Join(p, r)
		}
	}
	return p
}

func MakeWordDir(w string) {
	if err := os.MkdirAll(wordToPath(w), 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func IsOffline(w string) (bool, error) {
	if _, err := os.Stat(wordToPath(w)); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			fmt.Fprintln(os.Stderr, err)
			return false, err
		}
	} else {
		return true, nil
	}
}
