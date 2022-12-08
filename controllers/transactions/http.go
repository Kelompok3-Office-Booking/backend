package transactions

import (
	transactions "backend/businesses/transactions"
	ctrl "backend/controllers"
	"backend/controllers/transactions/request"
	"backend/controllers/transactions/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUsecase transactions.Usecase
}

func NewTransactionController(tc transactions.Usecase) *TransactionController {
	return &TransactionController{
		TransactionUsecase: tc,
	}
}

func (t *TransactionController) GetAll(c echo.Context) error {
	TransactionsData := t.TransactionUsecase.GetAll()

	Transactions := []response.Transaction{}

	for _, trans := range TransactionsData {
		Transactions = append(Transactions, response.FromDomain(trans))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all transactions", Transactions)
}

func (t *TransactionController) Create(c echo.Context) error {
	input := request.Transaction{}

	if err := c.Bind(&input); err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "validation failed")
	}

	err := input.Validate()

	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "validation failed")
	}

	trans := t.TransactionUsecase.Create(input.ToDomain())

	if trans.ID == 0 {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "transaction failed")
	}

	return ctrl.NewResponse(c, http.StatusCreated, "success", "transaction created", response.FromDomain(trans))
}
