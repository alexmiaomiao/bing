package main

import (
	"fmt"
	"strings"
	"log"
	"flag"
	"github.com/bing/jsonapi"
	"github.com/bing/sox"
)

var verbose = flag.Bool("v", false, "查看例句")
var ame = flag.Bool("a", false, "美式发音")
var bre = flag.Bool("b", false, "美式发音")

var complete chan int = make(chan int)

func printres(res *jsonapi.SearchResult) {
	if res.Word != "" {
		fmt.Printf("%s\t", res.Word)
	} else {
		log.Fatal("no such word")
	}

	if res.Pronunciation != nil {
		fmt.Printf("[美 %s]\t[英 %s]", res.Pronunciation.AmE, res.Pronunciation.BrE)
	}

	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
	
	for _, item := range res.Defs {
		fmt.Println(item.Pos, "\t", item.Def)
	}
	
	fmt.Println("---------------------------------------------------------------------------")
	
	if *verbose {
		for _, item := range res.Sams {
			fmt.Println(item.Eng)
			fmt.Println(item.Chn)
			fmt.Println()
		}
	} else {
		fmt.Printf("使用 bing -v %s 可以查看例句\n", strings.Join(flag.Args(), " "))
	}
	complete <- 0
}

func pronres(res *jsonapi.SearchResult) {
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
	
	go printres(result)
	go pronres(result)
	
	<- complete
	<- complete
}
