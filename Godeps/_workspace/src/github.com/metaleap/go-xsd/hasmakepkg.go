package xsd
package xsd


func (me *hasElemAll) makePkg(bag *PkgBag) {
func (me *hasElemAll) makePkg(bag *PkgBag) {
	if me.All != nil {
	if me.All != nil {
		me.All.makePkg(bag)
		me.All.makePkg(bag)
type XsdtString struct{ string }


}
}


func (me *hasElemAnnotation) makePkg(bag *PkgBag) {
func (me *hasElemAnnotation) makePkg(bag *PkgBag) {
	if me.Annotation != nil {
	if me.Annotation != nil {
type XsdtString struct{ string }		me.Annotation.makePkg(bag)
	}
	}
}
}


func (me *hasElemsAny) makePkg(bag *PkgBag) {
func (me *hasElemsAny) makePkg(bag *PkgBag) {
	for _, any := range me.Anys {
	for _, any := range me.Anys {
		any.makePkg(bag)
		any.makePkg(bag)
	}
	}
}
}


func (me *hasElemsAnyAttribute) makePkg(bag *PkgBag) {
func (me *hasElemsAnyAttribute) makePkg(bag *PkgBag) {
	for _, aa := range me.AnyAttributes {
	for _, aa := range me.AnyAttributes {
		aa.makePkg(bag)
		aa.makePkg(bag)
	}
	}
}
}


func (me *hasElemsAppInfo) makePkg(bag *PkgBag) {
func (me *hasElemsAppInfo) makePkg(bag *PkgBag) {
	for _, ai := range me.AppInfos {
	for _, ai := range me.AppInfos {
		ai.makePkg(bag)
		ai.makePkg(bag)
	}
	}
}
}


func (me *hasElemsAttribute) makePkg(bag *PkgBag) {
func (me *hasElemsAttribute) makePkg(bag *PkgBag) {
	for _, ea := range me.Attributes {
	for _, ea := range me.Attributes {
		ea.makePkg(bag)
		ea.makePkg(bag)
	}
	}
}
}


func (me *hasElemsAttributeGroup) makePkg(bag *PkgBag) {
func (me *hasElemsAttributeGroup) makePkg(bag *PkgBag) {
	for _, ag := range me.AttributeGroups {
	for _, ag := range me.AttributeGroups {
		ag.makePkg(bag)
		ag.makePkg(bag)
	}
	}
}
}


func (me *hasElemChoice) makePkg(bag *PkgBag) {
func (me *hasElemChoice) makePkg(bag *PkgBag) {
	if me.Choice != nil {
	if me.Choice != nil {
		me.Choice.makePkg(bag)
		me.Choice.makePkg(bag)
	}
	}
}
}


func (me *hasElemsChoice) makePkg(bag *PkgBag) {
func (me *hasElemsChoice) makePkg(bag *PkgBag) {
	for _, ch := range me.Choices {
	for _, ch := range me.Choices {
		ch.makePkg(bag)
		ch.makePkg(bag)
	}
	}
}
}


func (me *hasElemComplexContent) makePkg(bag *PkgBag) {
func (me *hasElemComplexContent) makePkg(bag *PkgBag) {
	if me.ComplexContent != nil {
	if me.ComplexContent != nil {
		me.ComplexContent.makePkg(bag)
		me.ComplexContent.makePkg(bag)
	}
	}
}
}


func (me *hasElemComplexType) makePkg(bag *PkgBag) {
func (me *hasElemComplexType) makePkg(bag *PkgBag) {
	if me.ComplexType != nil {
	if me.ComplexType != nil {
		me.ComplexType.makePkg(bag)
		me.ComplexType.makePkg(bag)
	}
	}
}
}


func (me *hasElemsComplexType) makePkg(bag *PkgBag) {
func (me *hasElemsComplexType) makePkg(bag *PkgBag) {
	for _, ct := range me.ComplexTypes {
	for _, ct := range me.ComplexTypes {
		ct.makePkg(bag)
		ct.makePkg(bag)
	}
	}
}
}


func (me *hasElemsDocumentation) makePkg(bag *PkgBag) {
func (me *hasElemsDocumentation) makePkg(bag *PkgBag) {
	for _, doc := range me.Documentations {
	for _, doc := range me.Documentations {
		doc.makePkg(bag)
		doc.makePkg(bag)
	}
	}
}
}


func (me *hasElemsElement) makePkg(bag *PkgBag) {
func (me *hasElemsElement) makePkg(bag *PkgBag) {
	for _, el := range me.Elements {
	for _, el := range me.Elements {
		el.makePkg(bag)
		el.makePkg(bag)
	}
	}
}
}


func (me *hasElemsEnumeration) makePkg(bag *PkgBag) {
func (me *hasElemsEnumeration) makePkg(bag *PkgBag) {
	for _, enum := range me.Enumerations {
	for _, enum := range me.Enumerations {
		enum.makePkg(bag)
		enum.makePkg(bag)
	}
	}
}
}


func (me *hasElemExtensionComplexContent) makePkg(bag *PkgBag) {
func (me *hasElemExtensionComplexContent) makePkg(bag *PkgBag) {
	if me.ExtensionComplexContent != nil {
	if me.ExtensionComplexContent != nil {
		me.ExtensionComplexContent.makePkg(bag)
		me.ExtensionComplexContent.makePkg(bag)
	}
	}
}
}


func (me *hasElemExtensionSimpleContent) makePkg(bag *PkgBag) {
func (me *hasElemExtensionSimpleContent) makePkg(bag *PkgBag) {
	if me.ExtensionSimpleContent != nil {
	if me.ExtensionSimpleContent != nil {
		me.ExtensionSimpleContent.makePkg(bag)
		me.ExtensionSimpleContent.makePkg(bag)
	}
	}
}
}


func (me *hasElemField) makePkg(bag *PkgBag) {
func (me *hasElemField) makePkg(bag *PkgBag) {
	if me.Field != nil {
	if me.Field != nil {
		me.Field.makePkg(bag)
		me.Field.makePkg(bag)
	}
	}
}
}


func (me *hasElemFractionDigits) makePkg(bag *PkgBag) {
func (me *hasElemFractionDigits) makePkg(bag *PkgBag) {
	if me.FractionDigits != nil {
	if me.FractionDigits != nil {
		me.FractionDigits.makePkg(bag)
		me.FractionDigits.makePkg(bag)
	}
	}
}
}


func (me *hasElemGroup) makePkg(bag *PkgBag) {
func (me *hasElemGroup) makePkg(bag *PkgBag) {
	if me.Group != nil {
	if me.Group != nil {
		me.Group.makePkg(bag)
		me.Group.makePkg(bag)
	}
	}
}
}


func (me *hasElemsGroup) makePkg(bag *PkgBag) {
func (me *hasElemsGroup) makePkg(bag *PkgBag) {
	for _, gr := range me.Groups {
	for _, gr := range me.Groups {
		gr.makePkg(bag)
		gr.makePkg(bag)
	}
	}
}
}


func (me *hasElemsImport) makePkg(bag *PkgBag) {
func (me *hasElemsImport) makePkg(bag *PkgBag) {
	for _, imp := range me.Imports {
	for _, imp := range me.Imports {
		imp.makePkg(bag)
		imp.makePkg(bag)
	}
	}
}
}


func (me *hasElemsKey) makePkg(bag *PkgBag) {
func (me *hasElemsKey) makePkg(bag *PkgBag) {
	for _, k := range me.Keys {
	for _, k := range me.Keys {
		k.makePkg(bag)
		k.makePkg(bag)
	}
	}
}
}


func (me *hasElemKeyRef) makePkg(bag *PkgBag) {
func (me *hasElemKeyRef) makePkg(bag *PkgBag) {
	if me.KeyRef != nil {
	if me.KeyRef != nil {
		me.KeyRef.makePkg(bag)
		me.KeyRef.makePkg(bag)
	}
	}
}
}


func (me *hasElemLength) makePkg(bag *PkgBag) {
func (me *hasElemLength) makePkg(bag *PkgBag) {
	if me.Length != nil {
	if me.Length != nil {
		me.Length.makePkg(bag)
		me.Length.makePkg(bag)
	}
	}
}
}


func (me *hasElemList) makePkg(bag *PkgBag) {
func (me *hasElemList) makePkg(bag *PkgBag) {
	if me.List != nil {
	if me.List != nil {
		me.List.makePkg(bag)
		me.List.makePkg(bag)
	}
	}
}
}


func (me *hasElemMaxExclusive) makePkg(bag *PkgBag) {
func (me *hasElemMaxExclusive) makePkg(bag *PkgBag) {
	if me.MaxExclusive != nil {
	if me.MaxExclusive != nil {
		me.MaxExclusive.makePkg(bag)
		me.MaxExclusive.makePkg(bag)
	}
	}
}
}


func (me *hasElemMaxInclusive) makePkg(bag *PkgBag) {
func (me *hasElemMaxInclusive) makePkg(bag *PkgBag) {
	if me.MaxInclusive != nil {
	if me.MaxInclusive != nil {
		me.MaxInclusive.makePkg(bag)
		me.MaxInclusive.makePkg(bag)
	}
	}
}
}


func (me *hasElemMaxLength) makePkg(bag *PkgBag) {
func (me *hasElemMaxLength) makePkg(bag *PkgBag) {
	if me.MaxLength != nil {
	if me.MaxLength != nil {
		me.MaxLength.makePkg(bag)
		me.MaxLength.makePkg(bag)
	}
	}
}
}


func (me *hasElemMinExclusive) makePkg(bag *PkgBag) {
func (me *hasElemMinExclusive) makePkg(bag *PkgBag) {
	if me.MinExclusive != nil {
	if me.MinExclusive != nil {
		me.MinExclusive.makePkg(bag)
		me.MinExclusive.makePkg(bag)
	}
	}
}
}


func (me *hasElemMinInclusive) makePkg(bag *PkgBag) {
func (me *hasElemMinInclusive) makePkg(bag *PkgBag) {
	if me.MinInclusive != nil {
	if me.MinInclusive != nil {
		me.MinInclusive.makePkg(bag)
		me.MinInclusive.makePkg(bag)
	}
	}
}
}


func (me *hasElemMinLength) makePkg(bag *PkgBag) {
func (me *hasElemMinLength) makePkg(bag *PkgBag) {
	if me.MinLength != nil {
	if me.MinLength != nil {
		me.MinLength.makePkg(bag)
		me.MinLength.makePkg(bag)
	}
	}
}
}


func (me *hasElemsNotation) makePkg(bag *PkgBag) {
func (me *hasElemsNotation) makePkg(bag *PkgBag) {
	for _, not := range me.Notations {
	for _, not := range me.Notations {
		not.makePkg(bag)
		not.makePkg(bag)
	}
	}
}
}


func (me *hasElemPattern) makePkg(bag *PkgBag) {
func (me *hasElemPattern) makePkg(bag *PkgBag) {
	if me.Pattern != nil {
	if me.Pattern != nil {
		me.Pattern.makePkg(bag)
		me.Pattern.makePkg(bag)
	}
	}
}
}


func (me *hasElemsRedefine) makePkg(bag *PkgBag) {
func (me *hasElemsRedefine) makePkg(bag *PkgBag) {
	for _, rd := range me.Redefines {
	for _, rd := range me.Redefines {
		rd.makePkg(bag)
		rd.makePkg(bag)
	}
	}
}
}


func (me *hasElemRestrictionComplexContent) makePkg(bag *PkgBag) {
func (me *hasElemRestrictionComplexContent) makePkg(bag *PkgBag) {
	if me.RestrictionComplexContent != nil {
	if me.RestrictionComplexContent != nil {
		me.RestrictionComplexContent.makePkg(bag)
		me.RestrictionComplexContent.makePkg(bag)
	}
	}
}
}


func (me *hasElemRestrictionSimpleContent) makePkg(bag *PkgBag) {
func (me *hasElemRestrictionSimpleContent) makePkg(bag *PkgBag) {
	if me.RestrictionSimpleContent != nil {
	if me.RestrictionSimpleContent != nil {
		me.RestrictionSimpleContent.makePkg(bag)
		me.RestrictionSimpleContent.makePkg(bag)
	}
	}
}
}


func (me *hasElemRestrictionSimpleType) makePkg(bag *PkgBag) {
func (me *hasElemRestrictionSimpleType) makePkg(bag *PkgBag) {
	if me.RestrictionSimpleType != nil {
	if me.RestrictionSimpleType != nil {
		me.RestrictionSimpleType.makePkg(bag)
		me.RestrictionSimpleType.makePkg(bag)
	}
	}
}
}


func (me *hasElemSelector) makePkg(bag *PkgBag) {
func (me *hasElemSelector) makePkg(bag *PkgBag) {
	if me.Selector != nil {
	if me.Selector != nil {
		me.Selector.makePkg(bag)
		me.Selector.makePkg(bag)
	}
	}
}
}


func (me *hasElemSequence) makePkg(bag *PkgBag) {
func (me *hasElemSequence) makePkg(bag *PkgBag) {
	if me.Sequence != nil {
	if me.Sequence != nil {
		me.Sequence.makePkg(bag)
		me.Sequence.makePkg(bag)
	}
	}
}
}


func (me *hasElemsSequence) makePkg(bag *PkgBag) {
func (me *hasElemsSequence) makePkg(bag *PkgBag) {
	for _, seq := range me.Sequences {
	for _, seq := range me.Sequences {
		seq.makePkg(bag)
		seq.makePkg(bag)
	}
	}
}
}


func (me *hasElemSimpleContent) makePkg(bag *PkgBag) {
func (me *hasElemSimpleContent) makePkg(bag *PkgBag) {
	if me.SimpleContent != nil {
	if me.SimpleContent != nil {
		me.SimpleContent.makePkg(bag)
		me.SimpleContent.makePkg(bag)
	}
	}
}
}


func (me *hasElemsSimpleType) makePkg(bag *PkgBag) {
func (me *hasElemsSimpleType) makePkg(bag *PkgBag) {
	for _, st := range me.SimpleTypes {
	for _, st := range me.SimpleTypes {
		st.makePkg(bag)
		st.makePkg(bag)
	}
	}
}
}


func (me *hasElemTotalDigits) makePkg(bag *PkgBag) {
func (me *hasElemTotalDigits) makePkg(bag *PkgBag) {
	if me.TotalDigits != nil {
	if me.TotalDigits != nil {
		me.TotalDigits.makePkg(bag)
		me.TotalDigits.makePkg(bag)
	}
	}
}
}


func (me *hasElemUnion) makePkg(bag *PkgBag) {
func (me *hasElemUnion) makePkg(bag *PkgBag) {
	if me.Union != nil {
	if me.Union != nil {
		me.Union.makePkg(bag)
		me.Union.makePkg(bag)
	}
	}
}
}


func (me *hasElemUnique) makePkg(bag *PkgBag) {
func (me *hasElemUnique) makePkg(bag *PkgBag) {
	if me.Unique != nil {
	if me.Unique != nil {
		me.Unique.makePkg(bag)
		me.Unique.makePkg(bag)
	}
	}
}
}


func (me *hasElemWhiteSpace) makePkg(bag *PkgBag) {
func (me *hasElemWhiteSpace) makePkg(bag *PkgBag) {
	if me.WhiteSpace != nil {
	if me.WhiteSpace != nil {
		me.WhiteSpace.makePkg(bag)
		me.WhiteSpace.makePkg(bag)
	}
	}
}
}


func (me *hasAttrAbstract) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrAbstract) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrBase) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrBase) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrBlock) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrBlock) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrDefault) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrDefault) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrFinal) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrFinal) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrFixed) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrFixed) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrForm) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrForm) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrId) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrId) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrLang) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrLang) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMixed) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrMixed) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrName) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrName) beforeMakePkg(bag *PkgBag) {
	bag.Stacks.Name.Push(me.Name)
	bag.Stacks.Name.Push(me.Name)
}
}


