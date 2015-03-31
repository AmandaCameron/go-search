package main

import (
	"amanda.camnet/search"
)

func main() {
	search.Run(func(inp string, result search.Results) {
		result.AddResult(search.Result{
			Title:    "Hello World",
			Subtitle: "Bacon is nom.",
		})

		result.AddResult(search.Result{
			Title:    "You inputted: " + inp,
			Subtitle: "Yay!",
		})
	})
}
