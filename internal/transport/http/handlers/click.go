package handlers

import (
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris/v12"
)

func (h *Handlers) ClickBanner(ctx iris.Context) {
	bodyUrl := ctx.Params().Get("bodyUrl")

	clickData, err := h.ClickerService.DecodeProtoClick(bodyUrl)
	if err != nil {
		ctx.StatusCode(400)
		_, _ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"code":    400,
				"message": err,
			},
			"data": nil,
		})
		return
	}

	uniqueness, err := h.ClicksRepository.CheckUniqueness(clickData.GetId())
	if err != nil {
		// TODO: remake
		logger.Error(err)
	}

	if !uniqueness || clickData.GetUa() != ctx.Request().UserAgent() {
		err := h.ClicksRepository.Create(clickData, true)
		if err != nil {
			// TODO: remake
			logger.Error(err)
		}

		ctx.StatusCode(301)
		ctx.Header("Location", clickData.RedirectUri)
		return
	}

	err = h.ClicksRepository.Create(clickData, false)
	if err != nil {
		// TODO: remake
		logger.Error(err)
	}

	ctx.StatusCode(301)
	ctx.Header("Location", clickData.RedirectUri)
	return
}
