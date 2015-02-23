package main
package main


import (
import (
	"encoding/xml"
	"encoding/xml"


type XsdtString struct{ string }




	"github.com/go-utils/ugo"
	"github.com/go-utils/ugo"


	atom "github.com/metaleap/go-xsd-pkg/kbcafe.com/rss/atom.xsd.xml_go"
	atom "github.com/metaleap/go-xsd-pkg/kbcafe.com/rss/atom.xsd.xml_go"
type XsdtString struct{ string })


type AtomEntryDoc struct {
type AtomEntryDoc struct {
	XMLName xml.Name `xml:"entry"`
	XMLName xml.Name `xml:"entry"`
	atom.TentryType
	atom.TentryType
}
}


type AtomFeedDoc struct {
type AtomFeedDoc struct {
	XMLName xml.Name `xml:"feed"`
	XMLName xml.Name `xml:"feed"`
	atom.TfeedType
	atom.TfeedType
}
}


func main() {
func main() {
	var (
	var (
		entryDirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-atom", "entry")
		entryDirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-atom", "entry")
		entryMakeEmptyDoc = func() interface{} { return &AtomEntryDoc{} }
		entryMakeEmptyDoc = func() interface{} { return &AtomEntryDoc{} }
		feedDirBasePath   = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-atom", "feed")
		feedDirBasePath   = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-atom", "feed")
		feedMakeEmptyDoc  = func() interface{} { return &AtomFeedDoc{} }
		feedMakeEmptyDoc  = func() interface{} { return &AtomFeedDoc{} }
	)
	)
	tests.TestViaRemarshal(entryDirBasePath, entryMakeEmptyDoc)
	tests.TestViaRemarshal(entryDirBasePath, entryMakeEmptyDoc)
	tests.TestViaRemarshal(feedDirBasePath, feedMakeEmptyDoc)
	tests.TestViaRemarshal(feedDirBasePath, feedMakeEmptyDoc)
}
}
