package main
package main


import (
import (
	"encoding/xml"
	"encoding/xml"


type XsdtString struct{ string }




	"github.com/go-utils/ugo"
	"github.com/go-utils/ugo"


	kml "github.com/metaleap/go-xsd-pkg/schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go"
	kml "github.com/metaleap/go-xsd-pkg/schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go"
type XsdtString struct{ string })


type KmlDoc struct {
type KmlDoc struct {
	XMLName xml.Name `xml:"http://www.opengis.net/kml/2.2 kml"`
	XMLName xml.Name `xml:"http://www.opengis.net/kml/2.2 kml"`
	kml.TKmlType
	kml.TKmlType
}
}


func main() {
func main() {
	var (
	var (
		dirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-kml")
		dirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-kml")
		makeEmptyDoc = func() interface{} { return &KmlDoc{} }
		makeEmptyDoc = func() interface{} { return &KmlDoc{} }
	)
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
}
