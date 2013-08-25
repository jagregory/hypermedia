package hypermedia

var host string

// Set a global hostname for use with hyperlinks. The
// string supplied will be prepended to any hyperlink
// urls.
func Host(h string) {
	host = h
}
