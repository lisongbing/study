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

type DecompositionVersion struct {
	md5 string //配置表md5数据
}

type DecompositionClass struct {
	DecompositionVersion
}

var Decompositioninstance *DecompositionClass

func DecompositionInstance() *DecompositionClass {
	if Decompositioninstance == nil {
		Decompositioninstance = &DecompositionClass{}
	}
	return Decompositioninstance
}

type Decomposition struct {
	Id int32 //编号
	Level int32 //装备等级
	Quality int32 //装备品质
	Experience int32 //经验
	Gold int32 //金币
	ItemID1 int32 //道具1
	Number1 int32 //道具1数量
	ItemID2 int32 //道具2
	Number2 int32 //道具2数量
	ItemID3 int32 //道具3
	Number3 int32 //道具3数量
	ItemID4 int32 //道具4
	Number4 int32 //道具4数量
	ItemID5 int32 //道具5
	Number5 int32 //道具5数量
}

func GetDecompositionByPk(id int32) (itm *Decomposition, ok bool) {
	mtxDecomposition.RLock()
	itm, ok = cnfDecomposition[id]
	mtxDecomposition.RUnlock()
	return
}

func SetDecompositionVersion(md5 string) string {
	DecompositionInstance().md5 = md5
	return ``
	}
func GetDecompositionVersion(md5 string) string {
	return DecompositionInstance().md5
}
func GetDecomposition() map[int32]*Decomposition{
	mtxDecomposition.RLock()
	cnf := cnfDecomposition
	mtxDecomposition.RUnlock()
	return cnf
}

func (this *Decomposition) getLevel() int32 {
	return this.Level 
}

func (this *Decomposition) getQuality() int32 {
	return this.Quality 
}

func (this *Decomposition) getExperience() int32 {
	return this.Experience 
}

func (this *Decomposition) getGold() int32 {
	return this.Gold 
}

func (this *Decomposition) getItemID1() int32 {
	return this.ItemID1 
}

func (this *Decomposition) getNumber1() int32 {
	return this.Number1 
}

func (this *Decomposition) getItemID2() int32 {
	return this.ItemID2 
}

func (this *Decomposition) getNumber2() int32 {
	return this.Number2 
}

func (this *Decomposition) getItemID3() int32 {
	return this.ItemID3 
}

func (this *Decomposition) getNumber3() int32 {
	return this.Number3 
}

func (this *Decomposition) getItemID4() int32 {
	return this.ItemID4 
}

func (this *Decomposition) getNumber4() int32 {
	return this.Number4 
}

func (this *Decomposition) getItemID5() int32 {
	return this.ItemID5 
}

func (this *Decomposition) getNumber5() int32 {
	return this.Number5 
}

