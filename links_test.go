package hypermedia

import "testing"

var hrefTests = []struct {
	name     string
	expected string
	url      string
}{
	{"normal", "/example", "/example"},
	{"parameterised", "/example", "/example{?q}"},
}

func TestHref(t *testing.T) {
	for _, test := range hrefTests {
		links := Links(Link(test.name, test.url))
		href, _ := links.Href(test.name)
		if href != test.expected {
			t.Errorf("%s: Expected href to be '%s', got '%s'", test.name, test.expected, href)
		}
	}
}

var hrefParamsTests = []struct {
	name     string
	expected string
	url      string
	params   Params
}{
	{"nil parameters", "/example", "/example{?q}", nil},
	{"empty parameters", "/example", "/example{?q}", Params{}},
	{"mismatched parameters", "/example", "/example{?q}", Params{"c": "test"}},
	{"single parameter", "/example?q=test", "/example{?q}", Params{"q": "test"}},
	{"multiple parameters", "/example?q=test&page=1", "/example{?q,page}", Params{"q": "test", "page": 1}},
}

func TestHrefParams(t *testing.T) {
	for _, test := range hrefParamsTests {
		links := Links(Link(test.name, test.url))
		href, _ := links.HrefParams(test.name, test.params)
		if href != test.expected {
			t.Errorf("%s: Expected href to be '%s', got '%s'", test.name, test.expected, href)
		}
	}
}
