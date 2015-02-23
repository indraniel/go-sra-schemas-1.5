package xsd
package xsd


import (
import (
	"bytes"
	"bytes"
	"encoding/xml"
	"encoding/xml"
type XsdtString struct{ string }


	"io/ioutil"
	"io/ioutil"
	"os"
	"os"
	"path"
	"path"
	"path/filepath"
	"path/filepath"
type XsdtString struct{ string }	"strings"


	"github.com/go-utils/ufs"
	"github.com/go-utils/ufs"
	"github.com/go-utils/unet"
	"github.com/go-utils/unet"
	"github.com/go-utils/ustr"
	"github.com/go-utils/ustr"
	"fmt"
	"fmt"
)
)


const (
const (
	goPkgPrefix     = ""
	goPkgPrefix     = ""
	goPkgSuffix     = "_go"
	goPkgSuffix     = "_go"
	protSep         = "://"
	protSep         = "://"
	xsdNamespaceUri = "http://www.w3.org/2001/XMLSchema"
	xsdNamespaceUri = "http://www.w3.org/2001/XMLSchema"
)
)


var (
var (
	loadedSchemas = map[string]*Schema{}
	loadedSchemas = map[string]*Schema{}
)
)


type Schema struct {
type Schema struct {
	elemBase
	elemBase
	XMLName            xml.Name          `xml:"schema"`
	XMLName            xml.Name          `xml:"schema"`
	XMLNamespacePrefix string            `xml:"-"`
	XMLNamespacePrefix string            `xml:"-"`
	XMLNamespaces      map[string]string `xml:"-"`
	XMLNamespaces      map[string]string `xml:"-"`
	XMLIncludedSchemas []*Schema         `xml:"-"`
	XMLIncludedSchemas []*Schema         `xml:"-"`
	XSDNamespacePrefix string            `xml:"-"`
	XSDNamespacePrefix string            `xml:"-"`
	XSDParentSchema    *Schema           `xml:"-"`
	XSDParentSchema    *Schema           `xml:"-"`


	hasAttrAttributeFormDefault
	hasAttrAttributeFormDefault
	hasAttrBlockDefault
	hasAttrBlockDefault
	hasAttrElementFormDefault
	hasAttrElementFormDefault
	hasAttrFinalDefault
	hasAttrFinalDefault
	hasAttrLang
	hasAttrLang
	hasAttrId
	hasAttrId
	hasAttrSchemaLocation
	hasAttrSchemaLocation
	hasAttrTargetNamespace
	hasAttrTargetNamespace
	hasAttrVersion
	hasAttrVersion
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsAttributeGroup
	hasElemsComplexType
	hasElemsComplexType
	hasElemsElement
	hasElemsElement
	hasElemsGroup
	hasElemsGroup
	hasElemsInclude
	hasElemsInclude
	hasElemsImport
	hasElemsImport
	hasElemsNotation
	hasElemsNotation
	hasElemsRedefine
	hasElemsRedefine
	hasElemsSimpleType
	hasElemsSimpleType


	loadLocalPath, loadUri string
	loadLocalPath, loadUri string
}
}


func (me *Schema) allSchemas(loadedSchemas map[string]bool) (schemas []*Schema) {
func (me *Schema) allSchemas(loadedSchemas map[string]bool) (schemas []*Schema) {
	schemas = append(schemas, me)
	schemas = append(schemas, me)
	loadedSchemas[me.loadUri] = true
	loadedSchemas[me.loadUri] = true
	for _, ss := range me.XMLIncludedSchemas {
	for _, ss := range me.XMLIncludedSchemas {
		if v, ok := loadedSchemas[ss.loadUri]; ok && v {
		if v, ok := loadedSchemas[ss.loadUri]; ok && v {
			continue
			continue
		}
		}
		schemas = append(schemas, ss.allSchemas(loadedSchemas)...)
		schemas = append(schemas, ss.allSchemas(loadedSchemas)...)
	}
	}
	return
	return
}
}


