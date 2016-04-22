// +build !js

package search

import (
	"fmt"
	"io/ioutil"
	"os"

	"image/png"

	"gopkg.in/yaml.v2"

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

var alfredCache string

// Run runs your searcher to get the results to be displayed to the user.
func Run(searcher Searcher) {
	if os.Getenv("alfred_workflow_uid") == "" {
		println("Expected to be run from Alfred.")

		os.Exit(1)
	}

	alfredCache = os.Getenv("alfred_workflow_cache")

	r := &alfredResults{}

	if _, err := os.Stat(alfredCache); os.IsNotExist(err) {
		if err := os.Mkdir(alfredCache, os.FileMode(0777)); err != nil {
			r.Error(err)
		}
	}

	if len(os.Args) != 2 {
		r.AddResult(Result{
			Title: "No Input.",
		})

		r.print()
	}

	searcher(os.Args[1], r)
	r.print()
}

func (results *alfredResults) LoadSettings(settings interface{}) error {
	file, err := os.Open("settings.yaml")
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, settings)
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

			Valid: false,
		},
	}).print()
}

func (results *alfredResults) print() {
	var ret alfredResponse

	for _, result := range *results {
		iconPath := ""
		icon := result.Icon

		if icon != nil {
			hash := hashIcon(icon)

			iconPath = alfredCache + "/icon-" + hash + ".png"

			if _, err := os.Stat(iconPath); os.IsNotExist(err) {
				file, _ := os.Create(iconPath)
				defer file.Close()

				if err := png.Encode(file, icon); err != nil {
					iconPath = "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns"
				}
			}
		}

		if result.Valid {
			ret.Items = append(ret.Items, alfredItem{
				Valid:    result.Valid,
				UID:      result.ID,
				Title:    result.Title,
				Subtitle: result.Subtitle,
				Arg:      result.URL,

				Icon: iconPath,
			})
		} else {
			ret.Items = append(ret.Items, alfredItem{
				Valid:        result.Valid,
				UID:          result.ID,
				Title:        result.Title,
				Subtitle:     result.Subtitle,
				Autocomplete: result.URL,

				Icon: iconPath,
			})
		}
	}

	fmt.Println(`<?xml version="1.0" encoding="UTF-8"?>`)
	xml.NewEncoder(os.Stdout).Encode(ret)

	os.Exit(0)
}