func (me *hasAttrNamespace) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrNamespace) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrNillable) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrNillable) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrPublic) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrPublic) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrRef) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrRef) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrRefer) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrRefer) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSource) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrSource) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSystem) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrSystem) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrType) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrType) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrUse) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrUse) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrValue) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrValue) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrVersion) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrVersion) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrXpath) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrXpath) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrBlockDefault) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrBlockDefault) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrFinalDefault) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrFinalDefault) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrItemType) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrItemType) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMaxOccurs) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrMaxOccurs) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMemberTypes) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrMemberTypes) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMinOccurs) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrMinOccurs) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrProcessContents) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrProcessContents) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSchemaLocation) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrSchemaLocation) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSubstitutionGroup) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrSubstitutionGroup) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrTargetNamespace) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrTargetNamespace) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrAttributeFormDefault) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrAttributeFormDefault) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrElementFormDefault) beforeMakePkg(bag *PkgBag) {
func (me *hasAttrElementFormDefault) beforeMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrAbstract) afterMakePkg(bag *PkgBag) {
func (me *hasAttrAbstract) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrBase) afterMakePkg(bag *PkgBag) {
func (me *hasAttrBase) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrBlock) afterMakePkg(bag *PkgBag) {
func (me *hasAttrBlock) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrDefault) afterMakePkg(bag *PkgBag) {
func (me *hasAttrDefault) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrFinal) afterMakePkg(bag *PkgBag) {
func (me *hasAttrFinal) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrFixed) afterMakePkg(bag *PkgBag) {
func (me *hasAttrFixed) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrForm) afterMakePkg(bag *PkgBag) {
func (me *hasAttrForm) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrId) afterMakePkg(bag *PkgBag) {
func (me *hasAttrId) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrLang) afterMakePkg(bag *PkgBag) {
func (me *hasAttrLang) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMixed) afterMakePkg(bag *PkgBag) {
func (me *hasAttrMixed) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrName) afterMakePkg(bag *PkgBag) {
func (me *hasAttrName) afterMakePkg(bag *PkgBag) {
	bag.Stacks.Name.Pop()
	bag.Stacks.Name.Pop()
}
}


