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

type SignInDataVersion struct {
	md5 string //配置表md5数据
}

type SignInDataClass struct {
	SignInDataVersion
}

var SignInDatainstance *SignInDataClass

func SignInDataInstance() *SignInDataClass {
	if SignInDatainstance == nil {
		SignInDatainstance = &SignInDataClass{}
	}
	return SignInDatainstance
}

type SignInData struct {
	Id int32 //天数
	Award string //注释
	VipLv int32 //vip等级要求
	VipBonus int32 //vip倍率
	SeqAward string //连续签到奖励
}

func GetSignInDataByPk(id int32) (itm *SignInData, ok bool) {
	mtxSignInData.RLock()
	itm, ok = cnfSignInData[id]
	mtxSignInData.RUnlock()
	return
}

func SetSignInDataVersion(md5 string) string {
	SignInDataInstance().md5 = md5
	return ``
	}
func GetSignInDataVersion(md5 string) string {
	return SignInDataInstance().md5
}
func GetSignInData() map[int32]*SignInData{
	mtxSignInData.RLock()
	cnf := cnfSignInData
	mtxSignInData.RUnlock()
	return cnf
}

func (this *SignInData) getAward() string {
	return this.Award 
}

func (this *SignInData) getVipLv() int32 {
	return this.VipLv 
}

func (this *SignInData) getVipBonus() int32 {
	return this.VipBonus 
}

func (this *SignInData) getSeqAward() string {
	return this.SeqAward 
}

func LoadSignInData(file string) string {
	var clen = []int32{5}
	sf := `svrsignInData.xlsx`
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
	cnf := make(map[int32]*SignInData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &SignInData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Award = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.VipLv, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.VipBonus, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.SeqAward = val
		cnf[itm.Id] = itm
	}
	mtxSignInData.Lock()
	cnfSignInData = cnf
	mtxSignInData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxSignInData = new(sync.RWMutex)
var cnfSignInData = map[int32]*SignInData{
	1: &SignInData{
		1, "1:10000", 0, 0, "", 
	},
	2: &SignInData{
		2, "1:20", 2, 20000, "", 
	},
	3: &SignInData{
		3, "101:150", 0, 0, "1:20000", 
	},
	4: &SignInData{
		4, "102:200", 3, 30000, "", 
	},
	5: &SignInData{
		5, "1001:10", 0, 0, "", 
	},
	6: &SignInData{
		6, "1002:20", 4, 40000, "", 
	},
	7: &SignInData{
		7, "1003:30", 0, 0, "1:40000", 
	},
	8: &SignInData{
		8, "1:20000", 5, 50000, "", 
	},
	9: &SignInData{
		9, "1:40", 0, 0, "", 
	},
	10: &SignInData{
		10, "101:250", 2, 20000, "", 
	},
	11: &SignInData{
		11, "102:300", 0, 0, "", 
	},
	12: &SignInData{
		12, "1001:40", 3, 30000, "", 
	},
	13: &SignInData{
		13, "1002:50", 0, 0, "", 
	},
	14: &SignInData{
		14, "1003:60", 4, 40000, "1:60000", 
	},
	15: &SignInData{
		15, "1:10000", 0, 0, "", 
	},
	16: &SignInData{
		16, "1:20", 5, 50000, "", 
	},
	17: &SignInData{
		17, "101:150", 0, 0, "", 
	},
	18: &SignInData{
		18, "102:200", 2, 20000, "", 
	},
	19: &SignInData{
		19, "1001:10", 0, 0, "", 
	},
	20: &SignInData{
		20, "1002:20", 3, 30000, "", 
	},
	21: &SignInData{
		21, "1003:30", 0, 0, "1:80000", 
	},
	22: &SignInData{
		22, "1:20000", 4, 40000, "", 
	},
	23: &SignInData{
		23, "1:40", 0, 0, "", 
	},
	24: &SignInData{
		24, "101:250", 5, 50000, "", 
	},
	25: &SignInData{
		25, "102:300", 0, 0, "", 
	},
	26: &SignInData{
		26, "1001:40", 2, 20000, "", 
	},
	27: &SignInData{
		27, "1002:50", 0, 0, "", 
	},
	28: &SignInData{
		28, "1003:60", 3, 30000, "1:100000", 
	},
	29: &SignInData{
		29, "1002:50", 0, 0, "", 
	},
	30: &SignInData{
		30, "1003:600", 4, 40000, "", 
	},
}

	func SignInData_hot() {
		for _, val := range cnfSignInData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}