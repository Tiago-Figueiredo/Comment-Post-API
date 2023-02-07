package bootstrap

import (
	"Mindera/service"
)

func Init(port int) error {
	api := service.NewRestApiService()
	return api.ServeContent(port)
}
