package types

import (
	"encoding/json"
	"strings"
)

// StringOrNumber 用于兼容上游既可能返回字符串也可能返回数值的字段
type StringOrNumber string

func (s *StringOrNumber) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && (b[0] == '"') {
		var v string
		if err := json.Unmarshal(b, &v); err != nil {
			return err
		}
		*s = StringOrNumber(v)
		return nil
	}
	*s = StringOrNumber(strings.TrimSpace(string(b)))
	return nil
}

func (s StringOrNumber) String() string {
	return string(s)
}
