package utils

type timeError struct{}
type stringOfIntSliceError struct{}
type latError struct{}
type lngError struct{}

func (e *timeError) Error() string {
	return "incorrect hour format, 'HH:MM' is the correct format"
}

func (e *stringOfIntSliceError) Error() string {
	return "input only number"
}

func (e *latError) Error() string {
	return "incorret latitude input, " + 
	"'-x.xxxxxxx' is the correct format, " + 
	"the precision of the number after the decimal point is seven digits"
}

func (e *lngError) Error() string {
	return "incorret longitude input, " + 
	"'xxx.xxxxxxx' is the correct format, " + 
	"the precision of the number after the decimal point is seven digits"
}