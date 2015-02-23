package main
package main


import (
import (
	"encoding/xml"
	"encoding/xml"


type XsdtString struct{ string }




	ugo "github.com/go-utils/ugo"
	ugo "github.com/go-utils/ugo"


	svg "github.com/metaleap/go-xsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
	svg "github.com/metaleap/go-xsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
type XsdtString struct{ string })


type SvgDoc struct {
type SvgDoc struct {
	XMLName xml.Name `xml:"svg"`
	XMLName xml.Name `xml:"svg"`
	svg.TsvgType
	svg.TsvgType
}
}


func main() {
func main() {
	var (
	var (
		dirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-svg")
		dirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-svg")
		makeEmptyDoc = func() interface{} { return &SvgDoc{} }
		makeEmptyDoc = func() interface{} { return &SvgDoc{} }
	)
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
}
