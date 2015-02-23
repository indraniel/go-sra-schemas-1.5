package xsd
package xsd


import (
import (
	xsdt "github.com/metaleap/go-xsd/types"
	xsdt "github.com/metaleap/go-xsd/types"
)
)
type XsdtString struct{ string }


type element interface {
type element interface {
	base() *elemBase
	base() *elemBase
	init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake)
	init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake)
	Parent() element
	Parent() element
type XsdtString struct{ string }}


type elemBase struct {
type elemBase struct {
	atts         []beforeAfterMake
	atts         []beforeAfterMake
	parent, self element // self is the struct that embeds elemBase, rather than the elemBase pseudo-field
	parent, self element // self is the struct that embeds elemBase, rather than the elemBase pseudo-field
	xsdName      xsdt.NCName
	xsdName      xsdt.NCName
	hasNameAttr  bool
	hasNameAttr  bool
}
}


func (me *elemBase) afterMakePkg(bag *PkgBag) {
func (me *elemBase) afterMakePkg(bag *PkgBag) {
	if !me.hasNameAttr {
	if !me.hasNameAttr {
		bag.Stacks.Name.Pop()
		bag.Stacks.Name.Pop()
	}
	}
	for _, a := range me.atts {
	for _, a := range me.atts {
		a.afterMakePkg(bag)
		a.afterMakePkg(bag)
	}
	}
}
}


func (me *elemBase) beforeMakePkg(bag *PkgBag) {
func (me *elemBase) beforeMakePkg(bag *PkgBag) {
	if !me.hasNameAttr {
	if !me.hasNameAttr {
		bag.Stacks.Name.Push(me.xsdName)
		bag.Stacks.Name.Push(me.xsdName)
	}
	}
	for _, a := range me.atts {
	for _, a := range me.atts {
		a.beforeMakePkg(bag)
		a.beforeMakePkg(bag)
	}
	}
}
}


func (me *elemBase) base() *elemBase { return me }
func (me *elemBase) base() *elemBase { return me }


func (me *elemBase) init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake) {
func (me *elemBase) init(parent, self element, xsdName xsdt.NCName, atts ...beforeAfterMake) {
	me.parent, me.self, me.xsdName, me.atts = parent, self, xsdName, atts
	me.parent, me.self, me.xsdName, me.atts = parent, self, xsdName, atts
	for _, a := range atts {
	for _, a := range atts {
		if _, me.hasNameAttr = a.(*hasAttrName); me.hasNameAttr {
		if _, me.hasNameAttr = a.(*hasAttrName); me.hasNameAttr {
			break
			break
		}
		}
	}
	}
}
}


func (me *elemBase) longSafeName(bag *PkgBag) (ln string) {
func (me *elemBase) longSafeName(bag *PkgBag) (ln string) {
	var els = []element{}
	var els = []element{}
	for el := me.self; (el != nil) && (el != bag.Schema); el = el.Parent() {
	for el := me.self; (el != nil) && (el != bag.Schema); el = el.Parent() {
		els = append(els, el)
		els = append(els, el)
	}
	}
	for i := len(els) - 1; i >= 0; i-- {
	for i := len(els) - 1; i >= 0; i-- {
		ln += bag.safeName(els[i].base().selfName().String())
		ln += bag.safeName(els[i].base().selfName().String())
	}
	}
	return
	return
}
}


func (me *elemBase) selfName() xsdt.NCName {
func (me *elemBase) selfName() xsdt.NCName {
	if me.hasNameAttr {
	if me.hasNameAttr {
		for _, at := range me.atts {
		for _, at := range me.atts {
			if an, ok := at.(*hasAttrName); ok {
			if an, ok := at.(*hasAttrName); ok {
				return an.Name
				return an.Name
			}
			}
		}
		}
	}
	}
	return me.xsdName
	return me.xsdName
}
}


func (me *elemBase) Parent() element { return me.parent }
func (me *elemBase) Parent() element { return me.parent }


type All struct {
type All struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"all"`
	//	XMLName xml.Name `xml:"all"`
	hasAttrId
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemAnnotation
	hasElemsElement
	hasElemsElement
}
}


type Annotation struct {
type Annotation struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"annotation"`
	//	XMLName xml.Name `xml:"annotation"`
	hasElemsAppInfo
	hasElemsAppInfo
	hasElemsDocumentation
	hasElemsDocumentation
}
}


