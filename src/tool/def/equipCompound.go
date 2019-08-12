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

type EquipCompoundVersion struct {
	md5 string //配置表md5数据
}

type EquipCompoundClass struct {
	EquipCompoundVersion
}

var EquipCompoundinstance *EquipCompoundClass

func EquipCompoundInstance() *EquipCompoundClass {
	if EquipCompoundinstance == nil {
		EquipCompoundinstance = &EquipCompoundClass{}
	}
	return EquipCompoundinstance
}

type EquipCompound struct {
	Id int32 //阶级
	TitleName string //名字
	GoldCost int32 //金币消耗
	UnlockType int32 //解锁方式
	UnlockParam int32 //参数
}

func GetEquipCompoundByPk(id int32) (itm *EquipCompound, ok bool) {
	mtxEquipCompound.RLock()
	itm, ok = cnfEquipCompound[id]
	mtxEquipCompound.RUnlock()
	return
}

func SetEquipCompoundVersion(md5 string) string {
	EquipCompoundInstance().md5 = md5
	return ``
	}
func GetEquipCompoundVersion(md5 string) string {
	return EquipCompoundInstance().md5
}
func GetEquipCompound() map[int32]*EquipCompound{
	mtxEquipCompound.RLock()
	cnf := cnfEquipCompound
	mtxEquipCompound.RUnlock()
	return cnf
}

func (this *EquipCompound) getTitleName() string {
	return this.TitleName 
}

func (this *EquipCompound) getGoldCost() int32 {
	return this.GoldCost 
}

func (this *EquipCompound) getUnlockType() int32 {
	return this.UnlockType 
}

func (this *EquipCompound) getUnlockParam() int32 {
	return this.UnlockParam 
}

func LoadEquipCompound(file string) string {
	var clen = []int32{5}
	sf := `equipCompound.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*EquipCompound)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &EquipCompound{}
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
		cnf[itm.Id] = itm
	}
	mtxEquipCompound.Lock()
	cnfEquipCompound = cnf
	mtxEquipCompound.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxEquipCompound = new(sync.RWMutex)
var cnfEquipCompound = map[int32]*EquipCompound{
	1: &EquipCompound{
		1, "一阶", 2500, 1, 10, 
	},
	2: &EquipCompound{
		2, "二阶", 5000, 1, 20, 
	},
	3: &EquipCompound{
		3, "三阶", 7500, 1, 30, 
	},
	4: &EquipCompound{
		4, "四阶", 10000, 1, 40, 
	},
	5: &EquipCompound{
		5, "五阶", 12500, 1, 50, 
	},
	6: &EquipCompound{
		6, "六阶", 15000, 1, 60, 
	},
	7: &EquipCompound{
		7, "七阶", 17500, 1, 70, 
	},
	8: &EquipCompound{
		8, "八阶", 20000, 1, 100, 
	},
	9: &EquipCompound{
		9, "九阶", 22500, 2, 1000, 
	},
	10: &EquipCompound{
		10, "十阶", 25000, 2, 1500, 
	},
	11: &EquipCompound{
		11, "十一阶", 27500, 2, 2000, 
	},
	12: &EquipCompound{
		12, "十二阶", 30000, 2, 2500, 
	},
	13: &EquipCompound{
		13, "十三阶", 32500, 2, 3000, 
	},
}

	func EquipCompound_hot() {
		for _, val := range cnfEquipCompound	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}