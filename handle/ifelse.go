package handle

// Q is an implementation of ternary operator.
// Be wary that positiveCase and negativeCase are determined before this function call.
func Q[T any](expr bool, positiveCase, negativeCase T) T {
	if expr {
		return positiveCase
	}

	return negativeCase
}

// QNil returns retValue if checkValue is not nil, otherwise returns zero value of R
// Be wary that retValue is determined before this function call.
func QNil[T any, R any](checkValue *T, retValue R) R {
	if checkValue == nil {
		var out R

		return out
	}

	return retValue
}

// Coalesce returns value if it's not nil, otherwise returns defaultValue.
// Be wary that value and defaultValue are determined before this function call.
func Coalesce[T any](value, defaultValue *T) *T {
	if value == nil {
		return defaultValue
	}

	return value
}
