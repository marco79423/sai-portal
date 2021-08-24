package app

import (
	"github.com/marco79423/sai-portal/service/app/feature"
	"github.com/marco79423/sai-portal/service/binder"
	"github.com/marco79423/sai-portal/service/utils"
	"golang.org/x/xerrors"
)

// enableFeatures 啟用所有應用功能
func (app *app) enableFeatures() error {
	logger := utils.GetCtxLogger(app.ctx)

	// 取得所有可用的功能
	availableFeatures, err := app.prepareAvailableFeatures()
	if err != nil {
		return xerrors.Errorf("啟用應用功能失敗: %w", err)
	}

	// 啟用應用功能
	conf := utils.GetCtxConfig(app.ctx)
	for _, availableFeature := range availableFeatures {
		// 檢查是否需啟動應用功能
		if !conf.IsEnabledFeature(availableFeature.GetID()) {
			continue
		}

		// 啟動服務
		logger.Info(app.ctx, "啟用應用功能 %s...", availableFeature.GetName())
		if err := availableFeature.Enable(app.ctx); err != nil {
			return xerrors.Errorf("啟用應用功能失敗: %w", err)
		}

		app.enabledFeatures = append(app.enabledFeatures, availableFeature)
		logger.Info(app.ctx, "啟用應用功能 %s 成功", availableFeature.GetName())
	}

	return nil
}

// disableFeatures 停用所有應用功能
func (app *app) disableFeatures() {
	logger := utils.GetCtxLogger(app.ctx)
	for _, enabledFeature := range app.enabledFeatures {
		logger.Info(app.ctx, "停用應用功能 %s...", enabledFeature.GetName())
		if err := enabledFeature.Disable(app.ctx); err != nil {
			logger.Error(app.ctx, xerrors.Errorf("停用應用功能 %s 失敗: %w", enabledFeature.GetName(), err))
			continue
		}
		logger.Info(app.ctx, "停用應用功能 %s 成功", enabledFeature.GetName())
	}
	app.enabledFeatures = nil
}

// prepareAvailableFeatures 準備所有可用的應用功能
func (app *app) prepareAvailableFeatures() ([]feature.IAppFeature, error) {
	// 建立 Binder
	b, err := binder.NewBinder()
	if err != nil {
		return nil, xerrors.Errorf("建立可用的應用功能失敗: %w", err)
	}

	// 建立服務
	var availableAppFeatures []feature.IAppFeature
	if err := b.Invoke(func(appFeatureSet feature.AppFeatureSetInterface) {
		availableAppFeatures = appFeatureSet.GetFeatures()
	}); err != nil {
		return nil, xerrors.Errorf("建立可用的應用功能失敗: %w", err)
	}

	return availableAppFeatures, nil
}
