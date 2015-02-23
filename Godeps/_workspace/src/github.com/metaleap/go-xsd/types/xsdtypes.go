package xsdt
package xsdt


import (
import (
	"strconv"
	"strconv"
)
)
type XsdtString struct{ string }


type notation struct {
type notation struct {
	Id, Name, Public, System string
	Id, Name, Public, System string
}
}


type XsdtString struct{ string }type Notations map[string]*notation


func (me Notations) Add(id, name, public, system string) {
func (me Notations) Add(id, name, public, system string) {
	me[name] = &notation{Id: id, Name: name, Public: public, System: system}
	me[name] = &notation{Id: id, Name: name, Public: public, System: system}
}
}


//	In XSD, the type xsd:anySimpleType is the base type from which all other built-in types are derived.
//	In XSD, the type xsd:anySimpleType is the base type from which all other built-in types are derived.
type AnySimpleType string
type AnySimpleType string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *AnySimpleType) Set(v string) {
func (me *AnySimpleType) Set(v string) {
	*me = AnySimpleType(v)
	*me = AnySimpleType(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me AnySimpleType) String() string {
func (me AnySimpleType) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to AnySimpleType.
//	A convenience interface that declares a type conversion to AnySimpleType.
type ToXsdtAnySimpleType interface {
type ToXsdtAnySimpleType interface {
	ToXsdtAnySimpleType() AnySimpleType
	ToXsdtAnySimpleType() AnySimpleType
}
}


//	In XSD, represents any simple or complex type. In Go, we hope no one schema ever uses it.
//	In XSD, represents any simple or complex type. In Go, we hope no one schema ever uses it.
type AnyType string
type AnyType string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *AnyType) Set(v string) {
func (me *AnyType) Set(v string) {
	*me = AnyType(v)
	*me = AnyType(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me AnyType) String() string {
func (me AnyType) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to AnyType.
//	A convenience interface that declares a type conversion to AnyType.
type ToXsdtAnyType interface {
type ToXsdtAnyType interface {
	ToXsdtAnyType() AnyType
	ToXsdtAnyType() AnyType
}
}


//	Represents a URI as defined by RFC 2396. An anyURI value can be absolute or relative, and may have an optional fragment identifier.
//	Represents a URI as defined by RFC 2396. An anyURI value can be absolute or relative, and may have an optional fragment identifier.
type AnyURI string
type AnyURI string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *AnyURI) Set(v string) {
func (me *AnyURI) Set(v string) {
	*me = AnyURI(v)
	*me = AnyURI(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me AnyURI) String() string {
func (me AnyURI) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to AnyURI.
//	A convenience interface that declares a type conversion to AnyURI.
type ToXsdtAnyURI interface {
type ToXsdtAnyURI interface {
	ToXsdtAnyURI() AnyURI
	ToXsdtAnyURI() AnyURI
}
}


//	Represents Base64-encoded arbitrary binary data. A base64Binary is the set of finite-length sequences of binary octets.
//	Represents Base64-encoded arbitrary binary data. A base64Binary is the set of finite-length sequences of binary octets.
type Base64Binary string // []byte
type Base64Binary string // []byte


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Base64Binary) Set(v string) {
func (me *Base64Binary) Set(v string) {
	*me = Base64Binary(v)
	*me = Base64Binary(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Base64Binary) String() string {
func (me Base64Binary) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Base64Binary.
//	A convenience interface that declares a type conversion to Base64Binary.
type ToXsdtBase64Binary interface {
type ToXsdtBase64Binary interface {
	ToXsdtBase64Binary() Base64Binary
	ToXsdtBase64Binary() Base64Binary
}
}


//	Represents Boolean values, which are either true or false.
//	Represents Boolean values, which are either true or false.
type Boolean bool
type Boolean bool


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Boolean) B() bool {
func (me Boolean) B() bool {
	return bool(me)
	return bool(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Boolean) Set(v string) {
func (me *Boolean) Set(v string) {
	//	most schemas use true and false but sadly, a very few rare ones *do* use "0" and "1"...
	//	most schemas use true and false but sadly, a very few rare ones *do* use "0" and "1"...
	switch v {
	switch v {
	case "0":
	case "0":
		*me = false
		*me = false
	case "1":
	case "1":
		*me = true
		*me = true
	default:
	default:
		b, _ := strconv.ParseBool(v)
		b, _ := strconv.ParseBool(v)
		*me = Boolean(b)
		*me = Boolean(b)
	}
	}
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Boolean) String() string {
func (me Boolean) String() string {
	return strconv.FormatBool(bool(me))
	return strconv.FormatBool(bool(me))
}
}


//	A convenience interface that declares a type conversion to Boolean.
//	A convenience interface that declares a type conversion to Boolean.
type ToXsdtBoolean interface {
type ToXsdtBoolean interface {
	ToXsdtBoolean() Boolean
	ToXsdtBoolean() Boolean
}
}


//	Represents an integer with a minimum value of -128 and maximum of 127.
//	Represents an integer with a minimum value of -128 and maximum of 127.
type Byte int8
type Byte int8


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Byte) N() int8 {
func (me Byte) N() int8 {
	return int8(me)
	return int8(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Byte) Set(s string) {
func (me *Byte) Set(s string) {
	v, _ := strconv.ParseInt(s, 0, 8)
	v, _ := strconv.ParseInt(s, 0, 8)
	*me = Byte(v)
	*me = Byte(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Byte) String() string {
func (me Byte) String() string {
	return strconv.FormatInt(int64(me), 10)
	return strconv.FormatInt(int64(me), 10)
}
}


//	A convenience interface that declares a type conversion to Byte.
//	A convenience interface that declares a type conversion to Byte.
type ToXsdtByte interface {
type ToXsdtByte interface {
	ToXsdtByte() Byte
	ToXsdtByte() Byte
}
}


//	Represents a calendar date.
//	Represents a calendar date.
//	The pattern for date is CCYY-MM-DD with optional time zone indicator as allowed for dateTime.
//	The pattern for date is CCYY-MM-DD with optional time zone indicator as allowed for dateTime.
type Date string // time.Time
type Date string // time.Time


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Date) Set(v string) {
func (me *Date) Set(v string) {
	*me = Date(v)
	*me = Date(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Date) String() string {
func (me Date) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Date.
//	A convenience interface that declares a type conversion to Date.
type ToXsdtDate interface {
type ToXsdtDate interface {
	ToXsdtDate() Date
	ToXsdtDate() Date
}
}


//	Represents a specific instance of time.
//	Represents a specific instance of time.
type DateTime string // time.Time
type DateTime string // time.Time


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *DateTime) Set(v string) {
func (me *DateTime) Set(v string) {
	*me = DateTime(v)
	*me = DateTime(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me DateTime) String() string {
func (me DateTime) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to DateTime.
//	A convenience interface that declares a type conversion to DateTime.
type ToXsdtDateTime interface {
type ToXsdtDateTime interface {
	ToXsdtDateTime() DateTime
	ToXsdtDateTime() DateTime
}
}


//	Represents a specific instance of time.
//	Represents a specific instance of time.
type Time string // time.Time
type Time string // time.Time


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Time) Set(v string) {
func (me *Time) Set(v string) {
	*me = Time(v)
	*me = Time(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Time) String() string {
func (me Time) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Time.
//	A convenience interface that declares a type conversion to Time.
type ToXsdtTime interface {
type ToXsdtTime interface {
	ToXsdtTime() Time
	ToXsdtTime() Time
}
}


//	Represents arbitrary precision numbers.
//	Represents arbitrary precision numbers.
type Decimal string // complex128
type Decimal string // complex128


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Decimal) Set(v string) {
func (me *Decimal) Set(v string) {
	*me = Decimal(v)
	*me = Decimal(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Decimal) String() string {
func (me Decimal) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Decimal.
//	A convenience interface that declares a type conversion to Decimal.
type ToXsdtDecimal interface {
type ToXsdtDecimal interface {
	ToXsdtDecimal() Decimal
	ToXsdtDecimal() Decimal
}
}


//	Represents double-precision 64-bit floating-point numbers.
//	Represents double-precision 64-bit floating-point numbers.
type Double float64
type Double float64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Double) N() float64 {
func (me Double) N() float64 {
	return float64(me)
	return float64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Double) Set(s string) {
func (me *Double) Set(s string) {
	v, _ := strconv.ParseFloat(s, 64)
	v, _ := strconv.ParseFloat(s, 64)
	*me = Double(v)
	*me = Double(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Double) String() string {
func (me Double) String() string {
	return strconv.FormatFloat(float64(me), 'f', 8, 64)
	return strconv.FormatFloat(float64(me), 'f', 8, 64)
}
}


//	A convenience interface that declares a type conversion to Double.
//	A convenience interface that declares a type conversion to Double.
type ToXsdtDouble interface {
type ToXsdtDouble interface {
	ToXsdtDouble() Double
	ToXsdtDouble() Double
}
}


//	Represents a duration of time.
//	Represents a duration of time.
type Duration string // time.Duration
type Duration string // time.Duration


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Duration) Set(v string) {
func (me *Duration) Set(v string) {
	*me = Duration(v)
	*me = Duration(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Duration) String() string {
func (me Duration) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Duration.
//	A convenience interface that declares a type conversion to Duration.
type ToXsdtDuration interface {
type ToXsdtDuration interface {
	ToXsdtDuration() Duration
	ToXsdtDuration() Duration
}
}


//	Represents the ENTITIES attribute type. Contains a set of values of type ENTITY.
//	Represents the ENTITIES attribute type. Contains a set of values of type ENTITY.
type Entities string
type Entities string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Entities) Set(v string) {
func (me *Entities) Set(v string) {
	*me = Entities(v)
	*me = Entities(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Entities) String() string {
func (me Entities) String() string {
	return string(me)
	return string(me)
}
}


//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
func (me Entities) Values() (list []Entity) {
func (me Entities) Values() (list []Entity) {
	spl := ListValues(string(me))
	spl := ListValues(string(me))
	list = make([]Entity, len(spl))
	list = make([]Entity, len(spl))
	for i, s := range spl {
	for i, s := range spl {
		list[i].Set(s)
		list[i].Set(s)
	}
	}
	return
	return
}
}


//	A convenience interface that declares a type conversion to Entities.
//	A convenience interface that declares a type conversion to Entities.
type ToXsdtEntities interface {
type ToXsdtEntities interface {
	ToXsdtEntities() Entities
	ToXsdtEntities() Entities
}
}


//	This is a reference to an unparsed entity with a name that matches the specified name.
//	This is a reference to an unparsed entity with a name that matches the specified name.
type Entity NCName
type Entity NCName


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Entity) Set(v string) {
func (me *Entity) Set(v string) {
	*me = Entity(v)
	*me = Entity(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Entity) String() string {
func (me Entity) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Entity.
//	A convenience interface that declares a type conversion to Entity.
type ToXsdtEntity interface {
type ToXsdtEntity interface {
	ToXsdtEntity() Entity
	ToXsdtEntity() Entity
}
}


//	Represents single-precision 32-bit floating-point numbers.
//	Represents single-precision 32-bit floating-point numbers.
type Float float32
type Float float32


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Float) N() float32 {
func (me Float) N() float32 {
	return float32(me)
	return float32(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Float) Set(s string) {
func (me *Float) Set(s string) {
	v, _ := strconv.ParseFloat(s, 32)
	v, _ := strconv.ParseFloat(s, 32)
	*me = Float(v)
	*me = Float(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Float) String() string {
func (me Float) String() string {
	return strconv.FormatFloat(float64(me), 'f', 8, 32)
	return strconv.FormatFloat(float64(me), 'f', 8, 32)
}
}


//	A convenience interface that declares a type conversion to Float.
//	A convenience interface that declares a type conversion to Float.
type ToXsdtFloat interface {
type ToXsdtFloat interface {
	ToXsdtFloat() Float
	ToXsdtFloat() Float
}
}


//	Represents a Gregorian day that recurs, specifically a day of the month such as the fifth day of the month. A gDay is the space of a set of calendar dates. Specifically, it is a set of one-day long, monthly periodic instances.
//	Represents a Gregorian day that recurs, specifically a day of the month such as the fifth day of the month. A gDay is the space of a set of calendar dates. Specifically, it is a set of one-day long, monthly periodic instances.
type GDay string
type GDay string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *GDay) Set(v string) {
func (me *GDay) Set(v string) {
	*me = GDay(v)
	*me = GDay(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me GDay) String() string {
func (me GDay) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to GDay.
//	A convenience interface that declares a type conversion to GDay.
type ToXsdtGDay interface {
type ToXsdtGDay interface {
	ToXsdtGDay() GDay
	ToXsdtGDay() GDay
}
}


//	Represents a Gregorian month that recurs every year. A gMonth is the space of a set of calendar months. Specifically, it is a set of one-month long, yearly periodic instances.
//	Represents a Gregorian month that recurs every year. A gMonth is the space of a set of calendar months. Specifically, it is a set of one-month long, yearly periodic instances.
type GMonth string
type GMonth string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *GMonth) Set(v string) {
func (me *GMonth) Set(v string) {
	*me = GMonth(v)
	*me = GMonth(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me GMonth) String() string {
func (me GMonth) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to GMonth.
//	A convenience interface that declares a type conversion to GMonth.
type ToXsdtGMonth interface {
type ToXsdtGMonth interface {
	ToXsdtGMonth() GMonth
	ToXsdtGMonth() GMonth
}
}


//	Represents a specific Gregorian date that recurs, specifically a day of the year such as the third of May. A gMonthDay is the set of calendar dates. Specifically, it is a set of one-day long, annually periodic instances.
//	Represents a specific Gregorian date that recurs, specifically a day of the year such as the third of May. A gMonthDay is the set of calendar dates. Specifically, it is a set of one-day long, annually periodic instances.
type GMonthDay string
type GMonthDay string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *GMonthDay) Set(v string) {
func (me *GMonthDay) Set(v string) {
	*me = GMonthDay(v)
	*me = GMonthDay(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me GMonthDay) String() string {
func (me GMonthDay) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to GMonthDay.
//	A convenience interface that declares a type conversion to GMonthDay.
type ToXsdtGMonthDay interface {
type ToXsdtGMonthDay interface {
	ToXsdtGMonthDay() GMonthDay
	ToXsdtGMonthDay() GMonthDay
}
}


//	Represents a Gregorian year. A set of one-year long, nonperiodic instances.
//	Represents a Gregorian year. A set of one-year long, nonperiodic instances.
type GYear string
type GYear string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *GYear) Set(v string) {
func (me *GYear) Set(v string) {
	*me = GYear(v)
	*me = GYear(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me GYear) String() string {
func (me GYear) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to GYear.
//	A convenience interface that declares a type conversion to GYear.
type ToXsdtGYear interface {
type ToXsdtGYear interface {
	ToXsdtGYear() GYear
	ToXsdtGYear() GYear
}
}


//	Represents a specific Gregorian month in a specific Gregorian year. A set of one-month long, nonperiodic instances.
//	Represents a specific Gregorian month in a specific Gregorian year. A set of one-month long, nonperiodic instances.
type GYearMonth string
type GYearMonth string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *GYearMonth) Set(v string) {
func (me *GYearMonth) Set(v string) {
	*me = GYearMonth(v)
	*me = GYearMonth(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me GYearMonth) String() string {
func (me GYearMonth) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to GYearMonth.
//	A convenience interface that declares a type conversion to GYearMonth.
type ToXsdtGYearMonth interface {
type ToXsdtGYearMonth interface {
	ToXsdtGYearMonth() GYearMonth
	ToXsdtGYearMonth() GYearMonth
}
}


//	Represents arbitrary hex-encoded binary data. A hexBinary is the set of finite-length sequences of binary octets. Each binary octet is encoded as a character tuple, consisting of two hexadecimal digits ([0-9a-fA-F]) representing the octet code.
//	Represents arbitrary hex-encoded binary data. A hexBinary is the set of finite-length sequences of binary octets. Each binary octet is encoded as a character tuple, consisting of two hexadecimal digits ([0-9a-fA-F]) representing the octet code.
type HexBinary string // []byte
type HexBinary string // []byte


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *HexBinary) Set(v string) {
func (me *HexBinary) Set(v string) {
	*me = HexBinary(v)
	*me = HexBinary(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me HexBinary) String() string {
func (me HexBinary) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to HexBinary.
//	A convenience interface that declares a type conversion to HexBinary.
type ToXsdtHexBinary interface {
type ToXsdtHexBinary interface {
	ToXsdtHexBinary() HexBinary
	ToXsdtHexBinary() HexBinary
}
}


//	The ID must be a no-colon-name (NCName) and must be unique within an XML document.
//	The ID must be a no-colon-name (NCName) and must be unique within an XML document.
type Id NCName
type Id NCName


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Id) Set(v string) {
func (me *Id) Set(v string) {
	*me = Id(v)
	*me = Id(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Id) String() string {
func (me Id) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Id.
//	A convenience interface that declares a type conversion to Id.
type ToXsdtId interface {
type ToXsdtId interface {
	ToXsdtId() Id
	ToXsdtId() Id
}
}


//	Represents a reference to an element that has an ID attribute that matches the specified ID. An IDREF must be an NCName and must be a value of an element or attribute of type ID within the XML document.
//	Represents a reference to an element that has an ID attribute that matches the specified ID. An IDREF must be an NCName and must be a value of an element or attribute of type ID within the XML document.
type Idref NCName
type Idref NCName


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Idref) Set(v string) {
func (me *Idref) Set(v string) {
	*me = Idref(v)
	*me = Idref(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Idref) String() string {
func (me Idref) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Idref.
//	A convenience interface that declares a type conversion to Idref.
type ToXsdtIdref interface {
type ToXsdtIdref interface {
	ToXsdtIdref() Idref
	ToXsdtIdref() Idref
}
}


//	Contains a set of values of type IDREF.
//	Contains a set of values of type IDREF.
type Idrefs string
type Idrefs string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Idrefs) Set(v string) {
func (me *Idrefs) Set(v string) {
	*me = Idrefs(v)
	*me = Idrefs(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Idrefs) String() string {
func (me Idrefs) String() string {
	return string(me)
	return string(me)
}
}


//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
func (me Idrefs) Values() (list []Idref) {
func (me Idrefs) Values() (list []Idref) {
	spl := ListValues(string(me))
	spl := ListValues(string(me))
	list = make([]Idref, len(spl))
	list = make([]Idref, len(spl))
	for i, s := range spl {
	for i, s := range spl {
		list[i].Set(s)
		list[i].Set(s)
	}
	}
	return
	return
}
}


//	A convenience interface that declares a type conversion to Idrefs.
//	A convenience interface that declares a type conversion to Idrefs.
type ToXsdtIdrefs interface {
type ToXsdtIdrefs interface {
	ToXsdtIdrefs() Idrefs
	ToXsdtIdrefs() Idrefs
}
}


//	Represents an integer with a minimum value of -2147483648 and maximum of 2147483647.
//	Represents an integer with a minimum value of -2147483648 and maximum of 2147483647.
type Int int32
type Int int32


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Int) N() int32 {
func (me Int) N() int32 {
	return int32(me)
	return int32(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Int) Set(s string) {
func (me *Int) Set(s string) {
	v, _ := strconv.ParseInt(s, 0, 32)
	v, _ := strconv.ParseInt(s, 0, 32)
	*me = Int(v)
	*me = Int(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Int) String() string {
func (me Int) String() string {
	return strconv.FormatInt(int64(me), 10)
	return strconv.FormatInt(int64(me), 10)
}
}


//	A convenience interface that declares a type conversion to Int.
//	A convenience interface that declares a type conversion to Int.
type ToXsdtInt interface {
type ToXsdtInt interface {
	ToXsdtInt() Int
	ToXsdtInt() Int
}
}


//	Represents a sequence of decimal digits with an optional leading sign (+ or -). 
//	Represents a sequence of decimal digits with an optional leading sign (+ or -). 
type Integer int64
type Integer int64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Integer) N() int64 {
func (me Integer) N() int64 {
	return int64(me)
	return int64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Integer) Set(s string) {
func (me *Integer) Set(s string) {
	v, _ := strconv.ParseInt(s, 0, 64)
	v, _ := strconv.ParseInt(s, 0, 64)
	*me = Integer(v)
	*me = Integer(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Integer) String() string {
func (me Integer) String() string {
	return strconv.FormatInt(int64(me), 10)
	return strconv.FormatInt(int64(me), 10)
}
}


//	A convenience interface that declares a type conversion to Integer.
//	A convenience interface that declares a type conversion to Integer.
type ToXsdtInteger interface {
type ToXsdtInteger interface {
	ToXsdtInteger() Integer
	ToXsdtInteger() Integer
}
}


//	Represents natural language identifiers (defined by RFC 1766).
//	Represents natural language identifiers (defined by RFC 1766).
type Language Token
type Language Token


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Language) Set(v string) {
func (me *Language) Set(v string) {
	*me = Language(v)
	*me = Language(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Language) String() string {
func (me Language) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Language.
//	A convenience interface that declares a type conversion to Language.
type ToXsdtLanguage interface {
type ToXsdtLanguage interface {
	ToXsdtLanguage() Language
	ToXsdtLanguage() Language
}
}


//	Represents an integer with a minimum value of -9223372036854775808 and maximum of 9223372036854775807.
//	Represents an integer with a minimum value of -9223372036854775808 and maximum of 9223372036854775807.
type Long int64
type Long int64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Long) N() int64 {
func (me Long) N() int64 {
	return int64(me)
	return int64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Long) Set(s string) {
func (me *Long) Set(s string) {
	v, _ := strconv.ParseInt(s, 0, 64)
	v, _ := strconv.ParseInt(s, 0, 64)
	*me = Long(v)
	*me = Long(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Long) String() string {
func (me Long) String() string {
	return strconv.FormatInt(int64(me), 10)
	return strconv.FormatInt(int64(me), 10)
}
}


//	A convenience interface that declares a type conversion to Long.
//	A convenience interface that declares a type conversion to Long.
type ToXsdtLong interface {
type ToXsdtLong interface {
	ToXsdtLong() Long
	ToXsdtLong() Long
}
}


//	Represents names in XML. A Name is a token that begins with a letter, underscore, or colon and continues with name characters (letters, digits, and other characters).
//	Represents names in XML. A Name is a token that begins with a letter, underscore, or colon and continues with name characters (letters, digits, and other characters).
type Name Token
type Name Token


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Name) Set(v string) {
func (me *Name) Set(v string) {
	*me = Name(v)
	*me = Name(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Name) String() string {
func (me Name) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Name.
//	A convenience interface that declares a type conversion to Name.
type ToXsdtName interface {
type ToXsdtName interface {
	ToXsdtName() Name
	ToXsdtName() Name
}
}


//	Represents noncolonized names. This data type is the same as Name, except it cannot begin with a colon.
//	Represents noncolonized names. This data type is the same as Name, except it cannot begin with a colon.
type NCName Name
type NCName Name


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *NCName) Set(v string) {
func (me *NCName) Set(v string) {
	*me = NCName(v)
	*me = NCName(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me NCName) String() string {
func (me NCName) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to NCName.
//	A convenience interface that declares a type conversion to NCName.
type ToXsdtNCName interface {
type ToXsdtNCName interface {
	ToXsdtNCName() NCName
	ToXsdtNCName() NCName
}
}


//	Represents an integer that is less than zero. Consists of a negative sign (-) and sequence of decimal digits.
//	Represents an integer that is less than zero. Consists of a negative sign (-) and sequence of decimal digits.
type NegativeInteger int64
type NegativeInteger int64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me NegativeInteger) N() int64 {
func (me NegativeInteger) N() int64 {
	return int64(me)
	return int64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *NegativeInteger) Set(s string) {
func (me *NegativeInteger) Set(s string) {
	v, _ := strconv.ParseInt(s, 0, 64)
	v, _ := strconv.ParseInt(s, 0, 64)
	*me = NegativeInteger(v)
	*me = NegativeInteger(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me NegativeInteger) String() string {
func (me NegativeInteger) String() string {
	return strconv.FormatInt(int64(me), 10)
	return strconv.FormatInt(int64(me), 10)
}
}


//	A convenience interface that declares a type conversion to NegativeInteger.
//	A convenience interface that declares a type conversion to NegativeInteger.
type ToXsdtNegativeInteger interface {
type ToXsdtNegativeInteger interface {
	ToXsdtNegativeInteger() NegativeInteger
	ToXsdtNegativeInteger() NegativeInteger
}
}


//	An NMTOKEN is set of name characters (letters, digits, and other characters) in any combination. Unlike Name and NCName, NMTOKEN has no restrictions on the starting character.
//	An NMTOKEN is set of name characters (letters, digits, and other characters) in any combination. Unlike Name and NCName, NMTOKEN has no restrictions on the starting character.
type Nmtoken Token
type Nmtoken Token


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Nmtoken) Set(v string) {
func (me *Nmtoken) Set(v string) {
	*me = Nmtoken(v)
	*me = Nmtoken(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Nmtoken) String() string {
func (me Nmtoken) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Nmtoken.
//	A convenience interface that declares a type conversion to Nmtoken.
type ToXsdtNmtoken interface {
type ToXsdtNmtoken interface {
	ToXsdtNmtoken() Nmtoken
	ToXsdtNmtoken() Nmtoken
}
}


//	Contains a set of values of type NMTOKEN.
//	Contains a set of values of type NMTOKEN.
type Nmtokens string
type Nmtokens string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Nmtokens) Set(v string) {
func (me *Nmtokens) Set(v string) {
	*me = Nmtokens(v)
	*me = Nmtokens(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Nmtokens) String() string {
func (me Nmtokens) String() string {
	return string(me)
	return string(me)
}
}


//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
func (me Nmtokens) Values() (list []Nmtoken) {
func (me Nmtokens) Values() (list []Nmtoken) {
	spl := ListValues(string(me))
	spl := ListValues(string(me))
	list = make([]Nmtoken, len(spl))
	list = make([]Nmtoken, len(spl))
	for i, s := range spl {
	for i, s := range spl {
		list[i].Set(s)
		list[i].Set(s)
	}
	}
	return
	return
}
}


//	A convenience interface that declares a type conversion to Nmtokens.
//	A convenience interface that declares a type conversion to Nmtokens.
type ToXsdtNmtokens interface {
type ToXsdtNmtokens interface {
	ToXsdtNmtokens() Nmtokens
	ToXsdtNmtokens() Nmtokens
}
}


//	Represents an integer that is greater than or equal to zero.
//	Represents an integer that is greater than or equal to zero.
type NonNegativeInteger uint64
type NonNegativeInteger uint64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me NonNegativeInteger) N() uint64 {
func (me NonNegativeInteger) N() uint64 {
	return uint64(me)
	return uint64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *NonNegativeInteger) Set(s string) {
func (me *NonNegativeInteger) Set(s string) {
	v, _ := strconv.ParseUint(s, 0, 64)
	v, _ := strconv.ParseUint(s, 0, 64)
	*me = NonNegativeInteger(v)
	*me = NonNegativeInteger(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me NonNegativeInteger) String() string {
func (me NonNegativeInteger) String() string {
	return strconv.FormatUint(uint64(me), 10)
	return strconv.FormatUint(uint64(me), 10)
}
}


//	A convenience interface that declares a type conversion to NonNegativeInteger.
//	A convenience interface that declares a type conversion to NonNegativeInteger.
type ToXsdtNonNegativeInteger interface {
type ToXsdtNonNegativeInteger interface {
	ToXsdtNonNegativeInteger() NonNegativeInteger
	ToXsdtNonNegativeInteger() NonNegativeInteger
}
}


//	Represents an integer that is less than or equal to zero. A nonPositiveIntegerconsists of a negative sign (-) and sequence of decimal digits.
//	Represents an integer that is less than or equal to zero. A nonPositiveIntegerconsists of a negative sign (-) and sequence of decimal digits.
type NonPositiveInteger int64
type NonPositiveInteger int64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me NonPositiveInteger) N() int64 {
func (me NonPositiveInteger) N() int64 {
	return int64(me)
	return int64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *NonPositiveInteger) Set(s string) {
func (me *NonPositiveInteger) Set(s string) {
	v, _ := strconv.ParseInt(s, 0, 64)
	v, _ := strconv.ParseInt(s, 0, 64)
	*me = NonPositiveInteger(v)
	*me = NonPositiveInteger(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me NonPositiveInteger) String() string {
func (me NonPositiveInteger) String() string {
	return strconv.FormatInt(int64(me), 10)
	return strconv.FormatInt(int64(me), 10)
}
}


//	A convenience interface that declares a type conversion to NonPositiveInteger.
//	A convenience interface that declares a type conversion to NonPositiveInteger.
type ToXsdtNonPositiveInteger interface {
type ToXsdtNonPositiveInteger interface {
	ToXsdtNonPositiveInteger() NonPositiveInteger
	ToXsdtNonPositiveInteger() NonPositiveInteger
}
}


//	Represents white space normalized strings.
//	Represents white space normalized strings.
type NormalizedString String
type NormalizedString String


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *NormalizedString) Set(v string) {
func (me *NormalizedString) Set(v string) {
	*me = NormalizedString(v)
	*me = NormalizedString(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me NormalizedString) String() string {
func (me NormalizedString) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to NormalizedString.
//	A convenience interface that declares a type conversion to NormalizedString.
type ToXsdtNormalizedString interface {
type ToXsdtNormalizedString interface {
	ToXsdtNormalizedS() NormalizedString
	ToXsdtNormalizedS() NormalizedString
}
}


//	A set of QNames.
//	A set of QNames.
type Notation string
type Notation string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Notation) Set(v string) {
func (me *Notation) Set(v string) {
	*me = Notation(v)
	*me = Notation(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Notation) String() string {
func (me Notation) String() string {
	return string(me)
	return string(me)
}
}


//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
//	This type declares a String containing a whitespace-separated list of values. This Values() method creates and returns a slice of all elements in that list.
func (me Notation) Values() (list []Qname) {
func (me Notation) Values() (list []Qname) {
	spl := ListValues(string(me))
	spl := ListValues(string(me))
	list = make([]Qname, len(spl))
	list = make([]Qname, len(spl))
	for i, s := range spl {
	for i, s := range spl {
		list[i].Set(s)
		list[i].Set(s)
	}
	}
	return
	return
}
}


//	A convenience interface that declares a type conversion to Notation.
//	A convenience interface that declares a type conversion to Notation.
type ToXsdtNotation interface {
type ToXsdtNotation interface {
	ToXsdtNotation() Notation
	ToXsdtNotation() Notation
}
}


//	Represents an integer that is greater than zero.
//	Represents an integer that is greater than zero.
type PositiveInteger uint64
type PositiveInteger uint64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me PositiveInteger) N() uint64 {
func (me PositiveInteger) N() uint64 {
	return uint64(me)
	return uint64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *PositiveInteger) Set(s string) {
func (me *PositiveInteger) Set(s string) {
	v, _ := strconv.ParseUint(s, 0, 64)
	v, _ := strconv.ParseUint(s, 0, 64)
	*me = PositiveInteger(v)
	*me = PositiveInteger(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me PositiveInteger) String() string {
func (me PositiveInteger) String() string {
	return strconv.FormatUint(uint64(me), 10)
	return strconv.FormatUint(uint64(me), 10)
}
}


//	A convenience interface that declares a type conversion to PositiveInteger.
//	A convenience interface that declares a type conversion to PositiveInteger.
type ToXsdtPositiveInteger interface {
type ToXsdtPositiveInteger interface {
	ToXsdtPositiveInteger() PositiveInteger
	ToXsdtPositiveInteger() PositiveInteger
}
}


//	Represents a qualified name. A qualified name is composed of a prefix and a local name separated by a colon. Both the prefix and local names must be an NCName. The prefix must be associated with a namespace URI reference, using a namespace declaration.
//	Represents a qualified name. A qualified name is composed of a prefix and a local name separated by a colon. Both the prefix and local names must be an NCName. The prefix must be associated with a namespace URI reference, using a namespace declaration.
type Qname string
type Qname string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Qname) Set(v string) {
func (me *Qname) Set(v string) {
	*me = Qname(v)
	*me = Qname(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Qname) String() string {
func (me Qname) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Qname.
//	A convenience interface that declares a type conversion to Qname.
type ToXsdtQname interface {
type ToXsdtQname interface {
	ToXsdtQname() Qname
	ToXsdtQname() Qname
}
}


//	Represents an integer with a minimum value of -32768 and maximum of 32767.
//	Represents an integer with a minimum value of -32768 and maximum of 32767.
type Short int16
type Short int16


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me Short) N() int16 {
func (me Short) N() int16 {
	return int16(me)
	return int16(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *Short) Set(s string) {
func (me *Short) Set(s string) {
	v, _ := strconv.ParseInt(s, 0, 16)
	v, _ := strconv.ParseInt(s, 0, 16)
	*me = Short(v)
	*me = Short(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me Short) String() string {
func (me Short) String() string {
	return strconv.FormatInt(int64(me), 10)
	return strconv.FormatInt(int64(me), 10)
}
}


//	A convenience interface that declares a type conversion to Short.
//	A convenience interface that declares a type conversion to Short.
type ToXsdtShort interface {
type ToXsdtShort interface {
	ToXsdtShort() Short
	ToXsdtShort() Short
}
}


//	Represents character strings.
//	Represents character strings.
type String string
type String string


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *String) Set(v string) {
func (me *String) Set(v string) {
	*me = String(v)
	*me = String(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me String) String() string {
func (me String) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to String.
//	A convenience interface that declares a type conversion to String.
type ToXsdtString interface {
type ToXsdtString interface {
	ToXsdtString() String
	ToXsdtString() String
}
}


//	Represents tokenized strings.
//	Represents tokenized strings.
type Token NormalizedString
type Token NormalizedString


//	Since this is just a simple String type, this merely sets the current value from the specified string.
//	Since this is just a simple String type, this merely sets the current value from the specified string.
func (me *Token) Set(v string) {
func (me *Token) Set(v string) {
	*me = Token(v)
	*me = Token(v)
}
}


//	Since this is just a simple String type, this merely returns its current string value.
//	Since this is just a simple String type, this merely returns its current string value.
func (me Token) String() string {
func (me Token) String() string {
	return string(me)
	return string(me)
}
}


//	A convenience interface that declares a type conversion to Token.
//	A convenience interface that declares a type conversion to Token.
type ToXsdtToken interface {
type ToXsdtToken interface {
	ToXsdtToken() Token
	ToXsdtToken() Token
}
}


//	Represents an integer with a minimum of zero and maximum of 255.
//	Represents an integer with a minimum of zero and maximum of 255.
type UnsignedByte uint8
type UnsignedByte uint8


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me UnsignedByte) N() uint8 {
func (me UnsignedByte) N() uint8 {
	return uint8(me)
	return uint8(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *UnsignedByte) Set(s string) {
func (me *UnsignedByte) Set(s string) {
	v, _ := strconv.ParseUint(s, 0, 8)
	v, _ := strconv.ParseUint(s, 0, 8)
	*me = UnsignedByte(v)
	*me = UnsignedByte(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me UnsignedByte) String() string {
func (me UnsignedByte) String() string {
	return strconv.FormatUint(uint64(me), 10)
	return strconv.FormatUint(uint64(me), 10)
}
}


//	A convenience interface that declares a type conversion to UnsignedByte.
//	A convenience interface that declares a type conversion to UnsignedByte.
type ToXsdtUnsignedByte interface {
type ToXsdtUnsignedByte interface {
	ToXsdtUnsignedByte() UnsignedByte
	ToXsdtUnsignedByte() UnsignedByte
}
}


//	Represents an integer with a minimum of zero and maximum of 4294967295.
//	Represents an integer with a minimum of zero and maximum of 4294967295.
type UnsignedInt uint32
type UnsignedInt uint32


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me UnsignedInt) N() uint32 {
func (me UnsignedInt) N() uint32 {
	return uint32(me)
	return uint32(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *UnsignedInt) Set(s string) {
func (me *UnsignedInt) Set(s string) {
	v, _ := strconv.ParseUint(s, 0, 32)
	v, _ := strconv.ParseUint(s, 0, 32)
	*me = UnsignedInt(v)
	*me = UnsignedInt(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me UnsignedInt) String() string {
func (me UnsignedInt) String() string {
	return strconv.FormatUint(uint64(me), 10)
	return strconv.FormatUint(uint64(me), 10)
}
}


//	A convenience interface that declares a type conversion to UnsignedInt.
//	A convenience interface that declares a type conversion to UnsignedInt.
type ToXsdtUnsignedInt interface {
type ToXsdtUnsignedInt interface {
	ToXsdtUnsignedInt() UnsignedInt
	ToXsdtUnsignedInt() UnsignedInt
}
}


//	Represents an integer with a minimum of zero and maximum of 18446744073709551615.
//	Represents an integer with a minimum of zero and maximum of 18446744073709551615.
type UnsignedLong uint64
type UnsignedLong uint64


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me UnsignedLong) N() uint64 {
func (me UnsignedLong) N() uint64 {
	return uint64(me)
	return uint64(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *UnsignedLong) Set(s string) {
func (me *UnsignedLong) Set(s string) {
	v, _ := strconv.ParseUint(s, 0, 64)
	v, _ := strconv.ParseUint(s, 0, 64)
	*me = UnsignedLong(v)
	*me = UnsignedLong(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me UnsignedLong) String() string {
func (me UnsignedLong) String() string {
	return strconv.FormatUint(uint64(me), 10)
	return strconv.FormatUint(uint64(me), 10)
}
}


//	A convenience interface that declares a type conversion to UnsignedLong.
//	A convenience interface that declares a type conversion to UnsignedLong.
type ToXsdtUnsignedLong interface {
type ToXsdtUnsignedLong interface {
	ToXsdtUnsignedLong() UnsignedLong
	ToXsdtUnsignedLong() UnsignedLong
}
}


//	Represents an integer with a minimum of zero and maximum of 65535.
//	Represents an integer with a minimum of zero and maximum of 65535.
type UnsignedShort uint16
type UnsignedShort uint16


//	Because littering your code with type conversions is a hassle...
//	Because littering your code with type conversions is a hassle...
func (me UnsignedShort) N() uint16 {
func (me UnsignedShort) N() uint16 {
	return uint16(me)
	return uint16(me)
}
}


//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
//	Since this is a non-string scalar type, sets its current value obtained from parsing the specified string.
func (me *UnsignedShort) Set(s string) {
func (me *UnsignedShort) Set(s string) {
	v, _ := strconv.ParseUint(s, 0, 16)
	v, _ := strconv.ParseUint(s, 0, 16)
	*me = UnsignedShort(v)
	*me = UnsignedShort(v)
}
}


//	Returns a string representation of its current non-string scalar value.
//	Returns a string representation of its current non-string scalar value.
func (me UnsignedShort) String() string {
func (me UnsignedShort) String() string {
	return strconv.FormatUint(uint64(me), 10)
	return strconv.FormatUint(uint64(me), 10)
}
}


//	A convenience interface that declares a type conversion to UnsignedShort.
//	A convenience interface that declares a type conversion to UnsignedShort.
type ToXsdtUnsignedShort interface {
type ToXsdtUnsignedShort interface {
	ToXsdtUnsignedShort() UnsignedShort
	ToXsdtUnsignedShort() UnsignedShort
}
}


// XSD "list" types are always space-separated strings. All generated Go types based on any XSD's list types get a Values() method, which will always resort to this function.
// XSD "list" types are always space-separated strings. All generated Go types based on any XSD's list types get a Values() method, which will always resort to this function.
func ListValues(v string) (spl []string) {
func ListValues(v string) (spl []string) {
	if len(v) == 0 {
	if len(v) == 0 {
		return
		return
	}
	}
	lastWs := true
	lastWs := true
	wsr := func(r rune) bool {
	wsr := func(r rune) bool {
		return (r == ' ') || (r == '\r') || (r == '\n') || (r == '\t')
		return (r == ' ') || (r == '\r') || (r == '\n') || (r == '\t')
	}
	}
	wss := func(r string) bool {
	wss := func(r string) bool {
		return (r == " ") || (r == "\r") || (r == "\n") || (r == "\t")
		return (r == " ") || (r == "\r") || (r == "\n") || (r == "\t")
	}
	}
	for wss(v[len(v)-1:]) {
	for wss(v[len(v)-1:]) {
		v = v[:len(v)-1]
		v = v[:len(v)-1]
	}
	}
	for wss(v[:1]) {
	for wss(v[:1]) {
		v = v[1:]
		v = v[1:]
	}
	}
	if len(v) > 0 {
	if len(v) > 0 {
		cur, num, i := "", 1, 0
		cur, num, i := "", 1, 0
		for _, r := range v {
		for _, r := range v {
			if wsr(r) {
			if wsr(r) {
				if !lastWs {
				if !lastWs {
					num++
					num++
					lastWs = true
					lastWs = true
				}
				}
			} else {
			} else {
				lastWs = false
				lastWs = false
			}
			}
		}
		}
		lastWs, spl = true, make([]string, num)
		lastWs, spl = true, make([]string, num)
		for _, r := range v {
		for _, r := range v {
			if wsr(r) {
			if wsr(r) {
				if !lastWs {
				if !lastWs {
					if len(cur) > 0 {
					if len(cur) > 0 {
						spl[i] = cur
						spl[i] = cur
						i++
						i++
					}
					}
					cur, lastWs = "", true
					cur, lastWs = "", true
				}
				}
			} else {
			} else {
				lastWs = false
				lastWs = false
				cur += string(r)
				cur += string(r)
			}
			}
		}
		}
		if len(cur) > 0 {
		if len(cur) > 0 {
			spl[i] = cur
			spl[i] = cur
		}
		}
	}
	}
	return
	return
}
}


func ListValuesBoolean(vals []Boolean) (sl []bool) {
func ListValuesBoolean(vals []Boolean) (sl []bool) {
	sl = make([]bool, len(vals))
	sl = make([]bool, len(vals))
	for i, b := range vals {
	for i, b := range vals {
		sl[i] = b.B()
		sl[i] = b.B()
	}
	}
	return
	return
}
}


func ListValuesDouble(vals []Double) (sl []float64) {
func ListValuesDouble(vals []Double) (sl []float64) {
	sl = make([]float64, len(vals))
	sl = make([]float64, len(vals))
	for i, d := range vals {
	for i, d := range vals {
		sl[i] = d.N()
		sl[i] = d.N()
	}
	}
	return
	return
}
}


func ListValuesLong(vals []Long) (sl []int64) {
func ListValuesLong(vals []Long) (sl []int64) {
	sl = make([]int64, len(vals))
	sl = make([]int64, len(vals))
	for i, l := range vals {
	for i, l := range vals {
		sl[i] = l.N()
		sl[i] = l.N()
	}
	}
	return
	return
}
}


//	A helper function for the Walk() functionality of generated wrapper packages.
//	A helper function for the Walk() functionality of generated wrapper packages.
func OnWalkError(err *error, slice *[]error, breakWalk bool, handler func(error)) (ret bool) {
func OnWalkError(err *error, slice *[]error, breakWalk bool, handler func(error)) (ret bool) {
	if e := *err; e != nil {
	if e := *err; e != nil {
		*slice = append(*slice, e)
		*slice = append(*slice, e)
		ret = breakWalk
		ret = breakWalk
		if handler != nil {
		if handler != nil {
			handler(e)
			handler(e)
		}
		}
	}
	}
	*err = nil
	*err = nil
	return
	return
}
}
