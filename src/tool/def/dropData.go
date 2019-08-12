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

type DropDataVersion struct {
	md5 string //配置表md5数据
}

type DropDataClass struct {
	DropDataVersion
}

var DropDatainstance *DropDataClass

func DropDataInstance() *DropDataClass {
	if DropDatainstance == nil {
		DropDatainstance = &DropDataClass{}
	}
	return DropDatainstance
}

type DropData struct {
	Id int32 //
	NoDrop int32 //不掉率
	Pick int32 //执行次数
	GoldMin int32 //最小值
	GoldMax int32 //最大值
	GoldWeight int32 //金币权重
	GoodsPack string //道具id
	EquipPack string //装备包
	DropPack string //掉落id
}

func GetDropDataByPk(id int32) (itm *DropData, ok bool) {
	mtxDropData.RLock()
	itm, ok = cnfDropData[id]
	mtxDropData.RUnlock()
	return
}

func SetDropDataVersion(md5 string) string {
	DropDataInstance().md5 = md5
	return ``
	}
func GetDropDataVersion(md5 string) string {
	return DropDataInstance().md5
}
func GetDropData() map[int32]*DropData{
	mtxDropData.RLock()
	cnf := cnfDropData
	mtxDropData.RUnlock()
	return cnf
}

func (this *DropData) getNoDrop() int32 {
	return this.NoDrop 
}

func (this *DropData) getPick() int32 {
	return this.Pick 
}

func (this *DropData) getGoldMin() int32 {
	return this.GoldMin 
}

func (this *DropData) getGoldMax() int32 {
	return this.GoldMax 
}

func (this *DropData) getGoldWeight() int32 {
	return this.GoldWeight 
}

func (this *DropData) getGoodsPack() string {
	return this.GoodsPack 
}

func (this *DropData) getEquipPack() string {
	return this.EquipPack 
}

func (this *DropData) getDropPack() string {
	return this.DropPack 
}

func LoadDropData(file string) string {
	var clen = []int32{12}
	sf := `dropData.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*DropData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &DropData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.NoDrop, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Pick, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.GoldMin, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.GoldMax, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.GoldWeight, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.GoodsPack = val
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.EquipPack = val
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.DropPack = val
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		cnf[itm.Id] = itm
	}
	mtxDropData.Lock()
	cnfDropData = cnf
	mtxDropData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxDropData = new(sync.RWMutex)
var cnfDropData = map[int32]*DropData{
	10000: &DropData{
		10000, 200, 5, 1000, 2000, 100, "1:100;2:200", "10:100;20:200", "1001:100", 
	},
	10001: &DropData{
		10001, 200, 2, 10000, 20000, 500, "", "10:100", "", 
	},
}

	func DropData_hot() {
		for _, val := range cnfDropData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}