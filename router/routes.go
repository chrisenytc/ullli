package router

import (
	"gopkg.in/kataras/iris.v6"

	"github.com/chrisenytc/ullli/endpoints"
)

func LoadRoutes(router *iris.Framework) {
	router.Get("/", endpoints.GetIndex)
	router.Get("/status", endpoints.GetStatus)
	router.Post("/shorten", endpoints.PostShorten)
	router.Get("/{shortCode}", endpoints.GetShortCode)
	router.Get("/{shortCode}/stats", endpoints.GetShortCodeStats)
}
