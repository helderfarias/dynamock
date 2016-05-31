package cli

type EndpointFactory struct {
	uri         string
	service     *Service
	latency     int
	contentType string
}

func (e *EndpointFactory) register(router RouterFactory) {
	for _, resp := range e.service.Responses {
		router.CreateGET(&RouterSettings{
			Uri:         e.uri,
			ContentType: e.getContentType(),
			Status:      resp.Status,
			Latency:     e.latency,
			Body:        resp.Body,
			BodyFile:    resp.BodyFile,
		})
	}
}

func (e *EndpointFactory) getContentType() string {
	var contentType string

	if len(e.service.ContentType) > 0 {
		contentType = e.service.ContentType
	} else {
		contentType = e.contentType
	}

	return contentType
}
