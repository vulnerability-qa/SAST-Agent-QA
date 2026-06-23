// CWE-611: XXE via encoding/xml (Go's stdlib processes external entities)
package main
import (
	"encoding/xml"
	"strings"
)
type Doc struct{ Content string }
func parseXML(xmlData string) (Doc, error) {
	var doc Doc
	err := xml.NewDecoder(strings.NewReader(xmlData)).Decode(&doc)
	return doc, err
}