type Any struct {
type Any struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"any"`
	//	XMLName xml.Name `xml:"any"`
	hasAttrId
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrMinOccurs
	hasAttrNamespace
	hasAttrNamespace
	hasAttrProcessContents
	hasAttrProcessContents
	hasElemAnnotation
	hasElemAnnotation
}
}


type AnyAttribute struct {
type AnyAttribute struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"anyAttribute"`
	//	XMLName xml.Name `xml:"anyAttribute"`
	hasAttrId
	hasAttrId
	hasAttrNamespace
	hasAttrNamespace
	hasAttrProcessContents
	hasAttrProcessContents
	hasElemAnnotation
	hasElemAnnotation
}
}


type AppInfo struct {
type AppInfo struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"appinfo"`
	//	XMLName xml.Name `xml:"appinfo"`
	hasAttrSource
	hasAttrSource
	hasCdata
	hasCdata
}
}


type Attribute struct {
type Attribute struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"attribute"`
	//	XMLName xml.Name `xml:"attribute"`
	hasAttrDefault
	hasAttrDefault
	hasAttrFixed
	hasAttrFixed
	hasAttrForm
	hasAttrForm
	hasAttrId
	hasAttrId
	hasAttrName
	hasAttrName
	hasAttrRef
	hasAttrRef
	hasAttrType
	hasAttrType
	hasAttrUse
	hasAttrUse
	hasElemAnnotation
	hasElemAnnotation
	hasElemsSimpleType
	hasElemsSimpleType
}
}


type AttributeGroup struct {
type AttributeGroup struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"attributeGroup"`
	//	XMLName xml.Name `xml:"attributeGroup"`
	hasAttrId
	hasAttrId
	hasAttrName
	hasAttrName
	hasAttrRef
	hasAttrRef
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsAttributeGroup
}
}


type Choice struct {
type Choice struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"choice"`
	//	XMLName xml.Name `xml:"choice"`
	hasAttrId
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAny
	hasElemsAny
	hasElemsChoice
	hasElemsChoice
	hasElemsElement
	hasElemsElement
	hasElemsGroup
	hasElemsGroup
	hasElemsSequence
	hasElemsSequence
}
}


type ComplexContent struct {
type ComplexContent struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"complexContent"`
	//	XMLName xml.Name `xml:"complexContent"`
	hasAttrId
	hasAttrId
	hasAttrMixed
	hasAttrMixed
	hasElemAnnotation
	hasElemAnnotation
	hasElemExtensionComplexContent
	hasElemExtensionComplexContent
	hasElemRestrictionComplexContent
	hasElemRestrictionComplexContent
}
}


type ComplexType struct {
type ComplexType struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"complexType"`
	//	XMLName xml.Name `xml:"complexType"`
	hasAttrAbstract
	hasAttrAbstract
	hasAttrBlock
	hasAttrBlock
	hasAttrFinal
	hasAttrFinal
	hasAttrId
	hasAttrId
	hasAttrMixed
	hasAttrMixed
	hasAttrName
	hasAttrName
	hasElemAll
	hasElemAll
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsAttributeGroup
	hasElemChoice
	hasElemChoice
	hasElemComplexContent
	hasElemComplexContent
	hasElemGroup
	hasElemGroup
	hasElemSequence
	hasElemSequence
	hasElemSimpleContent
	hasElemSimpleContent
}
}


