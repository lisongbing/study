package main

import (
	"encoding/json"
	"fmt"
)

type Gold int32

type GG struct{
	*Gold
}

func (this *Gold)setGold(gold Gold){
	*this = gold
}

func (this *Gold)getGold()Gold{
	return *this
}


type Diamond int32

type DD struct{
	*Diamond
}

func (this *Diamond)setGold(gold Diamond){
	*this = gold
}

func (this *Diamond)getGold()Diamond{
	return *this
}

type Mashel struct{
	Gold
	Itf []interface{}
}
func main() {
	var gold Gold
	gold.setGold(Gold(32))
	var d Diamond
	d.setGold(Diamond(32))
	md :=Mashel{}
	m :=&Mashel{Gold:gold,Itf:[]interface{}{DD{&d}},}
	bytes, _ := json.Marshal(m)
	json.Unmarshal(bytes,&md)
	fmt.Printf("%#v",md)
}