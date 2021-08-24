package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

type IConfig interface {
	GetName() string     // 名稱
	GetLogLevel() string // Log 層級

	IsEnabledFeature(serviceName string) bool // 是否為要啟動的功能

	GetRedirectsFeatureConfig() IRedirectsFeatureConfig // 探測器功能相關設定

	GetRawConfig() *rawConfig
}

func NewConfig() (IConfig, error) {
	// 要讀取的設定檔
	configPaths := []string{
		"conf.d/config.yaml",
	}

	// 讀取設定檔
	rawConfig := &rawConfig{}
	for _, configPath := range configPaths {
		if err := loadConfig(rawConfig, configPath); err != nil {
			return nil, xerrors.Errorf("無法取得設定檔: %w", err)
		}
	}

	return &config{
		rawConfig: rawConfig,
	}, nil
}

type config struct {
	rawConfig *rawConfig
}

// GetName 取得主應用的名稱
func (conf *config) GetName() string {
	return conf.rawConfig.App.Name
}

// GetLogLevel 取得顯示 Log 層級的設定
func (conf *config) GetLogLevel() string {
	return conf.rawConfig.App.LogLevel
}

// IsEnabledFeature 判斷是否為啟用的應用功能
func (conf *config) IsEnabledFeature(featureName string) bool {
	for _, enabledFeature := range conf.rawConfig.App.EnabledFeatures {
		if featureName == enabledFeature {
			return true
		}
	}
	return false
}

// GetRedirectsFeatureConfig 取得轉發功能的設定
func (conf *config) GetRedirectsFeatureConfig() IRedirectsFeatureConfig {
	return newRedirectsFeatureConfig(conf.rawConfig)
}

// loadConfig 讀取設定檔
func loadConfig(rawConfig interface{}, configPath string) error {
	absConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		return xerrors.Errorf("無法讀取設定檔 %s: %w", configPath, err)
	}

	if _, err := os.Stat(absConfigPath); os.IsNotExist(err) {
		return xerrors.Errorf("無法讀取設定檔 %s: %w", configPath, err)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(absConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		return xerrors.Errorf("無法讀取設定檔 %s: %w", configPath, err)
	}

	if err := viper.Unmarshal(rawConfig); err != nil {
		return xerrors.Errorf("無法讀取設定檔 %s: %w", configPath, err)
	}

	return nil
}

// GetRawConfig 直接取得原始設定檔內容
func (conf *config) GetRawConfig() *rawConfig {
	return conf.rawConfig
}
