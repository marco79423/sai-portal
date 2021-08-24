package feature

import (
	"context"

	notifier2 "github.com/marco79423/sai-portal/service/core/notifier"
	"go.uber.org/dig"
)

// IAppFeature 應用功能的介面
type IAppFeature interface {
	GetID() string
	GetName() string

	Enable(ctx context.Context) error
	Disable(ctx context.Context) error
}

// AppFeatureSetInterface 應用功能集合的介面
type AppFeatureSetInterface interface {
	GetFeatures() []IAppFeature
}

// NewAppFeatureSet 建立應用功能集合
func NewAppFeatureSet(in In) AppFeatureSetInterface {
	return &appFeatureSet{
		features: []IAppFeature{
			// HTTP 服務器
			newHTTPFeature(in),

			// 探測器功能：提供監控和通知的功能
			newRedirectsFeature(in),
		},
	}
}

// In AppFeatureSet 所需要的引入
type In struct {
	dig.In

	OddsDelayedNotifier  notifier2.OddsDelayedNotifierInterface
	ServiceErrorNotifier notifier2.ServiceErrorNotifierInterface
}

type appFeatureSet struct {
	features []IAppFeature
}

// GetFeatures 取得所有可用的應用功能
func (appFeatures *appFeatureSet) GetFeatures() []IAppFeature {
	return appFeatures.features
}
