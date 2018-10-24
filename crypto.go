package builtin

import (
	"encoding/base64"
	. "github.com/yatlabs/bigger"
)

func init() {


	Bigger.Crypto("percent", Map{
		"name": "百分比处理", "text": "百分比处理",
		"encode": func(value Any) Any {
			//text -> text
			if vv,ok := value.(float64); ok {
				return Bigger.Round(vv/100, 2)
			}
			return nil
		},
		"decode": func(value Any) Any {
			//data -> text
			if vv,ok := value.(float64); ok {
				return Bigger.Round(vv*100, 2)
			}
			return nil
		},
	}, false)


	Bigger.Crypto("base64", Map{
		"name": "BASE64加解密", "text": "BASE64加解密",
		"encode": func(value Any) Any {
			text := Bigger.ToString(value)
			return base64.StdEncoding.EncodeToString([]byte(text))
		},
		"decode": func(value Any) Any {
			text := Bigger.ToString(value)
			bytes,err := base64.StdEncoding.DecodeString(text)
			if err == nil {
				return string(bytes)
			}
			return nil
		},
	}, false)

}


