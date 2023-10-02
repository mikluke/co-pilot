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

// Coalesce returns first not nil value
// Be wary that all values are determined before this function call.
func Coalesce[T any](value ...*T) *T {
	for _, v := range value {
		if v != nil {
			return v
		}
	}

	return nil
}