func (me *Schema) collectGlobals(bag *PkgBag, loadedSchemas map[string]bool) {
func (me *Schema) collectGlobals(bag *PkgBag, loadedSchemas map[string]bool) {
	loadedSchemas[me.loadUri] = true
	loadedSchemas[me.loadUri] = true
	for _, att := range me.Attributes {
	for _, att := range me.Attributes {
		bag.allAtts = append(bag.allAtts, att)
		bag.allAtts = append(bag.allAtts, att)
	}
	}
	for _, agr := range me.AttributeGroups {
	for _, agr := range me.AttributeGroups {
		bag.allAttGroups = append(bag.allAttGroups, agr)
		bag.allAttGroups = append(bag.allAttGroups, agr)
	}
	}
	for _, el := range me.Elements {
	for _, el := range me.Elements {
		bag.allElems = append(bag.allElems, el)
		bag.allElems = append(bag.allElems, el)
	}
	}
	for _, egr := range me.Groups {
	for _, egr := range me.Groups {
		bag.allElemGroups = append(bag.allElemGroups, egr)
		bag.allElemGroups = append(bag.allElemGroups, egr)
	}
	}
	for _, not := range me.Notations {
	for _, not := range me.Notations {
		bag.allNotations = append(bag.allNotations, not)
		bag.allNotations = append(bag.allNotations, not)
	}
	}
	for _, ss := range me.XMLIncludedSchemas {
	for _, ss := range me.XMLIncludedSchemas {
		if v, ok := loadedSchemas[ss.loadUri]; ok && v {
		if v, ok := loadedSchemas[ss.loadUri]; ok && v {
			continue
			continue
		}
		}
		ss.collectGlobals(bag, loadedSchemas)
		ss.collectGlobals(bag, loadedSchemas)
	}
	}
}
}


func (me *Schema) globalComplexType(bag *PkgBag, name string, loadedSchemas map[string]bool) (ct *ComplexType) {
func (me *Schema) globalComplexType(bag *PkgBag, name string, loadedSchemas map[string]bool) (ct *ComplexType) {
	var imp string
	var imp string
	for _, ct = range me.ComplexTypes {
	for _, ct = range me.ComplexTypes {
		if bag.resolveQnameRef(ustr.PrefixWithSep(me.XMLNamespacePrefix, ":", ct.Name.String()), "T", &imp) == name {
		if bag.resolveQnameRef(ustr.PrefixWithSep(me.XMLNamespacePrefix, ":", ct.Name.String()), "T", &imp) == name {
			return
			return
		}
		}
	}
	}
	loadedSchemas[me.loadUri] = true
	loadedSchemas[me.loadUri] = true
	for _, ss := range me.XMLIncludedSchemas {
	for _, ss := range me.XMLIncludedSchemas {
		if v, ok := loadedSchemas[ss.loadUri]; ok && v {
		if v, ok := loadedSchemas[ss.loadUri]; ok && v {
			//fmt.Printf("Ignoring processed schema: %s\n", ss.loadUri)
			//fmt.Printf("Ignoring processed schema: %s\n", ss.loadUri)
			continue
			continue
		}
		}
		if ct = ss.globalComplexType(bag, name, loadedSchemas); ct != nil {
		if ct = ss.globalComplexType(bag, name, loadedSchemas); ct != nil {
			return
			return
		}
		}
	}
	}
	ct = nil
	ct = nil
	return
	return
}
}


func (me *Schema) globalElement(bag *PkgBag, name string) (el *Element) {
func (me *Schema) globalElement(bag *PkgBag, name string) (el *Element) {
	var imp string
	var imp string
	if len(name) > 0 {
	if len(name) > 0 {
		var rname = bag.resolveQnameRef(name, "", &imp)
		var rname = bag.resolveQnameRef(name, "", &imp)
		for _, el = range me.Elements {
		for _, el = range me.Elements {
			if bag.resolveQnameRef(ustr.PrefixWithSep(me.XMLNamespacePrefix, ":", el.Name.String()), "", &imp) == rname {
			if bag.resolveQnameRef(ustr.PrefixWithSep(me.XMLNamespacePrefix, ":", el.Name.String()), "", &imp) == rname {
				return
				return
			}
			}
		}
		}
		for _, ss := range me.XMLIncludedSchemas {
		for _, ss := range me.XMLIncludedSchemas {
			if el = ss.globalElement(bag, name); el != nil {
			if el = ss.globalElement(bag, name); el != nil {
				return
				return
			}
			}
		}
		}
	}
	}
	el = nil
	el = nil
	return
	return
}
}


