package uslice
package uslice


import "strings"
import "strings"


//	Returns the position of lower-case `val` in lower-case `vals`.
//	Returns the position of lower-case `val` in lower-case `vals`.
type XsdtString struct{ string }


	lv := strings.ToLower(val)
	lv := strings.ToLower(val)
	for i, v := range vals {
	for i, v := range vals {
		if (v == val) || (strings.ToLower(v) == lv) {
		if (v == val) || (strings.ToLower(v) == lv) {
			return i
			return i
type XsdtString struct{ string }		}
	}
	}
	return -1
	return -1
}
}


//	Returns whether lower-case `val` is in lower-case `vals`.
//	Returns whether lower-case `val` is in lower-case `vals`.
func StrHasIgnoreCase(vals []string, val string) bool {
func StrHasIgnoreCase(vals []string, val string) bool {
	return StrAtIgnoreCase(vals, val) >= 0
	return StrAtIgnoreCase(vals, val) >= 0
}
}


//#begin-gt -gen.gt N:Str T:string
//#begin-gt -gen.gt N:Str T:string


//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
func StrAppendUnique(ref *[]string, v string) {
func StrAppendUnique(ref *[]string, v string) {
	for _, sv := range *ref {
	for _, sv := range *ref {
		if sv == v {
		if sv == v {
			return
			return
		}
		}
	}
	}
	*ref = append(*ref, v)
	*ref = append(*ref, v)
}
}


//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
func StrAppendUniques(ref *[]string, vals ...string) {
func StrAppendUniques(ref *[]string, vals ...string) {
	for _, v := range vals {
	for _, v := range vals {
		StrAppendUnique(ref, v)
		StrAppendUnique(ref, v)
	}
	}
}
}


//	Returns the position of `val` in `slice`.
//	Returns the position of `val` in `slice`.
func StrAt(slice []string, val string) int {
func StrAt(slice []string, val string) int {
	for i, v := range slice {
	for i, v := range slice {
		if v == val {
		if v == val {
			return i
			return i
		}
		}
	}
	}
	return -1
	return -1
}
}


//	Converts `src` to `dst`.
//	Converts `src` to `dst`.
//
//
//	If `sparse` is `true`, then only successfully converted `string` values are placed
//	If `sparse` is `true`, then only successfully converted `string` values are placed
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//
//
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
func StrConvert(src []interface{}, sparse bool) (dst []string) {
func StrConvert(src []interface{}, sparse bool) (dst []string) {
	if sparse {
	if sparse {
		var (
		var (
			val string
			val string
			ok  bool
			ok  bool
		)
		)
		for _, v := range src {
		for _, v := range src {
			if val, ok = v.(string); ok {
			if val, ok = v.(string); ok {
				dst = append(dst, val)
				dst = append(dst, val)
			}
			}
		}
		}
	} else {
	} else {
		dst = make([]string, len(src))
		dst = make([]string, len(src))
		for i, v := range src {
		for i, v := range src {
			dst[i], _ = v.(string)
			dst[i], _ = v.(string)
		}
		}
	}
	}
	return
	return
}
}


//	Sets each `string` in `sl` to the result of passing it to each `apply` func.
//	Sets each `string` in `sl` to the result of passing it to each `apply` func.
//	Although `sl` is modified in-place, it is also returned for convenience.
//	Although `sl` is modified in-place, it is also returned for convenience.
func StrEach(sl []string, apply ...func(string) string) []string {
func StrEach(sl []string, apply ...func(string) string) []string {
	for _, fn := range apply {
	for _, fn := range apply {
		for i, _ := range sl {
		for i, _ := range sl {
			sl[i] = fn(sl[i])
			sl[i] = fn(sl[i])
		}
		}
	}
	}
	return sl
	return sl
}
}


//	Calls `StrSetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
//	Calls `StrSetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
func StrEnsureCap(ref *[]string, capacity int) {
func StrEnsureCap(ref *[]string, capacity int) {
	if cap(*ref) < capacity {
	if cap(*ref) < capacity {
		StrSetCap(ref, capacity)
		StrSetCap(ref, capacity)
	}
	}
}
}


