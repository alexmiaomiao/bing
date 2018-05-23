package online

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/bing/jsonapi"
)

const bingURL = "http://xtk.azurewebsites.net/BingDictService.aspx"

func Search(terms []string) (*jsonapi.SearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(bingURL + "?Word=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result jsonapi.SearchResult
	if err := jsonapi.Decode(resp.Body, &result); err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}
