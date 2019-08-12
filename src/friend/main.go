package main

import (
	"diabloserver/models"
	"fmt"
	"friend/gen"
	"github.com/aoeu/mgo/bson"
	"gopkg.in/mgo.v2"
	"time"
)

type Frdbase struct {
	Openid   string `bson:"openid"`
	RoleName string `bson:"rolename"`
}

var myself models.FriendInfos
var friends []Frdbase

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Println("dial:", err)
		return
	}
	database := mgo.Database{session, "cwsl_friend"}
	c := database.C("friends")

	err1 := c.Find(bson.M{}).Sort("-lv").All(&friends)
	if err1 != nil {
		fmt.Println("err1", err1)
		return
	}

	err2 := c.Find(bson.M{"rolename": "奥立芙·沃伦"}).One(&myself)
	if err2 != nil {
		fmt.Println("err2", err2)
	}
	fmt.Printf("%#v,\n%v\n", friends, len(friends))
	var current []int
	var recommendList []Frdbase
	r := time.Now().Unix()
	Start:
	current = gen.GenerateRandomNumber(&r, 0, len(friends), 10, current)
	fmt.Println(current, len(friends))

	fmt.Printf("myself:%#v\n", myself)
	for _,index :=range current{
		if friends[index].RoleName == myself.RoleName{
			goto END
		}
		for _,apply :=range myself.Applys{
			if apply.RoleName == friends[index].RoleName{
				goto END
			}
		}
		for _,friend :=range myself.Friends{
			if friend.RoleName == friends[index].RoleName{
				goto END
			}
		}
		recommendList = append(recommendList,friends[index])
		if len(recommendList) == 5{
			break
		}
		END:
	}
	if len(recommendList) < 5{
		goto Start
	}
	fmt.Printf("recommendList:%#v",recommendList)
}