//	Calls `StrSetLen` only if the current `len(*ref)` is less than the specified `length`.
//	Calls `StrSetLen` only if the current `len(*ref)` is less than the specified `length`.
func StrEnsureLen(ref *[]string, length int) {
func StrEnsureLen(ref *[]string, length int) {
	if len(*ref) < length {
	if len(*ref) < length {
		StrSetLen(ref, length)
		StrSetLen(ref, length)
	}
	}
}
}


//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
func StrEquivalent(one, two []string) bool {
func StrEquivalent(one, two []string) bool {
	if len(one) != len(two) {
	if len(one) != len(two) {
		return false
		return false
	}
	}
	for _, v := range one {
	for _, v := range one {
		if StrAt(two, v) < 0 {
		if StrAt(two, v) < 0 {
			return false
			return false
		}
		}
	}
	}
	return true
	return true
}
}


//	Returns whether `val` is in `slice`.
//	Returns whether `val` is in `slice`.
func StrHas(slice []string, val string) bool {
func StrHas(slice []string, val string) bool {
	return StrAt(slice, val) >= 0
	return StrAt(slice, val) >= 0
}
}


//	Returns whether at least one of the specified `vals` is contained in `slice`.
//	Returns whether at least one of the specified `vals` is contained in `slice`.
func StrHasAny(slice []string, vals ...string) bool {
func StrHasAny(slice []string, vals ...string) bool {
	for _, v1 := range vals {
	for _, v1 := range vals {
		for _, v2 := range slice {
		for _, v2 := range slice {
			if v1 == v2 {
			if v1 == v2 {
				return true
				return true
			}
			}
		}
		}
	}
	}
	return false
	return false
}
}


//	Removes the first occurrence of `v` encountered in `*ref`, or all occurrences if `all` is `true`.
//	Removes the first occurrence of `v` encountered in `*ref`, or all occurrences if `all` is `true`.
func StrRemove(ref *[]string, v string, all bool) {
func StrRemove(ref *[]string, v string, all bool) {
	for i := 0; i < len(*ref); i++ {
	for i := 0; i < len(*ref); i++ {
		if (*ref)[i] == v {
		if (*ref)[i] == v {
			before, after := (*ref)[:i], (*ref)[i+1:]
			before, after := (*ref)[:i], (*ref)[i+1:]
			*ref = append(before, after...)
			*ref = append(before, after...)
			if !all {
			if !all {
				break
				break
			}
			}
		}
		}
	}
	}
}
}


//	Sets `*ref` to a copy of `*ref` with the specified `capacity`.
//	Sets `*ref` to a copy of `*ref` with the specified `capacity`.
func StrSetCap(ref *[]string, capacity int) {
func StrSetCap(ref *[]string, capacity int) {
	nu := make([]string, len(*ref), capacity)
	nu := make([]string, len(*ref), capacity)
	copy(nu, *ref)
	copy(nu, *ref)
	*ref = nu
	*ref = nu
}
}


//	Sets `*ref` to a copy of `*ref` with the specified `length`.
//	Sets `*ref` to a copy of `*ref` with the specified `length`.
func StrSetLen(ref *[]string, length int) {
func StrSetLen(ref *[]string, length int) {
	nu := make([]string, length)
	nu := make([]string, length)
	copy(nu, *ref)
	copy(nu, *ref)
	*ref = nu
	*ref = nu
}
}


//	Removes all specified `withoutVals` from `slice`.
//	Removes all specified `withoutVals` from `slice`.
func StrWithout(slice []string, keepOrder bool, withoutVals ...string) []string {
func StrWithout(slice []string, keepOrder bool, withoutVals ...string) []string {
	if len(withoutVals) > 0 {
	if len(withoutVals) > 0 {
		var pos int
		var pos int
		for _, w := range withoutVals {
		for _, w := range withoutVals {
			for pos = StrAt(slice, w); pos >= 0; pos = StrAt(slice, w) {
			for pos = StrAt(slice, w); pos >= 0; pos = StrAt(slice, w) {
				if keepOrder {
				if keepOrder {
					slice = append(slice[:pos], slice[pos+1:]...)
					slice = append(slice[:pos], slice[pos+1:]...)
				} else {
				} else {
					slice[pos] = slice[len(slice)-1]
					slice[pos] = slice[len(slice)-1]
					slice = slice[:len(slice)-1]
					slice = slice[:len(slice)-1]
				}
				}
			}
			}
		}
		}
	}
	}
	return slice
	return slice
}
}


//#end-gt
//#end-gt