func (me *Schema) globalSubstitutionElems(el *Element, loadedSchemas map[string]bool) (els []*Element) {
func (me *Schema) globalSubstitutionElems(el *Element, loadedSchemas map[string]bool) (els []*Element) {
	var elName = el.Ref.String()
	var elName = el.Ref.String()
	if len(elName) == 0 {
	if len(elName) == 0 {
		elName = el.Name.String()
		elName = el.Name.String()
	}
	}
	for _, tle := range me.Elements {
	for _, tle := range me.Elements {
		if (tle != el) && (len(tle.SubstitutionGroup) > 0) {
		if (tle != el) && (len(tle.SubstitutionGroup) > 0) {
			if tle.SubstitutionGroup.String()[(strings.Index(tle.SubstitutionGroup.String(), ":")+1):] == elName {
			if tle.SubstitutionGroup.String()[(strings.Index(tle.SubstitutionGroup.String(), ":")+1):] == elName {
				els = append(els, tle)
				els = append(els, tle)
			}
			}
		}
		}
	}
	}
	loadedSchemas[me.loadUri] = true
	loadedSchemas[me.loadUri] = true
	for _, inc := range me.XMLIncludedSchemas {
	for _, inc := range me.XMLIncludedSchemas {
		if v, ok := loadedSchemas[inc.loadUri]; ok && v {
		if v, ok := loadedSchemas[inc.loadUri]; ok && v {
			//fmt.Printf("Ignoring processed schema: %s\n", inc.loadUri)
			//fmt.Printf("Ignoring processed schema: %s\n", inc.loadUri)
			continue
			continue
		}
		}
		els = append(els, inc.globalSubstitutionElems(el, loadedSchemas)...)
		els = append(els, inc.globalSubstitutionElems(el, loadedSchemas)...)
	}
	}
	return
	return
}
}


func (me *Schema) MakeGoPkgSrcFile() (goOutFilePath string, err error) {
func (me *Schema) MakeGoPkgSrcFile() (goOutFilePath string, err error) {
	var goOutDirPath = filepath.Join(filepath.Dir(me.loadLocalPath), goPkgPrefix+filepath.Base(me.loadLocalPath)+goPkgSuffix)
	var goOutDirPath = filepath.Join(filepath.Dir(me.loadLocalPath), goPkgPrefix+filepath.Base(me.loadLocalPath)+goPkgSuffix)
	goOutFilePath = filepath.Join(goOutDirPath, path.Base(me.loadUri)+".go")
	goOutFilePath = filepath.Join(goOutDirPath, path.Base(me.loadUri)+".go")
	var bag = newPkgBag(me)
	var bag = newPkgBag(me)
	loadedSchemas := make(map[string]bool)
	loadedSchemas := make(map[string]bool)
	for _, inc := range me.allSchemas(loadedSchemas) {
	for _, inc := range me.allSchemas(loadedSchemas) {
		bag.Schema = inc
		bag.Schema = inc
		inc.makePkg(bag)
		inc.makePkg(bag)
	}
	}
	bag.Schema = me
	bag.Schema = me
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	bag.appendFmt(true, "")
	bag.appendFmt(true, "")
	me.makePkg(bag)
	me.makePkg(bag)
	if err = ufs.EnsureDirExists(filepath.Dir(goOutFilePath)); err == nil {
	if err = ufs.EnsureDirExists(filepath.Dir(goOutFilePath)); err == nil {
		err = ufs.WriteTextFile(goOutFilePath, bag.assembleSource())
		err = ufs.WriteTextFile(goOutFilePath, bag.assembleSource())
	}
	}
	return
	return
}
}


