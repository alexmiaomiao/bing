package offline

import (
	"os"
	"path"
	"strings"
	"io/ioutil"
	"sort"
	"time"
	"github.com/bing/jsonapi"
)

var bingPath string

type byMtime []os.FileInfo
func (x byMtime) Len() int { return len(x) }
func (x byMtime) Less(i, j int) bool { return x[i].ModTime().Before(x[j].ModTime()) }
func (x byMtime) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

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
	q := path.Join(bingPath, strings.Join(terms, "_"))

	f, err :=os.Open(q)	
	if err != nil {
		return nil, err
	}

	defer os.Chtimes(q, time.Now(), time.Now())
	defer f.Close()

	var result jsonapi.SearchResult
	if err := jsonapi.Decode(f, &result); err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func ListWords() ([]os.FileInfo, error) {
	if dent, err := ioutil.ReadDir(bingPath); err != nil {
		return nil, err
	} else {
		sort.Sort(byMtime(dent))
		return dent, nil
	}
}
