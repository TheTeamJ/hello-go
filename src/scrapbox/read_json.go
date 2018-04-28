package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("./home/data/daiiz.json")
	var projects Project
	json.Unmarshal(file, &projects)
	fmt.Println(file)
	fmt.Println(projects.Pages[0].Lines[0])
}
