package binder

import (
	"github.com/marco79423/sai-portal/service/app/feature"
	"go.uber.org/dig"
	"golang.org/x/xerrors"
)

func NewBinder() (*dig.Container, error) {
	binder := dig.New()

	// App
	if err := binder.Provide(feature.NewAppFeatureSet); err != nil {
		return nil, xerrors.Errorf("無法正常建立依賴: %w", err)
	}

	return binder, nil
}
