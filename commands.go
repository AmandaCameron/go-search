package search

import (
	"strings"
)

type command struct {
	Result
	searcher Searcher
}

type Commands map[string]command

func (commands Commands) Add(name, title, subtitle string, searcher Searcher) {
	if subtitle == "" {
		subtitle = "Command"
	}

	commands[name] = command{
		Result: Result{
			Title:    title,
			Subtitle: subtitle,
			Valid:    false,
			URL:      name,
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
