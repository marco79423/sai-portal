package config

type rawConfig struct {
	App      rawAppConfig      `mapstructure:"app"`
	Features rawFeaturesConfig `mapstructure:"features"`
}

type rawAppConfig struct {
	Name     string `mapstructure:"name"`      // 名稱
	LogLevel string `mapstructure:"log_level"` // Log 層級
	Address  string `mapstructure:"address"`

	EnabledFeatures []string `mapstructure:"enabled_features"` // 啟動的功能
}

type rawFeaturesConfig struct {
	RedirectsFeature rawRedirectsFeatureConfig `mapstructure:"redirects"`
}

type rawRedirectsFeatureConfig struct {
	OriginHost string `mapstructure:"origin_host"`
	TargetHost string `mapstructure:"target_host"`
}
