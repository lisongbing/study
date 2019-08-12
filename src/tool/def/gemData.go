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

type GemDataVersion struct {
	md5 string //配置表md5数据
}

type GemDataClass struct {
	GemDataVersion
}

var GemDatainstance *GemDataClass

func GemDataInstance() *GemDataClass {
	if GemDatainstance == nil {
		GemDatainstance = &GemDataClass{}
	}
	return GemDatainstance
}

type GemData struct {
	Id int32 //id
	VipExp string //注释
	Attr string //属性
	LengendSkill string //传奇宝石技能
	Note string //注释
}

func GetGemDataByPk(id int32) (itm *GemData, ok bool) {
	mtxGemData.RLock()
	itm, ok = cnfGemData[id]
	mtxGemData.RUnlock()
	return
}

func SetGemDataVersion(md5 string) string {
	GemDataInstance().md5 = md5
	return ``
	}
func GetGemDataVersion(md5 string) string {
	return GemDataInstance().md5
}
func GetGemData() map[int32]*GemData{
	mtxGemData.RLock()
	cnf := cnfGemData
	mtxGemData.RUnlock()
	return cnf
}

func (this *GemData) getVipExp() string {
	return this.VipExp 
}

func (this *GemData) getAttr() string {
	return this.Attr 
}

func (this *GemData) getLengendSkill() string {
	return this.LengendSkill 
}

func (this *GemData) getNote() string {
	return this.Note 
}

