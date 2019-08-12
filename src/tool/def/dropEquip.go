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

type DropEquipVersion struct {
	md5 string //配置表md5数据
}

type DropEquipClass struct {
	DropEquipVersion
}

var DropEquipinstance *DropEquipClass

func DropEquipInstance() *DropEquipClass {
	if DropEquipinstance == nil {
		DropEquipinstance = &DropEquipClass{}
	}
	return DropEquipinstance
}

type DropEquip struct {
	Id int32 //
	DropType int32 //掉落方式
	SpecifyClass int32 //指定职业
	DropSource string //掉落来源
	WhiteWeight int32 //白品质权重
	BlueWeight int32 //蓝品质权重
	YellowWeight int32 //黄品质权重
	GlodenWeight int32 //暗金品质权重
	GreenWeight int32 //绿装品质权重
	AncientPr int32 //远古几率
	PrimalAncientPr int32 //太古几率
	SpecifyEquip string //指定装备
}

func GetDropEquipByPk(id int32) (itm *DropEquip, ok bool) {
	mtxDropEquip.RLock()
	itm, ok = cnfDropEquip[id]
	mtxDropEquip.RUnlock()
	return
}

const (
	DropType_RandDrop = 0  // 随机掉落
	DropType_SpecifyDrop = 1 //指定id掉落
)

const (
	SpecifyClass_AllClass = 0 // 全职业
	SpecifyClass_Crusader = 1 //圣教军
	SpecifyClass_DemonHunter = 2  //猎魔人
	SpecifyClass_Wizard = 3  //法师
)

func SetDropEquipVersion(md5 string) string {
	DropEquipInstance().md5 = md5
	return ``
	}
func GetDropEquipVersion(md5 string) string {
	return DropEquipInstance().md5
}
func GetDropEquip() map[int32]*DropEquip{
	mtxDropEquip.RLock()
	cnf := cnfDropEquip
	mtxDropEquip.RUnlock()
	return cnf
}

func (this *DropEquip) getDropType() int32 {
	return this.DropType 
}

func (this *DropEquip) getSpecifyClass() int32 {
	return this.SpecifyClass 
}

func (this *DropEquip) getDropSource() string {
	return this.DropSource 
}

func (this *DropEquip) getWhiteWeight() int32 {
	return this.WhiteWeight 
}

func (this *DropEquip) getBlueWeight() int32 {
	return this.BlueWeight 
}

func (this *DropEquip) getYellowWeight() int32 {
	return this.YellowWeight 
}

func (this *DropEquip) getGlodenWeight() int32 {
	return this.GlodenWeight 
}

func (this *DropEquip) getGreenWeight() int32 {
	return this.GreenWeight 
}

func (this *DropEquip) getAncientPr() int32 {
	return this.AncientPr 
}

func (this *DropEquip) getPrimalAncientPr() int32 {
	return this.PrimalAncientPr 
}

func (this *DropEquip) getSpecifyEquip() string {
	return this.SpecifyEquip 
}

func LoadDropEquip(file string) string {
	var clen = []int32{12}
	sf := `dropEquip.xlsx`
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
	cnf := make(map[int32]*DropEquip)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &DropEquip{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.DropType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.SpecifyClass, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.DropSource = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.WhiteWeight, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.BlueWeight, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.YellowWeight, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.GlodenWeight, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.GreenWeight, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.AncientPr, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 9, val)
		}
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.PrimalAncientPr, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 10, val)
		}
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		itm.SpecifyEquip = val
		cnf[itm.Id] = itm
	}
	mtxDropEquip.Lock()
	cnfDropEquip = cnf
	mtxDropEquip.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxDropEquip = new(sync.RWMutex)
var cnfDropEquip = map[int32]*DropEquip{
	10: &DropEquip{
		10, 0, 0, "1;2", 100, 200, 300, 50, 25, 10, 10, "", 
	},
	20: &DropEquip{
		20, 1, 0, "0", 0, 0, 0, 0, 0, 0, 0, "40001:0:100;40002:1:200;40003:2:300", 
	},
}

	func DropEquip_hot() {
		for _, val := range cnfDropEquip	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}