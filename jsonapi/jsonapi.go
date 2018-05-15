package jsonapi

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/bing/offline"
)

const bingURL = "http://xtk.azurewebsites.net/BingDictService.aspx"

type Pron struct {
	AmE    string
	AmEMP3 string
	BrE    string
	BrEMP3 string
}

type Definition struct {
	Pos string
	Def string
}

type Sample struct {
	Eng string
	Chn string
}

type SearchResult struct {
	Word          string
	Pronunciation *Pron
	Defs          []*Definition
	Sams          []*Sample
}

func Search(terms []string) (*SearchResult, error) {
	var result SearchResult

	if _, err := offline.Open(strings.Join(terms, " "), "json"); err != nil {
	
		q := url.QueryEscape(strings.Join(terms, " "))
		resp, err := http.Get(bingURL + "?Word=" + q)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("search query failed: %s", resp.Status)
		}

		if f, err := offline.Create(strings.Join(terms, " "), "json"); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			if d, err := ioutil.ReadAll(resp.Body); err == nil {
				f.Write(d)
				f.Close()
			}
		}
	}

	if f, err := offline.Open(strings.Join(terms, " "), "json"); err == nil {
		if err := json.NewDecoder(f).Decode(&result); err != nil {
			return nil, err
		}	
	}
	
	return &result, nil
}
