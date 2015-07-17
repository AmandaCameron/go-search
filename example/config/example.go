package main

import (
	"github.com/AmandaCameron/go-search"
)

func main() {
	search.Run(func(inp string, result search.Results) {
		var settings struct {
			Cake bool
		}
		
		if err := result.LoadSettings(&settings); err != nil {
			result.Error(err)
			return
		}

		result.AddResult(search.Result{
			Title: "Hello World",
			Subtitle: "Hello Bacon.",
		})

		if settings.Cake {
			result.AddResult(search.Result{
				Title: "Cake?",
				Subtitle: "Cake!",
			})
		}
	})
}
