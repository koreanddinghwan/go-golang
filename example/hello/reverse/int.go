package reverse

import "strconv"

func Int(iden int) int {
	iden, _ = strconv.Atoi(String(strconv.Itoa(iden)))
	return iden
}
