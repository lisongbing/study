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

type LoadTipsVersion struct {
	md5 string //配置表md5数据
}

type LoadTipsClass struct {
	LoadTipsVersion
}

var LoadTipsinstance *LoadTipsClass

func LoadTipsInstance() *LoadTipsClass {
	if LoadTipsinstance == nil {
		LoadTipsinstance = &LoadTipsClass{}
	}
	return LoadTipsinstance
}

type LoadTips struct {
	Id int32 //主键
	LowerLimit int32 //条件下限
	UpperLimit int32 //条件上限
	Content string //内容
}

func GetLoadTipsByPk(id int32) (itm *LoadTips, ok bool) {
	mtxLoadTips.RLock()
	itm, ok = cnfLoadTips[id]
	mtxLoadTips.RUnlock()
	return
}

func SetLoadTipsVersion(md5 string) string {
	LoadTipsInstance().md5 = md5
	return ``
	}
func GetLoadTipsVersion(md5 string) string {
	return LoadTipsInstance().md5
}
func GetLoadTips() map[int32]*LoadTips{
	mtxLoadTips.RLock()
	cnf := cnfLoadTips
	mtxLoadTips.RUnlock()
	return cnf
}

func (this *LoadTips) getLowerLimit() int32 {
	return this.LowerLimit 
}

func (this *LoadTips) getUpperLimit() int32 {
	return this.UpperLimit 
}

func (this *LoadTips) getContent() string {
	return this.Content 
}

func LoadLoadTips(file string) string {
	var clen = []int32{4}
	sf := `loadTips.xlsx`
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
    var shref = []int64{0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*LoadTips)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &LoadTips{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.LowerLimit, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.UpperLimit, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Content = val
		cnf[itm.Id] = itm
	}
	mtxLoadTips.Lock()
	cnfLoadTips = cnf
	mtxLoadTips.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxLoadTips = new(sync.RWMutex)
var cnfLoadTips = map[int32]*LoadTips{
	1001: &LoadTips{
		1001, 1, 30, "装备是增强人物属性的重要途径", 
	},
	1002: &LoadTips{
		1002, 1, 30, "挂机时间越长，收益越高哦", 
	},
	1003: &LoadTips{
		1003, 1, 70, "离线收益将在您本次进入游戏后发放", 
	},
	1004: &LoadTips{
		1004, 1, 30, "把不需要的装备分解后，获得的分解材料可以升级其他装备", 
	},
	1005: &LoadTips{
		1005, 1, 30, "部分物品可以进行出售，从而获得更多金币", 
	},
	1006: &LoadTips{
		1006, 1, 30, "经常在世界频道聊天，可以收获更多朋友哦", 
	},
}

	func LoadTips_hot() {
		for _, val := range cnfLoadTips	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}