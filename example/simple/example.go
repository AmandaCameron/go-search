package main

import (
	"github.com/AmandaCameron/go-search"
)

func main() {
	search.Run(func(inp string, result search.Results) {
		result.AddResult(search.Result{
			Title:    "Hello World",
			Subtitle: "Bacon is nom.",
			URL:      "https://github.com/bacon",
		})

		result.AddResult(search.Result{
			Title:    "You inputted: " + inp,
			Subtitle: "Yay!",
			URL:      "https://google.com/",
		})
	})
}
