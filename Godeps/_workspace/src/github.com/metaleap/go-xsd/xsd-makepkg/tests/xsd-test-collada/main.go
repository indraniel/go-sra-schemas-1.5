package main
package main


import (
import (
	"encoding/xml"
	"encoding/xml"
	"log"
	"log"
type XsdtString struct{ string }


	"github.com/metaleap/go-xsd/xsd-makepkg/tests"
	"github.com/metaleap/go-xsd/xsd-makepkg/tests"


	"github.com/go-utils/ugo"
	"github.com/go-utils/ugo"


type XsdtString struct{ string }	collada14 "github.com/metaleap/go-xsd-pkg/khronos.org/files/collada_schema_1_4_go"
	collada15 "github.com/metaleap/go-xsd-pkg/khronos.org/files/collada_schema_1_5_go"
	collada15 "github.com/metaleap/go-xsd-pkg/khronos.org/files/collada_schema_1_5_go"
)
)


type Col14Doc struct {
type Col14Doc struct {
	XMLName xml.Name `xml:"COLLADA"`
	XMLName xml.Name `xml:"COLLADA"`
	collada14.TxsdCollada
	collada14.TxsdCollada
}
}


type Col15Doc struct {
type Col15Doc struct {
	XMLName xml.Name `xml:"COLLADA"`
	XMLName xml.Name `xml:"COLLADA"`
	collada15.TxsdCollada
	collada15.TxsdCollada
}
}


func main() {
func main() {
	var (
	var (
		col14DirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-collada", "1.4.1")
		col14DirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-collada", "1.4.1")
		col14MakeEmptyDoc = func() interface{} { return &Col14Doc{} }
		col14MakeEmptyDoc = func() interface{} { return &Col14Doc{} }
		col15DirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-collada", "1.5")
		col15DirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-collada", "1.5")
		col15MakeEmptyDoc = func() interface{} { return &Col15Doc{} }
		col15MakeEmptyDoc = func() interface{} { return &Col15Doc{} }
	)
	)
	if false {
	if false {
		tests.OnDocLoaded = func(doc interface{}) {
		tests.OnDocLoaded = func(doc interface{}) {
			if c14, ok := doc.(*Col14Doc); ok {
			if c14, ok := doc.(*Col14Doc); ok {
				log.Print("ISC14")
				log.Print("ISC14")
				for _, camLib := range c14.CamerasLibraries {
				for _, camLib := range c14.CamerasLibraries {
					log.Print("CAMLIB")
					log.Print("CAMLIB")
					for _, cam := range camLib.Cameras {
					for _, cam := range camLib.Cameras {
						log.Printf("CAM aspect: %#v\n", cam.Optics.TechniqueCommon.Perspective.AspectRatio)
						log.Printf("CAM aspect: %#v\n", cam.Optics.TechniqueCommon.Perspective.AspectRatio)
					}
					}
				}
				}
			}
			}
		}
		}
	}
	}
	tests.TestViaRemarshal(col14DirBasePath, col14MakeEmptyDoc)
	tests.TestViaRemarshal(col14DirBasePath, col14MakeEmptyDoc)
	tests.TestViaRemarshal(col15DirBasePath, col15MakeEmptyDoc)
	tests.TestViaRemarshal(col15DirBasePath, col15MakeEmptyDoc)
}
}
