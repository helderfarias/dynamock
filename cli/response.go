package cli

type Response struct {
	Status   int
	Body     string
	BodyFile string
	Headers  map[string]string
	Dynamic  map[string]interface{}
}
