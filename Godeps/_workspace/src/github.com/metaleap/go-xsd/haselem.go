package xsd
package xsd


type hasCdata struct {
type hasCdata struct {
	CDATA string `xml:",chardata"`
	CDATA string `xml:",chardata"`
}
}
type XsdtString struct{ string }


type hasElemAll struct {
type hasElemAll struct {
	All *All `xml:"all"`
	All *All `xml:"all"`
}
}


type XsdtString struct{ string }type hasElemAnnotation struct {
	Annotation *Annotation `xml:"annotation"`
	Annotation *Annotation `xml:"annotation"`
}
}


type hasElemsAny struct {
type hasElemsAny struct {
	Anys []*Any `xml:"any"`
	Anys []*Any `xml:"any"`
}
}


type hasElemsAnyAttribute struct {
type hasElemsAnyAttribute struct {
	AnyAttributes []*AnyAttribute `xml:"anyAttribute"`
	AnyAttributes []*AnyAttribute `xml:"anyAttribute"`
}
}


type hasElemsAppInfo struct {
type hasElemsAppInfo struct {
	AppInfos []*AppInfo `xml:"appinfo"`
	AppInfos []*AppInfo `xml:"appinfo"`
}
}


type hasElemsAttribute struct {
type hasElemsAttribute struct {
	Attributes []*Attribute `xml:"attribute"`
	Attributes []*Attribute `xml:"attribute"`
}
}


type hasElemsAttributeGroup struct {
type hasElemsAttributeGroup struct {
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
}
}


type hasElemChoice struct {
type hasElemChoice struct {
	Choice *Choice `xml:"choice"`
	Choice *Choice `xml:"choice"`
}
}


type hasElemsChoice struct {
type hasElemsChoice struct {
	Choices []*Choice `xml:"choice"`
	Choices []*Choice `xml:"choice"`
}
}


type hasElemComplexContent struct {
type hasElemComplexContent struct {
	ComplexContent *ComplexContent `xml:"complexContent"`
	ComplexContent *ComplexContent `xml:"complexContent"`
}
}


type hasElemComplexType struct {
type hasElemComplexType struct {
	ComplexType *ComplexType `xml:"complexType"`
	ComplexType *ComplexType `xml:"complexType"`
}
}


type hasElemsComplexType struct {
type hasElemsComplexType struct {
	ComplexTypes []*ComplexType `xml:"complexType"`
	ComplexTypes []*ComplexType `xml:"complexType"`
}
}


type hasElemsDocumentation struct {
type hasElemsDocumentation struct {
	Documentations []*Documentation `xml:"documentation"`
	Documentations []*Documentation `xml:"documentation"`
}
}


type hasElemsElement struct {
type hasElemsElement struct {
	Elements []*Element `xml:"element"`
	Elements []*Element `xml:"element"`
}
}


type hasElemsEnumeration struct {
type hasElemsEnumeration struct {
	Enumerations []*RestrictionSimpleEnumeration `xml:"enumeration"`
	Enumerations []*RestrictionSimpleEnumeration `xml:"enumeration"`
}
}


type hasElemExtensionComplexContent struct {
type hasElemExtensionComplexContent struct {
	ExtensionComplexContent *ExtensionComplexContent `xml:"extension"`
	ExtensionComplexContent *ExtensionComplexContent `xml:"extension"`
}
}


type hasElemExtensionSimpleContent struct {
type hasElemExtensionSimpleContent struct {
	ExtensionSimpleContent *ExtensionSimpleContent `xml:"extension"`
	ExtensionSimpleContent *ExtensionSimpleContent `xml:"extension"`
}
}


type hasElemField struct {
type hasElemField struct {
	Field *Field `xml:"field"`
	Field *Field `xml:"field"`
}
}


type hasElemFractionDigits struct {
type hasElemFractionDigits struct {
	FractionDigits *RestrictionSimpleFractionDigits `xml:"fractionDigits"`
	FractionDigits *RestrictionSimpleFractionDigits `xml:"fractionDigits"`
}
}


type hasElemGroup struct {
type hasElemGroup struct {
	Group *Group `xml:"group"`
	Group *Group `xml:"group"`
}
}


type hasElemsGroup struct {
type hasElemsGroup struct {
	Groups []*Group `xml:"group"`
	Groups []*Group `xml:"group"`
}
}


type hasElemsImport struct {
type hasElemsImport struct {
	Imports []*Import `xml:"import"`
	Imports []*Import `xml:"import"`
}
}


type hasElemsInclude struct {
type hasElemsInclude struct {
	Includes []*Include `xml:"include"`
	Includes []*Include `xml:"include"`
}
}


type hasElemsKey struct {
type hasElemsKey struct {
	Keys []*Key `xml:"key"`
	Keys []*Key `xml:"key"`
}
}


type hasElemKeyRef struct {
type hasElemKeyRef struct {
	KeyRef *KeyRef `xml:"keyref"`
	KeyRef *KeyRef `xml:"keyref"`
}
}


