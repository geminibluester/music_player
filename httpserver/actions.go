package httpserver

import (
	"context"
	"music_player/app"
	"music_player/model"
	"music_player/pkg/contant"
	"music_player/pkg/e"
	"music_player/service"
	"net/http"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func Find(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	n := ctx.DefaultQuery("n", contant.DEFAULT_CHAR)
	v := ctx.DefaultQuery("v", contant.DEFAULT_CHAR)
	svr := service.InfoService{}
	result, err := svr.FindByKey(n, v)
	if err != nil {
		appG.Response(http.StatusGone, e.ERROR, err.Error())
	}
	appG.Success(result)

}
func Create(ctx *gin.Context) {}
func Modifan(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is the edit page")
}
func Delete(ctx *gin.Context) {}
func Reload(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	appG.Success("sadjfhj")
}
func NoRouteFound(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_ARTICLE, nil)
}
func ChatWithAi(ctx *gin.Context) {
	client := openai.NewClient(model.AK)
	appG := app.Gin{C: ctx}
	q := ctx.DefaultQuery("q", "hello")
	if q == "" || q == "hello" {
		appG.Response(http.StatusConflict, e.ERROR, "param error")
		return
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			MaxTokens: 2048,
			Model:     openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: q,
				},
			},
		},
	)

	if err != nil {
		appG.Response(http.StatusConflict, e.ERROR, err.Error())
		return
	}
	appG.Success(resp.Choices[0].Message.Content)
}
