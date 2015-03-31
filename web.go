// +build js,angularjs

package search

import (
	"github.com/gopherjs/gopherjs/js"

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
		scope.Set("selected", 0)

		scope.Set("keyDown", func(keyCode int) {
			println("got keyDown:", keyCode)

			if keyCode == 38 {
				scope.Set("selected", scope.Get("selected").Int()-1)
			} else if keyCode == 40 {
				scope.Set("selected", scope.Get("selected").Int()+1)
			} else if keyCode == 13 {
				url := scope.Get("results").Get(scope.Get("selected").String()).Get("URL")
				js.Global.Get("window").Call("open", url)
			}

			if scope.Get("selected").Int() < 0 {
				scope.Set("selected", scope.Get("results").Get("length"))
			} else if scope.Get("selected").Int() > scope.Get("results").Get("length").Int() {
				scope.Set("selected", 0)
			}
		})

		scope.Call("$watch", "input",
			func(inp string, oldValue string) {
				go func() {
					r := &angularResults{Scope: scope}
					searcher(inp, r)

					scope.Apply(r.print)
				}()
			})
	})
}

func (results *angularResults) Len() int {
	return len(results.results)
}

func (results *angularResults) AddResult(r Result) {
	results.results = append(results.results, r)
}

func (results *angularResults) Error(err error) {
	results.results = nil

	results.AddResult(Result{
		Title:    "Error Searching",
		Subtitle: err.Error(),
		Valid:    false,
	})
}

func (results *angularResults) print() {
	results.Scope.Set("selected", 0)

	results.Scope.Set("results", results.results)
}
