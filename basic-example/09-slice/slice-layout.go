package main

import "reflect"

var sum int64

func addUpDirect(s []int64) {
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
}

func addUpViaInterface(s []interface{}) {
	for i := 0; i < len(s); i++ {
		sum += s[i].(int64)
	}
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func main() {
	is := []int64{0x55, 0x22, 0xab, 0x9}

	addUpDirect(is)

	iis := make([]interface{}, len(is))
	for i := 0; i < len(is); i++ {
		iis[i] = is[i]
	}

	addUpViaInterface(iis)
}
