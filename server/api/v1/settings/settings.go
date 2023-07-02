package settings

import (
	"github.com/gin-gonic/gin"
	"keyu.tech/website/models/common/response"
)

type SettingsApi struct{}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	response.OKWithData(map[string]string{
		"id": "aaaa",
	}, c)
}
