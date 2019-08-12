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

type InnateAffixesVersion struct {
	md5 string //配置表md5数据
}

type InnateAffixesClass struct {
	InnateAffixesVersion
}

var InnateAffixesinstance *InnateAffixesClass

func InnateAffixesInstance() *InnateAffixesClass {
	if InnateAffixesinstance == nil {
		InnateAffixesinstance = &InnateAffixesClass{}
	}
	return InnateAffixesinstance
}

type InnateAffixes struct {
	Id int32 //
	Descr string //策划看
	Affix string //
}

func SetInnateAffixesVersion(md5 string) string {
	InnateAffixesInstance().md5 = md5
	return ``
	}
func GetInnateAffixesVersion(md5 string) string {
	return InnateAffixesInstance().md5
}
func GetInnateAffixes() []*InnateAffixes{
	mtxInnateAffixes.RLock()
	cnf := cnfInnateAffixes
	mtxInnateAffixes.RUnlock()
	return cnf
}

func (this *InnateAffixes) getDescr() string {
	return this.Descr 
}

func (this *InnateAffixes) getAffix() string {
	return this.Affix 
}

func LoadInnateAffixes(file string) string {
	var clen = []int32{4}
	sf := `innateAffixes.xlsx`
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
var cnf []*InnateAffixes
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &InnateAffixes{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Descr = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Affix = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		cnf = append(cnf, itm)
	}
	mtxInnateAffixes.Lock()
	cnfInnateAffixes = cnf
	mtxInnateAffixes.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxInnateAffixes = new(sync.RWMutex)
var cnfInnateAffixes = []*InnateAffixes{
	&InnateAffixes{
		1, "天生1", "1;2", 
	},
	&InnateAffixes{
		1, "天生2", "3;4", 
	},
}

	func InnateAffixes_hot() {
		for _, val := range cnfInnateAffixes	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}