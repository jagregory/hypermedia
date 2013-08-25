package hypermedia

import "time"

// Collection of Resources, with an hyperlink collection.
type Collection struct {
	Links      HyperlinkSet `json:"links"`
	Collection []Resource   `json:"collection"`
	modtime    *time.Time
}

func (c Collection) Modtime() *time.Time {
	return c.modtime
}

// Create a new collection of Resources, with a set of links.
func NewCollection(resources []Resource, links ...Hyperlink) Collection {
	var modtime *time.Time = nil

	for _, res := range resources {
		imod := res.Modtime()

		if imod != nil && (modtime == nil || modtime.Before(*imod)) {
			modtime = imod
		}
	}

	return Collection{
		Links:      Links(links...),
		Collection: resources,
		modtime:    modtime,
	}
}