type Documentation struct {
type Documentation struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"documentation"`
	//	XMLName xml.Name `xml:"documentation"`
	hasAttrLang
	hasAttrLang
	hasAttrSource
	hasAttrSource
	hasCdata
	hasCdata
}
}


type Element struct {
type Element struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"element"`
	//	XMLName xml.Name `xml:"element"`
	hasAttrAbstract
	hasAttrAbstract
	hasAttrBlock
	hasAttrBlock
	hasAttrDefault
	hasAttrDefault
	hasAttrFinal
	hasAttrFinal
	hasAttrFixed
	hasAttrFixed
	hasAttrForm
	hasAttrForm
	hasAttrId
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrMinOccurs
	hasAttrName
	hasAttrName
	hasAttrNillable
	hasAttrNillable
	hasAttrRef
	hasAttrRef
	hasAttrSubstitutionGroup
	hasAttrSubstitutionGroup
	hasAttrType
	hasAttrType
	hasElemAnnotation
	hasElemAnnotation
	hasElemComplexType
	hasElemComplexType
	hasElemsKey
	hasElemsKey
	hasElemKeyRef
	hasElemKeyRef
	hasElemsSimpleType
	hasElemsSimpleType
	hasElemUnique
	hasElemUnique
}
}


type ExtensionComplexContent struct {
type ExtensionComplexContent struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"extension"`
	//	XMLName xml.Name `xml:"extension"`
	hasAttrBase
	hasAttrBase
	hasAttrId
	hasAttrId
	hasElemAll
	hasElemAll
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsAttributeGroup
	hasElemsChoice
	hasElemsChoice
	hasElemsGroup
	hasElemsGroup
	hasElemsSequence
	hasElemsSequence
}
}


type ExtensionSimpleContent struct {
type ExtensionSimpleContent struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"extension"`
	//	XMLName xml.Name `xml:"extension"`
	hasAttrBase
	hasAttrBase
	hasAttrId
	hasAttrId
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsAttributeGroup
}
}


type Field struct {
type Field struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"field"`
	//	XMLName xml.Name `xml:"field"`
	hasAttrId
	hasAttrId
	hasAttrXpath
	hasAttrXpath
	hasElemAnnotation
	hasElemAnnotation
}
}


type Group struct {
type Group struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"group"`
	//	XMLName xml.Name `xml:"group"`
	hasAttrId
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrMinOccurs
	hasAttrName
	hasAttrName
	hasAttrRef
	hasAttrRef
	hasElemAll
	hasElemAll
	hasElemAnnotation
	hasElemAnnotation
	hasElemChoice
	hasElemChoice
	hasElemSequence
	hasElemSequence
}
}


type Include struct {
type Include struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"include"`
	//	XMLName xml.Name `xml:"include"`
	hasAttrId
	hasAttrId
	hasAttrSchemaLocation
	hasAttrSchemaLocation
	hasElemAnnotation
	hasElemAnnotation
}
}


type Import struct {
type Import struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"import"`
	//	XMLName xml.Name `xml:"import"`
	hasAttrId
	hasAttrId
	hasAttrNamespace
	hasAttrNamespace
	hasAttrSchemaLocation
	hasAttrSchemaLocation
	hasElemAnnotation
	hasElemAnnotation
}
}


type Key struct {
type Key struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"key"`
	//	XMLName xml.Name `xml:"key"`
	hasAttrId
	hasAttrId
	hasAttrName
	hasAttrName
	hasElemAnnotation
	hasElemAnnotation
	hasElemField
	hasElemField
	hasElemSelector
	hasElemSelector
}
}


type KeyRef struct {
type KeyRef struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"keyref"`
	//	XMLName xml.Name `xml:"keyref"`
	hasAttrId
	hasAttrId
	hasAttrName
	hasAttrName
	hasAttrRefer
	hasAttrRefer
	hasElemAnnotation
	hasElemAnnotation
	hasElemField
	hasElemField
	hasElemSelector
	hasElemSelector
}
}


type List struct {
type List struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"list"`
	//	XMLName xml.Name `xml:"list"`
	hasAttrId
	hasAttrId
	hasAttrItemType
	hasAttrItemType
	hasElemAnnotation
	hasElemAnnotation
	hasElemsSimpleType
	hasElemsSimpleType
}
}


type Notation struct {
type Notation struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"notation"`
	//	XMLName xml.Name `xml:"notation"`
	hasAttrId
	hasAttrId
	hasAttrName
	hasAttrName
	hasAttrPublic
	hasAttrPublic
	hasAttrSystem
	hasAttrSystem
	hasElemAnnotation
	hasElemAnnotation
}
}


