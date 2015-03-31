// +build !js

package search

import (
	"fmt"
	"os"

	"encoding/xml"
)

type alfredResults []Result

type alfredItem struct {
	Valid        bool   `xml:"valid,attr"`
	Title        string `xml:"title"`
	Subtitle     string `xml:"subtitle,omitempty"`
	Icon         string `xml:"icon,omitempty"`
	UID          string `xml:"uid,attr,omitempty"`
	Autocomplete string `xml:"autocomplete,attr,omitempty"`
	Arg          string `xml:"arg,attr,omitempty"`

	XMLName struct{} `xml:"item"`
}

type alfredResponse struct {
	Items   []alfredItem
	XMLName struct{} `xml:"items"`
}

// Run runs your searcher to get the results to be displayed to the user.
func Run(searcher Searcher) {
	r := &alfredResults{}
	searcher(os.Args[1], r)

	r.print()
}

func (results *alfredResults) Len() int {
	return len(*results)
}

func (results *alfredResults) AddResult(r Result) {
	*results = append(*results, r)
}

func (results *alfredResults) Error(err error) {
	(&alfredResults{
		{
			Title:    "Error",
			Subtitle: err.Error(),
			Icon:     "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns",
			Valid:    false,
		},
	}).print()
}

func (results *alfredResults) print() {
	var ret alfredResponse

	for _, result := range *results {
		if result.Valid {
			ret.Items = append(ret.Items, alfredItem{
				Valid:    result.Valid,
				UID:      result.ID,
				Title:    result.Title,
				Subtitle: result.Subtitle,
				Arg:      result.URL,
			})
		} else {
			ret.Items = append(ret.Items, alfredItem{
				Valid:        result.Valid,
				UID:          result.ID,
				Title:        result.Title,
				Subtitle:     result.Subtitle,
				Autocomplete: result.URL,
			})
		}
	}

	fmt.Println(`<?xml version="1.0" encoding="UTF-8"?>`)
	xml.NewEncoder(os.Stdout).Encode(ret)

	os.Exit(0)
}
