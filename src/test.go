package main

import (
	"encoding/json"
	"fmt"
)

/*func main() {
	const jsonStream = `{"name":"ethancai", "fansCount": 9223372036854775807}`
	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	// UseNumber causes the Decoder to unmarshal a number into an interface{} as a Number instead of as a float64.
	decoder.UseNumber()
	var user interface{}
	if err := decoder.Decode(&user); err != nil {
		fmt.Println("error:", err)
		return
	}
	m := user.(map[string]interface{})
	fmt.Printf("map:%#v",m)
	fansCount := m["fansCount"]
	fmt.Printf("%+v \n", reflect.TypeOf(fansCount).PkgPath() + "." + reflect.TypeOf(fansCount).Name())
	v, err := fansCount.(json.Number).Int64()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%+v \n", v)
}*/
type RecvMsg struct {
	MsgType int32       `json:"msgtype"` //消息类型
	//MsgData interface{} `json:"msgdata"` //消息数据(暂时用json格式encode/decode)
}
func main() {
	r :=&RecvMsg{1003}
	bytes, _ := json.Marshal(r)
	fmt.Println(bytes)
}