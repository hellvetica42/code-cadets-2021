package controllers

import (
	"encoding/json"
	"github.com/superbet-group/code-cadets-2021/homework4/01_bet_api/internal/api/controllers/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
)

const idKey = "id"
const statusKey = "status"

// Controller implements handlers for web server requests.
type Controller struct {
	betRequestValidator BetRequestValidator
	betRepository BetRepository
	dtoMapper DtoBetMapper
}

// NewController creates a new instance of Controller
func NewController(betRequestValidator BetRequestValidator, betRepository BetRepository, dtoMapper DtoBetMapper) *Controller {
	return &Controller{
		betRequestValidator: betRequestValidator,
		betRepository: betRepository,
		dtoMapper: dtoMapper,
	}
}

func (e *Controller) GetBetByID() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id := ctx.Param(idKey)

		if !e.betRequestValidator.BetRequestIdIsValid(id) {
			ctx.String(http.StatusBadRequest, "bet id is invalid")
			return
		}

		bet, exists, err := e.betRepository.GetBetByID(ctx, id)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "error fetching bet")
			return
		}

		if !exists {
			ctx.String(http.StatusNotFound, "no such bet in repository")
			return
		}

		dto := e.dtoMapper.MapDomainBetToDtoBet(bet)

		dtoJson, err := json.Marshal(dto)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "error converting to json")
			return
		}

		ctx.Data(http.StatusOK, "application/json", dtoJson)
		return
	}
}

func (e *Controller) GetUserBets() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		userId := ctx.Param(idKey)

		if !e.betRequestValidator.BetRequestUserIdIsValid(userId) {
			ctx.String(http.StatusBadRequest, "user id is invalid")
			return
		}

		bets, exists, err := e.betRepository.GetBetsByUserID(ctx, userId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "error fetching bets")
			return
		}

		if !exists {
			ctx.String(http.StatusNotFound, "no such bets in repository")
			return
		}

		var dtoBets []models.BetResponseDto

		for _, bet := range bets {
			dtoBet := e.dtoMapper.MapDomainBetToDtoBet(bet)
			dtoBets = append(dtoBets, dtoBet)
		}

		dtoJson, err := json.Marshal(dtoBets)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "error converting to json")
			return
		}

		ctx.Data(http.StatusOK, "application/json", dtoJson)
		return

	}
}

func (e *Controller) GetBetsByStatus() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		status := ctx.Query(statusKey)
		if len(status) == 0 {
			ctx.String(http.StatusBadRequest, "status not specified")
			return
		}

		if !e.betRequestValidator.BetRequestStatusIsValid(status){
			ctx.String(http.StatusBadRequest, "status is invalid")
			return
		}

		bets, exists, err := e.betRepository.GetBetsByStatus(ctx, status)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "error fetching bets")
			return
		}

		if !exists {
			ctx.String(http.StatusNotFound, "no such bets in repository")
			return
		}

		var dtoBets []models.BetResponseDto

		for _, bet := range bets {
			dtoBet := e.dtoMapper.MapDomainBetToDtoBet(bet)
			dtoBets = append(dtoBets, dtoBet)
		}

		dtoJson, err := json.Marshal(dtoBets)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "error converting to json")
			return
		}

		ctx.Data(http.StatusOK, "application/json", dtoJson)
		return
	}
}
