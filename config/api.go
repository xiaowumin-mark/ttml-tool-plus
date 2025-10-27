package config

import userdata "ttml-tool-plus/user-data"

func (c *ConfigApiService) GetConfig() map[string]any {
	return Config
}

func (c *ConfigApiService) GetConfigValue(key string) any {
	return Config[key]
}

func (c *ConfigApiService) SaveConfig() {
	Save(&Config)
}

func (c *ConfigApiService) ChangeConfig(key string, value any) map[string]any {
	Config[key] = value
	return Config
}
func (c *ConfigApiService) ChangeConfigAndSave(key string, value any) map[string]any {
	Config[key] = value
	Save(&Config)
	return Config
}

func (c *ConfigApiService) UserData() {
	GetUsetData()
}

// 退出登录
func (c *ConfigApiService) Logout() {
	userdata.DelToken()
}
func (c *ConfigApiService) IsLogin() bool {
	t, err := userdata.GetToken()
	return t != "" && err == nil
}
