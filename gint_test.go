package gint_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janek-bieser/gint"
)

func TestNonExistingLayoutFile(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()

	htmlRender := newHTMLRenderForTesting("no_layout")
	htmlRender.Instance("index", nil)
}

func TestInvalidTemplate(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()

	htmlRender := newHTMLRenderForTesting("only_layout")
	// expect this to panic becuase the template does not exist
	htmlRender.Instance("doesnotexist", nil)
}

func TestValidTemplateLoading(t *testing.T) {
	htmlRender := newHTMLRenderForTesting("valid")
	instance := htmlRender.Instance("index", map[string]string{
		"title": "Hello, World",
	})
	res := httptest.NewRecorder()
	instance.Render(res)

	expectStatusCode(t, http.StatusOK, res.Code)

	expectedResult := []byte("<h1>Hello, World</h1>\n<div>Index</div>")
	resBytes := res.Body.Bytes()

	if !bytes.Equal(expectedResult, resBytes) {
		t.Errorf("Expected body to be: '%s' but got: '%s'",
			string(expectedResult), string(resBytes))
	}
}

func TestPartialsLoading(t *testing.T) {
	htmlRender := newHTMLRenderForTesting("partials")
	instance := htmlRender.Instance("index", map[string]string{"title": "Test"})
	res := httptest.NewRecorder()
	instance.Render(res)

	expectStatusCode(t, http.StatusOK, res.Code)

	expectedResult := []byte("<h1>Test</h1>\n<div>Hello World</div>")
	resBytes := res.Body.Bytes()

	if !bytes.Equal(expectedResult, resBytes) {
		t.Errorf("Expected body to be: '%s' but got: '%s'",
			string(expectedResult), string(resBytes))
	}
}

func newHTMLRenderForTesting(path string) *gint.HTMLRender {
	r := gint.NewHTMLRender()
	r.TemplateDir = "test_templates/" + path
	return r
}

func expectStatusCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected status to be %d, but is %d.", expected, actual)
	}
}
