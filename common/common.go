package common

import (
	"github.com/BurntSushi/toml"
	json "github.com/json-iterator/go"
)

func JsonBytesToMap(d []byte) map[string]interface{} {
	m := make(map[string]interface{})
	json.Unmarshal(d, &m)
	return m
}

func MapToJsonBytes(m map[string]interface{}) []byte {
	d, _ := json.Marshal(m)
	return d
}

func UnmarshalTomlFromFile(filePath string, dst interface{}) error {
	_, err := toml.DecodeFile(filePath, dst)
	return err
}

func UnmarshalToml(data []byte, dst interface{}) error {
	return toml.Unmarshal(data, dst)
}
