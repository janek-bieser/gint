package gint_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/janek-bieser/gint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTMLRender", func() {

	var (
		htmlRender *gint.HTMLRender
	)

	BeforeEach(func() {
		htmlRender = gint.NewHTMLRender()
	})

	Context("Layout file does not exist", func() {

		BeforeEach(func() {
			htmlRender.TemplateDir = "test_templates/no_layout"
		})

		It("should panic when loading a template", func() {
			f := func() { htmlRender.Instance("index", nil) }
			Expect(f).To(Panic())
		})

	})

	Context("Template file does not exist", func() {

		BeforeEach(func() {
			htmlRender.TemplateDir = "test_templates/only_layout"
		})

		It("should panic when loading non existing template", func() {
			f := func() { htmlRender.Instance("index", nil) }
			Expect(f).To(Panic())
		})

	})

	Context("Valid template files exist", func() {

		BeforeEach(func() {
			htmlRender.TemplateDir = "test_templates/valid"
		})

		It("Should load the template without partials", func() {
			instance := htmlRender.Instance("index", map[string]string{"title": "Hello, World"})
			res := httptest.NewRecorder()
			instance.Render(res)

			Expect(res.Code).To(Equal(http.StatusOK))

			expectedBody := "<h1>Hello, World</h1>\n<div>Index</div>"
			Expect(res.Body.String()).To(Equal(expectedBody))
		})

		It("Should load the template with partials", func() {
			instance := htmlRender.Instance("partials", map[string]string{"title": "Test"})
			res := httptest.NewRecorder()
			instance.Render(res)

			Expect(res.Code).To(Equal(http.StatusOK))

			expectedBody := "<h1>Test</h1>\n<div>Hello World</div>"
			Expect(res.Body.String()).To(Equal(expectedBody))
		})

	})

})
