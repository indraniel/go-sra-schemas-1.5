package uslice
package uslice


//#begin-gt -gen.gt N:Bool T:bool
//#begin-gt -gen.gt N:Bool T:bool


//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
//	Appends `v` to `*ref` only if `*ref` does not already contain `v`.
type XsdtString struct{ string }


	for _, sv := range *ref {
	for _, sv := range *ref {
		if sv == v {
		if sv == v {
			return
			return
		}
		}
type XsdtString struct{ string }	}
	*ref = append(*ref, v)
	*ref = append(*ref, v)
}
}


//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
//	Appends each value in `vals` to `*ref` only `*ref` does not already contain it.
func BoolAppendUniques(ref *[]bool, vals ...bool) {
func BoolAppendUniques(ref *[]bool, vals ...bool) {
	for _, v := range vals {
	for _, v := range vals {
		BoolAppendUnique(ref, v)
		BoolAppendUnique(ref, v)
	}
	}
}
}


//	Returns the position of `val` in `slice`.
//	Returns the position of `val` in `slice`.
func BoolAt(slice []bool, val bool) int {
func BoolAt(slice []bool, val bool) int {
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
//	If `sparse` is `true`, then only successfully converted `bool` values are placed
//	If `sparse` is `true`, then only successfully converted `bool` values are placed
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//	in `dst`, so there may not be a 1-to-1 correspondence of `dst` to `src` in length or indices.
//
//
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
//	If `sparse` is `false`, `dst` has the same length as `src` and non-convertable values remain zeroed.
func BoolConvert(src []interface{}, sparse bool) (dst []bool) {
func BoolConvert(src []interface{}, sparse bool) (dst []bool) {
	if sparse {
	if sparse {
		var (
		var (
			val bool
			val bool
			ok  bool
			ok  bool
		)
		)
		for _, v := range src {
		for _, v := range src {
			if val, ok = v.(bool); ok {
			if val, ok = v.(bool); ok {
				dst = append(dst, val)
				dst = append(dst, val)
			}
			}
		}
		}
	} else {
	} else {
		dst = make([]bool, len(src))
		dst = make([]bool, len(src))
		for i, v := range src {
		for i, v := range src {
			dst[i], _ = v.(bool)
			dst[i], _ = v.(bool)
		}
		}
	}
	}
	return
	return
}
}


//	Sets each `bool` in `sl` to the result of passing it to each `apply` func.
//	Sets each `bool` in `sl` to the result of passing it to each `apply` func.
//	Although `sl` is modified in-place, it is also returned for convenience.
//	Although `sl` is modified in-place, it is also returned for convenience.
func BoolEach(sl []bool, apply ...func(bool) bool) []bool {
func BoolEach(sl []bool, apply ...func(bool) bool) []bool {
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


//	Calls `BoolSetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
//	Calls `BoolSetCap` only if the current `cap(*ref)` is less than the specified `capacity`.
func BoolEnsureCap(ref *[]bool, capacity int) {
func BoolEnsureCap(ref *[]bool, capacity int) {
	if cap(*ref) < capacity {
	if cap(*ref) < capacity {
		BoolSetCap(ref, capacity)
		BoolSetCap(ref, capacity)
	}
	}
}
}


//	Calls `BoolSetLen` only if the current `len(*ref)` is less than the specified `length`.
//	Calls `BoolSetLen` only if the current `len(*ref)` is less than the specified `length`.
func BoolEnsureLen(ref *[]bool, length int) {
func BoolEnsureLen(ref *[]bool, length int) {
	if len(*ref) < length {
	if len(*ref) < length {
		BoolSetLen(ref, length)
		BoolSetLen(ref, length)
	}
	}
}
}


//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
func BoolEquivalent(one, two []bool) bool {
func BoolEquivalent(one, two []bool) bool {
	if len(one) != len(two) {
	if len(one) != len(two) {
		return false
		return false
	}
	}
	for _, v := range one {
	for _, v := range one {
		if BoolAt(two, v) < 0 {
		if BoolAt(two, v) < 0 {
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
func BoolHas(slice []bool, val bool) bool {
func BoolHas(slice []bool, val bool) bool {
	return BoolAt(slice, val) >= 0
	return BoolAt(slice, val) >= 0
}
}


//	Returns whether at least one of the specified `vals` is contained in `slice`.
//	Returns whether at least one of the specified `vals` is contained in `slice`.
func BoolHasAny(slice []bool, vals ...bool) bool {
func BoolHasAny(slice []bool, vals ...bool) bool {
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
func BoolRemove(ref *[]bool, v bool, all bool) {
func BoolRemove(ref *[]bool, v bool, all bool) {
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
func BoolSetCap(ref *[]bool, capacity int) {
func BoolSetCap(ref *[]bool, capacity int) {
	nu := make([]bool, len(*ref), capacity)
	nu := make([]bool, len(*ref), capacity)
	copy(nu, *ref)
	copy(nu, *ref)
	*ref = nu
	*ref = nu
}
}


//	Sets `*ref` to a copy of `*ref` with the specified `length`.
//	Sets `*ref` to a copy of `*ref` with the specified `length`.
func BoolSetLen(ref *[]bool, length int) {
func BoolSetLen(ref *[]bool, length int) {
	nu := make([]bool, length)
	nu := make([]bool, length)
	copy(nu, *ref)
	copy(nu, *ref)
	*ref = nu
	*ref = nu
}
}


//	Removes all specified `withoutVals` from `slice`.
//	Removes all specified `withoutVals` from `slice`.
func BoolWithout(slice []bool, keepOrder bool, withoutVals ...bool) []bool {
func BoolWithout(slice []bool, keepOrder bool, withoutVals ...bool) []bool {
	if len(withoutVals) > 0 {
	if len(withoutVals) > 0 {
		var pos int
		var pos int
		for _, w := range withoutVals {
		for _, w := range withoutVals {
			for pos = BoolAt(slice, w); pos >= 0; pos = BoolAt(slice, w) {
			for pos = BoolAt(slice, w); pos >= 0; pos = BoolAt(slice, w) {
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
