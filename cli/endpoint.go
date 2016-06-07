package cli

type EndpointFactory struct {
	uri         string
	service     *Service
	latency     int
	contentType string
}

func (e *EndpointFactory) register(router RouterFactory) {
	for method, resp := range e.service.Responses {
		settings := &RouterSettings{
			Uri:         e.uri,
			ContentType: e.getContentType(),
			Status:      resp.Status,
			Latency:     e.latency,
			Body:        resp.Body,
			BodyFile:    resp.BodyFile,
			Dynamic:     resp.Dynamic,
			Headers:     e.getHeaders(&resp),
		}

		if method == "get" {
			router.CreateGET(settings)
		}

		if method == "post" {
			router.CreatePOST(settings)
		}

		if method == "put" {
			router.CreatePUT(settings)
		}

		if method == "delete" {
			router.CreateDELETE(settings)
		}
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

func (e *EndpointFactory) getHeaders(r *Response) map[string]string {
	var headers map[string]string

	if len(e.service.Headers) > 0 {
		headers = e.service.Headers
	} else {
		headers = r.Headers
	}

	return headers
}
