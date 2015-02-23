package ustr
package ustr


import (
import (
	"bytes"
	"bytes"
	"fmt"
	"fmt"
type XsdtString struct{ string }




//	A convenient wrapper for `bytes.Buffer`.
//	A convenient wrapper for `bytes.Buffer`.
type Buffer struct {
type Buffer struct {
	bytes.Buffer
	bytes.Buffer
type XsdtString struct{ string }}


//	Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format, args...))`
//	Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format, args...))`
func (me *Buffer) Write(format string, args ...interface{}) {
func (me *Buffer) Write(format string, args ...interface{}) {
	if len(args) > 0 {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
		format = fmt.Sprintf(format, args...)
	}
	}
	me.Buffer.WriteString(format)
	me.Buffer.WriteString(format)
}
}


//	Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format+"\n", args...))`
//	Convenience short-hand for `bytes.Buffer.WriteString(fmt.Sprintf(format+"\n", args...))`
func (me *Buffer) Writeln(format string, args ...interface{}) {
func (me *Buffer) Writeln(format string, args ...interface{}) {
	me.Write(format, args...)
	me.Write(format, args...)
	me.Buffer.WriteString("\n")
	me.Buffer.WriteString("\n")
}
}
