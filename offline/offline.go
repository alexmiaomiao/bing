package offline

import (
	"os"
	"path"
	"strings"
	"github.com/bing/jsonapi"
)

var bingPath string

func init() {
	if gp := os.Getenv("GOPATH"); gp != "" {
		bingPath = path.Join(os.Getenv("GOPATH"), "bin", "offline")
	} else {
		bingPath = path.Join(os.Getenv("HOME"), "offline")
	}

	if b, _ := pathExists(bingPath); !b {
		os.MkdirAll(bingPath, 0755)
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Backup(s *jsonapi.SearchResult) error {
	q := strings.Replace(s.Word, " ", "_", -1)
	
	f, err :=os.Create(path.Join(bingPath, q))	
	if err != nil {
		return err
	}
	defer f.Close()

	if err := jsonapi.Encode(f, s); err != nil {
		return err
	} else {
		return nil
	}
}

func Search(terms []string) (*jsonapi.SearchResult, error) {
	q := strings.Join(terms, "_")

	f, err :=os.Open(path.Join(bingPath, q))	
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var result jsonapi.SearchResult
	if err := jsonapi.Decode(f, &result); err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}
