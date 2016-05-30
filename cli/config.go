package cli

type Configuration struct {
	Port     string
	Latency  int
	Services map[string]Service
}
