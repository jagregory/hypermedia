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
	Content interface{}
}

func (e Entity) Modtime() *time.Time {
	if c, ok := e.Content.(Modtime); ok {
		return c.Modtime()
	}

	return nil
}

// Optional interface for exposing a Modtime through a resource.
type Modtime interface {
	Modtime() *time.Time
}

func (m Entity) MarshalJSON() ([]byte, error) {
	if marshaler, ok := m.Content.(json.Marshaler); ok {
		return marshaler.MarshalJSON()
	} else {
		return json.Marshal(m.Content)
	}
}

func (m Entity) UnmarshalJSON(d []byte) error {
	if marshaler, ok := m.Content.(json.Unmarshaler); ok {
		return marshaler.UnmarshalJSON(d)
	} else {
		return json.Unmarshal(d, m.Content)
	}
}

func NewResource(entity interface{}, links ...Hyperlink) Resource {
	return Resource{
		Links:  Links(links...),
		Entity: Entity{entity},
	}
}