func (me *Schema) onLoad(rootAtts []xml.Attr, loadUri, localPath string) (err error) {
func (me *Schema) onLoad(rootAtts []xml.Attr, loadUri, localPath string) (err error) {
	var tmpUrl string
	var tmpUrl string
	var sd *Schema
	var sd *Schema
	loadedSchemas[loadUri] = me
	loadedSchemas[loadUri] = me
	me.loadLocalPath, me.loadUri = localPath, loadUri
	me.loadLocalPath, me.loadUri = localPath, loadUri
	me.XMLNamespaces = map[string]string{}
	me.XMLNamespaces = map[string]string{}
	for _, att := range rootAtts {
	for _, att := range rootAtts {
		if att.Name.Space == "xmlns" {
		if att.Name.Space == "xmlns" {
			me.XMLNamespaces[att.Name.Local] = att.Value
			me.XMLNamespaces[att.Name.Local] = att.Value
		} else if len(att.Name.Space) > 0 {
		} else if len(att.Name.Space) > 0 {


		} else if att.Name.Local == "xmlns" {
		} else if att.Name.Local == "xmlns" {
			me.XMLNamespaces[""] = att.Value
			me.XMLNamespaces[""] = att.Value
		}
		}
	}
	}
	for k, v := range me.XMLNamespaces {
	for k, v := range me.XMLNamespaces {
		if v == xsdNamespaceUri {
		if v == xsdNamespaceUri {
			me.XSDNamespacePrefix = k
			me.XSDNamespacePrefix = k
		} else if v == me.TargetNamespace.String() {
		} else if v == me.TargetNamespace.String() {
			me.XMLNamespacePrefix = k
			me.XMLNamespacePrefix = k
		}
		}
	}
	}
	if len(me.XMLNamespaces["xml"]) == 0 {
	if len(me.XMLNamespaces["xml"]) == 0 {
		me.XMLNamespaces["xml"] = "http://www.w3.org/XML/1998/namespace"
		me.XMLNamespaces["xml"] = "http://www.w3.org/XML/1998/namespace"
	}
	}
	me.XMLIncludedSchemas = []*Schema{}
	me.XMLIncludedSchemas = []*Schema{}
	for _, inc := range me.Includes {
	for _, inc := range me.Includes {
		if tmpUrl = inc.SchemaLocation.String(); strings.Index(tmpUrl, protSep) < 0 {
		if tmpUrl = inc.SchemaLocation.String(); strings.Index(tmpUrl, protSep) < 0 {
			tmpUrl = path.Join(path.Dir(loadUri), tmpUrl)
			tmpUrl = path.Join(path.Dir(loadUri), tmpUrl)
		}
		}
		var ok bool
		var ok bool
		var toLoadUri string
		var toLoadUri string
		if pos := strings.Index(tmpUrl, protSep); pos >= 0 {
		if pos := strings.Index(tmpUrl, protSep); pos >= 0 {
			toLoadUri = tmpUrl[pos+len(protSep):]
			toLoadUri = tmpUrl[pos+len(protSep):]
		} else {
		} else {
			toLoadUri = tmpUrl
			toLoadUri = tmpUrl
		}
		}
		if sd, ok = loadedSchemas[toLoadUri]; !ok {
		if sd, ok = loadedSchemas[toLoadUri]; !ok {
			if sd, err = LoadSchema(tmpUrl, len(localPath) > 0); err != nil {
			if sd, err = LoadSchema(tmpUrl, len(localPath) > 0); err != nil {
				return
				return
			}
			}
		}
		}
		sd.XSDParentSchema = me
		sd.XSDParentSchema = me
		me.XMLIncludedSchemas = append(me.XMLIncludedSchemas, sd)
		me.XMLIncludedSchemas = append(me.XMLIncludedSchemas, sd)
	}
	}
	me.initElement(nil)
	me.initElement(nil)
	return
	return
}
}


func (me *Schema) RootSchema(pathSchemas []string) *Schema {
func (me *Schema) RootSchema(pathSchemas []string) *Schema {
	if me.XSDParentSchema != nil {
	if me.XSDParentSchema != nil {
		for _, sch := range pathSchemas {
		for _, sch := range pathSchemas {
			if me.XSDParentSchema.loadUri == sch {
			if me.XSDParentSchema.loadUri == sch {
				fmt.Printf("schema loop detected %+v - > %s!\n", pathSchemas, me.XSDParentSchema.loadUri)
				fmt.Printf("schema loop detected %+v - > %s!\n", pathSchemas, me.XSDParentSchema.loadUri)
				return me
				return me
			}
			}
	  }
	  }
		pathSchemas = append(pathSchemas, me.loadUri)
		pathSchemas = append(pathSchemas, me.loadUri)
		return me.XSDParentSchema.RootSchema(pathSchemas)
		return me.XSDParentSchema.RootSchema(pathSchemas)
	}
	}
	return me
	return me
}
}


