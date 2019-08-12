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

type DropGoodsVersion struct {
	md5 string //配置表md5数据
}

type DropGoodsClass struct {
	DropGoodsVersion
}

var DropGoodsinstance *DropGoodsClass

func DropGoodsInstance() *DropGoodsClass {
	if DropGoodsinstance == nil {
		DropGoodsinstance = &DropGoodsClass{}
	}
	return DropGoodsinstance
}

type DropGoods struct {
	Id int32 //
	SpecifyClass int32 //指定职业
	GoodsGet string //获得的道具
}

func GetDropGoodsByPk(id int32) (itm *DropGoods, ok bool) {
	mtxDropGoods.RLock()
	itm, ok = cnfDropGoods[id]
	mtxDropGoods.RUnlock()
	return
}

func SetDropGoodsVersion(md5 string) string {
	DropGoodsInstance().md5 = md5
	return ``
	}
func GetDropGoodsVersion(md5 string) string {
	return DropGoodsInstance().md5
}
func GetDropGoods() map[int32]*DropGoods{
	mtxDropGoods.RLock()
	cnf := cnfDropGoods
	mtxDropGoods.RUnlock()
	return cnf
}

func (this *DropGoods) getSpecifyClass() int32 {
	return this.SpecifyClass 
}

func (this *DropGoods) getGoodsGet() string {
	return this.GoodsGet 
}

func LoadDropGoods(file string) string {
	var clen = []int32{7}
	sf := `dropGoods.xlsx`
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
	cnf := make(map[int32]*DropGoods)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &DropGoods{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.SpecifyClass, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.GoodsGet = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		cnf[itm.Id] = itm
	}
	mtxDropGoods.Lock()
	cnfDropGoods = cnf
	mtxDropGoods.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxDropGoods = new(sync.RWMutex)
var cnfDropGoods = map[int32]*DropGoods{
	1: &DropGoods{
		1, 0, "40003:10:20:100;40004:20:30:200;40005:20:40:400", 
	},
	2: &DropGoods{
		2, 0, "40005:5:10:200;40006:80:100:900", 
	},
}

	func DropGoods_hot() {
		for _, val := range cnfDropGoods	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}