package hypermedia

type root struct {
	Links HyperlinkSet `json:"links"`
}

// Create a root resource representation, which contains
// a series of hyperlinks to other resources with no content.
func Root(links ...Hyperlink) root {
	return root{Links(links...)}
}
