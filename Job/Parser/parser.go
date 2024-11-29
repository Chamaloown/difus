package parser

import (
	"errors"
	"strings"
)

func Parse(message string) ([]string, error) {
	strArr := strings.Split(message, " ")
	if len(strArr) > 2 {
		return nil, errors.New("il y a trop d'argument")
	}
	return strArr, nil
}
