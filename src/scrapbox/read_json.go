package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	file, _ := ioutil.ReadFile("./home/data/daiiz.json")
	var project Project
	json.Unmarshal(file, &project)
	// fmt.Println(file)
	sampleTitle(project)
	countLines(project)
}

func countLines(project Project) {
	lines := 0
	var pages = project.Pages
	for i := 0; i < len(pages); i++ {
		lines += len(pages[i].Lines)
	}
	fmt.Printf("total lines: %d\n", lines)
}

func sampleTitle(project Project) {
	pageLength := len(project.Pages)

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(pageLength)

	// ランダムに記事タイトルを表示する
	fmt.Println(project.Pages[idx].Lines[0].Text)
}
