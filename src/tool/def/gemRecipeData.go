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

type GemRecipeDataVersion struct {
	md5 string //配置表md5数据
}

type GemRecipeDataClass struct {
	GemRecipeDataVersion
}

var GemRecipeDatainstance *GemRecipeDataClass

func GemRecipeDataInstance() *GemRecipeDataClass {
	if GemRecipeDatainstance == nil {
		GemRecipeDatainstance = &GemRecipeDataClass{}
	}
	return GemRecipeDatainstance
}

type GemRecipeData struct {
	Id int32 //id
	Type int32 //配方类型
	SubType int32 //subType
	TitleName string //titleName
	Note string //注释
	CraftItemNeed string //合成材料
	GoldCost int32 //金币消耗
	CraftResult int32 //合成出的道具
	UnlockType int32 //解锁方式
	UnlockParam int32 //参数
}

type Sheet1Version struct {
	md5 string //配置表md5数据
}

type Sheet1Class struct {
	Sheet1Version
}

var Sheet1instance *Sheet1Class

func Sheet1Instance() *Sheet1Class {
	if Sheet1instance == nil {
		Sheet1instance = &Sheet1Class{}
	}
	return Sheet1instance
}

type Sheet1 struct {
}

func GetGemRecipeDataByPk(id int32) (itm *GemRecipeData, ok bool) {
	mtxGemRecipeData.RLock()
	itm, ok = cnfGemRecipeData[id]
	mtxGemRecipeData.RUnlock()
	return
}

const (
	SubType_LegendGem = 0
	SubType_WhiteGem = 1
	SubType_PurpleGem =2
	SubType_RedGem =3
	SubType_GreenGem = 4
	SubType_YellowGem = 5
	SubType_ColorAGem=6
	SubType_ColorBGem = 7
	SubType_ColorCGem = 8
)

func SetGemRecipeDataVersion(md5 string) string {
	GemRecipeDataInstance().md5 = md5
	return ``
	}
func GetGemRecipeDataVersion(md5 string) string {
	return GemRecipeDataInstance().md5
}
func GetGemRecipeData() map[int32]*GemRecipeData{
	mtxGemRecipeData.RLock()
	cnf := cnfGemRecipeData
	mtxGemRecipeData.RUnlock()
	return cnf
}

func (this *GemRecipeData) getType() int32 {
	return this.Type 
}

func (this *GemRecipeData) getSubType() int32 {
	return this.SubType 
}

func (this *GemRecipeData) getTitleName() string {
	return this.TitleName 
}

func (this *GemRecipeData) getNote() string {
	return this.Note 
}

func (this *GemRecipeData) getCraftItemNeed() string {
	return this.CraftItemNeed 
}

func (this *GemRecipeData) getGoldCost() int32 {
	return this.GoldCost 
}

func (this *GemRecipeData) getCraftResult() int32 {
	return this.CraftResult 
}

func (this *GemRecipeData) getUnlockType() int32 {
	return this.UnlockType 
}

func (this *GemRecipeData) getUnlockParam() int32 {
	return this.UnlockParam 
}

