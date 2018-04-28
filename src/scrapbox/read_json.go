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
	var projects Project
	json.Unmarshal(file, &projects)
	// fmt.Println(file)
	sample(projects)
}

func sample(projects Project) {
	pageLength := len(projects.Pages)

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(pageLength)

	// ランダムに記事タイトルを表示する
	fmt.Println(projects.Pages[idx].Lines[0].Text)
}
