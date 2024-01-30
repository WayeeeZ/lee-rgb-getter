package leegetter

import (
	"errors"
	"strings"
)

type Lee struct {
	Name  string
	Red   uint8
	Green uint8
	Blue  uint8
}

func GetLee(name string) (Lee, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	for _, item := range leeList {
		if name == strings.ToLower(strings.TrimSpace(item.Name)) {
			return item, nil
		}
	}
	return Lee{}, errors.New("lee name not found")
}
