package tests
package tests


import (
import (
	"encoding/xml"
	"encoding/xml"
	"fmt"
	"fmt"
type XsdtString struct{ string }


	"path/filepath"
	"path/filepath"
	"strings"
	"strings"


	"github.com/go-utils/ufs"
	"github.com/go-utils/ufs"
type XsdtString struct{ string }
	xmlx "github.com/go-forks/go-pkg-xmlx"
	xmlx "github.com/go-forks/go-pkg-xmlx"
)
)


var (
var (
	OnDocLoaded func(interface{})
	OnDocLoaded func(interface{})
)
)


func verifyDocs(origData, faksData []byte) (errs []error) {
func verifyDocs(origData, faksData []byte) (errs []error) {
	orig, faks := xmlx.New(), xmlx.New()
	orig, faks := xmlx.New(), xmlx.New()
	err := orig.LoadBytes(origData, nil)
	err := orig.LoadBytes(origData, nil)
	if err == nil {
	if err == nil {
		if err = faks.LoadBytes(faksData, nil); err == nil {
		if err = faks.LoadBytes(faksData, nil); err == nil {
			errs = verifyNode(orig.Root, faks.Root)
			errs = verifyNode(orig.Root, faks.Root)
		}
		}
	}
	}
	return
	return
}
}


func verifyNode(orig, faks *xmlx.Node) (errs []error) {
func verifyNode(orig, faks *xmlx.Node) (errs []error) {
	type both struct {
	type both struct {
		origNodes, faksNodes []*xmlx.Node
		origNodes, faksNodes []*xmlx.Node
	}
	}
	var (
	var (
		curBoth *both
		curBoth *both
		cn      *xmlx.Node
		cn      *xmlx.Node
		tmp     string
		tmp     string
		i       int
		i       int
		subErrs []error
		subErrs []error
	)
	)
	attVal := func(xn *xmlx.Node, att *xmlx.Attr) (v string) {
	attVal := func(xn *xmlx.Node, att *xmlx.Attr) (v string) {
		if v = xn.As("", att.Name.Local); len(v) == 0 {
		if v = xn.As("", att.Name.Local); len(v) == 0 {
			v = xn.As(att.Name.Space, att.Name.Local)
			v = xn.As(att.Name.Space, att.Name.Local)
		}
		}
		return
		return
	}
	}
	cleanNodes := func(xns ...*xmlx.Node) {
	cleanNodes := func(xns ...*xmlx.Node) {
		for _, xn := range xns {
		for _, xn := range xns {
			for _, cn = range xn.Children {
			for _, cn = range xn.Children {
				if cn.Type != xmlx.NT_ELEMENT {
				if cn.Type != xmlx.NT_ELEMENT {
					xn.RemoveChild(cn)
					xn.RemoveChild(cn)
				}
				}
			}
			}
		}
		}
	}
	}
	cleanNodes(orig, faks)
	cleanNodes(orig, faks)
	for _, a := range orig.Attributes {
	for _, a := range orig.Attributes {
		if tmp = attVal(faks, a); tmp != a.Value {
		if tmp = attVal(faks, a); tmp != a.Value {
			errs = append(errs, fmt.Errorf("Attribute '%s:%s' of <%s> element: different values (orig='%s' faks='%s')", a.Name.Space, a.Name.Local, orig.Name.Local, a.Value, tmp))
			errs = append(errs, fmt.Errorf("Attribute '%s:%s' of <%s> element: different values (orig='%s' faks='%s')", a.Name.Space, a.Name.Local, orig.Name.Local, a.Value, tmp))
		}
		}
	}
	}
	if len(orig.Children) > len(faks.Children) {
	if len(orig.Children) > len(faks.Children) {
		errs = append(errs, fmt.Errorf("Orig <%s> element has %v children, but faks has %v.", orig.Name.Local, len(orig.Children), len(faks.Children)))
		errs = append(errs, fmt.Errorf("Orig <%s> element has %v children, but faks has %v.", orig.Name.Local, len(orig.Children), len(faks.Children)))
	}
	}
	if orig.Value != faks.Value {
	if orig.Value != faks.Value {
		errs = append(errs, fmt.Errorf("Orig <%s> element value differs from faks value.", orig.Name.Local))
		errs = append(errs, fmt.Errorf("Orig <%s> element value differs from faks value.", orig.Name.Local))
	}
	}
	namedNodes := map[string]*both{}
	namedNodes := map[string]*both{}
	for _, cn = range orig.Children {
	for _, cn = range orig.Children {
		if curBoth = namedNodes[cn.Name.Local]; curBoth == nil {
		if curBoth = namedNodes[cn.Name.Local]; curBoth == nil {
			curBoth = &both{}
			curBoth = &both{}
			namedNodes[cn.Name.Local] = curBoth
			namedNodes[cn.Name.Local] = curBoth
		}
		}
		curBoth.origNodes = append(curBoth.origNodes, cn)
		curBoth.origNodes = append(curBoth.origNodes, cn)
	}
	}
	for _, cn = range faks.Children {
	for _, cn = range faks.Children {
		if curBoth = namedNodes[cn.Name.Local]; curBoth != nil {
		if curBoth = namedNodes[cn.Name.Local]; curBoth != nil {
			curBoth.faksNodes = append(curBoth.faksNodes, cn)
			curBoth.faksNodes = append(curBoth.faksNodes, cn)
		}
		}
	}
	}
	for tmp, curBoth = range namedNodes {
	for tmp, curBoth = range namedNodes {
		if len(curBoth.origNodes) != len(curBoth.faksNodes) {
		if len(curBoth.origNodes) != len(curBoth.faksNodes) {
			errs = append(errs, fmt.Errorf("Orig <%s> element has %v <%s> elements but faks <%s> element has %v.", orig.Name.Local, len(curBoth.origNodes), tmp, faks.Name.Local, len(curBoth.faksNodes)))
			errs = append(errs, fmt.Errorf("Orig <%s> element has %v <%s> elements but faks <%s> element has %v.", orig.Name.Local, len(curBoth.origNodes), tmp, faks.Name.Local, len(curBoth.faksNodes)))
		} else if len(curBoth.origNodes) == 1 {
		} else if len(curBoth.origNodes) == 1 {
			errs = append(errs, verifyNode(curBoth.origNodes[0], curBoth.faksNodes[0])...)
			errs = append(errs, verifyNode(curBoth.origNodes[0], curBoth.faksNodes[0])...)
		} else {
		} else {
			for i, cn = range curBoth.origNodes {
			for i, cn = range curBoth.origNodes {
				if subErrs = verifyNode(cn, curBoth.faksNodes[i]); len(subErrs) > 0 {
				if subErrs = verifyNode(cn, curBoth.faksNodes[i]); len(subErrs) > 0 {
					errs = append(errs, subErrs...)
					errs = append(errs, subErrs...)
				}
				}
			}
			}
		}
		}
	}
	}
	return
	return
}
}


