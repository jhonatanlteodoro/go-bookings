package forms

import (
	"fmt"
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors     errors
	FieldNames []string
	request    *http.Request
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	for _, field := range f.FieldNames {
		f.Has(field)
	}
	return len(f.Errors) == 0
}

func New(data url.Values, field_names []string, request *http.Request) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
		field_names,
		request,
	}
}

func (f *Form) Has(field_name string) bool {
	x := f.request.Form.Get(field_name)
	fmt.Println(x)
	if x == "" {
		f.Errors.Add(field_name, "This field cannot be blank")
		return false
	}
	return true

}
