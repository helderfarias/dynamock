package cli

import (
	"log"
	"math/rand"
	"time"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	qrcode "github.com/skip2/go-qrcode"
)

type RandomPlugin struct {
	Status   []int
	Body     []string
	BodyFile []string
	MockDir  string
}

type SwitchPlugin struct {
	Input   interface{}
	Context echo.Context
	MockDir string
}

type QrCodePlugin struct {
	Status      int
	Content     string
	Quality     string
	Size        int
	Context     echo.Context
	MockDir     string
	ContentType string
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
	body := ""
	bodyFile := ""

	if len(r.Body) > 0 {
		body = r.Body[rand.Intn(len(r.Body))]
	} else if len(r.BodyFile) > 0 {
		bodyFile = r.BodyFile[rand.Intn(len(r.BodyFile))]
	}

	return createSingleResult(&RouterSettings{
		Status:   status,
		Body:     body,
		BodyFile: makeBodyFile(r.MockDir, bodyFile),
	})
}

func (s *SwitchPlugin) Create() (int, interface{}) {
	mapper := make(map[string][]Action, 0)
	params := make(map[string]string, 0)

	for key, value := range s.Input.(map[string]interface{}) {
		var actions []Action

		for _, data := range value.([]interface{}) {
			var action Action
			mapstructure.Decode(data, &action)
			actions = append(actions, action)
		}

		mapper[key] = actions
	}

	for key := range mapper {
		query := s.Context.QueryParam(key)
		if len(query) > 0 {
			params[key] = query
		}

		param := s.Context.Param(key)
		if len(param) > 0 {
			params[key] = param
		}
	}

	for key, value := range params {
		for _, action := range mapper[key] {
			if action.If == value {
				return createSingleResult(&RouterSettings{
					Status:   action.When.Status,
					Body:     action.When.Body,
					BodyFile: makeBodyFile(s.MockDir, action.When.BodyFile),
				})
			}
		}
	}

	return 404, "action not found"
}

func (q *QrCodePlugin) Create() (int, interface{}) {
	if q.ContentType != "image/png" {
		log.Fatalln("ContentType invalid. The qrcode plugin supports 'image/png' only.")
		return 0, nil
	}

	levels := map[string]qrcode.RecoveryLevel{
		"low":    qrcode.Low,
		"medium": qrcode.Medium,
		"high":   qrcode.Highest,
	}

	status := q.Status
	content := q.Content
	quality := levels["medium"]
	size := 256

	if q.Quality != "" {
		quality = levels[q.Quality]
	}

	if q.Size >= 56 {
		size = q.Size
	}

	png, err := qrcode.Encode(content, quality, size)
	if err != nil {
		log.Println(err)
	}

	return status, png
}
