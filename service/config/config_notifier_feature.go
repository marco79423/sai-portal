package config

type IRedirectsFeatureConfig interface {
	GetOriginHost() string
	GetTargetHost() string
}

func newRedirectsFeatureConfig(rawConfig *rawConfig) IRedirectsFeatureConfig {
	return &redirectFeatureConfig{
		rawConfig: rawConfig,
	}
}

type redirectFeatureConfig struct {
	rawConfig *rawConfig
}

// GetOriginHost 取得來源 Host
func (conf *redirectFeatureConfig) GetOriginHost() string {
	return conf.rawConfig.Features.RedirectsFeature.OriginHost
}

// GetTargetHost 取得目標 Host
func (conf *redirectFeatureConfig) GetTargetHost() string {
	return conf.rawConfig.Features.RedirectsFeature.TargetHost
}
