package paymentmethods

import (
	// "github.com/PICH-IO/admin-api/pkg/constants"
	util_common "github.com/PICH-IO/admin-api/pkg/utils/common"
	util_response "github.com/PICH-IO/admin-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type PaymentHandler struct {
	s *PaymentService
}

func NewPaymentHandler(db *sqlx.DB) *PaymentHandler {
	return &PaymentHandler{
		s: NewPaymentService(db),
	}
}

func (p *PaymentHandler) Show(c *fiber.Ctx) error {
	var data, err = p.s.Show()
	// fmt.Println("error:", err)
	if err != nil {
		var msg, err_msg = util_common.TranslateWithError(c, "DataNotFound")
		if err_msg != nil {
			err := util_response.HttpResponse(
				false,
				err_msg.Error(),
				400, //constants.ErrorTranslate,
				nil,
			)
			return c.Status(500).JSON(err)
		}
		err_response := util_response.HttpResponse(
			false,
			msg,
			400, //constants.GetPaymentDataFailed,
			nil,
		)
		return c.Status(500).JSON(err_response)
	}
	// fmt.Println("hi")
	var success = util_common.Translate(c, "QueryPaymentDataSuccess")
	dataResponse := util_response.HttpResponse(
		true,
		success,
		400, //constants.PaymentQueryDataSuccessfully,
		data,
	)
	return c.Status(200).JSON(dataResponse)
}
