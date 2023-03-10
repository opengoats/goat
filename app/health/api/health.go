package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/opengoats/goat/app/health"
	"github.com/opengoats/goat/http/restful/response"
)

func (h *handler) Check(r *restful.Request, w *restful.Response) {
	req := health.NewHealthCheckRequest()
	resp, err := h.service.Check(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, NewHealth(resp))
}
