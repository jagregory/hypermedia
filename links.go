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
	linkmap := make(map[string]Hyperlink, len(links))
	for _, link := range links {
		linkmap[link.Rel] = link
	}

	return HyperlinkSet{linkmap}
}

// Set of hyperlinks
type HyperlinkSet struct {
	links map[string]Hyperlink
}

func (l HyperlinkSet) MarshalJSON() ([]byte, error) {
	out := make(map[string]map[string]string)

	if l.links != nil {
		for rel, link := range l.links {
			out[rel] = map[string]string{
				"href": link.Href,
			}
		}
	}

	return json.Marshal(out)
}

func (l *HyperlinkSet) UnmarshalJSON(d []byte) error {
	var out map[string]map[string]string

	if err := json.Unmarshal(d, &out); err != nil {
		return err
	}

	l.links = make(map[string]Hyperlink, len(out))

	for rel, link := range out {
		l.links[rel] = Hyperlink{rel, link["href"]}
	}

	return nil
}

// Find the href of a link by its relationship. Returns
// "" if a link doesn't exist.
func (l HyperlinkSet) Href(rel string) string {
	link, found := l.links[rel]

	if found {
		return link.Href
	}

	return ""
}
