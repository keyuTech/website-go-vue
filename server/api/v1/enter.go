package v1

import "keyu.tech/website/api/v1/settings"

type ApiGroup struct {
	SettingsApi settings.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
