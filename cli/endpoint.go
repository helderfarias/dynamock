package cli

type EndpointFactory struct {
	uri         string
	service     *Service
	latency     int
	contentType string
	mockDir     string
	cors        *Cors
}

func (e *EndpointFactory) register(router RouterFactory) {
	for method, resp := range e.service.Responses {
		settings := &RouterSettings{
			Uri:         e.uri,
			ContentType: e.getContentType(),
			Status:      resp.Status,
			Latency:     e.getLatency(),
			Body:        resp.Body,
			BodyFile:    resp.BodyFile,
			Dynamic:     resp.Dynamic,
			Headers:     e.getHeaders(&resp, e.cors),
			MockDir:     e.mockDir,
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

func (e *EndpointFactory) getLatency() int {
	latency := e.latency

	if e.service.Latency != 0 {
		latency = e.service.Latency
	}

	return latency
}

func (e *EndpointFactory) getContentType() string {
	contentType := "text/plain"

	if len(e.service.ContentType) > 0 {
		contentType = e.service.ContentType
	} else {
		contentType = e.contentType
	}

	return contentType
}

func (e *EndpointFactory) getHeaders(r *Response, c *Cors) map[string]string {
	headers := make(map[string]string, 0)

	if len(e.service.Headers) > 0 {
		for k, v := range e.service.Headers {
			headers[k] = v
		}
	} else {
		for k, v := range r.Headers {
			headers[k] = v
		}
	}

	if c != nil {
		headers["Access-Control-Allow-Origin"] = c.AllowOrigin
		headers["Access-Control-Allow-Headers"] = c.AllowHeaders
		headers["Access-Control-Allow-Methods"] = c.AllowMethods
		headers["Access-Control-Allow-Credentials"] = c.AllowCredentials
		headers["Access-Control-Expose-Headers"] = c.ExposeHeaders
	}

	return headers
}
