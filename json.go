package main

import (
	"net/http"
	"go-simplejson"
	"fmt"
	"io"
)

type Person struct{
	Name string
	Sex  bool
}
func main() {
/*	testjson :=[]string{"li","song","bing"}

	fmt.Println(testjson)
	person :=Person{"lisongb",true}
	b:= byte(person)
	const myjson =`{"name":"lisong","sex":"true"}`
	fmt.Println(simplejson.Version())
	jsondata := simplejson.NewFromReader()
	jsondata.UnmarshalJSON()
	fmt.Printf("%#v\n",jsondata)
	if v,ok := jsondata.Interface().(map[string]interface{});ok{
		fmt.Println(v)
	}
	json, e := simplejson.NewFromReader(strings.NewReader(myjson))
	fmt.Println("json,e",json,e)
	jsonMap := json.Interface().(map[string]interface{})
	fmt.Println(jsonMap["name"])

	newJson, _ := simplejson.NewJson([]byte(myjson))
	fmt.Println("newJson",newJson)
	m, err := newJson.Map()
	fmt.Println("array",m,err)
	get, b := newJson.CheckGet("name")
	fmt.Println("get",get,b)

	namejson := newJson.Get("name")
	marshalJSON, _ := namejson.MarshalJSON()
	fmt.Println(marshalJSON)
	fmt.Println("namejson",namejson)
	newJson.Del("name")
	fmt.Println("newJson",newJson)

	bytes, _ := newJson.Encode()
	fmt.Println(bytes,string(bytes))

	pretty, _ := newJson.EncodePretty()
	fmt.Println(string(pretty))

	f, _ := newJson.Float64()
	fmt.Println("float",f)*/
	http.HandleFunc("/json",json)
	http.ListenAndServe("127.0.0.1:8080",nil)

}
func json(w http.ResponseWriter, r *http.Request){
	p :=make([]byte,1024)
	n, err := r.Body.Read(p)
	if err !=nil&&err !=io.EOF{
		fmt.Println("err0:",err)
		return
	}
	b :=p[0:n]
	newJson, err := simplejson.NewJson(b)
	if err !=nil{
		fmt.Println("err1:",err)
	}
	fmt.Println("newJson",newJson)
	data, _ := newJson.Map()

	for key,_ :=range data{
		array, _ := newJson.Get(key).Array()
		fmt.Println(len(array))
		for _,a :=range array{
			_ = a.(map[string]interface{})
		}
	}

/*	for key,value:=range  data{
		fmt.Printf("%T,%#v",key,key)
		fmt.Printf("%T,%#v",value,value)
	}*/
	fmt.Println(data)
	bytes, err := newJson.MarshalJSON()
	if err !=nil{
		fmt.Println("err2:",err)
	}
	w.Write(bytes)
}