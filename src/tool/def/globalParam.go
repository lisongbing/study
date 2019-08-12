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

type GlobalParamVersion struct {
	md5 string //配置表md5数据
}

type GlobalParamClass struct {
	GlobalParamVersion
}

var GlobalParaminstance *GlobalParamClass

func GlobalParamInstance() *GlobalParamClass {
	if GlobalParaminstance == nil {
		GlobalParaminstance = &GlobalParamClass{}
	}
	return GlobalParaminstance
}

type GlobalParam struct {
	Id int32 //编号
	Name string //读名字别读id
	Value string //变量值
	Des string //描述
}

func GetGlobalParamByPk(id int32) (itm *GlobalParam, ok bool) {
	mtxGlobalParam.RLock()
	itm, ok = cnfGlobalParam[id]
	mtxGlobalParam.RUnlock()
	return
}

func SetGlobalParamVersion(md5 string) string {
	GlobalParamInstance().md5 = md5
	return ``
	}
func GetGlobalParamVersion(md5 string) string {
	return GlobalParamInstance().md5
}
func GetGlobalParam() map[int32]*GlobalParam{
	mtxGlobalParam.RLock()
	cnf := cnfGlobalParam
	mtxGlobalParam.RUnlock()
	return cnf
}

func (this *GlobalParam) getName() string {
	return this.Name 
}

func (this *GlobalParam) getValue() string {
	return this.Value 
}

func (this *GlobalParam) getDes() string {
	return this.Des 
}

func LoadGlobalParam(file string) string {
	var clen = []int32{4}
	sf := `globalParam.xlsx`
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
	cnf := make(map[int32]*GlobalParam)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &GlobalParam{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Name = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Value = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Des = val
		cnf[itm.Id] = itm
	}
	mtxGlobalParam.Lock()
	cnfGlobalParam = cnf
	mtxGlobalParam.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxGlobalParam = new(sync.RWMutex)
var cnfGlobalParam = map[int32]*GlobalParam{
	1: &GlobalParam{
		1, "playerMaxLv", "70", "格式：数值（角色等级上限）", 
	},
	2: &GlobalParam{
		2, "atkLvArmor", "100", "攻击者等级护甲常数", 
	},
	3: &GlobalParam{
		3, "atkLvRes", "10", "攻击者等级抗性常数", 
	},
	4: &GlobalParam{
		4, "whiteBonus", "5000;0;0;0;0", "白装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	5: &GlobalParam{
		5, "blueBonus", "10000;2;1;1;0", "蓝装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	6: &GlobalParam{
		6, "yellowBonus", "20000;4;2;2;1", "黄装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	7: &GlobalParam{
		7, "goldenBonus", "40000;4;4;2;2", "暗金装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	8: &GlobalParam{
		8, "greenBonus", "40000;4;4;2;2", "绿装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	9: &GlobalParam{
		9, "ancientGoldenBonus", "50000;4;4;2;2", "远古暗金装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	10: &GlobalParam{
		10, "ancientGreenBonus", "50000;4;4;2;2", "远古绿装装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	11: &GlobalParam{
		11, "primeGoldenBonus", "60000;4;4;2;2", "远古暗金装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	12: &GlobalParam{
		12, "primeGreenBonus", "60000;4;4;2;2", "远古绿色装属性加成参数,主要词条max,主要词条min,次要词条max,次要词条min", 
	},
	13: &GlobalParam{
		13, "whiteEquipFusion", "5000;5000;0;0", "白装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	14: &GlobalParam{
		14, "blueEquipFusion", "10000;10000;5000;0", "蓝装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	15: &GlobalParam{
		15, "yellowEquipItem1", "15000;15000;15000;10000", "黄装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	16: &GlobalParam{
		16, "goldenEquipItem1", "20000;20000;20000;20000", "暗金装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	17: &GlobalParam{
		17, "greenEquipItem1", "20000;20000;20000;20000", "绿装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	18: &GlobalParam{
		18, "ancientGoldenEquipItem1", "20000;20000;20000;20000", "远古暗金装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	19: &GlobalParam{
		19, "ancientGreenEquipItem1", "20000;20000;20000;20000", "远古绿装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	20: &GlobalParam{
		20, "primeGoldenEquipItem1", "20000;20000;20000;20000", "远古暗金装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	21: &GlobalParam{
		21, "primeGreenEquipItem1", "20000;20000;20000;20000", "远古绿装道具1参数，道具2参数，道具3参数，道具4参数", 
	},
	86: &GlobalParam{
		86, "reSignInPrice", "10", "补充签到钻石价格 ", 
	},
	87: &GlobalParam{
		87, "ancientBOSS1", "10", "远古BOSS刷新参数1", 
	},
	88: &GlobalParam{
		88, "ancientBOSS2", "20", "远古BOSS刷新参数2", 
	},
	89: &GlobalParam{
		89, "ancientBOSS3", "30", "远古BOSS刷新参数3", 
	},
	90: &GlobalParam{
		90, "ancientBOSS4", "40", "远古BOSS刷新参数4", 
	},
	91: &GlobalParam{
		91, "ancientBOSS5", "50", "远古BOSS刷新参数5", 
	},
	92: &GlobalParam{
		92, "ancientBOSS6", "60", "远古BOSS刷新参数6", 
	},
	93: &GlobalParam{
		93, "ancientBOSS7", "70", "远古BOSS刷新参数7", 
	},
	94: &GlobalParam{
		94, "ancientBOSS8", "80", "远古BOSS刷新参数8", 
	},
}

	func GlobalParam_hot() {
		for _, val := range cnfGlobalParam	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}