// +build js,angularjs

package search

import (
	"encoding/json"

	"github.com/gopherjs/gopherjs/js"

	"github.com/AmandaCameron/go-angularjs"
	"github.com/AmandaCameron/go-angularjs/directive"
)

type angularResults struct {
	*js.Object `ajs-service:"Search"`

	config string `js:"config"`
	selected int `js:"selected"`

	results []Result
}

func (scope *angularResults) New() {
	scope.config = ""
	scope.selected = 0
	scope.results = nil
}

// Run provides an AngularJS Module named "search-app" with a controller "Search"
func Run(searcher Searcher) {
	defer angularjs.LogErrors()
	module := angularjs.NewModule("search-app", nil)

	module.Factory("Search", func() interface{} {
		return &angularResults{
			Object: js.Global.Get("Object").New(),
		}
	})

	err := module.Directive("searchConfig", func(search *angularResults) *js.Object {
		defer angularjs.LogErrors()
		
		return directive.New(
			directive.Link(func(_ angularjs.Scope, element angularjs.JQueryElement) {
				search.config = element.Text();
				
				element.SetText("");
			}))
	})

	if err != nil {
		print("Error registering directive: ", err.Error())
		
		return
	}

	err = module.Directive("searchView", func () *js.Object {
		return directive.New(
			directive.Transclude(),
			directive.Template(searchViewTempl),
			
			directive.Controller(func(search *angularResults, scope *angularjs.Scope) {
				defer angularjs.LogErrors()
				
				scope.Set("keyDown", func(keyCode int) {
					defer angularjs.LogErrors()

					if keyCode == 38 {
						search.selected = search.selected - 1
					} else if keyCode == 40 {
						search.selected = search.selected + 1
					} else if keyCode == 13 {
						url := search.results[search.selected].URL
						
						js.Global.Get("window").Call("open", url)
					}

					if search.selected < 0 {
						search.selected = len(search.results)
					} else if search.selected > len(search.results) {
						search.selected = 0
					}

					search.updateSelected()
					scope.Set("results", search.results)
				})
				
				scope.Call("$watch", "input",
					func(inp, oldValue string) {
						go func() {
							search.results = nil;
							searcher(inp, search)

							search.updateSelected()
							scope.Apply(func() {
								scope.Set("results", search.results)
							})
						}()
					})
			}))
	})
	
	if err != nil {
		print("Error registering directive: ", err.Error())
		
		return
	}
	
	print("Done.")
}

var searchViewTempl = `
<div class="search-input-box">
  <input class="search-input" ng-model="input" ng-keydown="keyDown($event.keyCode)">
</div>

<div ng-if="results.length == 0">
  <div class="no-results">
    No results found for <i>{{ input }}</i>
  </div>
</div>

<div ng-repeat="result in results">
  <div class="result" ng-selected="result.Selected">
    <div class="title">{{result.Title}}</div>
    <div class="subtitle">{{result.Subtitle}}</div>
  </div>
</div>`

func (results *angularResults) LoadSettings(settings interface{}) error {
	println("Config: ", results.config)
	
	return json.Unmarshal([]byte(results.config), settings)
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

func (results *angularResults) updateSelected() {
	for i := 0; i < results.Len(); i++ {
		result := results.results[i]
		
		result.Selected = (i == results.selected)
		
		results.results[i] = result
	}
}
