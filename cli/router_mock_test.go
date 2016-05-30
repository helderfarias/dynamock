package cli

type routerFactoryMock struct {
	invokeGet bool
}

func (r *routerFactoryMock) CreateGET(data *RouterSettings) {
	r.invokeGet = true
}

func (r *routerFactoryMock) isInvokeCreateGET() bool {
	return r.invokeGet
}
