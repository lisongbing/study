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

type EquipUpgradeVersion struct {
	md5 string //配置表md5数据
}

type EquipUpgradeClass struct {
	EquipUpgradeVersion
}

var EquipUpgradeinstance *EquipUpgradeClass

func EquipUpgradeInstance() *EquipUpgradeClass {
	if EquipUpgradeinstance == nil {
		EquipUpgradeinstance = &EquipUpgradeClass{}
	}
	return EquipUpgradeinstance
}

type EquipUpgrade struct {
	Id int32 //阶级
	TitleName string //名字
	GoldCost int32 //金币消耗
	UnlockType int32 //解锁方式
	UnlockParam int32 //参数
}

func GetEquipUpgradeByPk(id int32) (itm *EquipUpgrade, ok bool) {
	mtxEquipUpgrade.RLock()
	itm, ok = cnfEquipUpgrade[id]
	mtxEquipUpgrade.RUnlock()
	return
}

func SetEquipUpgradeVersion(md5 string) string {
	EquipUpgradeInstance().md5 = md5
	return ``
	}
func GetEquipUpgradeVersion(md5 string) string {
	return EquipUpgradeInstance().md5
}
func GetEquipUpgrade() map[int32]*EquipUpgrade{
	mtxEquipUpgrade.RLock()
	cnf := cnfEquipUpgrade
	mtxEquipUpgrade.RUnlock()
	return cnf
}

func (this *EquipUpgrade) getTitleName() string {
	return this.TitleName 
}

func (this *EquipUpgrade) getGoldCost() int32 {
	return this.GoldCost 
}

func (this *EquipUpgrade) getUnlockType() int32 {
	return this.UnlockType 
}

func (this *EquipUpgrade) getUnlockParam() int32 {
	return this.UnlockParam 
}

func LoadEquipUpgrade(file string) string {
	var clen = []int32{11}
	sf := `equipUpgrade.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*EquipUpgrade)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &EquipUpgrade{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.TitleName = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.GoldCost, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.UnlockType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.UnlockParam, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		cnf[itm.Id] = itm
	}
	mtxEquipUpgrade.Lock()
	cnfEquipUpgrade = cnf
	mtxEquipUpgrade.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxEquipUpgrade = new(sync.RWMutex)
var cnfEquipUpgrade = map[int32]*EquipUpgrade{
	8: &EquipUpgrade{
		8, "八阶", 2500, 1, 1000, 
	},
	9: &EquipUpgrade{
		9, "九阶", 5000, 2, 1200, 
	},
}

	func EquipUpgrade_hot() {
		for _, val := range cnfEquipUpgrade	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}