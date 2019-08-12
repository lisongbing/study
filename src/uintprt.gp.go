package main

import (
	"fmt"
	"unsafe"
	"io"
	"crypto/rand"
)

type BsonItem struct {
	TpId    int32  `bson:"tpid"  json:"tpid" `
	Pos     int32  `bson:"pos"  json:"pos" `
	BagType int32  `bson:"bagType"  json:"bagType" `
	Counts  int32  `bson:"counts"  json:"counts" ` //数量
	Lv      int32  `bson:"lv"  json:"lv" `         //等级
	Color   int32  `bson:"color"  json:"color" `   //颜色
	Times   int64  `bson:"times"  json:"times" `   //有效时间 截止时间
	ItemId  string `bson:"itemid"  json:"itemid" ` //道具id
}

func main() {
	var b []byte = []byte{'a', 'b', 'c'}
	var c *byte = &b[0]
	bitem := BsonItem{

		55000,
		65535,
		0,
		1,
		0,
		0,
		2,
		"5cdccd8017d1687946ed57e9",
	}

	var s = "5cdccd8017d1687946ed57e9"
	var ss = "5"
	i := []byte(s)
	ii := []byte(ss)
	fmt.Printf("%c\n", *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1))))
	fmt.Println(unsafe.Sizeof(bitem))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(len(i))
	fmt.Println(unsafe.Sizeof(ii))
	fmt.Println(len(ii))

	id :=make([]byte,3)
	_, err2 := io.ReadFull(rand.Reader, id)
	if err2 != nil {
		panic(fmt.Errorf("cannot get hostname: %v; %v", err2))
	}
	fmt.Println(id)
}
