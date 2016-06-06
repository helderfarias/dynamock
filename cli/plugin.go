package cli

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type RandomPlugin struct {
	Status   []int
	Body     []string
	BodyFile []string
}

type SwitchPlugin struct {
	Input   map[string][]Action
	Context *gin.Context
}

type Action struct {
	If   string
	When *Response
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (r *RandomPlugin) Create() (int, interface{}) {
	status := r.Status[rand.Intn(len(r.Status))]

	var data interface{}
	if len(r.Body) > 0 {
		data = r.Body[rand.Intn(len(r.Body))]
	} else if len(r.BodyFile) > 0 {
		bodyFile := r.BodyFile[rand.Intn(len(r.BodyFile))]
		data = parseFile(bodyFile)
	}

	return status, data
}

func (s *SwitchPlugin) Create() (int, interface{}) {
	for key, actions := range s.Input {
		// param := s.Context.Query(key)
		log.Println(key, actions)
	}

	log.Println(s.Input)

	return 0, nil
}