type hasElemLength struct {
type hasElemLength struct {
	Length *RestrictionSimpleLength `xml:"length"`
	Length *RestrictionSimpleLength `xml:"length"`
}
}


type hasElemList struct {
type hasElemList struct {
	List *List `xml:"list"`
	List *List `xml:"list"`
}
}


type hasElemMaxExclusive struct {
type hasElemMaxExclusive struct {
	MaxExclusive *RestrictionSimpleMaxExclusive `xml:"maxExclusive"`
	MaxExclusive *RestrictionSimpleMaxExclusive `xml:"maxExclusive"`
}
}


type hasElemMaxInclusive struct {
type hasElemMaxInclusive struct {
	MaxInclusive *RestrictionSimpleMaxInclusive `xml:"maxInclusive"`
	MaxInclusive *RestrictionSimpleMaxInclusive `xml:"maxInclusive"`
}
}


type hasElemMaxLength struct {
type hasElemMaxLength struct {
	MaxLength *RestrictionSimpleMaxLength `xml:"maxLength"`
	MaxLength *RestrictionSimpleMaxLength `xml:"maxLength"`
}
}


type hasElemMinExclusive struct {
type hasElemMinExclusive struct {
	MinExclusive *RestrictionSimpleMinExclusive `xml:"minExclusive"`
	MinExclusive *RestrictionSimpleMinExclusive `xml:"minExclusive"`
}
}


type hasElemMinInclusive struct {
type hasElemMinInclusive struct {
	MinInclusive *RestrictionSimpleMinInclusive `xml:"minInclusive"`
	MinInclusive *RestrictionSimpleMinInclusive `xml:"minInclusive"`
}
}


type hasElemMinLength struct {
type hasElemMinLength struct {
	MinLength *RestrictionSimpleMinLength `xml:"minLength"`
	MinLength *RestrictionSimpleMinLength `xml:"minLength"`
}
}


type hasElemsNotation struct {
type hasElemsNotation struct {
	Notations []*Notation `xml:"notation"`
	Notations []*Notation `xml:"notation"`
}
}


type hasElemPattern struct {
type hasElemPattern struct {
	Pattern *RestrictionSimplePattern `xml:"pattern"`
	Pattern *RestrictionSimplePattern `xml:"pattern"`
}
}


type hasElemsRedefine struct {
type hasElemsRedefine struct {
	Redefines []*Redefine `xml:"redefine"`
	Redefines []*Redefine `xml:"redefine"`
}
}


type hasElemRestrictionComplexContent struct {
type hasElemRestrictionComplexContent struct {
	RestrictionComplexContent *RestrictionComplexContent `xml:"restriction"`
	RestrictionComplexContent *RestrictionComplexContent `xml:"restriction"`
}
}


type hasElemRestrictionSimpleContent struct {
type hasElemRestrictionSimpleContent struct {
	RestrictionSimpleContent *RestrictionSimpleContent `xml:"restriction"`
	RestrictionSimpleContent *RestrictionSimpleContent `xml:"restriction"`
}
}


type hasElemRestrictionSimpleType struct {
type hasElemRestrictionSimpleType struct {
	RestrictionSimpleType *RestrictionSimpleType `xml:"restriction"`
	RestrictionSimpleType *RestrictionSimpleType `xml:"restriction"`
}
}


type hasElemSelector struct {
type hasElemSelector struct {
	Selector *Selector `xml:"selector"`
	Selector *Selector `xml:"selector"`
}
}


type hasElemSequence struct {
type hasElemSequence struct {
	Sequence *Sequence `xml:"sequence"`
	Sequence *Sequence `xml:"sequence"`
}
}


type hasElemsSequence struct {
type hasElemsSequence struct {
	Sequences []*Sequence `xml:"sequence"`
	Sequences []*Sequence `xml:"sequence"`
}
}


type hasElemSimpleContent struct {
type hasElemSimpleContent struct {
	SimpleContent *SimpleContent `xml:"simpleContent"`
	SimpleContent *SimpleContent `xml:"simpleContent"`
}
}


type hasElemsSimpleType struct {
type hasElemsSimpleType struct {
	SimpleTypes []*SimpleType `xml:"simpleType"`
	SimpleTypes []*SimpleType `xml:"simpleType"`
}
}


type hasElemTotalDigits struct {
type hasElemTotalDigits struct {
	TotalDigits *RestrictionSimpleTotalDigits `xml:"totalDigits"`
	TotalDigits *RestrictionSimpleTotalDigits `xml:"totalDigits"`
}
}


type hasElemUnion struct {
type hasElemUnion struct {
	Union *Union `xml:"union"`
	Union *Union `xml:"union"`
}
}


type hasElemUnique struct {
type hasElemUnique struct {
	Unique *Unique `xml:"unique"`
	Unique *Unique `xml:"unique"`
}
}


type hasElemWhiteSpace struct {
type hasElemWhiteSpace struct {
	WhiteSpace *RestrictionSimpleWhiteSpace `xml:"whiteSpace"`
	WhiteSpace *RestrictionSimpleWhiteSpace `xml:"whiteSpace"`
}
}
