package hypermedia

import (
	"encoding/json"
	"time"
)

// A Resource representation. Contains an "entity" and a
// hyperlink collection.
type Resource struct {
	Links  HyperlinkSet `json:"links"`
	Entity Entity       `json:"entity"`
}

func (r Resource) Modtime() *time.Time {
	return r.Entity.Modtime()
}

// Wrapper for the resource content. Handles marshaling
// the content to JSON.
type Entity struct {
	content interface{}
}

func (e Entity) Modtime() *time.Time {
	if c, ok := e.content.(Modtime); ok {
		return c.Modtime()
	}

	return nil
}

// Optional interface for exposing a Modtime through a resource.
type Modtime interface {
	Modtime() *time.Time
}

func (m Entity) MarshalJSON() ([]byte, error) {
	if marshaler, ok := m.content.(json.Marshaler); ok {
		return marshaler.MarshalJSON()
	} else {
		return json.Marshal(m.content)
	}
}

func NewResource(entity interface{}, links ...Hyperlink) Resource {
	return Resource{
		Links:  Links(links...),
		Entity: Entity{entity},
	}
}