//	Attempts to xml.Unmarshal() all files in the "infiles" sub-directory of the specified directory path into the interface{} structure returned by the specified constructor.
//	Attempts to xml.Unmarshal() all files in the "infiles" sub-directory of the specified directory path into the interface{} structure returned by the specified constructor.
//	For each such input file, then attempts to xml.MarshalIndent() said structure back into a new output XML file with the same name, in the "outfiles" sub-directory of the specified directory path.
//	For each such input file, then attempts to xml.MarshalIndent() said structure back into a new output XML file with the same name, in the "outfiles" sub-directory of the specified directory path.
func TestViaRemarshal(dirPath string, makeEmptyDoc func() interface{}) {
func TestViaRemarshal(dirPath string, makeEmptyDoc func() interface{}) {
	var dirPathInFiles = filepath.Join(dirPath, "infiles")
	var dirPathInFiles = filepath.Join(dirPath, "infiles")
	var dirPathOutFiles = filepath.Join(dirPath, "outfiles")
	var dirPathOutFiles = filepath.Join(dirPath, "outfiles")
	var loadXmlDocFile = func(filename string) bool {
	var loadXmlDocFile = func(filename string) bool {
		log.Printf("Loading %s", filename)
		log.Printf("Loading %s", filename)
		doc, dataOrig := makeEmptyDoc(), ufs.ReadBinaryFile(filename, true)
		doc, dataOrig := makeEmptyDoc(), ufs.ReadBinaryFile(filename, true)
		err := xml.Unmarshal(dataOrig, doc)
		err := xml.Unmarshal(dataOrig, doc)
		if err != nil {
		if err != nil {
			panic(err)
			panic(err)
		}
		}
		if OnDocLoaded != nil {
		if OnDocLoaded != nil {
			OnDocLoaded(doc)
			OnDocLoaded(doc)
		}
		}
		outFileName := filepath.Join(dirPathOutFiles, filepath.Base(filename))
		outFileName := filepath.Join(dirPathOutFiles, filepath.Base(filename))
		log.Printf("Writing %s", outFileName)
		log.Printf("Writing %s", outFileName)
		dataFaks, err := xml.MarshalIndent(doc, "", "\t")
		dataFaks, err := xml.MarshalIndent(doc, "", "\t")
		if err != nil {
		if err != nil {
			panic(err)
			panic(err)
		}
		}
		ufs.WriteTextFile(outFileName, strings.Trim(string(dataFaks), " \r\n\t"))
		ufs.WriteTextFile(outFileName, strings.Trim(string(dataFaks), " \r\n\t"))
		log.Printf("Verifying...")
		log.Printf("Verifying...")
		if errs := verifyDocs(dataOrig, dataFaks); len(errs) > 0 {
		if errs := verifyDocs(dataOrig, dataFaks); len(errs) > 0 {
			for _, err = range errs {
			for _, err = range errs {
				log.Printf("%v", err)
				log.Printf("%v", err)
			}
			}
		}
		}
		return true
		return true
	}
	}
	if errs := ufs.NewDirWalker(false, nil, loadXmlDocFile).Walk(dirPathInFiles); len(errs) > 0 {
	if errs := ufs.NewDirWalker(false, nil, loadXmlDocFile).Walk(dirPathInFiles); len(errs) > 0 {
		panic(errs[0])
		panic(errs[0])
	}
	}
}
}
