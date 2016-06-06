package cli

type Service struct {
	ContentType string
	Headers     map[string]string
	Responses   map[string]Response
}
