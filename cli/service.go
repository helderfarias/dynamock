package cli

type Service struct {
	Verbs       []string
	ContentType string
	Responses   map[string]Response
}