type Redefine struct {
type Redefine struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"redefine"`
	//	XMLName xml.Name `xml:"redefine"`
	hasAttrId
	hasAttrId
	hasAttrSchemaLocation
	hasAttrSchemaLocation
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAttributeGroup
	hasElemsAttributeGroup
	hasElemsComplexType
	hasElemsComplexType
	hasElemsGroup
	hasElemsGroup
	hasElemsSimpleType
	hasElemsSimpleType
}
}


type RestrictionComplexContent struct {
type RestrictionComplexContent struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"restriction"`
	//	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrBase
	hasAttrId
	hasAttrId
	hasElemAll
	hasElemAll
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsAttributeGroup
	hasElemsChoice
	hasElemsChoice
	hasElemsSequence
	hasElemsSequence
}
}


type RestrictionSimpleContent struct {
type RestrictionSimpleContent struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"restriction"`
	//	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrBase
	hasAttrId
	hasAttrId
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsAttributeGroup
	hasElemsEnumeration
	hasElemsEnumeration
	hasElemFractionDigits
	hasElemFractionDigits
	hasElemLength
	hasElemLength
	hasElemMaxExclusive
	hasElemMaxExclusive
	hasElemMaxInclusive
	hasElemMaxInclusive
	hasElemMaxLength
	hasElemMaxLength
	hasElemMinExclusive
	hasElemMinExclusive
	hasElemMinInclusive
	hasElemMinInclusive
	hasElemMinLength
	hasElemMinLength
	hasElemPattern
	hasElemPattern
	hasElemsSimpleType
	hasElemsSimpleType
	hasElemTotalDigits
	hasElemTotalDigits
	hasElemWhiteSpace
	hasElemWhiteSpace
}
}


type RestrictionSimpleEnumeration struct {
type RestrictionSimpleEnumeration struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"enumeration"`
	//	XMLName xml.Name `xml:"enumeration"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleFractionDigits struct {
type RestrictionSimpleFractionDigits struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"fractionDigits"`
	//	XMLName xml.Name `xml:"fractionDigits"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleLength struct {
type RestrictionSimpleLength struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"length"`
	//	XMLName xml.Name `xml:"length"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleMaxExclusive struct {
type RestrictionSimpleMaxExclusive struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"maxExclusive"`
	//	XMLName xml.Name `xml:"maxExclusive"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleMaxInclusive struct {
type RestrictionSimpleMaxInclusive struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"maxInclusive"`
	//	XMLName xml.Name `xml:"maxInclusive"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleMaxLength struct {
type RestrictionSimpleMaxLength struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"maxLength"`
	//	XMLName xml.Name `xml:"maxLength"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleMinExclusive struct {
type RestrictionSimpleMinExclusive struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"minExclusive"`
	//	XMLName xml.Name `xml:"minExclusive"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleMinInclusive struct {
type RestrictionSimpleMinInclusive struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"minInclusive"`
	//	XMLName xml.Name `xml:"minInclusive"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleMinLength struct {
type RestrictionSimpleMinLength struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"minLength"`
	//	XMLName xml.Name `xml:"minLength"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimplePattern struct {
type RestrictionSimplePattern struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"pattern"`
	//	XMLName xml.Name `xml:"pattern"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleTotalDigits struct {
type RestrictionSimpleTotalDigits struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"totalDigits"`
	//	XMLName xml.Name `xml:"totalDigits"`
	hasAttrValue
	hasAttrValue
}
}


type RestrictionSimpleType struct {
type RestrictionSimpleType struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"restriction"`
	//	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrBase
	hasAttrId
	hasAttrId
	hasElemAnnotation
	hasElemAnnotation
	hasElemsEnumeration
	hasElemsEnumeration
	hasElemFractionDigits
	hasElemFractionDigits
	hasElemLength
	hasElemLength
	hasElemMaxExclusive
	hasElemMaxExclusive
	hasElemMaxInclusive
	hasElemMaxInclusive
	hasElemMaxLength
	hasElemMaxLength
	hasElemMinExclusive
	hasElemMinExclusive
	hasElemMinInclusive
	hasElemMinInclusive
	hasElemMinLength
	hasElemMinLength
	hasElemPattern
	hasElemPattern
	hasElemsSimpleType
	hasElemsSimpleType
	hasElemTotalDigits
	hasElemTotalDigits
	hasElemWhiteSpace
	hasElemWhiteSpace
}
}


type RestrictionSimpleWhiteSpace struct {
type RestrictionSimpleWhiteSpace struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"whiteSpace"`
	//	XMLName xml.Name `xml:"whiteSpace"`
	hasAttrValue
	hasAttrValue
}
}


type Selector struct {
type Selector struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"selector"`
	//	XMLName xml.Name `xml:"selector"`
	hasAttrId
	hasAttrId
	hasAttrXpath
	hasAttrXpath
	hasElemAnnotation
	hasElemAnnotation
}
}


type Sequence struct {
type Sequence struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"sequence"`
	//	XMLName xml.Name `xml:"sequence"`
	hasAttrId
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemAnnotation
	hasElemsAny
	hasElemsAny
	hasElemsChoice
	hasElemsChoice
	hasElemsElement
	hasElemsElement
	hasElemsGroup
	hasElemsGroup
	hasElemsSequence
	hasElemsSequence
}
}


type SimpleContent struct {
type SimpleContent struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"simpleContent"`
	//	XMLName xml.Name `xml:"simpleContent"`
	hasAttrId
	hasAttrId
	hasElemAnnotation
	hasElemAnnotation
	hasElemExtensionSimpleContent
	hasElemExtensionSimpleContent
	hasElemRestrictionSimpleContent
	hasElemRestrictionSimpleContent
}
}


type SimpleType struct {
type SimpleType struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"simpleType"`
	//	XMLName xml.Name `xml:"simpleType"`
	hasAttrFinal
	hasAttrFinal
	hasAttrId
	hasAttrId
	hasAttrName
	hasAttrName
	hasElemAnnotation
	hasElemAnnotation
	hasElemList
	hasElemList
	hasElemRestrictionSimpleType
	hasElemRestrictionSimpleType
	hasElemUnion
	hasElemUnion
}
}


type Union struct {
type Union struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"union"`
	//	XMLName xml.Name `xml:"union"`
	hasAttrId
	hasAttrId
	hasAttrMemberTypes
	hasAttrMemberTypes
	hasElemAnnotation
	hasElemAnnotation
	hasElemsSimpleType
	hasElemsSimpleType
}
}


type Unique struct {
type Unique struct {
	elemBase
	elemBase
	//	XMLName xml.Name `xml:"unique"`
	//	XMLName xml.Name `xml:"unique"`
	hasAttrId
	hasAttrId
	hasAttrName
	hasAttrName
	hasElemAnnotation
	hasElemAnnotation
	hasElemField
	hasElemField
	hasElemSelector
	hasElemSelector
}
}


func Flattened(choices []*Choice, seqs []*Sequence) (allChoices []*Choice, allSeqs []*Sequence) {
func Flattened(choices []*Choice, seqs []*Sequence) (allChoices []*Choice, allSeqs []*Sequence) {
	var tmpChoices []*Choice
	var tmpChoices []*Choice
	var tmpSeqs []*Sequence
	var tmpSeqs []*Sequence
	for _, ch := range choices {
	for _, ch := range choices {
		if ch != nil {
		if ch != nil {
			allChoices = append(allChoices, ch)
			allChoices = append(allChoices, ch)
			tmpChoices, tmpSeqs = Flattened(ch.Choices, ch.Sequences)
			tmpChoices, tmpSeqs = Flattened(ch.Choices, ch.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
		}
	}
	}
	for _, seq := range seqs {
	for _, seq := range seqs {
		if seq != nil {
		if seq != nil {
			allSeqs = append(allSeqs, seq)
			allSeqs = append(allSeqs, seq)
			tmpChoices, tmpSeqs = Flattened(seq.Choices, seq.Sequences)
			tmpChoices, tmpSeqs = Flattened(seq.Choices, seq.Sequences)
			allChoices = append(allChoices, tmpChoices...)
			allChoices = append(allChoices, tmpChoices...)
			allSeqs = append(allSeqs, tmpSeqs...)
			allSeqs = append(allSeqs, tmpSeqs...)
		}
		}
	}
	}
	return
	return
}
}
