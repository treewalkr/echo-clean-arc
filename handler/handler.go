package handler

type HttpHandlers struct {
	HelloHandler *HelloHandler
	UserHandler  *UserHandler
}

func New(hh *HelloHandler, uh *UserHandler) *HttpHandlers {
	return &HttpHandlers{hh, uh}
}
