package cli

type Cors struct {
	AllowOrigin   string `json:"allow_origin"`
	AllowHeaders  string `json:"allow_headers"`
	AllowMethods  string `json:"allow_methods"`
	ExposeHeaders string `json:"expose_headers"`
}

type Configuration struct {
	Port        string
	Latency     int
	ContentType string
	Services    map[string]Service
	MockDir     string
	Cors        *Cors
}
