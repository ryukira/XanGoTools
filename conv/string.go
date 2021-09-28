package conv

import "strconv"

func StringNumberToNumber(Number string) (int, error) {
	return strconv.Atoi(Number)
}
