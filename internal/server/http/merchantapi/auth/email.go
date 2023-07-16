package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"https://github.com/stream-innovations/streampayments-platfrom/internal/server/http/common"
	"https://github.com/stream-innovations/streampayments-platfrom/internal/server/http/middleware"
	"https://github.com/stream-innovations/streampayments-platfrom/internal/service/user"
	"https://github.com/stream-innovations/streampayments-platfrom/pkg/api-dashboard/v1/model"
	"github.com/pkg/errors"
)

func (h *Handler) PostLogin(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.LoginRequest
	if !common.BindAndValidateRequest(c, &req) {
		return nil
	}

	// already logged in
	if u := middleware.ResolveUser(c); u != nil {
		return c.NoContent(http.StatusNoContent)
	}

	person, err := h.users.GetByEmailWithPasswordCheck(ctx, req.Email.String(), req.Password)
	switch {
	case errors.Is(err, user.ErrNotFound), errors.Is(err, user.ErrWrongPassword):
		return common.ValidationErrorItemResponse(c, "email", "User with provided email or password not found")
	case err != nil:
		return errors.Wrap(err, "unable to resolve user")
	}

	if err := h.persistSessionUserID(c, person.ID, "email"); err != nil {
		return common.ErrorResponse(c, "internal error")
	}

	return c.NoContent(http.StatusNoContent)
}
