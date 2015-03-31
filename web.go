// +build js,angularjs

package search

import (
	"github.com/neelance/go-angularjs"
)

type angularResults struct {
	*angularjs.Scope
	results []Result
}

// Run provides an AngularJS Module named "search-app" with a controller "Search"
func Run(searcher Searcher) {
	module := angularjs.NewModule("search-app", nil, func() {})
	module.NewController("Search", func(scope *angularjs.Scope) {
		scope.Set("results", []Result{})
		scope.Set("input", "")

		scope.Call("$watch", "input",
			func(inp string, oldValue string) {
				go func() {
					r := &angularResults{Scope: scope}
					searcher(inp, r)

					r.print()
				}()
			})
	})
}

func (results *angularResults) AddResult(r Result) {
	results.results = append(results.results, r)
}

func (results *angularResults) Error(err error) {
	results.results = []Result{
		{
			Title:    "Error Searching",
			Subtitle: err.Error(),
			Valid:    false,
		},
	}
}

func (results *angularResults) print() {
	results.Scope.Set("results", results.results)
}
