/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 2.0.1
 * 
 * This file is not intended to be easily readable and contains a number of 
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG 
 * interface file instead. 
 * ----------------------------------------------------------------------------- */

package example


type _swig_fnptr *byte
type _swig_memberptr *byte


func _swig_allocatememory(int) *byte
func _swig_internal_allocate(len int) *byte {
	return _swig_allocatememory(len)
}

func _swig_allocatestring(*byte, int) string
func _swig_internal_makegostring(p *byte, l int) string {
	return _swig_allocatestring(p, l)
}

func _swig_internal_gopanic(p *byte, l int) {
	panic(_swig_allocatestring(p, l))
}

func _swig_wrap_add(*int, *int, *int)

func Add(arg1 *int, arg2 *int, arg3 *int) {
	_swig_wrap_add(arg1, arg2, arg3)
}

func New_intp() *int
func Copy_intp(int) *int
func _swig_wrap_delete_intp(*int)

func Delete_intp(arg1 *int) {
	_swig_wrap_delete_intp(arg1)
}

func _swig_wrap_intp_assign(*int, int)

func Intp_assign(arg1 *int, arg2 int) {
	_swig_wrap_intp_assign(arg1, arg2)
}

func Intp_value(*int) int
func _swig_wrap_sub(int, int, []int)

func Sub(arg1 int, arg2 int, arg3 []int) {
	_swig_wrap_sub(arg1, arg2, arg3)
}

func Divide(int, int, []int) int

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}
