package utils

import "strings"

func RemoveTopName(fieldErr map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fieldErr {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
