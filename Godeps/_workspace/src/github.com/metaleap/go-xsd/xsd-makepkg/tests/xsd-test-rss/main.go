package main
package main


import (
import (
	"encoding/xml"
	"encoding/xml"


type XsdtString struct{ string }




	ugo "github.com/go-utils/ugo"
	ugo "github.com/go-utils/ugo"


	rss "github.com/metaleap/go-xsd-pkg/thearchitect.co.uk/schemas/rss-2_0.xsd_go"
	rss "github.com/metaleap/go-xsd-pkg/thearchitect.co.uk/schemas/rss-2_0.xsd_go"
type XsdtString struct{ string })


type RssDoc struct {
type RssDoc struct {
	XMLName xml.Name `xml:"rss"`
	XMLName xml.Name `xml:"rss"`
	rss.TxsdRss
	rss.TxsdRss
}
}


func main() {
func main() {
	var (
	var (
		dirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-rss")
		dirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-rss")
		makeEmptyDoc = func() interface{} { return &RssDoc{} }
		makeEmptyDoc = func() interface{} { return &RssDoc{} }
	)
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
}
