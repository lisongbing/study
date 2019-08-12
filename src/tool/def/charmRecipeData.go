package def

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/tealeg/xlsx"
	"reflect"
	"os/exec"
	"path/filepath"
)

type CharmsRecipeDataVersion struct {
	md5 string //配置表md5数据
}

type CharmsRecipeDataClass struct {
	CharmsRecipeDataVersion
}

var CharmsRecipeDatainstance *CharmsRecipeDataClass

func CharmsRecipeDataInstance() *CharmsRecipeDataClass {
	if CharmsRecipeDatainstance == nil {
		CharmsRecipeDatainstance = &CharmsRecipeDataClass{}
	}
	return CharmsRecipeDatainstance
}

type CharmsRecipeData struct {
	Id int32 //id
	Type int32 //类型
	TitleName string //titleName
	Note string //note
	CraftItemNeed string //craftItemNeed
	GoldCost int32 //goldCost
	CraftResult int32 //craftResult
	UnlockType int32 //unlockType
	UnlockParam int32 //unlockParam
}

func GetCharmsRecipeDataByPk(id int32) (itm *CharmsRecipeData, ok bool) {
	mtxCharmsRecipeData.RLock()
	itm, ok = cnfCharmsRecipeData[id]
	mtxCharmsRecipeData.RUnlock()
	return
}

const (
	Type_SmallRecipe=1 // 小护身符
	Type_BigRecipe =2 //大护身符
	Type_HugeRecipe =3 //超大护身符
)

const (
	UnlockType_None = 0 //无需解锁
	UnlockType_CharLv = 1 //角色等级
)

func SetCharmsRecipeDataVersion(md5 string) string {
	CharmsRecipeDataInstance().md5 = md5
	return ``
	}
func GetCharmsRecipeDataVersion(md5 string) string {
	return CharmsRecipeDataInstance().md5
}
func GetCharmsRecipeData() map[int32]*CharmsRecipeData{
	mtxCharmsRecipeData.RLock()
	cnf := cnfCharmsRecipeData
	mtxCharmsRecipeData.RUnlock()
	return cnf
}

func (this *CharmsRecipeData) getType() int32 {
	return this.Type 
}

func (this *CharmsRecipeData) getTitleName() string {
	return this.TitleName 
}

func (this *CharmsRecipeData) getNote() string {
	return this.Note 
}

func (this *CharmsRecipeData) getCraftItemNeed() string {
	return this.CraftItemNeed 
}

func (this *CharmsRecipeData) getGoldCost() int32 {
	return this.GoldCost 
}

func (this *CharmsRecipeData) getCraftResult() int32 {
	return this.CraftResult 
}

func (this *CharmsRecipeData) getUnlockType() int32 {
	return this.UnlockType 
}

func (this *CharmsRecipeData) getUnlockParam() int32 {
	return this.UnlockParam 
}

func LoadCharmsRecipeData(file string) string {
	var clen = []int32{9}
	sf := `charmRecipeData.xlsx`
	fi, _ := exec.LookPath(os.Args[0])
	pa, _ := filepath.Abs(fi)
	rst := filepath.Dir(pa)
	sf = rst + "/" +"res"+"/" + filepath.Base(sf)
    if file != "" {
        sf = file
    } else {
        _, err := os.Lstat(sf)
        if err != nil && os.IsNotExist(err) {
            sf = "xlsx/" + sf
        }
    }
    f, err := xlsx.OpenFile(sf)
    if err != nil {
        return err.Error()
    }
    if len(f.Sheets[0].Rows) < 5 {
        return sf + " 没有配置内容"
    }
    //slen := int32(len(f.Sheets))
    var rlen []int32
    for _, sheet := range f.Sheets {
		if sheet.Name == "注释页" {
			continue
		}
        rlen = append(rlen, int32(len(sheet.Rows)))
    }
    //flen := int32(len(f.Sheets[0].Rows[3].Cells))
    var val string
    var ok bool
    var r64 int32
    var f64 float64
    _, _, _ = ok, r64, f64
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*CharmsRecipeData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &CharmsRecipeData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Type, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.TitleName = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Note = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.CraftItemNeed = val
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.GoldCost, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.CraftResult, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.UnlockType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.UnlockParam, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		cnf[itm.Id] = itm
	}
	mtxCharmsRecipeData.Lock()
	cnfCharmsRecipeData = cnf
	mtxCharmsRecipeData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxCharmsRecipeData = new(sync.RWMutex)
var cnfCharmsRecipeData = map[int32]*CharmsRecipeData{
	1: &CharmsRecipeData{
		1, 1, "小护身符", "", "1001:1;1002:1;1003:1;1004:1", 2500, 15001, 1, 15, 
	},
	2: &CharmsRecipeData{
		2, 2, "大护身符", "", "1001:1;1002:1;1003:1;1005:1", 5000, 15002, 1, 30, 
	},
	3: &CharmsRecipeData{
		3, 3, "超大护身符", "", "1001:1;1002:1;1003:1;1006:1", 10000, 15002, 1, 60, 
	},
}

	func CharmsRecipeData_hot() {
		for _, val := range cnfCharmsRecipeData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}