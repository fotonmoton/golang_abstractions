package serializers

import "encoding/json"

type StringSerializer struct{}

func (s *StringSerializer) Serialize(obj []string) string {
	str, _ := json.Marshal(obj)

	return string(str)
}
