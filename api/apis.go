package api

import (
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func Get(url string) {
	Router.GET(url)
}

func Post(url string) {
	Router.POST(url)
}

func Delete(url string) {
	Router.DELETE(url)
}
