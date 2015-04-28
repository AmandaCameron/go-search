package search

// Searcher is what the function passed to Run() should look like.
type Searcher func(string, Results)

// Results is passed to your Searcher function to allow you to return results.
type Results interface {
	// Indicates that the results should be aborted, and that there was an error.
	// Your searcher should return after calling this.
	Error(error)

	// Adds the given Result to this search's result set.
	AddResult(Result)

	// Loads settings for this handler.
	LoadSettings(interface{}) error

	// Returns the total number of results.
	Len() int
}

// Result holds the data for the results you will be returning.
type Result struct {
	// If this is a valid target for a Result.
	Valid bool

	// The stuff to show to the user.
	Title    string
	Subtitle string
	Icon     string

	// The result's target URL
	URL string

	// The result's unique ID
	ID string
}
