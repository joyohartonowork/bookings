package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	has := form.Has("whatever")

	if has {
		t.Error("form show has field when it  doesnt")
	}
	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("show form doesnt have field when it should")
	}
}

func TestMinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.MinLength("x", 3)
	if form.Valid() {
		t.Error("form show min length for non existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but didnt get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("form show min length 100 when data shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("anotherfield", "abc123")
	form = New(postedValues)

	form.MinLength("anotherfield", 1)
	if !form.Valid() {
		t.Error("show minlegth 1 not meet when it is ")
	}

	isError = form.Errors.Get("anotherfield")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestIsEmail(t *testing.T) {
	//r := httptest.NewRequest("POST", "/whatever", nil)
	postedValues := url.Values{}
	//form := New(r.PostForm)
	form := New(postedValues)
	form.IsEmail("abc")
	if form.Valid() {
		t.Error("form show valid email for non existent field")
	}
	postedValues = url.Values{}
	postedValues.Add("email", "aaa@a.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got valid for invalid email ")
	}
}
