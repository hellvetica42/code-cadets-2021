package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/api/controllers/models"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/homework_4/02_bet_accept_api/internal/infrastructure/rabbitmq/models"
	"net/http"
)

const customerIdKey = "customer_id"
const selectionIdKey = "selection_id"
const selectionCoeffKey = "selection_coefficient"
const paymentKey = "payment"

// Controller implements handlers for web server requests.
type Controller struct {
	betReceivedValidator BetReceivedValidator
	betReceivedService   BetReceivedService
}

// NewController creates a new instance of Controller
func NewController(betReceivedValidator BetReceivedValidator, betReceivedService BetReceivedService) *Controller {
	return &Controller{
		betReceivedValidator: betReceivedValidator,
		betReceivedService:   betReceivedService,
	}
}

func (e *Controller) CreateBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betReceivedRequestDto models.BetReceivedRequestDto
		err := ctx.ShouldBind(&betReceivedRequestDto)
		if err != nil {
			ctx.String(http.StatusBadRequest, "new bet request is invalid")
			fmt.Println(err)
			fmt.Println(betReceivedRequestDto)
			return
		}

		if !e.betReceivedValidator.BetReceivedIsValid(betReceivedRequestDto) {
			ctx.String(http.StatusBadRequest, "new bet request validation failed")
		}

		//should create a separate mapper for this but it's 4am
		id, err := uuid.NewV4()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "failed to create new bet")
		}
		betReceivedDto := rabbitmqmodels.BetReceivedDto{
			Id:                   id.String(),
			CustomerId:           betReceivedRequestDto.CustomerId,
			SelectionId:          betReceivedRequestDto.SelectionId,
			SelectionCoefficient: betReceivedRequestDto.SelectionCoefficient,
			Payment:              betReceivedRequestDto.Payment,
		}

		err = e.betReceivedService.PublishBet(betReceivedDto)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "failed processing the bet request")
		}

		ctx.Status(http.StatusOK)
	}
}
