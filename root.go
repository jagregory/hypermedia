package hypermedia

import "time"

type Root struct {
	Links   HyperlinkSet `json:"links"`
	modtime *time.Time
}

func (r Root) Modtime() *time.Time {
	return r.modtime
}

// Create a root resource representation, which contains
// a series of hyperlinks to other resources with no content.
func NewRoot(links ...Hyperlink) Root {
	return Root{Links(links...), nil}
}

// Create a root resource representation, which contains
// a series of hyperlinks to other resources with no content
// and a Last-Modified time.
func NewMRoot(modtime time.Time, links ...Hyperlink) Root {
	return Root{Links(links...), &modtime}
}
