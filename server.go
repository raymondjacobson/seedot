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

  m.Get("/", func(r render.Render) {
    random_string := RandomString(10, 30)
    r.HTML(200, "main", "/d/" + random_string)
  })

  m.Get("/d/**", func(params martini.Params, r render.Render) {
    r.HTML(200, "d", params["_1"])
  })

  m.Run()
}
