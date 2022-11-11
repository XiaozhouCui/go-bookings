package forms

// create type to hold errors
type errors map[string][]string

// Add adds an error message for a given form field
func (e errors) Add(field, message string) {
	// has a receiver e to tie to the errors type
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	// create an error string
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
