package feature

import (
	"sync"
)

// baseFeature 基礎應用功能
type baseFeature struct {
	enabled bool
	mux     sync.RWMutex
}

// isEnabled 檢查應用功能是否啟用
func (feature *baseFeature) isEnabled() bool {
	feature.mux.RLock()
	defer feature.mux.RUnlock()

	return feature.enabled
}

// setEnabledVar 設定啟用狀態
func (feature *baseFeature) setEnabledVar(enabled bool) {
	feature.mux.Lock()
	defer feature.mux.Unlock()

	feature.enabled = enabled
}
