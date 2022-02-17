package xcodec

import "encoding/json"

type Json struct{}

func (Json)Decode(data []byte, i interface{}) error {
	return json.Unmarshal(data,i)
}
func (Json)Encode(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}



const XcodecJson = "json"
func init()  {
	Register(XcodecJson,Json{})
}