func (me *hasAttrNamespace) afterMakePkg(bag *PkgBag) {
func (me *hasAttrNamespace) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrNillable) afterMakePkg(bag *PkgBag) {
func (me *hasAttrNillable) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrPublic) afterMakePkg(bag *PkgBag) {
func (me *hasAttrPublic) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrRef) afterMakePkg(bag *PkgBag) {
func (me *hasAttrRef) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrRefer) afterMakePkg(bag *PkgBag) {
func (me *hasAttrRefer) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSource) afterMakePkg(bag *PkgBag) {
func (me *hasAttrSource) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSystem) afterMakePkg(bag *PkgBag) {
func (me *hasAttrSystem) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrType) afterMakePkg(bag *PkgBag) {
func (me *hasAttrType) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrUse) afterMakePkg(bag *PkgBag) {
func (me *hasAttrUse) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrValue) afterMakePkg(bag *PkgBag) {
func (me *hasAttrValue) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrVersion) afterMakePkg(bag *PkgBag) {
func (me *hasAttrVersion) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrXpath) afterMakePkg(bag *PkgBag) {
func (me *hasAttrXpath) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrBlockDefault) afterMakePkg(bag *PkgBag) {
func (me *hasAttrBlockDefault) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrFinalDefault) afterMakePkg(bag *PkgBag) {
func (me *hasAttrFinalDefault) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrItemType) afterMakePkg(bag *PkgBag) {
func (me *hasAttrItemType) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMaxOccurs) afterMakePkg(bag *PkgBag) {
func (me *hasAttrMaxOccurs) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMemberTypes) afterMakePkg(bag *PkgBag) {
func (me *hasAttrMemberTypes) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrMinOccurs) afterMakePkg(bag *PkgBag) {
func (me *hasAttrMinOccurs) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrProcessContents) afterMakePkg(bag *PkgBag) {
func (me *hasAttrProcessContents) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSchemaLocation) afterMakePkg(bag *PkgBag) {
func (me *hasAttrSchemaLocation) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrSubstitutionGroup) afterMakePkg(bag *PkgBag) {
func (me *hasAttrSubstitutionGroup) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrTargetNamespace) afterMakePkg(bag *PkgBag) {
func (me *hasAttrTargetNamespace) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrAttributeFormDefault) afterMakePkg(bag *PkgBag) {
func (me *hasAttrAttributeFormDefault) afterMakePkg(bag *PkgBag) {
}
}


func (me *hasAttrElementFormDefault) afterMakePkg(bag *PkgBag) {
func (me *hasAttrElementFormDefault) afterMakePkg(bag *PkgBag) {
}
}
