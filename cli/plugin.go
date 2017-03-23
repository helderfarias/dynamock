package cli

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	qrcode "github.com/skip2/go-qrcode"
)

type RandomPlugin struct {
	Status         []int
	Body           []string
	BodyFile       []string
	MockDir        string
	TemplateTokens map[string]string
}

type SwitchPlugin struct {
	Input          interface{}
	Context        echo.Context
	MockDir        string
	TemplateTokens map[string]string
}

type QrCodePlugin struct {
	Status         int
	Content        string
	Quality        string
	Size           string
	Context        echo.Context
	MockDir        string
	ContentType    string
	TemplateTokens map[string]string
}

type JWTPlugin struct {
	Status         int
	Alg            string
	Payload        string
	Secret         string
	MockDir        string
	Output         string
	Context        echo.Context
	TemplateTokens map[string]string
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
		Status:         status,
		Body:           body,
		TemplateTokens: r.TemplateTokens,
		BodyFile:       makeBodyFile(r.MockDir, bodyFile),
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
					Status:         action.When.Status,
					Body:           action.When.Body,
					TemplateTokens: s.TemplateTokens,
					BodyFile:       makeBodyFile(s.MockDir, action.When.BodyFile),
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
	content := parseTemplate(q.TemplateTokens, q.Content)
	quality := levels["medium"]
	size := 256

	parseQuality := parseTemplate(q.TemplateTokens, q.Quality)
	if levels[parseQuality] != 0 {
		quality = levels[parseQuality]
	}

	parseSize := parseTemplate(q.TemplateTokens, q.Size)
	if val, err := strconv.Atoi(parseSize); err == nil && val >= 256 {
		size = val
	}

	png, err := qrcode.Encode(content, quality, size)
	if err != nil {
		log.Println(err)
	}

	return status, png
}

func (q *JWTPlugin) Create() (int, interface{}) {
	mapAlgs := map[string]jwt.SigningMethod{
		"HS256": jwt.SigningMethodHS256,
		"HS512": jwt.SigningMethodHS512,
		"ES256": jwt.SigningMethodES256,
		"ES512": jwt.SigningMethodES512,
	}

	claims := jwt.MapClaims{}
	claims["iat"] = time.Now().UTC().Unix()

	if val := parseTemplate(q.TemplateTokens, q.Payload); val != "" {
		payload := fromJsonToMap(val)

		if maps, ok := payload.(map[string]interface{}); ok {
			for key, val := range maps {
				claims[key] = val
			}
		}
	}

	secret := "secret"
	if val := parseTemplate(q.TemplateTokens, q.Secret); val != "" {
		secret = val
	}

	alg := "HS256"
	if val := parseTemplate(q.TemplateTokens, q.Alg); val != "" {
		alg = val
	}

	if val := mapAlgs[alg]; val == nil {
		log.Fatalf("alg is not defined")
		return 0, nil
	}

	signer := jwt.NewWithClaims(mapAlgs[alg].(jwt.SigningMethod), claims)
	token, err := signer.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}

	var output interface{}

	q.TemplateTokens["@token"] = token

	if val := parseTemplate(q.TemplateTokens, q.Output); val != "" {
		output = fromJsonToMap(val)
	} else {
		output = map[string]string{"access_token": token}
	}

	return q.Status, output
}
