package main

import (
	ace "go.bytebuilders.dev/installer/apis/installer/v1alpha1"
	"gopkg.in/macaron.v1"

	"github.com/go-macaron/binding"
)

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Post("/generate", binding.Json(ace.AceOptionsSpec{}), GenerateInstaller)
	m.Run()
}

func GenerateInstaller(ctx *macaron.Context, opts ace.AceOptionsSpec) {

}
