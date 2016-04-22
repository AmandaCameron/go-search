package search

import (
	"strings"
)

type command struct {
	Result
	searcher Searcher
}

type Commands map[string]command

func (commands Commands) Add(name, desc string, searcher Searcher) {
	commands[name] = command{
		Result: Result{
			Title:    name,
			Subtitle: desc,
			
			URL:      name,

			Valid:    false,
		},
		searcher: searcher,
	}
}

func (commands Commands) Process(inp string, results Results) bool {
	if commands == nil {
		return false
	}

	for cmd, res := range commands {
		if strings.HasPrefix(cmd, inp) {
			results.AddResult(res.Result)
		}

		if strings.HasPrefix(inp, cmd+" ") {
			res.searcher(strings.TrimPrefix(inp, cmd+" "), results)

			return true
		}
	}

	return false
}
