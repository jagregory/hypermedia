package hypermedia

type Root struct {
	Links HyperlinkSet `json:"links"`
}

// Create a root resource representation, which contains
// a series of hyperlinks to other resources with no content.
func NewRoot(links ...Hyperlink) Root {
	return Root{Links(links...)}
}
