package leegetter

import (
	"encoding/json"
	"errors"
	"strings"
)

var leeList []Lee

type Lee struct {
	Name string "json:\"name\""
	Rgb  string "json:\"rgb\""
}

func init() {
	_ = json.Unmarshal([]byte(LeeJsonString), &leeList)
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
