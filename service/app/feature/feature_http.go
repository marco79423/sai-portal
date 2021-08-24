package feature

import (
	"context"
	"net"
	"net/http"

	"github.com/marco79423/sai-portal/service/utils"
	"golang.org/x/xerrors"
)

func newHTTPFeature(in In) IAppFeature {
	return &httpFeature{
		in: in,
	}
}

type httpFeature struct {
	in In

	httpServer *http.Server

	baseFeature // 應用功能基礎功能
}

// GetID 取得應用功能的 ID
func (feature *httpFeature) GetID() string {
	return "http_feature"
}

// GetName 取得應用功能的名稱
func (feature *httpFeature) GetName() string {
	return "HTTP 服務器"
}

// Enable 啟動應用功能
func (feature *httpFeature) Enable(ctx context.Context) error {
	// 檢查是否重覆啟用
	if feature.isEnabled() {
		return xerrors.Errorf("啟用 %s 失敗: 不支援重覆啟動", feature.GetName())
	}

	// 始初化
	conf := utils.GetCtxConfig(ctx)
	feature.httpServer = &http.Server{
		Addr:    conf.GetHTTPFeatureConfig().GetAddress(),
		Handler: feature.getRouteHandler(ctx),
	}

	// 啟動
	go func() {
		ln, err := net.Listen("tcp", conf.GetHTTPFeatureConfig().GetAddress())
		if err != nil {
			logger := utils.GetCtxLogger(ctx)
			logger.Error(ctx, xerrors.Errorf("啟用 %s 失敗: 啟動服務器失敗: %w", feature.GetName(), err))
		}

		_ = feature.httpServer.Serve(ln)
	}()

	// 更新應用功能狀態
	feature.setEnabledVar(true)
	return nil
}

// Disable 關閉應用功能
func (feature *httpFeature) Disable(ctx context.Context) error {
	err := feature.httpServer.Shutdown(ctx)
	if err != nil {
		return xerrors.Errorf("停用 %s 失敗: %w", feature.GetName(), err)
	}
	feature.setEnabledVar(false)
	return nil
}
