package cli

func (t *CliSuite) TestShouldBeCreateEnpointWithMethodGet() {
	routerMock := &routerFactoryMock{}

	api := &EndpointFactory{
		uri: "api/ping",
		service: &Service{
			Verbs: []string{"get"},
			Responses: map[string]Response{
				"get": Response{Status: 200},
			},
		},
		latency: 0,
	}

	api.register(routerMock)

	t.Equal(routerMock.isInvokeCreateGET(), true)
}
