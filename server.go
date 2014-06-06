package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()

  m.Get("/", func() string {
    return "Welcome to seedot"
  })

  m.Get("/d/:query", func(params martini.Params) string {
    return params["query"]
  })

  m.Run()
}
