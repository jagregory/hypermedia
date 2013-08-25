package hypermedia

import (
	"encoding/json"
	"fmt"
)

// A hyperlink with a href/URL and a relationship
type Hyperlink struct {
	Rel  string
	Href string
}

// Create a hyperlink to a URL
func Link(rel string, url string) Hyperlink {
	return Hyperlink{rel, host + url}
}

// Create a hyperlink to a URL with a format string
func Linkf(rel string, format string, args ...interface{}) Hyperlink {
	return Link(rel, fmt.Sprintf(format, args...))
}

// Create a rel:self hyperlink to a url
func Self(url string) Hyperlink {
	return Link("self", url)
}

// Create a rel:self hyperlink with a format string
func Selff(format string, args ...interface{}) Hyperlink {
	return Self(fmt.Sprintf(format, args...))
}

// Create a set of hyperlinks
func Links(links ...Hyperlink) HyperlinkSet {
	return HyperlinkSet{links}
}

// Set of hyperlinks
type HyperlinkSet struct {
	links []Hyperlink
}

func (l HyperlinkSet) MarshalJSON() ([]byte, error) {
	out := make(map[string]map[string]string)

	if l.links != nil {
		for _, link := range l.links {
			out[link.Rel] = map[string]string{
				"href": link.Href,
			}
		}
	}

	return json.Marshal(out)
}
