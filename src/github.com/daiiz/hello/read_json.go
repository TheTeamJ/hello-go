package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type project struct {
	Name        string
	DisplayName string
	Exported    uint
	Pages       []PagesType
}

type PagesType struct {
	Title   string
	Created uint
	Updated uint
	Lines   []LinesType
}

type LinesType struct {
	Text    string
	Created uint
	Updated uint
}

func main() {
	file, _ := ioutil.ReadFile("./home/data/daiiz.json")
	var projects project
	json.Unmarshal(file, &projects)
	fmt.Println(file)
	fmt.Println(projects.Pages[0].Lines[0])
}
