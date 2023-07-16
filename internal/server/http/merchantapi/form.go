package merchantapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"https://github.com/stream-innovations/streampayments-platfrom/internal/bus"
	"https://github.com/stream-innovations/streampayments-platfrom/internal/server/http/common"
	"https://github.com/stream-innovations/streampayments-platfrom/internal/server/http/middleware"
	"https://github.com/stream-innovations/streampayments-platfrom/pkg/api-dashboard/v1/model"
	"github.com/pkg/errors"
)

func (h *Handler) CreateFormSubmission(c echo.Context) error {
	var req model.FormRequest
	if !common.BindAndValidateRequest(c, &req) {
		return nil
	}

	u := middleware.ResolveUser(c)
	if u == nil {
		return errors.New("unable to resolve user")
	}

	event := bus.FormSubmittedEvent{
		RequestType: *req.Topic,
		Message:     *req.Message,
		UserID:      u.ID,
	}

	if err := h.publisher.Publish(bus.TopicFormSubmissions, event); err != nil {
		return errors.Wrap(err, "unable to publish event")
	}

	return c.NoContent(http.StatusCreated)
}
