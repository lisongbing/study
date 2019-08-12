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

type RuneRecipeDataVersion struct {
	md5 string //配置表md5数据
}

type RuneRecipeDataClass struct {
	RuneRecipeDataVersion
}

var RuneRecipeDatainstance *RuneRecipeDataClass

func RuneRecipeDataInstance() *RuneRecipeDataClass {
	if RuneRecipeDatainstance == nil {
		RuneRecipeDatainstance = &RuneRecipeDataClass{}
	}
	return RuneRecipeDatainstance
}

type RuneRecipeData struct {
	Id int32 //
	Note string //注释
	CraftItemNeed string //合成材料
	GoldCost int32 //金币消耗
	CraftResult int32 //合成出的道具
	UnlockType int32 //解锁方式
	UnlockParam int32 //参数
}

func GetRuneRecipeDataByPk(id int32) (itm *RuneRecipeData, ok bool) {
	mtxRuneRecipeData.RLock()
	itm, ok = cnfRuneRecipeData[id]
	mtxRuneRecipeData.RUnlock()
	return
}

func SetRuneRecipeDataVersion(md5 string) string {
	RuneRecipeDataInstance().md5 = md5
	return ``
	}
func GetRuneRecipeDataVersion(md5 string) string {
	return RuneRecipeDataInstance().md5
}
func GetRuneRecipeData() map[int32]*RuneRecipeData{
	mtxRuneRecipeData.RLock()
	cnf := cnfRuneRecipeData
	mtxRuneRecipeData.RUnlock()
	return cnf
}

func (this *RuneRecipeData) getNote() string {
	return this.Note 
}

func (this *RuneRecipeData) getCraftItemNeed() string {
	return this.CraftItemNeed 
}

func (this *RuneRecipeData) getGoldCost() int32 {
	return this.GoldCost 
}

func (this *RuneRecipeData) getCraftResult() int32 {
	return this.CraftResult 
}

func (this *RuneRecipeData) getUnlockType() int32 {
	return this.UnlockType 
}

func (this *RuneRecipeData) getUnlockParam() int32 {
	return this.UnlockParam 
}

func LoadRuneRecipeData(file string) string {
	var clen = []int32{7}
	sf := `runeRecipeData.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*RuneRecipeData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &RuneRecipeData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Note = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.CraftItemNeed = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.GoldCost, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.CraftResult, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.UnlockType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.UnlockParam, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		cnf[itm.Id] = itm
	}
	mtxRuneRecipeData.Lock()
	cnfRuneRecipeData = cnf
	mtxRuneRecipeData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxRuneRecipeData = new(sync.RWMutex)
var cnfRuneRecipeData = map[int32]*RuneRecipeData{
	1: &RuneRecipeData{
		1, "2号神符", "10001:3", 2500, 10002, 1, 5, 
	},
	2: &RuneRecipeData{
		2, "3号神符", "10002:3", 5000, 10003, 1, 10, 
	},
	3: &RuneRecipeData{
		3, "4号神符", "10003:3", 10000, 10004, 1, 15, 
	},
	4: &RuneRecipeData{
		4, "5号神符", "10004:3", 15000, 10005, 1, 20, 
	},
	5: &RuneRecipeData{
		5, "6号神符", "10005:3", 20000, 10006, 1, 25, 
	},
	6: &RuneRecipeData{
		6, "7号神符", "10006:3", 25000, 10007, 1, 30, 
	},
	7: &RuneRecipeData{
		7, "8号神符", "10007:3", 30000, 10008, 1, 35, 
	},
	8: &RuneRecipeData{
		8, "9号神符", "10008:3", 35000, 10009, 1, 40, 
	},
	9: &RuneRecipeData{
		9, "10号神符", "10009:3", 40000, 10010, 1, 45, 
	},
}

	func RuneRecipeData_hot() {
		for _, val := range cnfRuneRecipeData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}