func LoadDecomposition(file string) string {
	var clen = []int32{15}
	sf := `decomposition.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*Decomposition)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &Decomposition{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Level, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Quality, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Experience, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Gold, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.ItemID1, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.Number1, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.ItemID2, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.Number2, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.ItemID3, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 9, val)
		}
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.Number3, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 10, val)
		}
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		itm.ItemID4, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 11, val)
		}
		val = strings.Replace(row.Cells[12].String(), " \t\r\n", ``, -1)
		itm.Number4, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 12, val)
		}
		val = strings.Replace(row.Cells[13].String(), " \t\r\n", ``, -1)
		itm.ItemID5, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 13, val)
		}
		val = strings.Replace(row.Cells[14].String(), " \t\r\n", ``, -1)
		itm.Number5, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 14, val)
		}
		cnf[itm.Id] = itm
	}
	mtxDecomposition.Lock()
	cnfDecomposition = cnf
	mtxDecomposition.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxDecomposition = new(sync.RWMutex)
var cnfDecomposition = map[int32]*Decomposition{
	10001: &Decomposition{
		10001, 100, 1, 100, 100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	10002: &Decomposition{
		10002, 100, 2, 200, 200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	10003: &Decomposition{
		10003, 100, 3, 300, 300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	10004: &Decomposition{
		10004, 100, 4, 400, 400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	10005: &Decomposition{
		10005, 100, 5, 500, 500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	10006: &Decomposition{
		10006, 100, 6, 600, 600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	20001: &Decomposition{
		20001, 200, 1, 700, 700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	20002: &Decomposition{
		20002, 200, 2, 800, 800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	20003: &Decomposition{
		20003, 200, 3, 900, 900, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	20004: &Decomposition{
		20004, 200, 4, 1000, 1000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	20005: &Decomposition{
		20005, 200, 5, 1100, 1100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	20006: &Decomposition{
		20006, 200, 6, 1200, 1200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	30001: &Decomposition{
		30001, 300, 1, 1300, 1300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	30002: &Decomposition{
		30002, 300, 2, 1400, 1400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	30003: &Decomposition{
		30003, 300, 3, 1500, 1500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	30004: &Decomposition{
		30004, 300, 4, 1600, 1600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	30005: &Decomposition{
		30005, 300, 5, 1700, 1700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	30006: &Decomposition{
		30006, 300, 6, 1800, 1800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	40001: &Decomposition{
		40001, 400, 1, 1900, 1900, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	40002: &Decomposition{
		40002, 400, 2, 2000, 2000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	40003: &Decomposition{
		40003, 400, 3, 2100, 2100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	40004: &Decomposition{
		40004, 400, 4, 2200, 2200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	40005: &Decomposition{
		40005, 400, 5, 2300, 2300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	40006: &Decomposition{
		40006, 400, 6, 2400, 2400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	50001: &Decomposition{
		50001, 500, 1, 2500, 2500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	50002: &Decomposition{
		50002, 500, 2, 2600, 2600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	50003: &Decomposition{
		50003, 500, 3, 2700, 2700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	50004: &Decomposition{
		50004, 500, 4, 2800, 2800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	50005: &Decomposition{
		50005, 500, 5, 2900, 2900, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	50006: &Decomposition{
		50006, 500, 6, 3000, 3000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	60001: &Decomposition{
		60001, 600, 1, 3100, 3100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	60002: &Decomposition{
		60002, 600, 2, 3200, 3200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	60003: &Decomposition{
		60003, 600, 3, 3300, 3300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	60004: &Decomposition{
		60004, 600, 4, 3400, 3400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	60005: &Decomposition{
		60005, 600, 5, 3500, 3500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	60006: &Decomposition{
		60006, 600, 6, 3600, 3600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	70001: &Decomposition{
		70001, 700, 1, 3700, 3700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	70002: &Decomposition{
		70002, 700, 2, 3800, 3800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	70003: &Decomposition{
		70003, 700, 3, 3900, 3900, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	70004: &Decomposition{
		70004, 700, 4, 4000, 4000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	70005: &Decomposition{
		70005, 700, 5, 4100, 4100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	70006: &Decomposition{
		70006, 700, 6, 4200, 4200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	80001: &Decomposition{
		80001, 800, 1, 4300, 4300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	80002: &Decomposition{
		80002, 800, 2, 4400, 4400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	80003: &Decomposition{
		80003, 800, 3, 4500, 4500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	80004: &Decomposition{
		80004, 800, 4, 4600, 4600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	80005: &Decomposition{
		80005, 800, 5, 4700, 4700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	80006: &Decomposition{
		80006, 800, 6, 4800, 4800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	90001: &Decomposition{
		90001, 900, 1, 4900, 4900, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	90002: &Decomposition{
		90002, 900, 2, 5000, 5000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	90003: &Decomposition{
		90003, 900, 3, 5100, 5100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	90004: &Decomposition{
		90004, 900, 4, 5200, 5200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	90005: &Decomposition{
		90005, 900, 5, 5300, 5300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	90006: &Decomposition{
		90006, 900, 6, 5400, 5400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	100001: &Decomposition{
		100001, 1000, 1, 5500, 5500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	100002: &Decomposition{
		100002, 1000, 2, 5600, 5600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	100003: &Decomposition{
		100003, 1000, 3, 5700, 5700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	100004: &Decomposition{
		100004, 1000, 4, 5800, 5800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	100005: &Decomposition{
		100005, 1000, 5, 5900, 5900, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	100006: &Decomposition{
		100006, 1000, 6, 6000, 6000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	110001: &Decomposition{
		110001, 1100, 1, 6100, 6100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	110002: &Decomposition{
		110002, 1100, 2, 6200, 6200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	110003: &Decomposition{
		110003, 1100, 3, 6300, 6300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	110004: &Decomposition{
		110004, 1100, 4, 6400, 6400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	110005: &Decomposition{
		110005, 1100, 5, 6500, 6500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	110006: &Decomposition{
		110006, 1100, 6, 6600, 6600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	120001: &Decomposition{
		120001, 1200, 1, 6700, 6700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	120002: &Decomposition{
		120002, 1200, 2, 6800, 6800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	120003: &Decomposition{
		120003, 1200, 3, 6900, 6900, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	120004: &Decomposition{
		120004, 1200, 4, 7000, 7000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	120005: &Decomposition{
		120005, 1200, 5, 7100, 7100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	120006: &Decomposition{
		120006, 1200, 6, 7200, 7200, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	130001: &Decomposition{
		130001, 1300, 1, 7300, 7300, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	130002: &Decomposition{
		130002, 1300, 2, 7400, 7400, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	130003: &Decomposition{
		130003, 1300, 3, 7500, 7500, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	130004: &Decomposition{
		130004, 1300, 4, 7600, 7600, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	130005: &Decomposition{
		130005, 1300, 5, 7700, 7700, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
	130006: &Decomposition{
		130006, 1300, 6, 7800, 7800, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	},
}

	func Decomposition_hot() {
		for _, val := range cnfDecomposition	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}