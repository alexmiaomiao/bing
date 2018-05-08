package main

import (
	"log"
	"flag"
	"os"
	"github.com/bing/jsonapi"
	"github.com/bing/sox"
	"github.com/bing/templ"
)

var verbose = flag.Bool("v", false, "查看例句")
var ame = flag.Bool("a", false, "美式发音")
var bre = flag.Bool("b", false, "美式发音")

var complete chan int = make(chan int)

func bingText(res *jsonapi.SearchResult) {
	s := &struct {
		Res *jsonapi.SearchResult
		Verbose bool
	}{res, *verbose}
	
	if err := templ.CommandLine.Execute(os.Stdout, s); err != nil {
		log.Println(err)
	}
	
	complete <- 0
}

func bingSound(res *jsonapi.SearchResult) {
	if res.Pronunciation != nil {
		if *ame {
			sox.Play(res.Pronunciation.AmEMP3)
		}
		if *bre {
			sox.Play(res.Pronunciation.BrEMP3)
		}
	}
	
	complete <- 0
}

func main() {
	flag.Parse()
	result, err := jsonapi.Search(flag.Args())
	
	if err != nil {
		log.Fatal(err)
	}

	if "" == result.Word {
		log.Fatal("no such word")
	}
	
	go bingText(result)
	go bingSound(result)
	
	<- complete
	<- complete
}
