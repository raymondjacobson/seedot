package main

import (
  "github.com/go-martini/martini"
  "github.com/codegangsta/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Directory: "template",
    Layout: "layout",
    Extensions: []string{".html"},
  }))

  m.Get("/", func() string {
    return "Welcome to seedot"
  })

  m.Get("/d/:query", func(params martini.Params) string {
    return params["query"]
  })

  m.Run()
}
