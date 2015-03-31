package main

import (
	"github.com/AmandaCameron/go-search"
)

func main() {
	cmds := search.Commands{}

	cmds.Add("hello", "Say Hello", "", func(inp string, result search.Results) {
		result.AddResult(search.Result{
			Title: "Hello " + inp,
		})
	})

	cmds.Add("goodbye", "Say goodbye", "", func(inp string, result search.Results) {
		result.AddResult(search.Result{
			Title: "Gooebye " + inp,
		})
	})

	search.Run(func(inp string, result search.Results) {
		if cmds.Process(inp, result) {
			return
		}

		if result.Len() == 0 {
			result.AddResult(search.Result{
				Title: "Unknown command '" + inp + "'",
			})
		}
	})
}
