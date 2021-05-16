package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//Form create custom form struct, embeds url.values object
type Form struct {
	url.Values
	Errors errors
}

//New init a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Has  check field in post not empty
func (f *Form) Has(field string) bool {
	//x := r.Form.Get(field)
	x := f.Get(field)
	if x == "" {
		//f.Errors.Add(field, "This Field cant blank")
		return false
	}
	return true
}

//Required check field(s) need filled
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cant blank")
		}
	}
}

//check minimum length
func (f *Form) MinLength(field string, length int) bool {
	// x := r.Form.Get(field)
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least #%d chars long", length))
		return false
	}
	return true
}

//Valid check data, true if no errors otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//check email format
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "not valid email")
	}

}
