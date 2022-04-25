package forms

type errors map[string][]string

// Add adds an error message for a given form field
func (e errors) Add(field_name, message string) {
	e[field_name] = append(e[field_name], message)

}

// Get return the first error message
func (e errors) Get(field_name string) string {
	es := e[field_name]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
