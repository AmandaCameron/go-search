package main

import (
	"github.com/AmandaCameron/go-search"
)

//go:generate embed-result-image test-image.png ResultImage

func main() {
	search.Run(func(inp string, ctx search.Results) {
		ctx.AddResult(search.Result{
			Title: "Test Search Result",
			Subtitle: "Search Results are fun.",
			
			Icon: ResultImage,
		})
	})
}

