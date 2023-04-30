package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"

)

type Form struct {
	url.Values
	Errors errors
}

func New (data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has (field string)bool{
	x := f.Get(field)
	if x != "" {
		return true
	}

	f.Errors.Add(field,fmt.Sprintf("the field '%s' has no data inputed",field ))
	return false
}
     

func (f *Form) Required(fields ...string){
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field ,fmt.Sprintf("the field '%s'is required", field))
		}
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors)==0
 }

 func (f *Form)MinLength(field string , length int) bool {
	x:=f.Get(field)
	if len(x) < length {
		f.Errors.Add(field,fmt.Sprintf("this field mus be at least %d character long",length))
		return false
	}
	return true
 }

 // IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
