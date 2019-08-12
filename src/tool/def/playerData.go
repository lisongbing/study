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

type PlayerDataVersion struct {
	md5 string //配置表md5数据
}

type PlayerDataClass struct {
	PlayerDataVersion
}

var PlayerDatainstance *PlayerDataClass

func PlayerDataInstance() *PlayerDataClass {
	if PlayerDatainstance == nil {
		PlayerDatainstance = &PlayerDataClass{}
	}
	return PlayerDatainstance
}

type PlayerData struct {
	Id int32 //等级
	Exp int32 //升下级经验
	Attribute string //属性
}

func GetPlayerDataByPk(id int32) (itm *PlayerData, ok bool) {
	mtxPlayerData.RLock()
	itm, ok = cnfPlayerData[id]
	mtxPlayerData.RUnlock()
	return
}

func SetPlayerDataVersion(md5 string) string {
	PlayerDataInstance().md5 = md5
	return ``
	}
func GetPlayerDataVersion(md5 string) string {
	return PlayerDataInstance().md5
}
func GetPlayerData() map[int32]*PlayerData{
	mtxPlayerData.RLock()
	cnf := cnfPlayerData
	mtxPlayerData.RUnlock()
	return cnf
}

func (this *PlayerData) getExp() int32 {
	return this.Exp 
}

func (this *PlayerData) getAttribute() string {
	return this.Attribute 
}

