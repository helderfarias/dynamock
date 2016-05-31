package cli

type Response struct {
	Status   int
	Body     string
	BodyFile string
	Dynamic  map[string]interface{}
}
