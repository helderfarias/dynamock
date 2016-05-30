package cli

type EndpointFactory struct {
	uri     string
	service *Service
	latency int
}

func (e *EndpointFactory) register(router RouterFactory) {
	for _, method := range e.service.Verbs {
		resp := e.service.Responses[method]

		if resp.Status != 0 {
			router.CreateGET(&RouterSettings{
				Uri:         e.uri,
				ContentType: e.service.ContentType,
				Status:      resp.Status,
				Result:      resp.Result,
				Latency:     e.latency,
				File:        resp.File,
			})
		}
	}
}
