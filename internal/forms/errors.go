package forms

type errors map[string][]string

//ADD adds an error message for a given form fields
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//GET return the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
