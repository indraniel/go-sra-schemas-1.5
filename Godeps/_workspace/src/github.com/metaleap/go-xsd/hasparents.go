package xsd
package xsd


func (me *hasElemAll) initChildren(p element) {
func (me *hasElemAll) initChildren(p element) {
	if me.All != nil {
	if me.All != nil {
		me.All.initElement(p)
		me.All.initElement(p)
type XsdtString struct{ string }


}
}


func (me *hasElemAnnotation) initChildren(p element) {
func (me *hasElemAnnotation) initChildren(p element) {
	if me.Annotation != nil {
	if me.Annotation != nil {
type XsdtString struct{ string }		me.Annotation.initElement(p)
	}
	}
}
}


func (me *hasElemsAny) initChildren(p element) {
func (me *hasElemsAny) initChildren(p element) {
	for _, any := range me.Anys {
	for _, any := range me.Anys {
		any.initElement(p)
		any.initElement(p)
	}
	}
}
}


func (me *hasElemsAnyAttribute) initChildren(p element) {
func (me *hasElemsAnyAttribute) initChildren(p element) {
	for _, aa := range me.AnyAttributes {
	for _, aa := range me.AnyAttributes {
		aa.initElement(p)
		aa.initElement(p)
	}
	}
}
}


func (me *hasElemsAppInfo) initChildren(p element) {
func (me *hasElemsAppInfo) initChildren(p element) {
	for _, ai := range me.AppInfos {
	for _, ai := range me.AppInfos {
		ai.initElement(p)
		ai.initElement(p)
	}
	}
}
}


func (me *hasElemsAttribute) initChildren(p element) {
func (me *hasElemsAttribute) initChildren(p element) {
	for _, ea := range me.Attributes {
	for _, ea := range me.Attributes {
		ea.initElement(p)
		ea.initElement(p)
	}
	}
}
}


func (me *hasElemsAttributeGroup) initChildren(p element) {
func (me *hasElemsAttributeGroup) initChildren(p element) {
	for _, ag := range me.AttributeGroups {
	for _, ag := range me.AttributeGroups {
		ag.initElement(p)
		ag.initElement(p)
	}
	}
}
}


func (me *hasElemChoice) initChildren(p element) {
func (me *hasElemChoice) initChildren(p element) {
	if me.Choice != nil {
	if me.Choice != nil {
		me.Choice.initElement(p)
		me.Choice.initElement(p)
	}
	}
}
}


func (me *hasElemsChoice) initChildren(p element) {
func (me *hasElemsChoice) initChildren(p element) {
	for _, ch := range me.Choices {
	for _, ch := range me.Choices {
		ch.initElement(p)
		ch.initElement(p)
	}
	}
}
}


func (me *hasElemComplexContent) initChildren(p element) {
func (me *hasElemComplexContent) initChildren(p element) {
	if me.ComplexContent != nil {
	if me.ComplexContent != nil {
		me.ComplexContent.initElement(p)
		me.ComplexContent.initElement(p)
	}
	}
}
}


func (me *hasElemComplexType) initChildren(p element) {
func (me *hasElemComplexType) initChildren(p element) {
	if me.ComplexType != nil {
	if me.ComplexType != nil {
		me.ComplexType.initElement(p)
		me.ComplexType.initElement(p)
	}
	}
}
}


func (me *hasElemsComplexType) initChildren(p element) {
func (me *hasElemsComplexType) initChildren(p element) {
	for _, ct := range me.ComplexTypes {
	for _, ct := range me.ComplexTypes {
		ct.initElement(p)
		ct.initElement(p)
	}
	}
}
}


func (me *hasElemsDocumentation) initChildren(p element) {
func (me *hasElemsDocumentation) initChildren(p element) {
	for _, doc := range me.Documentations {
	for _, doc := range me.Documentations {
		doc.initElement(p)
		doc.initElement(p)
	}
	}
}
}


func (me *hasElemsElement) initChildren(p element) {
func (me *hasElemsElement) initChildren(p element) {
	for _, el := range me.Elements {
	for _, el := range me.Elements {
		el.initElement(p)
		el.initElement(p)
	}
	}
}
}


func (me *hasElemsEnumeration) initChildren(p element) {
func (me *hasElemsEnumeration) initChildren(p element) {
	for _, enum := range me.Enumerations {
	for _, enum := range me.Enumerations {
		enum.initElement(p)
		enum.initElement(p)
	}
	}
}
}


func (me *hasElemExtensionComplexContent) initChildren(p element) {
func (me *hasElemExtensionComplexContent) initChildren(p element) {
	if me.ExtensionComplexContent != nil {
	if me.ExtensionComplexContent != nil {
		me.ExtensionComplexContent.initElement(p)
		me.ExtensionComplexContent.initElement(p)
	}
	}
}
}


func (me *hasElemExtensionSimpleContent) initChildren(p element) {
func (me *hasElemExtensionSimpleContent) initChildren(p element) {
	if me.ExtensionSimpleContent != nil {
	if me.ExtensionSimpleContent != nil {
		me.ExtensionSimpleContent.initElement(p)
		me.ExtensionSimpleContent.initElement(p)
	}
	}
}
}


func (me *hasElemField) initChildren(p element) {
func (me *hasElemField) initChildren(p element) {
	if me.Field != nil {
	if me.Field != nil {
		me.Field.initElement(p)
		me.Field.initElement(p)
	}
	}
}
}


func (me *hasElemFractionDigits) initChildren(p element) {
func (me *hasElemFractionDigits) initChildren(p element) {
	if me.FractionDigits != nil {
	if me.FractionDigits != nil {
		me.FractionDigits.initElement(p)
		me.FractionDigits.initElement(p)
	}
	}
}
}


func (me *hasElemGroup) initChildren(p element) {
func (me *hasElemGroup) initChildren(p element) {
	if me.Group != nil {
	if me.Group != nil {
		me.Group.initElement(p)
		me.Group.initElement(p)
	}
	}
}
}


func (me *hasElemsGroup) initChildren(p element) {
func (me *hasElemsGroup) initChildren(p element) {
	for _, gr := range me.Groups {
	for _, gr := range me.Groups {
		gr.initElement(p)
		gr.initElement(p)
	}
	}
}
}


func (me *hasElemsImport) initChildren(p element) {
func (me *hasElemsImport) initChildren(p element) {
	for _, imp := range me.Imports {
	for _, imp := range me.Imports {
		imp.initElement(p)
		imp.initElement(p)
	}
	}
}
}


func (me *hasElemsKey) initChildren(p element) {
func (me *hasElemsKey) initChildren(p element) {
	for _, k := range me.Keys {
	for _, k := range me.Keys {
		k.initElement(p)
		k.initElement(p)
	}
	}
}
}


func (me *hasElemKeyRef) initChildren(p element) {
func (me *hasElemKeyRef) initChildren(p element) {
	if me.KeyRef != nil {
	if me.KeyRef != nil {
		me.KeyRef.initElement(p)
		me.KeyRef.initElement(p)
	}
	}
}
}


func (me *hasElemLength) initChildren(p element) {
func (me *hasElemLength) initChildren(p element) {
	if me.Length != nil {
	if me.Length != nil {
		me.Length.initElement(p)
		me.Length.initElement(p)
	}
	}
}
}


func (me *hasElemList) initChildren(p element) {
func (me *hasElemList) initChildren(p element) {
	if me.List != nil {
	if me.List != nil {
		me.List.initElement(p)
		me.List.initElement(p)
	}
	}
}
}


func (me *hasElemMaxExclusive) initChildren(p element) {
func (me *hasElemMaxExclusive) initChildren(p element) {
	if me.MaxExclusive != nil {
	if me.MaxExclusive != nil {
		me.MaxExclusive.initElement(p)
		me.MaxExclusive.initElement(p)
	}
	}
}
}


func (me *hasElemMaxInclusive) initChildren(p element) {
func (me *hasElemMaxInclusive) initChildren(p element) {
	if me.MaxInclusive != nil {
	if me.MaxInclusive != nil {
		me.MaxInclusive.initElement(p)
		me.MaxInclusive.initElement(p)
	}
	}
}
}


func (me *hasElemMaxLength) initChildren(p element) {
func (me *hasElemMaxLength) initChildren(p element) {
	if me.MaxLength != nil {
	if me.MaxLength != nil {
		me.MaxLength.initElement(p)
		me.MaxLength.initElement(p)
	}
	}
}
}


func (me *hasElemMinExclusive) initChildren(p element) {
func (me *hasElemMinExclusive) initChildren(p element) {
	if me.MinExclusive != nil {
	if me.MinExclusive != nil {
		me.MinExclusive.initElement(p)
		me.MinExclusive.initElement(p)
	}
	}
}
}


func (me *hasElemMinInclusive) initChildren(p element) {
func (me *hasElemMinInclusive) initChildren(p element) {
	if me.MinInclusive != nil {
	if me.MinInclusive != nil {
		me.MinInclusive.initElement(p)
		me.MinInclusive.initElement(p)
	}
	}
}
}


func (me *hasElemMinLength) initChildren(p element) {
func (me *hasElemMinLength) initChildren(p element) {
	if me.MinLength != nil {
	if me.MinLength != nil {
		me.MinLength.initElement(p)
		me.MinLength.initElement(p)
	}
	}
}
}


func (me *hasElemsNotation) initChildren(p element) {
func (me *hasElemsNotation) initChildren(p element) {
	for _, not := range me.Notations {
	for _, not := range me.Notations {
		not.initElement(p)
		not.initElement(p)
	}
	}
}
}


func (me *hasElemPattern) initChildren(p element) {
func (me *hasElemPattern) initChildren(p element) {
	if me.Pattern != nil {
	if me.Pattern != nil {
		me.Pattern.initElement(p)
		me.Pattern.initElement(p)
	}
	}
}
}


func (me *hasElemsRedefine) initChildren(p element) {
func (me *hasElemsRedefine) initChildren(p element) {
	for _, rd := range me.Redefines {
	for _, rd := range me.Redefines {
		rd.initElement(p)
		rd.initElement(p)
	}
	}
}
}


func (me *hasElemRestrictionComplexContent) initChildren(p element) {
func (me *hasElemRestrictionComplexContent) initChildren(p element) {
	if me.RestrictionComplexContent != nil {
	if me.RestrictionComplexContent != nil {
		me.RestrictionComplexContent.initElement(p)
		me.RestrictionComplexContent.initElement(p)
	}
	}
}
}


func (me *hasElemRestrictionSimpleContent) initChildren(p element) {
func (me *hasElemRestrictionSimpleContent) initChildren(p element) {
	if me.RestrictionSimpleContent != nil {
	if me.RestrictionSimpleContent != nil {
		me.RestrictionSimpleContent.initElement(p)
		me.RestrictionSimpleContent.initElement(p)
	}
	}
}
}


func (me *hasElemRestrictionSimpleType) initChildren(p element) {
func (me *hasElemRestrictionSimpleType) initChildren(p element) {
	if me.RestrictionSimpleType != nil {
	if me.RestrictionSimpleType != nil {
		me.RestrictionSimpleType.initElement(p)
		me.RestrictionSimpleType.initElement(p)
	}
	}
}
}


func (me *hasElemSelector) initChildren(p element) {
func (me *hasElemSelector) initChildren(p element) {
	if me.Selector != nil {
	if me.Selector != nil {
		me.Selector.initElement(p)
		me.Selector.initElement(p)
	}
	}
}
}


func (me *hasElemSequence) initChildren(p element) {
func (me *hasElemSequence) initChildren(p element) {
	if me.Sequence != nil {
	if me.Sequence != nil {
		me.Sequence.initElement(p)
		me.Sequence.initElement(p)
	}
	}
}
}


func (me *hasElemsSequence) initChildren(p element) {
func (me *hasElemsSequence) initChildren(p element) {
	for _, seq := range me.Sequences {
	for _, seq := range me.Sequences {
		seq.initElement(p)
		seq.initElement(p)
	}
	}
}
}


func (me *hasElemSimpleContent) initChildren(p element) {
func (me *hasElemSimpleContent) initChildren(p element) {
	if me.SimpleContent != nil {
	if me.SimpleContent != nil {
		me.SimpleContent.initElement(p)
		me.SimpleContent.initElement(p)
	}
	}
}
}


func (me *hasElemsSimpleType) initChildren(p element) {
func (me *hasElemsSimpleType) initChildren(p element) {
	for _, st := range me.SimpleTypes {
	for _, st := range me.SimpleTypes {
		st.initElement(p)
		st.initElement(p)
	}
	}
}
}


func (me *hasElemTotalDigits) initChildren(p element) {
func (me *hasElemTotalDigits) initChildren(p element) {
	if me.TotalDigits != nil {
	if me.TotalDigits != nil {
		me.TotalDigits.initElement(p)
		me.TotalDigits.initElement(p)
	}
	}
}
}


func (me *hasElemUnion) initChildren(p element) {
func (me *hasElemUnion) initChildren(p element) {
	if me.Union != nil {
	if me.Union != nil {
		me.Union.initElement(p)
		me.Union.initElement(p)
	}
	}
}
}


func (me *hasElemUnique) initChildren(p element) {
func (me *hasElemUnique) initChildren(p element) {
	if me.Unique != nil {
	if me.Unique != nil {
		me.Unique.initElement(p)
		me.Unique.initElement(p)
	}
	}
}
}


func (me *hasElemWhiteSpace) initChildren(p element) {
func (me *hasElemWhiteSpace) initChildren(p element) {
	if me.WhiteSpace != nil {
	if me.WhiteSpace != nil {
		me.WhiteSpace.initElement(p)
		me.WhiteSpace.initElement(p)
	}
	}
}
}
