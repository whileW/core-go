package xcodec

import "gopkg.in/yaml.v3"

type Yaml struct{}

func (Yaml)Decode(data []byte, i interface{}) error {
	return yaml.Unmarshal(data,i)
}
func (Yaml)Encode(i interface{}) ([]byte, error) {
	return yaml.Marshal(i)
}



const XcodecYaml = "yaml"
func init()  {
	Register(XcodecYaml,Yaml{})
}
