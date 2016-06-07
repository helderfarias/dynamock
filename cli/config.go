package cli

type Configuration struct {
	Port        string
	Latency     int
	ContentType string
	Services    map[string]Service
	MockDir     string
}
