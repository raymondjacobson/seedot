package main

import (
  "github.com/go-martini/martini"
  "github.com/codegangsta/martini-contrib/render"
  "math/rand"
  "time"
  "./lib"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Directory: "template",
    Layout: "layout",
    Extensions: []string{".html"},
  }))

  m.Get("/", func(r render.Render) {
    random_string := randomString(10, 30)
    r.HTML(200, "main", "/d/" + random_string)
  })

  m.Get("/d/:query", func(params martini.Params, r render.Render) {
    r.HTML(200, "d", params["query"])
  })

  m.Run()
}
