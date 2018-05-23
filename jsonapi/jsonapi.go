package jsonapi

import (
	"encoding/json"
	"io"
)

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

func Decode(r io.Reader, s *SearchResult) error {
	if err := json.NewDecoder(r).Decode(s); err != nil {
		return err
	} else {
		return nil		
	}
}

func Encode(w io.Writer, s *SearchResult) error {
	if err := json.NewEncoder(w).Encode(s); err != nil {
		return err
	} else {
		return nil
	}
}
