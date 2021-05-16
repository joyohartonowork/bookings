package render

import (
	"net/http"
	"testing"

	"github.com/joyohartonowork/bookings/internal/models"
)

func testAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("failed 123 not in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateChache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error write template to browser")
	}

	err = RenderTemplate(&ww, r, "non-exist.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("render template that not exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestNewTemplate(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
