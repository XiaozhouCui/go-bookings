package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initialises a form struct
func New(data url.Values) *Form {
	// return a pointer to Form
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	// check request
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
