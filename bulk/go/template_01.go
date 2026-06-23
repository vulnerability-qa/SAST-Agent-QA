// CWE-94: Template injection via html/template with user-controlled template string
package main
import (
	"html/template"
	"os"
)
func renderTemplate(tmplStr string, data interface{}) error {
	t, err := template.New("").Parse(tmplStr) // user-controlled template
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, data)
}
