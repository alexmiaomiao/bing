package main

import (
	"log"
	"flag"
	"os"
	"fmt"
	"github.com/bing/jsonapi"
	"github.com/bing/sox"
	"github.com/bing/templ"
	"github.com/bing/online"
	"github.com/bing/offline"
)

var verbose = flag.Bool("v", false, "查看例句")
var ame = flag.Bool("a", false, "美式发音")
var bre = flag.Bool("b", false, "美式发音")
var list = flag.Int("l", 0, "查看历史单词")

var complete chan int = make(chan int)

func bingText(res *jsonapi.SearchResult) {
	s := &struct {
		Res *jsonapi.SearchResult
		Verbose bool
	}{res, *verbose}

	if err := templ.CommandLine.Execute(os.Stdout, s); err != nil {
		log.Fatal(err)
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

func bingBackup(res *jsonapi.SearchResult) {
	if err := offline.Backup(res); err != nil {
		log.Print(err)
	}

	complete <- 0
}

func main() {
	flag.Parse()

	if *list > 0 {
		if dent, err := offline.ListWords(); err != nil {
			log.Fatal(err)
		} else {
			l := len(dent) - *list
			for i, w := range dent {
				if i >= l {
					fmt.Printf("%d\t%s\n", i+1, w.Name())
				}
			}
			os.Exit(0)
		}
	}

	var result *jsonapi.SearchResult
	var err1, err2 error
	
	result, err1 = offline.Search(flag.Args())
	if err1 != nil {
		result, err2 = online.Search(flag.Args())
	}
	
	if err1 != nil && err2 != nil {
		log.Fatal(err1, err2)
	}

	if "" == result.Word {
		log.Fatal("no such word")
	}

	go bingText(result)
	go bingSound(result)
	if err1 != nil {
		go bingBackup(result)
		<- complete
	}
	
	<- complete
	<- complete
}
