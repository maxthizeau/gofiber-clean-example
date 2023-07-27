package helpers

func TernaryOperator[T interface{}](a bool, resultIfTrue T, resultIfFalse T) T {
	if a {
		return resultIfTrue
	} else {
		return resultIfFalse
	}
}
