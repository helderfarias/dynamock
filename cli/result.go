package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func makeBodyFile(mockDir string, bodyFile string) string {
	separator := fmt.Sprintf("%c", os.PathSeparator)
	file := fmt.Sprintf("%s%s%s", mockDir, separator, bodyFile)

	if strings.HasPrefix(file, separator) {
		return file[1:]
	}

	return file
}

func createSingleResult(data *RouterSettings) (int, interface{}) {
	var result interface{}

	if len(data.Body) == 0 {
		result = parseFile(makeBodyFile(data.MockDir, data.BodyFile))
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
