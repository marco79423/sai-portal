package feature

import (
	"context"

	"github.com/marco79423/sai-portal/service/config"
	"github.com/marco79423/sai-portal/service/core/notifier"
	"github.com/marco79423/sai-portal/service/utils"
	"golang.org/x/xerrors"
)

// newRedirectsFeature 建立探測器服務功能
func newRedirectsFeature(in In) IAppFeature {
	return &notifierFeature{
		notifiers: []notifier.NotifierInterface{
			// 賠率延遲探測器
			in.OddsDelayedNotifier,
			// 服務異常探測器
			in.ServiceErrorNotifier,
		},
	}
}

type notifierFeature struct {
	notifiers []notifier.NotifierInterface

	baseFeature // 應用功能基礎功能
}

// GetID 取得應用功能的 ID
func (feature *notifierFeature) GetID() string {
	return "notifier_feature"
}

// GetName 取得應用功能的名稱
func (feature *notifierFeature) GetName() string {
	return "探測器功能"
}

// Enable 啟動應用功能
func (feature *notifierFeature) Enable(ctx context.Context) error {
	// 檢查是否重覆啟用
	if feature.isEnabled() {
		return xerrors.Errorf("啟用 %s 失敗: 不支援重覆啟動", feature.GetName())
	}

	// 啟動 Notifier
	logger := utils.GetCtxLogger(ctx)
	notifierConfig := feature.getNotifierConfig(ctx)
	for _, n := range feature.notifiers {
		if notifierConfig.IsEnabledNotifier(n.GetID()) {
			if err := n.Start(ctx); err != nil {
				return xerrors.Errorf("啟用 %s 失敗: %w", feature.GetName(), err)
			}
			logger.Info(ctx, "啟動 %s 成功", n.GetName())
		}
	}

	// 更新應用功能狀態
	feature.setEnabledVar(true)
	return nil
}

// Disable 關閉應用功能
func (feature *notifierFeature) Disable(ctx context.Context) error {
	logger := utils.GetCtxLogger(ctx)
	for _, n := range feature.notifiers {
		if err := n.Stop(ctx); err != nil {
			logger.Error(ctx, xerrors.Errorf("停用 %s 失敗: %w", n.GetName(), err))
			continue
		}
		logger.Info(ctx, "停用 %s 成功", n.GetName())
	}
	return nil
}

// getNotifierConfig 取得探測器的設定資訊
func (feature *notifierFeature) getNotifierConfig(ctx context.Context) config.IRedirectsFeatureConfig {
	return utils.GetCtxConfig(ctx).GetRedirectsFeatureConfig()
}