func LoadGemRecipeData(file string) string {
	var clen = []int32{10, 0}
	sf := `gemRecipeData.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*GemRecipeData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &GemRecipeData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Type, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.SubType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.TitleName = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Note = val
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.CraftItemNeed = val
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.GoldCost, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.CraftResult, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.UnlockType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.UnlockParam, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 9, val)
		}
		cnf[itm.Id] = itm
	}
	mtxGemRecipeData.Lock()
	cnfGemRecipeData = cnf
	mtxGemRecipeData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxGemRecipeData = new(sync.RWMutex)
var cnfGemRecipeData = map[int32]*GemRecipeData{
	1: &GemRecipeData{
		1, 1, 1, "白宝石", "白宝石2级", "1001:2", 2500, 1002, 1, 10, 
	},
	2: &GemRecipeData{
		2, 1, 1, "白宝石", "白宝石3级", "1002:2", 5000, 1003, 1, 20, 
	},
	3: &GemRecipeData{
		3, 1, 1, "白宝石", "白宝石4级", "1003:2", 10000, 1004, 1, 30, 
	},
	4: &GemRecipeData{
		4, 1, 1, "白宝石", "白宝石5级", "1004:3", 25000, 1005, 1, 40, 
	},
	5: &GemRecipeData{
		5, 1, 1, "白宝石", "白宝石6级", "1005:3", 50000, 1006, 1, 50, 
	},
	6: &GemRecipeData{
		6, 1, 1, "白宝石", "白宝石7级", "1006:3", 100000, 1007, 1, 60, 
	},
	7: &GemRecipeData{
		7, 1, 1, "白宝石", "白宝石8级", "1007:3", 150000, 1008, 1, 70, 
	},
	8: &GemRecipeData{
		8, 1, 1, "白宝石", "白宝石9级", "1008:3", 250000, 1009, 1, 80, 
	},
	9: &GemRecipeData{
		9, 1, 1, "白宝石", "白宝石10级", "1009:3", 400000, 1010, 1, 90, 
	},
	10: &GemRecipeData{
		10, 1, 2, "紫宝石", "紫宝石2级", "1011:2", 2500, 1012, 1, 10, 
	},
	11: &GemRecipeData{
		11, 1, 2, "紫宝石", "紫宝石3级", "1012:2", 5000, 1013, 1, 20, 
	},
	12: &GemRecipeData{
		12, 1, 2, "紫宝石", "紫宝石4级", "1013:2", 10000, 1014, 1, 30, 
	},
	13: &GemRecipeData{
		13, 1, 2, "紫宝石", "紫宝石5级", "1014:3", 25000, 1015, 1, 40, 
	},
	14: &GemRecipeData{
		14, 1, 2, "紫宝石", "紫宝石6级", "1015:3", 50000, 1016, 1, 50, 
	},
	15: &GemRecipeData{
		15, 1, 2, "紫宝石", "紫宝石7级", "1016:3", 100000, 1017, 1, 60, 
	},
	16: &GemRecipeData{
		16, 1, 2, "紫宝石", "紫宝石8级", "1017:3", 150000, 1018, 1, 70, 
	},
	17: &GemRecipeData{
		17, 1, 2, "紫宝石", "紫宝石9级", "1018:3", 250000, 1019, 1, 80, 
	},
	18: &GemRecipeData{
		18, 1, 2, "紫宝石", "紫宝石10级", "1019:3", 400000, 1020, 1, 90, 
	},
	19: &GemRecipeData{
		19, 1, 3, "红宝石", "红宝石2级", "1021:2", 2500, 1022, 1, 10, 
	},
	20: &GemRecipeData{
		20, 1, 3, "红宝石", "红宝石3级", "1022:2", 5000, 1023, 1, 20, 
	},
	21: &GemRecipeData{
		21, 1, 3, "红宝石", "红宝石4级", "1023:2", 10000, 1024, 1, 30, 
	},
	22: &GemRecipeData{
		22, 1, 3, "红宝石", "红宝石5级", "1024:3", 25000, 1025, 1, 40, 
	},
	23: &GemRecipeData{
		23, 1, 3, "红宝石", "红宝石6级", "1025:3", 50000, 1026, 1, 50, 
	},
	24: &GemRecipeData{
		24, 1, 3, "红宝石", "红宝石7级", "1026:3", 100000, 1027, 1, 60, 
	},
	25: &GemRecipeData{
		25, 1, 3, "红宝石", "红宝石8级", "1027:3", 150000, 1028, 1, 70, 
	},
	26: &GemRecipeData{
		26, 1, 3, "红宝石", "红宝石9级", "1028:3", 250000, 1029, 1, 80, 
	},
	27: &GemRecipeData{
		27, 1, 3, "红宝石", "红宝石10级", "1029:3", 400000, 1030, 1, 90, 
	},
	28: &GemRecipeData{
		28, 1, 4, "绿宝石", "绿宝石2级", "1031:2", 2500, 1032, 1, 10, 
	},
	29: &GemRecipeData{
		29, 1, 4, "绿宝石", "绿宝石3级", "1032:2", 5000, 1033, 1, 20, 
	},
	30: &GemRecipeData{
		30, 1, 4, "绿宝石", "绿宝石4级", "1033:2", 10000, 1034, 1, 30, 
	},
	31: &GemRecipeData{
		31, 1, 4, "绿宝石", "绿宝石5级", "1034:3", 25000, 1035, 1, 40, 
	},
	32: &GemRecipeData{
		32, 1, 4, "绿宝石", "绿宝石6级", "1035:3", 50000, 1036, 1, 50, 
	},
	33: &GemRecipeData{
		33, 1, 4, "绿宝石", "绿宝石7级", "1036:3", 100000, 1037, 1, 60, 
	},
	34: &GemRecipeData{
		34, 1, 4, "绿宝石", "绿宝石8级", "1037:3", 150000, 1038, 1, 70, 
	},
	35: &GemRecipeData{
		35, 1, 4, "绿宝石", "绿宝石9级", "1038:3", 250000, 1039, 1, 80, 
	},
	36: &GemRecipeData{
		36, 1, 4, "绿宝石", "绿宝石10级", "1039:3", 400000, 1040, 1, 90, 
	},
	37: &GemRecipeData{
		37, 1, 5, "黄宝石", "黄宝石2级", "1041:2", 2500, 1042, 1, 10, 
	},
	38: &GemRecipeData{
		38, 1, 5, "黄宝石", "黄宝石3级", "1042:2", 5000, 1043, 1, 20, 
	},
	39: &GemRecipeData{
		39, 1, 5, "黄宝石", "黄宝石4级", "1043:2", 10000, 1044, 1, 30, 
	},
	40: &GemRecipeData{
		40, 1, 5, "黄宝石", "黄宝石5级", "1044:3", 25000, 1045, 1, 40, 
	},
	41: &GemRecipeData{
		41, 1, 5, "黄宝石", "黄宝石6级", "1045:3", 50000, 1046, 1, 50, 
	},
	42: &GemRecipeData{
		42, 1, 5, "黄宝石", "黄宝石7级", "1046:3", 100000, 1047, 1, 60, 
	},
	43: &GemRecipeData{
		43, 1, 5, "黄宝石", "黄宝石8级", "1047:3", 150000, 1048, 1, 70, 
	},
	44: &GemRecipeData{
		44, 1, 5, "黄宝石", "黄宝石9级", "1048:3", 250000, 1049, 1, 80, 
	},
	45: &GemRecipeData{
		45, 1, 5, "黄宝石", "黄宝石10级", "1049:3", 400000, 1050, 1, 90, 
	},
	46: &GemRecipeData{
		46, 2, 6, "多彩宝石", "多彩宝石2级", "1051:3", 2500, 1052, 1, 10, 
	},
	47: &GemRecipeData{
		47, 2, 6, "多彩宝石", "多彩宝石3级", "1052:3", 5000, 1053, 1, 20, 
	},
	48: &GemRecipeData{
		48, 2, 6, "多彩宝石", "多彩宝石4级", "1053:3", 10000, 1054, 1, 30, 
	},
	49: &GemRecipeData{
		49, 2, 6, "多彩宝石", "多彩宝石5级", "1054:3", 25000, 1055, 1, 40, 
	},
	50: &GemRecipeData{
		50, 2, 6, "多彩宝石", "多彩宝石6级", "1055:3", 50000, 1056, 1, 50, 
	},
	51: &GemRecipeData{
		51, 2, 6, "多彩宝石", "多彩宝石7级", "1056:3", 100000, 1057, 1, 60, 
	},
	52: &GemRecipeData{
		52, 2, 6, "多彩宝石", "多彩宝石8级", "1057:3", 150000, 1058, 1, 70, 
	},
	53: &GemRecipeData{
		53, 2, 6, "多彩宝石", "多彩宝石9级", "1058:3", 250000, 1059, 1, 80, 
	},
	54: &GemRecipeData{
		54, 2, 6, "多彩宝石", "多彩宝石10级", "1059:3", 400000, 1060, 1, 90, 
	},
	55: &GemRecipeData{
		55, 2, 7, "炫彩宝石", "炫彩宝石2级", "1061:3", 2500, 1062, 1, 10, 
	},
	56: &GemRecipeData{
		56, 2, 7, "炫彩宝石", "炫彩宝石3级", "1062:3", 5000, 1063, 1, 20, 
	},
	57: &GemRecipeData{
		57, 2, 7, "炫彩宝石", "炫彩宝石4级", "1063:3", 10000, 1064, 1, 30, 
	},
	58: &GemRecipeData{
		58, 2, 7, "炫彩宝石", "炫彩宝石5级", "1064:3", 25000, 1065, 1, 40, 
	},
	59: &GemRecipeData{
		59, 2, 7, "炫彩宝石", "炫彩宝石6级", "1065:3", 50000, 1066, 1, 50, 
	},
	60: &GemRecipeData{
		60, 2, 7, "炫彩宝石", "炫彩宝石7级", "1066:3", 100000, 1067, 1, 60, 
	},
	61: &GemRecipeData{
		61, 2, 7, "炫彩宝石", "炫彩宝石8级", "1067:3", 150000, 1068, 1, 70, 
	},
	62: &GemRecipeData{
		62, 2, 7, "炫彩宝石", "炫彩宝石9级", "1068:3", 250000, 1069, 1, 80, 
	},
	63: &GemRecipeData{
		63, 2, 7, "炫彩宝石", "炫彩宝石10级", "1069:3", 400000, 1070, 1, 90, 
	},
	64: &GemRecipeData{
		64, 2, 8, "彩虹宝石", "彩虹宝石2级", "1071:3", 2500, 1072, 1, 10, 
	},
	65: &GemRecipeData{
		65, 2, 8, "彩虹宝石", "彩虹宝石3级", "1072:3", 5000, 1073, 1, 20, 
	},
	66: &GemRecipeData{
		66, 2, 8, "彩虹宝石", "彩虹宝石4级", "1073:3", 10000, 1074, 1, 30, 
	},
	67: &GemRecipeData{
		67, 2, 8, "彩虹宝石", "彩虹宝石5级", "1074:3", 25000, 1075, 1, 40, 
	},
	68: &GemRecipeData{
		68, 2, 8, "彩虹宝石", "彩虹宝石6级", "1075:3", 50000, 1076, 1, 50, 
	},
	69: &GemRecipeData{
		69, 2, 8, "彩虹宝石", "彩虹宝石7级", "1076:3", 100000, 1077, 1, 60, 
	},
	70: &GemRecipeData{
		70, 2, 8, "彩虹宝石", "彩虹宝石8级", "1077:3", 150000, 1078, 1, 70, 
	},
	71: &GemRecipeData{
		71, 2, 8, "彩虹宝石", "彩虹宝石9级", "1078:3", 250000, 1079, 1, 80, 
	},
	72: &GemRecipeData{
		72, 2, 8, "彩虹宝石", "彩虹宝石10级", "1079:3", 400000, 1080, 1, 90, 
	},
	73: &GemRecipeData{
		73, 3, 0, "", "活力宝石", "102:1", 100000, 1081, 0, 0, 
	},
	74: &GemRecipeData{
		74, 3, 0, "", "至简之力", "102:1", 100000, 1082, 0, 0, 
	},
	75: &GemRecipeData{
		75, 3, 0, "", "受罚者之灾", "102:1", 100000, 1083, 0, 0, 
	},
	76: &GemRecipeData{
		76, 3, 0, "", "困者之灾", "102:1", 100000, 1084, 0, 0, 
	},
	77: &GemRecipeData{
		77, 3, 0, "", "闪电华冠", "102:1", 100000, 1085, 0, 0, 
	},
	78: &GemRecipeData{
		78, 3, 0, "", "火牛羚砂囊", "102:1", 100000, 1086, 0, 0, 
	},
	79: &GemRecipeData{
		79, 3, 0, "", "银河，阿尔塔夏之泪", "102:1", 100000, 1087, 0, 0, 
	},
	80: &GemRecipeData{
		80, 3, 0, "", "强者之灾", "102:1", 100000, 1088, 0, 0, 
	},
	81: &GemRecipeData{
		81, 3, 0, "", "毁伤", "102:1", 100000, 1089, 0, 0, 
	},
	82: &GemRecipeData{
		82, 3, 0, "", "魄罗芯片", "102:1", 100000, 1090, 0, 0, 
	},
	83: &GemRecipeData{
		83, 3, 0, "", "复仇", "102:1", 100000, 1091, 0, 0, 
	},
	84: &GemRecipeData{
		84, 3, 0, "", "冰晶", "102:1", 100000, 1092, 0, 0, 
	},
	85: &GemRecipeData{
		85, 3, 0, "", "挚诚", "102:1", 100000, 1093, 0, 0, 
	},
	86: &GemRecipeData{
		86, 3, 0, "", "剧毒", "102:1", 100000, 1094, 0, 0, 
	},
	87: &GemRecipeData{
		87, 3, 0, "", "重伤", "102:1", 100000, 1095, 0, 0, 
	},
	88: &GemRecipeData{
		88, 3, 0, "", "免死", "102:1", 100000, 1096, 0, 0, 
	},
	89: &GemRecipeData{
		89, 3, 0, "", "太极", "102:1", 100000, 1097, 0, 0, 
	},
}

	func GemRecipeData_hot() {
		for _, val := range cnfGemRecipeData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}