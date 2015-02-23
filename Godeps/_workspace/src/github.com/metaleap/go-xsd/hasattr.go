package xsd
package xsd


import (
import (
	xsdt "github.com/metaleap/go-xsd/types"
	xsdt "github.com/metaleap/go-xsd/types"
)
)
type XsdtString struct{ string }


type hasAttrAbstract struct {
type hasAttrAbstract struct {
	Abstract bool `xml:"abstract,attr"`
	Abstract bool `xml:"abstract,attr"`
}
}


type XsdtString struct{ string }type hasAttrAttributeFormDefault struct {
	AttributeFormDefault string `xml:"attributeFormDefault,attr"`
	AttributeFormDefault string `xml:"attributeFormDefault,attr"`
}
}


type hasAttrBase struct {
type hasAttrBase struct {
	Base xsdt.Qname `xml:"base,attr"`
	Base xsdt.Qname `xml:"base,attr"`
}
}


type hasAttrBlock struct {
type hasAttrBlock struct {
	Block string `xml:"block,attr"`
	Block string `xml:"block,attr"`
}
}


type hasAttrBlockDefault struct {
type hasAttrBlockDefault struct {
	BlockDefault string `xml:"blockDefault,attr"`
	BlockDefault string `xml:"blockDefault,attr"`
}
}


type hasAttrDefault struct {
type hasAttrDefault struct {
	Default string `xml:"default,attr"`
	Default string `xml:"default,attr"`
}
}


type hasAttrFinal struct {
type hasAttrFinal struct {
	Final string `xml:"final,attr"`
	Final string `xml:"final,attr"`
}
}


type hasAttrFinalDefault struct {
type hasAttrFinalDefault struct {
	FinalDefault string `xml:"finalDefault,attr"`
	FinalDefault string `xml:"finalDefault,attr"`
}
}


type hasAttrFixed struct {
type hasAttrFixed struct {
	Fixed string `xml:"fixed,attr"`
	Fixed string `xml:"fixed,attr"`
}
}


type hasAttrForm struct {
type hasAttrForm struct {
	Form string `xml:"form,attr"`
	Form string `xml:"form,attr"`
}
}


type hasAttrElementFormDefault struct {
type hasAttrElementFormDefault struct {
	ElementFormDefault string `xml:"elementFormDefault,attr"`
	ElementFormDefault string `xml:"elementFormDefault,attr"`
}
}


type hasAttrId struct {
type hasAttrId struct {
	Id xsdt.Id `xml:"id,attr"`
	Id xsdt.Id `xml:"id,attr"`
}
}


type hasAttrItemType struct {
type hasAttrItemType struct {
	ItemType xsdt.Qname `xml:"itemType,attr"`
	ItemType xsdt.Qname `xml:"itemType,attr"`
}
}


type hasAttrLang struct {
type hasAttrLang struct {
	Lang xsdt.Language `xml:"lang,attr"`
	Lang xsdt.Language `xml:"lang,attr"`
}
}


type hasAttrMaxOccurs struct {
type hasAttrMaxOccurs struct {
	MaxOccurs string `xml:"maxOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
}
}


func (me *hasAttrMaxOccurs) Value() (l xsdt.Long) {
func (me *hasAttrMaxOccurs) Value() (l xsdt.Long) {
	if len(me.MaxOccurs) == 0 {
	if len(me.MaxOccurs) == 0 {
		l = 1
		l = 1
	} else if me.MaxOccurs == "unbounded" {
	} else if me.MaxOccurs == "unbounded" {
		l = -1
		l = -1
	} else {
	} else {
		l.Set(me.MaxOccurs)
		l.Set(me.MaxOccurs)
	}
	}
	return
	return
}
}


type hasAttrMemberTypes struct {
type hasAttrMemberTypes struct {
	MemberTypes string `xml:"memberTypes,attr"`
	MemberTypes string `xml:"memberTypes,attr"`
}
}


type hasAttrMinOccurs struct {
type hasAttrMinOccurs struct {
	MinOccurs uint64 `xml:"minOccurs,attr"`
	MinOccurs uint64 `xml:"minOccurs,attr"`
}
}


type hasAttrMixed struct {
type hasAttrMixed struct {
	Mixed bool `xml:"mixed,attr"`
	Mixed bool `xml:"mixed,attr"`
}
}


type hasAttrName struct {
type hasAttrName struct {
	Name xsdt.NCName `xml:"name,attr"`
	Name xsdt.NCName `xml:"name,attr"`
}
}


type hasAttrNamespace struct {
type hasAttrNamespace struct {
	Namespace string `xml:"namespace,attr"`
	Namespace string `xml:"namespace,attr"`
}
}


type hasAttrNillable struct {
type hasAttrNillable struct {
	Nillable bool `xml:"nillable,attr"`
	Nillable bool `xml:"nillable,attr"`
}
}


type hasAttrProcessContents struct {
type hasAttrProcessContents struct {
	ProcessContents string `xml:"processContents,attr"`
	ProcessContents string `xml:"processContents,attr"`
}
}


type hasAttrPublic struct {
type hasAttrPublic struct {
	Public string `xml:"public,attr"`
	Public string `xml:"public,attr"`
}
}


type hasAttrRef struct {
type hasAttrRef struct {
	Ref xsdt.Qname `xml:"ref,attr"`
	Ref xsdt.Qname `xml:"ref,attr"`
}
}


type hasAttrRefer struct {
type hasAttrRefer struct {
	Refer xsdt.Qname `xml:"refer,attr"`
	Refer xsdt.Qname `xml:"refer,attr"`
}
}


type hasAttrSchemaLocation struct {
type hasAttrSchemaLocation struct {
	SchemaLocation xsdt.AnyURI `xml:"schemaLocation,attr"`
	SchemaLocation xsdt.AnyURI `xml:"schemaLocation,attr"`
}
}


type hasAttrSource struct {
type hasAttrSource struct {
	Source xsdt.AnyURI `xml:"source,attr"`
	Source xsdt.AnyURI `xml:"source,attr"`
}
}


type hasAttrSubstitutionGroup struct {
type hasAttrSubstitutionGroup struct {
	SubstitutionGroup xsdt.Qname `xml:"substitutionGroup,attr"`
	SubstitutionGroup xsdt.Qname `xml:"substitutionGroup,attr"`
}
}


type hasAttrSystem struct {
type hasAttrSystem struct {
	System xsdt.AnyURI `xml:"system,attr"`
	System xsdt.AnyURI `xml:"system,attr"`
}
}


type hasAttrTargetNamespace struct {
type hasAttrTargetNamespace struct {
	TargetNamespace xsdt.AnyURI `xml:"targetNamespace,attr"`
	TargetNamespace xsdt.AnyURI `xml:"targetNamespace,attr"`
}
}


type hasAttrType struct {
type hasAttrType struct {
	Type xsdt.Qname `xml:"type,attr"`
	Type xsdt.Qname `xml:"type,attr"`
}
}


type hasAttrUse struct {
type hasAttrUse struct {
	Use string `xml:"use,attr"`
	Use string `xml:"use,attr"`
}
}


type hasAttrValue struct {
type hasAttrValue struct {
	Value string `xml:"value,attr"`
	Value string `xml:"value,attr"`
}
}


type hasAttrVersion struct {
type hasAttrVersion struct {
	Version xsdt.Token `xml:"version,attr"`
	Version xsdt.Token `xml:"version,attr"`
}
}


type hasAttrXpath struct {
type hasAttrXpath struct {
	Xpath string `xml:"xpath,attr"`
	Xpath string `xml:"xpath,attr"`
}
}
