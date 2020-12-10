package configuration

import (
	"github.com/containerssh/http"
	"github.com/containerssh/log"
)

type handler struct {
	handler ConfigRequestHandler
	logger  log.Logger
}

func (h *handler) OnRequest(request http.ServerRequest, response http.ServerResponse) error {
	requestObject := ConfigRequest{}
	if err := request.Decode(&requestObject); err != nil {
		return err
	}
	appConfig, err := h.handler.OnConfig(requestObject)
	if err != nil {
		return err
	}
	responseObject := ConfigResponseBody{
		Config: appConfig,
	}
	response.SetBody(responseObject)
	response.SetStatus(200)
	return nil
}
