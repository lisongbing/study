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

type CreateroleVersion struct {
	md5 string //配置表md5数据
}

type CreateroleClass struct {
	CreateroleVersion
}

var Createroleinstance *CreateroleClass

func CreateroleInstance() *CreateroleClass {
	if Createroleinstance == nil {
		Createroleinstance = &CreateroleClass{}
	}
	return Createroleinstance
}

type Createrole struct {
	Id int32 //职业ID
	BirthMap int32 //出生地
	XCoordinates int32 //X坐标
	YCoordinates int32 //Y坐标
	Level int32 //初始等级
	Task int32 //初始任务ID
	EquipId1 string //初始装备
	Item1 int32 //包裹物品1ID
	Number1 int32 //物品1数量
	Item2 int32 //包裹物品2ID
	Number2 int32 //物品2数量
	Item3 int32 //包裹物品3ID
	Number3 int32 //物品3数量
	Item4 int32 //包裹物品4ID
	Number4 int32 //物品4数量
	Item5 int32 //包裹物品5ID
	Number5 int32 //物品5数量
}

func GetCreateroleByPk(id int32) (itm *Createrole, ok bool) {
	mtxCreaterole.RLock()
	itm, ok = cnfCreaterole[id]
	mtxCreaterole.RUnlock()
	return
}

func SetCreateroleVersion(md5 string) string {
	CreateroleInstance().md5 = md5
	return ``
	}
func GetCreateroleVersion(md5 string) string {
	return CreateroleInstance().md5
}
func GetCreaterole() map[int32]*Createrole{
	mtxCreaterole.RLock()
	cnf := cnfCreaterole
	mtxCreaterole.RUnlock()
	return cnf
}

func (this *Createrole) getBirthMap() int32 {
	return this.BirthMap 
}

func (this *Createrole) getXCoordinates() int32 {
	return this.XCoordinates 
}

func (this *Createrole) getYCoordinates() int32 {
	return this.YCoordinates 
}

func (this *Createrole) getLevel() int32 {
	return this.Level 
}

func (this *Createrole) getTask() int32 {
	return this.Task 
}

func (this *Createrole) getEquipId1() string {
	return this.EquipId1 
}

func (this *Createrole) getItem1() int32 {
	return this.Item1 
}

func (this *Createrole) getNumber1() int32 {
	return this.Number1 
}

func (this *Createrole) getItem2() int32 {
	return this.Item2 
}

func (this *Createrole) getNumber2() int32 {
	return this.Number2 
}

func (this *Createrole) getItem3() int32 {
	return this.Item3 
}

func (this *Createrole) getNumber3() int32 {
	return this.Number3 
}

func (this *Createrole) getItem4() int32 {
	return this.Item4 
}

func (this *Createrole) getNumber4() int32 {
	return this.Number4 
}

func (this *Createrole) getItem5() int32 {
	return this.Item5 
}

func (this *Createrole) getNumber5() int32 {
	return this.Number5 
}

func LoadCreaterole(file string) string {
	var clen = []int32{17}
	sf := `createrole.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*Createrole)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &Createrole{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.BirthMap, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.XCoordinates, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.YCoordinates, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Level, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.Task, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.EquipId1 = val
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.Item1, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.Number1, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.Item2, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 9, val)
		}
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.Number2, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 10, val)
		}
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		itm.Item3, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 11, val)
		}
		val = strings.Replace(row.Cells[12].String(), " \t\r\n", ``, -1)
		itm.Number3, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 12, val)
		}
		val = strings.Replace(row.Cells[13].String(), " \t\r\n", ``, -1)
		itm.Item4, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 13, val)
		}
		val = strings.Replace(row.Cells[14].String(), " \t\r\n", ``, -1)
		itm.Number4, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 14, val)
		}
		val = strings.Replace(row.Cells[15].String(), " \t\r\n", ``, -1)
		itm.Item5, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 15, val)
		}
		val = strings.Replace(row.Cells[16].String(), " \t\r\n", ``, -1)
		itm.Number5, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 16, val)
		}
		cnf[itm.Id] = itm
	}
	mtxCreaterole.Lock()
	cnfCreaterole = cnf
	mtxCreaterole.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxCreaterole = new(sync.RWMutex)
var cnfCreaterole = map[int32]*Createrole{
	1: &Createrole{
		1, 11001, 1318, 1927, 1, 10001, "40101:1;40102:1", 40003, 1, 40004, 1, 40005, 1, 40006, 1, 0, 0, 
	},
	2: &Createrole{
		2, 11001, 1318, 1927, 1, 10001, "40103:1", 40003, 1, 40004, 1, 40005, 1, 40006, 1, 0, 0, 
	},
	3: &Createrole{
		3, 11001, 1318, 1927, 1, 10001, "40100:1", 40003, 1, 40004, 1, 40005, 1, 40006, 1, 0, 0, 
	},
}

	func Createrole_hot() {
		for _, val := range cnfCreaterole	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}