package feature

import (
	"context"
	"net/http"
	"runtime"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/marco79423/sai-portal/service/model/dto"
	"github.com/marco79423/sai-portal/service/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (feature *httpFeature) getRouteHandler(ctx context.Context) http.Handler {
	handler := gin.New()
	handler.Use(
		gin.Recovery(),

		// 注入基本資訊
		func(ginCtx *gin.Context) {
			ginCtx.Set("logger", utils.GetCtxLogger(ctx))
			ginCtx.Set("config", utils.GetCtxConfig(ctx))
		},
	)

	feature.setAdminRoutes(handler)
	feature.setInternalRoutes(handler)

	return handler
}

func (feature *httpFeature) setAdminRoutes(handler *gin.Engine) {
	privateRouteGroup := handler.Group("/admin-api")
	privateRouteGroup.POST("/ctrl/claim-system-stability", func(ctx *gin.Context) {
		err := feature.in.ServiceErrorNotifier.ChangeStateBackToNormalDirectly(ctx)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				dto.APIResponse{
					Meta: dto.APIResponseMeta{
						RequestID: ctx.Query("requestID"),
					},
					Code:    "0",
					Message: "確認系統恢復正常失敗",
				},
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			dto.APIResponse{
				Meta: dto.APIResponseMeta{
					RequestID: ctx.Query("requestID"),
				},
				Code:    "0",
				Message: "確認系統恢復正常成功",
			},
		)
	})
}

func (feature *httpFeature) setInternalRoutes(handler *gin.Engine) {
	privateRouteGroup := handler.Group("/_")

	// config
	privateRouteGroup.GET("/config", func(ctx *gin.Context) {
		config := utils.GetCtxConfig(ctx)
		ctx.IndentedJSON(http.StatusOK, config.GetRawConfig())
	})

	// health check
	privateRouteGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// prometheus
	prometheusHandler := promhttp.Handler()
	privateRouteGroup.GET("/metrics", func(c *gin.Context) {
		prometheusHandler.ServeHTTP(c.Writer, c.Request)
	})

	// pprof
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	pprof.Register(handler, "/_/debug")
}