func LoadPlayerData(file string) string {
	var clen = []int32{3}
	sf := `playerData.xlsx`
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
    var shref = []int64{0, 0, 0}
	_ = shref
	cnf := make(map[int32]*PlayerData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &PlayerData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Exp, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Attribute = val
		cnf[itm.Id] = itm
	}
	mtxPlayerData.Lock()
	cnfPlayerData = cnf
	mtxPlayerData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxPlayerData = new(sync.RWMutex)
var cnfPlayerData = map[int32]*PlayerData{
	1: &PlayerData{
		1, 100, "2:100000000;5:1000;45:200", 
	},
	2: &PlayerData{
		2, 200, "2:100000000;5:1000;45:200", 
	},
	3: &PlayerData{
		3, 300, "2:100000000;5:1000;45:200", 
	},
	4: &PlayerData{
		4, 400, "2:100000000;5:1000;45:200", 
	},
	5: &PlayerData{
		5, 500, "2:100000000;5:1000;45:200", 
	},
	6: &PlayerData{
		6, 600, "2:100000000;5:1000;45:200", 
	},
	7: &PlayerData{
		7, 700, "2:100000000;5:1000;45:200", 
	},
	8: &PlayerData{
		8, 800, "2:100000000;5:1000;45:200", 
	},
	9: &PlayerData{
		9, 900, "2:100000000;5:1000;45:200", 
	},
	10: &PlayerData{
		10, 1000, "2:100000000;5:1000;45:200", 
	},
	11: &PlayerData{
		11, 1100, "2:100000000;5:1000;45:200", 
	},
	12: &PlayerData{
		12, 1200, "2:100000000;5:1000;45:200", 
	},
	13: &PlayerData{
		13, 1300, "2:100000000;5:1000;45:200", 
	},
	14: &PlayerData{
		14, 1400, "2:100000000;5:1000;45:200", 
	},
	15: &PlayerData{
		15, 1500, "2:100000000;5:1000;45:200", 
	},
	16: &PlayerData{
		16, 1600, "2:100000000;5:1000;45:200", 
	},
	17: &PlayerData{
		17, 1700, "2:100000000;5:1000;45:200", 
	},
	18: &PlayerData{
		18, 1800, "2:100000000;5:1000;45:200", 
	},
	19: &PlayerData{
		19, 1900, "2:100000000;5:1000;45:200", 
	},
	20: &PlayerData{
		20, 2000, "2:100000000;5:1000;45:200", 
	},
	21: &PlayerData{
		21, 2100, "2:100000000;5:1000;45:200", 
	},
	22: &PlayerData{
		22, 2200, "2:100000000;5:1000;45:200", 
	},
	23: &PlayerData{
		23, 2300, "2:100000000;5:1000;45:200", 
	},
	24: &PlayerData{
		24, 2400, "2:100000000;5:1000;45:200", 
	},
	25: &PlayerData{
		25, 2500, "2:100000000;5:1000;45:200", 
	},
	26: &PlayerData{
		26, 2600, "2:100000000;5:1000;45:200", 
	},
	27: &PlayerData{
		27, 2700, "2:100000000;5:1000;45:200", 
	},
	28: &PlayerData{
		28, 2800, "2:100000000;5:1000;45:200", 
	},
	29: &PlayerData{
		29, 2900, "2:100000000;5:1000;45:200", 
	},
	30: &PlayerData{
		30, 3000, "2:100000000;5:1000;45:200", 
	},
	31: &PlayerData{
		31, 3100, "2:100000000;5:1000;45:200", 
	},
	32: &PlayerData{
		32, 3200, "2:100000000;5:1000;45:200", 
	},
	33: &PlayerData{
		33, 3300, "2:100000000;5:1000;45:200", 
	},
	34: &PlayerData{
		34, 3400, "2:100000000;5:1000;45:200", 
	},
	35: &PlayerData{
		35, 3500, "2:100000000;5:1000;45:200", 
	},
	36: &PlayerData{
		36, 3600, "2:100000000;5:1000;45:200", 
	},
	37: &PlayerData{
		37, 3700, "2:100000000;5:1000;45:200", 
	},
	38: &PlayerData{
		38, 3800, "2:100000000;5:1000;45:200", 
	},
	39: &PlayerData{
		39, 3900, "2:100000000;5:1000;45:200", 
	},
	40: &PlayerData{
		40, 4000, "2:100000000;5:1000;45:200", 
	},
	41: &PlayerData{
		41, 4100, "2:100000000;5:1000;45:200", 
	},
	42: &PlayerData{
		42, 4200, "2:100000000;5:1000;45:200", 
	},
	43: &PlayerData{
		43, 4300, "2:100000000;5:1000;45:200", 
	},
	44: &PlayerData{
		44, 4400, "2:100000000;5:1000;45:200", 
	},
	45: &PlayerData{
		45, 4500, "2:100000000;5:1000;45:200", 
	},
	46: &PlayerData{
		46, 4600, "2:100000000;5:1000;45:200", 
	},
	47: &PlayerData{
		47, 4700, "2:100000000;5:1000;45:200", 
	},
	48: &PlayerData{
		48, 4800, "2:100000000;5:1000;45:200", 
	},
	49: &PlayerData{
		49, 4900, "2:100000000;5:1000;45:200", 
	},
	50: &PlayerData{
		50, 5000, "2:100000000;5:1000;45:200", 
	},
	51: &PlayerData{
		51, 5100, "2:100000000;5:1000;45:200", 
	},
	52: &PlayerData{
		52, 5200, "2:100000000;5:1000;45:200", 
	},
	53: &PlayerData{
		53, 5300, "2:100000000;5:1000;45:200", 
	},
	54: &PlayerData{
		54, 5400, "2:100000000;5:1000;45:200", 
	},
	55: &PlayerData{
		55, 5500, "2:100000000;5:1000;45:200", 
	},
	56: &PlayerData{
		56, 5600, "2:100000000;5:1000;45:200", 
	},
	57: &PlayerData{
		57, 5700, "2:100000000;5:1000;45:200", 
	},
	58: &PlayerData{
		58, 5800, "2:100000000;5:1000;45:200", 
	},
	59: &PlayerData{
		59, 5900, "2:100000000;5:1000;45:200", 
	},
	60: &PlayerData{
		60, 6000, "2:100000000;5:1000;45:200", 
	},
}

	func PlayerData_hot() {
		for _, val := range cnfPlayerData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}