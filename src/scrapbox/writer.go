package main

import (
	"encoding/json"
	"io/ioutil"
)

func saveLog(log []byte) {
	log = append(log, '\n')
	_ = ioutil.WriteFile("./log.txt", log, 0644)
}

func saveAsJson(project Project) {
	outputJson, _ := json.Marshal(project)
	_ = ioutil.WriteFile("./log.json", outputJson, 0644)
}
