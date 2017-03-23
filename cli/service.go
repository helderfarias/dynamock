package cli

type Service struct {
	ContentType string
	Headers     map[string]string
	Cors        bool
	Responses   map[string]Response
	Latency     int
}
