package jsonapi

import (
	"encoding/json"
	"fmt"
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
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(bingURL + "?Word=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result SearchResult
	if offline.IsOffline
	
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	
	return &result, nil
}
