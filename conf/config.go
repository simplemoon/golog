package conf

type Config struct {
	Path             string   `json:"path"`               // log配置加载的位置
	Dir              string   `json:"dir"`                // 配置文件的路径
	RemoteProviders  []string `json:"remote_providers"`   // 配置的提供商
	EndPoints        string   `json:"end_points"`         // 配置的端点
	ReloadWhenChange bool     `json:"reload_when_change"` // 配置改变是是否重置
}