func LoadGemData(file string) string {
	var clen = []int32{5}
	sf := `gemData.xlsx`
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
	cnf := make(map[int32]*GemData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &GemData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.VipExp = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Attr = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.LengendSkill = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Note = val
		cnf[itm.Id] = itm
	}
	mtxGemData.Lock()
	cnfGemData = cnf
	mtxGemData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxGemData = new(sync.RWMutex)
var cnfGemData = map[int32]*GemData{
	1001: &GemData{
		1001, "白宝石1级", "1021:5", "", "全抗性", 
	},
	1002: &GemData{
		1002, "白宝石2级", "1021:10", "", "全抗性", 
	},
	1003: &GemData{
		1003, "白宝石3级", "1021:15", "", "全抗性", 
	},
	1004: &GemData{
		1004, "白宝石4级", "1021:20", "", "全抗性", 
	},
	1005: &GemData{
		1005, "白宝石5级", "1021:25", "", "全抗性", 
	},
	1006: &GemData{
		1006, "白宝石6级", "1021:30", "", "全抗性", 
	},
	1007: &GemData{
		1007, "白宝石7级", "1021:35", "", "全抗性", 
	},
	1008: &GemData{
		1008, "白宝石8级", "1021:40", "", "全抗性", 
	},
	1009: &GemData{
		1009, "白宝石9级", "1021:45", "", "全抗性", 
	},
	1010: &GemData{
		1010, "白宝石10级", "1021:50", "", "全抗性", 
	},
	1011: &GemData{
		1011, "紫宝石1级", "1002:500", "", "生命值", 
	},
	1012: &GemData{
		1012, "紫宝石2级", "1002:1000", "", "生命值", 
	},
	1013: &GemData{
		1013, "紫宝石3级", "1002:1500", "", "生命值", 
	},
	1014: &GemData{
		1014, "紫宝石4级", "1002:2000", "", "生命值", 
	},
	1015: &GemData{
		1015, "紫宝石5级", "1002:2500", "", "生命值", 
	},
	1016: &GemData{
		1016, "紫宝石6级", "1002:3000", "", "生命值", 
	},
	1017: &GemData{
		1017, "紫宝石7级", "1002:3500", "", "生命值", 
	},
	1018: &GemData{
		1018, "紫宝石8级", "1002:4000", "", "生命值", 
	},
	1019: &GemData{
		1019, "紫宝石9级", "1002:4500", "", "生命值", 
	},
	1020: &GemData{
		1020, "紫宝石10级", "1002:5000", "", "生命值", 
	},
	1021: &GemData{
		1021, "红宝石1级", "1005:10", "", "攻击力", 
	},
	1022: &GemData{
		1022, "红宝石2级", "1005:20", "", "攻击力", 
	},
	1023: &GemData{
		1023, "红宝石3级", "1005:30", "", "攻击力", 
	},
	1024: &GemData{
		1024, "红宝石4级", "1005:40", "", "攻击力", 
	},
	1025: &GemData{
		1025, "红宝石5级", "1005:50", "", "攻击力", 
	},
	1026: &GemData{
		1026, "红宝石6级", "1005:60", "", "攻击力", 
	},
	1027: &GemData{
		1027, "红宝石7级", "1005:70", "", "攻击力", 
	},
	1028: &GemData{
		1028, "红宝石8级", "1005:80", "", "攻击力", 
	},
	1029: &GemData{
		1029, "红宝石9级", "1005:90", "", "攻击力", 
	},
	1030: &GemData{
		1030, "红宝石10级", "1005:100", "", "攻击力", 
	},
	1031: &GemData{
		1031, "绿宝石1级", "1032：3", "", "生命恢复", 
	},
	1032: &GemData{
		1032, "绿宝石2级", "1032：6", "", "生命恢复", 
	},
	1033: &GemData{
		1033, "绿宝石3级", "1032：9", "", "生命恢复", 
	},
	1034: &GemData{
		1034, "绿宝石4级", "1032：12", "", "生命恢复", 
	},
	1035: &GemData{
		1035, "绿宝石5级", "1032：15", "", "生命恢复", 
	},
	1036: &GemData{
		1036, "绿宝石6级", "1032：18", "", "生命恢复", 
	},
	1037: &GemData{
		1037, "绿宝石7级", "1032：21", "", "生命恢复", 
	},
	1038: &GemData{
		1038, "绿宝石8级", "1032：24", "", "生命恢复", 
	},
	1039: &GemData{
		1039, "绿宝石9级", "1032：27", "", "生命恢复", 
	},
	1040: &GemData{
		1040, "绿宝石10级", "1032：30", "", "生命恢复", 
	},
	1041: &GemData{
		1041, "黄宝石1级", "1050:100", "", "荆棘", 
	},
	1042: &GemData{
		1042, "黄宝石2级", "1050:200", "", "荆棘", 
	},
	1043: &GemData{
		1043, "黄宝石3级", "1050:300", "", "荆棘", 
	},
	1044: &GemData{
		1044, "黄宝石4级", "1050:400", "", "荆棘", 
	},
	1045: &GemData{
		1045, "黄宝石5级", "1050:500", "", "荆棘", 
	},
	1046: &GemData{
		1046, "黄宝石6级", "1050:600", "", "荆棘", 
	},
	1047: &GemData{
		1047, "黄宝石7级", "1050:700", "", "荆棘", 
	},
	1048: &GemData{
		1048, "黄宝石8级", "1050:800", "", "荆棘", 
	},
	1049: &GemData{
		1049, "黄宝石9级", "1050:900", "", "荆棘", 
	},
	1050: &GemData{
		1050, "黄宝石10级", "1050:1000", "", "荆棘", 
	},
	1051: &GemData{
		1051, "多彩宝石1级", "4002:1100", "", "最大生命值比", 
	},
	1052: &GemData{
		1052, "多彩宝石2级", "4002:1200", "", "最大生命值比", 
	},
	1053: &GemData{
		1053, "多彩宝石3级", "4002:1300", "", "最大生命值比", 
	},
	1054: &GemData{
		1054, "多彩宝石4级", "4002:1400", "", "最大生命值比", 
	},
	1055: &GemData{
		1055, "多彩宝石5级", "4002:1500", "", "最大生命值比", 
	},
	1056: &GemData{
		1056, "多彩宝石6级", "4002:1600", "", "最大生命值比", 
	},
	1057: &GemData{
		1057, "多彩宝石7级", "4002:1700", "", "最大生命值比", 
	},
	1058: &GemData{
		1058, "多彩宝石8级", "4002:1800", "", "最大生命值比", 
	},
	1059: &GemData{
		1059, "多彩宝石9级", "4002:1900", "", "最大生命值比", 
	},
	1060: &GemData{
		1060, "多彩宝石10级", "4002:2000", "", "最大生命值比", 
	},
	1061: &GemData{
		1061, "炫彩宝石1级", "4043:800", "", "技能加速", 
	},
	1062: &GemData{
		1062, "炫彩宝石2级", "4043:850", "", "技能加速", 
	},
	1063: &GemData{
		1063, "炫彩宝石3级", "4043:900", "", "技能加速", 
	},
	1064: &GemData{
		1064, "炫彩宝石4级", "4043:950", "", "技能加速", 
	},
	1065: &GemData{
		1065, "炫彩宝石5级", "4043:1000", "", "技能加速", 
	},
	1066: &GemData{
		1066, "炫彩宝石6级", "4043:1050", "", "技能加速", 
	},
	1067: &GemData{
		1067, "炫彩宝石7级", "4043:1100", "", "技能加速", 
	},
	1068: &GemData{
		1068, "炫彩宝石8级", "4043:1150", "", "技能加速", 
	},
	1069: &GemData{
		1069, "炫彩宝石9级", "4043:1200", "", "技能加速", 
	},
	1070: &GemData{
		1070, "炫彩宝石10级", "4043:1250", "", "技能加速", 
	},
	1071: &GemData{
		1071, "彩虹宝石1级", "4052:600", "", "对精英怪伤害", 
	},
	1072: &GemData{
		1072, "彩虹宝石2级", "4052:700", "", "对精英怪伤害", 
	},
	1073: &GemData{
		1073, "彩虹宝石3级", "4052:800", "", "对精英怪伤害", 
	},
	1074: &GemData{
		1074, "彩虹宝石4级", "4052:900", "", "对精英怪伤害", 
	},
	1075: &GemData{
		1075, "彩虹宝石5级", "4052:1000", "", "对精英怪伤害", 
	},
	1076: &GemData{
		1076, "彩虹宝石6级", "4052:1100", "", "对精英怪伤害", 
	},
	1077: &GemData{
		1077, "彩虹宝石7级", "4052:1200", "", "对精英怪伤害", 
	},
	1078: &GemData{
		1078, "彩虹宝石8级", "4052:1300", "", "对精英怪伤害", 
	},
	1079: &GemData{
		1079, "彩虹宝石9级", "4052:1400", "", "对精英怪伤害", 
	},
	1080: &GemData{
		1080, "彩虹宝石10级", "4052:1500", "", "对精英怪伤害", 
	},
	1081: &GemData{
		1081, "活力宝石", "", "", "", 
	},
	1082: &GemData{
		1082, "至简之力", "", "", "", 
	},
	1083: &GemData{
		1083, "受罚者之灾", "", "", "", 
	},
	1084: &GemData{
		1084, "困者之灾", "", "", "", 
	},
	1085: &GemData{
		1085, "闪电华冠", "", "", "", 
	},
	1086: &GemData{
		1086, "火牛羚砂囊", "", "", "", 
	},
	1087: &GemData{
		1087, "银河，阿尔塔夏之泪", "", "", "", 
	},
	1088: &GemData{
		1088, "强者之灾", "", "", "", 
	},
	1089: &GemData{
		1089, "毁伤", "", "", "", 
	},
	1090: &GemData{
		1090, "魄罗芯片", "", "", "", 
	},
	1091: &GemData{
		1091, "复仇", "", "", "", 
	},
	1092: &GemData{
		1092, "冰晶", "", "", "", 
	},
	1093: &GemData{
		1093, "挚诚", "", "", "", 
	},
	1094: &GemData{
		1094, "剧毒", "", "", "", 
	},
	1095: &GemData{
		1095, "重伤", "", "", "", 
	},
	1096: &GemData{
		1096, "免死", "", "", "", 
	},
	1097: &GemData{
		1097, "太极", "", "", "", 
	},
	40003: &GemData{
		40003, "太极123", "4052:1500", "", "对精英怪伤害", 
	},
}

	func GemData_hot() {
		for _, val := range cnfGemData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}