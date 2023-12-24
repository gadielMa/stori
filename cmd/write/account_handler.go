package write

import (
	"context"
	"github.com/satori/go.uuid"
	"net/http"
	"stori/cmd/logger"
)

type AccountHandler struct {
	AccountService AccountService
}

func NewAccountHandler(accountService AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: accountService,
	}
}

// Handle Execute godoc
// @Summary Account Handler
// @Description Account handler
// @Tags Account Handler
// @ID account_handler
// @Accept json
// @Produce json
// @Success 200 {object} account.Response "OK"
// @Failure 400 {object} Error "Bad Request"
// @Failure 500 {object} web.Error "Internal Server Error"
// @Router /stori/summary [post]
func (h *AccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), logger.RequestID, uuid.NewV4().String())
	err := h.AccountService.ProcessSummary(ctx)
	if err != nil {
		logger.Error(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	logger.Info(ctx, "success")
	w.WriteHeader(http.StatusOK)
}
