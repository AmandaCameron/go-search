// +build chrome,js

package search

import (
	"fmt"

	"github.com/fabioberger/chrome"
)

type chromeResults []Result

// Run runs your searcher to get the results to be displayed to the user.
func Run(searcher Searcher) {
	omnibox := chrome.NewChrome().Omnibox

	// omnibox.OnInputStarted(func(results func() {
	// 	fmt.Println("onInputStarted")

	// 	r := &chromeResults{}

	// 	searcher("", r)

	// 	r.print(results)
	// })

	omnibox.OnInputChanged(func(inp string, results func([]chrome.SuggestResult)) {
		fmt.Println("onInputChanged:", inp)

		r := &chromeResults{}

		searcher(inp, r)

		r.print(results)
	})
}

func (results *chromeResults) Len() int {
	return len(*results)
}

func (results *chromeResults) AddResult(r Result) {
	*results = append(*results, r)
}

func (results *chromeResults) Error(err error) {
	// TODO
}

func (results *chromeResults) print(suggest func([]chrome.SuggestResult)) {
	fmt.Println("chromeResults.print: ", results)

	var ret []chrome.SuggestResult

	for _, res := range *results {
		ret = append(ret, chrome.SuggestResult{
			Content:     res.Url,
			Description: res.Title,
		})
	}

	suggest(ret)
}
