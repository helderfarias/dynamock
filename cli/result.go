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
	var result string

	if len(data.Body) == 0 {
		result = loadContentFromFile(makeBodyFile(data.MockDir, data.BodyFile))
	} else {
		result = data.Body
	}

	buffer := parseTemplate(data.TemplateTokens, result)

	return data.Status, toJSON(buffer)
}

func loadContentFromFile(f string) string {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(content)
}

func toJSON(content string) interface{} {
	var d interface{}
	err := json.Unmarshal([]byte(content), &d)
	if err != nil {
		log.Println(err)
		return ""
	}
	return d
}
