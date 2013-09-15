package hypermedia

// Collection of Resources, with an hyperlink collection.
type Collection struct {
	Links      HyperlinkSet `json:"links"`
	Collection interface{}  `json:"collection"`
}

// Create a new collection of Resources, with a set of links.
func NewCollection(collection interface{}, links ...Hyperlink) Collection {
	return Collection{
		Links:      Links(links...),
		Collection: collection,
	}
}
