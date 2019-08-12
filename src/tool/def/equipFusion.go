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

type EquipFusionVersion struct {
	md5 string //配置表md5数据
}

type EquipFusionClass struct {
	EquipFusionVersion
}

var EquipFusioninstance *EquipFusionClass

func EquipFusionInstance() *EquipFusionClass {
	if EquipFusioninstance == nil {
		EquipFusioninstance = &EquipFusionClass{}
	}
	return EquipFusioninstance
}

type EquipFusion struct {
	Id int32 //装备等级
	GetItem string //获得道具
}

func GetEquipFusionByPk(id int32) (itm *EquipFusion, ok bool) {
	mtxEquipFusion.RLock()
	itm, ok = cnfEquipFusion[id]
	mtxEquipFusion.RUnlock()
	return
}

func SetEquipFusionVersion(md5 string) string {
	EquipFusionInstance().md5 = md5
	return ``
	}
func GetEquipFusionVersion(md5 string) string {
	return EquipFusionInstance().md5
}
func GetEquipFusion() map[int32]*EquipFusion{
	mtxEquipFusion.RLock()
	cnf := cnfEquipFusion
	mtxEquipFusion.RUnlock()
	return cnf
}

func (this *EquipFusion) getGetItem() string {
	return this.GetItem 
}

func LoadEquipFusion(file string) string {
	var clen = []int32{2}
	sf := `equipFusion.xlsx`
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
    var shref = []int64{0, 0}
	_ = shref
	cnf := make(map[int32]*EquipFusion)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &EquipFusion{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.GetItem = val
		cnf[itm.Id] = itm
	}
	mtxEquipFusion.Lock()
	cnfEquipFusion = cnf
	mtxEquipFusion.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxEquipFusion = new(sync.RWMutex)
var cnfEquipFusion = map[int32]*EquipFusion{
	1: &EquipFusion{
		1, "1:100:120;4:200:400", 
	},
	2: &EquipFusion{
		2, "1:110:140;4:400:600", 
	},
	3: &EquipFusion{
		3, "1:120:160;4:600:800", 
	},
	4: &EquipFusion{
		4, "1:130:180;4:800:1000", 
	},
	5: &EquipFusion{
		5, "1:140:200;4:1000:1200", 
	},
	6: &EquipFusion{
		6, "1:150:220;4:1200:1400", 
	},
	7: &EquipFusion{
		7, "1:160:240;4:1400:1600", 
	},
	8: &EquipFusion{
		8, "1:170:260;4:1600:1800", 
	},
	9: &EquipFusion{
		9, "1:180:280;4:1800:2000", 
	},
	10: &EquipFusion{
		10, "1:190:300;4:2000:2200", 
	},
	11: &EquipFusion{
		11, "1:200:320;4:2200:2400;101:200:320", 
	},
	12: &EquipFusion{
		12, "1:210:340;4:2400:2600;101:210:340", 
	},
	13: &EquipFusion{
		13, "1:220:360;4:2600:2800;101:220:360", 
	},
	14: &EquipFusion{
		14, "1:230:380;4:2800:3000;101:230:380", 
	},
	15: &EquipFusion{
		15, "1:240:400;4:3000:3200;101:240:400", 
	},
	16: &EquipFusion{
		16, "1:250:420;4:3200:3400;101:250:420", 
	},
	17: &EquipFusion{
		17, "1:260:440;4:3400:3600;101:260:440", 
	},
	18: &EquipFusion{
		18, "1:270:460;4:3600:3800;101:270:460", 
	},
	19: &EquipFusion{
		19, "1:280:480;4:3800:4000;101:280:480", 
	},
	20: &EquipFusion{
		20, "1:290:500;4:4000:4200;101:290:500", 
	},
	21: &EquipFusion{
		21, "1:300:520;4:4200:4400;101:300:520;102:300:520", 
	},
	22: &EquipFusion{
		22, "1:310:540;4:4400:4600;101:310:540;102:310:540", 
	},
	23: &EquipFusion{
		23, "1:320:560;4:4600:4800;101:320:560;102:320:560", 
	},
	24: &EquipFusion{
		24, "1:330:580;4:4800:5000;101:330:580;102:330:580", 
	},
	25: &EquipFusion{
		25, "1:340:600;4:5000:5200;101:340:600;102:340:600", 
	},
	26: &EquipFusion{
		26, "1:350:620;4:5200:5400;101:350:620;102:350:620", 
	},
	27: &EquipFusion{
		27, "1:360:640;4:5400:5600;101:360:640;102:360:640", 
	},
	28: &EquipFusion{
		28, "1:370:660;4:5600:5800;101:370:660;102:370:660", 
	},
	29: &EquipFusion{
		29, "1:380:680;4:5800:6000;101:380:680;102:380:680", 
	},
	30: &EquipFusion{
		30, "1:390:700;4:6000:6200;101:390:700;102:390:700", 
	},
}

	func EquipFusion_hot() {
		for _, val := range cnfEquipFusion	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}