func ClearLoadedSchemasCache() {
func ClearLoadedSchemasCache() {
	loadedSchemas = map[string]*Schema{}
	loadedSchemas = map[string]*Schema{}
}
}


func loadSchema(r io.Reader, loadUri, localPath string) (sd *Schema, err error) {
func loadSchema(r io.Reader, loadUri, localPath string) (sd *Schema, err error) {
	var data []byte
	var data []byte
	var rootAtts []xml.Attr
	var rootAtts []xml.Attr
	if data, err = ioutil.ReadAll(r); err == nil {
	if data, err = ioutil.ReadAll(r); err == nil {
		var t xml.Token
		var t xml.Token
		sd = new(Schema)
		sd = new(Schema)
		for xd := xml.NewDecoder(bytes.NewReader(data)); err == nil; {
		for xd := xml.NewDecoder(bytes.NewReader(data)); err == nil; {
			if t, err = xd.Token(); err == nil {
			if t, err = xd.Token(); err == nil {
				if startEl, ok := t.(xml.StartElement); ok {
				if startEl, ok := t.(xml.StartElement); ok {
					rootAtts = startEl.Attr
					rootAtts = startEl.Attr
					break
					break
				}
				}
			}
			}
		}
		}
		if err = xml.Unmarshal(data, sd); err == nil {
		if err = xml.Unmarshal(data, sd); err == nil {
			err = sd.onLoad(rootAtts, loadUri, localPath)
			err = sd.onLoad(rootAtts, loadUri, localPath)
		}
		}
		if err != nil {
		if err != nil {
			sd = nil
			sd = nil
		}
		}
	}
	}
	return
	return
}
}


func loadSchemaFile(filename string, loadUri string) (sd *Schema, err error) {
func loadSchemaFile(filename string, loadUri string) (sd *Schema, err error) {
	var file *os.File
	var file *os.File
	if file, err = os.Open(filename); err == nil {
	if file, err = os.Open(filename); err == nil {
		defer file.Close()
		defer file.Close()
		sd, err = loadSchema(file, loadUri, filename)
		sd, err = loadSchema(file, loadUri, filename)
	}
	}
	return
	return
}
}


func LoadSchema(uri string, localCopy bool) (sd *Schema, err error) {
func LoadSchema(uri string, localCopy bool) (sd *Schema, err error) {
	var protocol, localPath string
	var protocol, localPath string
	var rc io.ReadCloser
	var rc io.ReadCloser


	if pos := strings.Index(uri, protSep); pos < 0 {
	if pos := strings.Index(uri, protSep); pos < 0 {
		protocol = "http" + protSep
		protocol = "http" + protSep
	} else {
	} else {
		protocol = uri[:pos+len(protSep)]
		protocol = uri[:pos+len(protSep)]
		uri = uri[pos+len(protSep):]
		uri = uri[pos+len(protSep):]
	}
	}
	if localCopy {
	if localCopy {
		if localPath = filepath.Join(PkgGen.BaseCodePath, uri); !ufs.FileExists(localPath) {
		if localPath = filepath.Join(PkgGen.BaseCodePath, uri); !ufs.FileExists(localPath) {
			if err = ufs.EnsureDirExists(filepath.Dir(localPath)); err == nil {
			if err = ufs.EnsureDirExists(filepath.Dir(localPath)); err == nil {
				err = unet.DownloadFile(protocol+uri, localPath)
				err = unet.DownloadFile(protocol+uri, localPath)
			}
			}
		}
		}
		if err == nil {
		if err == nil {
			if sd, err = loadSchemaFile(localPath, uri); sd != nil {
			if sd, err = loadSchemaFile(localPath, uri); sd != nil {
				sd.loadLocalPath = localPath
				sd.loadLocalPath = localPath
			}
			}
		}
		}
	} else if rc, err = unet.OpenRemoteFile(protocol + uri); err == nil {
	} else if rc, err = unet.OpenRemoteFile(protocol + uri); err == nil {
		defer rc.Close()
		defer rc.Close()
		sd, err = loadSchema(rc, uri, "")
		sd, err = loadSchema(rc, uri, "")
	}
	}
	return
	return
}
}
