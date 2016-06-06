package cli

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func createSingleResult(data *RouterSettings) (int, interface{}) {
	var result interface{}

	if len(data.BodyFile) > 0 {
		result = parseFile(data.BodyFile)
	} else {
		result = data.Body
	}

	return data.Status, result
}

func parseFile(f string) interface{} {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return ""
	}

	var d interface{}
	err = json.Unmarshal(content, &d)
	if err != nil {
		log.Println(err)
		return ""
	}

	return d
}
