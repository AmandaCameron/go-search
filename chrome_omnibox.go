//+build chrome,js

package search

import (
	"errors"
	"strings"

	"github.com/fabioberger/chrome"
	"github.com/gopherjs/gopherjs/js"
)

type chromeResults []Result

// Run runs your searcher to get the results to be displayed to the user.
func Run(searcher Searcher) {
	omnibox := chrome.NewChrome().Omnibox

	search := func(inp string, results func([]chrome.SuggestResult)) {
		r := &chromeResults{}
		searcher("", r)

		r.print(results)
	}

	omnibox.OnInputChanged(func(inp string, results func([]chrome.SuggestResult)) {
		go search(inp, results)
	})

	omnibox.OnInputEntered(func(text, disposition string) {
		if strings.HasPrefix(text, "command:") {
			return
		}

		js.Global.Get("window").Call("open", text)
	})
}

func (results *chromeResults) LoadSettings(settings interface{}) error {
	return errors.New("Not implementes.")
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
	var ret []chrome.SuggestResult

	for _, res := range *results {
		r := chrome.SuggestResult{
			Object: js.MakeWrapper(chrome.SuggestResult{}),
		}

		r.Set("content", res.URL)
		r.Set("description", strings.Replace(res.Title, "&", "&amp;", -1))

		ret = append(ret, r)
	}

	suggest(ret)
}
