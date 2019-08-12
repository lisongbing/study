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

type EquipAffixesVersion struct {
	md5 string //配置表md5数据
}

type EquipAffixesClass struct {
	EquipAffixesVersion
}

var EquipAffixesinstance *EquipAffixesClass

func EquipAffixesInstance() *EquipAffixesClass {
	if EquipAffixesinstance == nil {
		EquipAffixesinstance = &EquipAffixesClass{}
	}
	return EquipAffixesinstance
}

type EquipAffixes struct {
	Id int32 //
	Note string //注释
	Affixtype int32 //词条
	AffixLevel int32 //词条等级
	Conflict int32 //冲突
	Slot int32 //部位限定
	Class int32 //职业限定
	Content int32 //属性
	ContentID int32 //属性ID
	Value string //数值
	MinStep int32 //步长
}

const (
	Affixtype_PrimeAffix = 0 //主要词条
	Affixtype_MinorAffix = 1 //次要词条
)

const (
	Class_None = 0 // 全职业
	Class_Crusader = 1 //圣教军
	Class_DemonHunter = 2 // 猎魔人
	Class_Wizard = 3 // 法师
)

const (
	Content_Attr =0 //属性
	Content_Skill =1 //技能
)

func SetEquipAffixesVersion(md5 string) string {
	EquipAffixesInstance().md5 = md5
	return ``
	}
func GetEquipAffixesVersion(md5 string) string {
	return EquipAffixesInstance().md5
}
func GetEquipAffixes() []*EquipAffixes{
	mtxEquipAffixes.RLock()
	cnf := cnfEquipAffixes
	mtxEquipAffixes.RUnlock()
	return cnf
}

func (this *EquipAffixes) getNote() string {
	return this.Note 
}

func (this *EquipAffixes) getAffixtype() int32 {
	return this.Affixtype 
}

func (this *EquipAffixes) getAffixLevel() int32 {
	return this.AffixLevel 
}

func (this *EquipAffixes) getConflict() int32 {
	return this.Conflict 
}

func (this *EquipAffixes) getSlot() int32 {
	return this.Slot 
}

func (this *EquipAffixes) getClass() int32 {
	return this.Class 
}

func (this *EquipAffixes) getContent() int32 {
	return this.Content 
}

func (this *EquipAffixes) getContentID() int32 {
	return this.ContentID 
}

func (this *EquipAffixes) getValue() string {
	return this.Value 
}

func (this *EquipAffixes) getMinStep() int32 {
	return this.MinStep 
}

func LoadEquipAffixes(file string) string {
	var clen = []int32{11}
	sf := `equipAffixes.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
var cnf []*EquipAffixes
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &EquipAffixes{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Note = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Affixtype, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.AffixLevel, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Conflict, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.Slot, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.Class, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.Content, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.ContentID, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.Value = val
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.MinStep, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 10, val)
		}
		cnf = append(cnf, itm)
	}
	mtxEquipAffixes.Lock()
	cnfEquipAffixes = cnf
	mtxEquipAffixes.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxEquipAffixes = new(sync.RWMutex)
var cnfEquipAffixes = []*EquipAffixes{
	&EquipAffixes{
		1, "生命", 0, 10, 0, 0, 0, 0, 1002, "5:10;11:20;21:30;40:40", 1, 
	},
	&EquipAffixes{
		1, "生命", 0, 20, 0, 0, 0, 0, 1002, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		2, "攻击", 0, 10, 0, 0, 0, 0, 1005, "5:20;11:30;21:40;50:50", 3, 
	},
	&EquipAffixes{
		3, "物理伤害增强", 0, 10, 0, 1, 0, 0, 1006, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		4, "火焰伤害增强", 0, 10, 0, 1, 0, 0, 1007, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		5, "冰霜伤害增强", 0, 10, 0, 1, 0, 0, 1008, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		6, "闪电伤害增强", 0, 10, 0, 1, 0, 0, 1009, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		7, "毒素伤害增强", 0, 10, 0, 1, 0, 0, 1010, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		8, "神圣伤害增强", 0, 10, 0, 1, 0, 0, 1011, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		9, "护甲", 0, 10, 0, 0, 0, 0, 1014, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		10, "全抗性", 0, 10, 0, 0, 0, 0, 1021, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		11, "全伤害减免", 0, 10, 0, 0, 0, 0, 4030, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		12, "生命恢复", 0, 10, 0, 0, 0, 0, 1032, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		13, "暴击率", 0, 10, 0, 0, 0, 0, 1037, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		14, "暴击伤害", 0, 10, 0, 0, 0, 0, 1038, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		15, "闪避率", 0, 10, 0, 0, 0, 0, 4039, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		16, "格挡率", 0, 10, 0, 0, 0, 0, 1040, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		17, "无视防御", 0, 10, 0, 0, 0, 0, 1042, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		18, "冷却速度", 0, 10, 0, 0, 0, 0, 4043, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		19, "坚韧", 0, 10, 0, 0, 0, 0, 4051, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		20, "对精英怪伤害", 0, 10, 0, 2, 0, 0, 1052, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		21, "对恶魔系伤害加成", 0, 10, 0, 2, 0, 0, 1054, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		22, "对不死系伤害加成", 0, 10, 0, 2, 0, 0, 1055, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		23, "击中时造成流血效果", 0, 10, 0, 3, 0, 0, 1061, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		24, "击中时造成燃烧效果", 0, 10, 0, 3, 0, 0, 1062, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		25, "击中时造成冰冷效果", 0, 10, 0, 3, 0, 0, 1063, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		26, "击中时造成电击效果", 0, 10, 0, 3, 0, 0, 1064, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		27, "击中时造成中毒效果", 0, 10, 0, 3, 0, 0, 1065, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		28, "击中时造成惩戒效果", 0, 10, 0, 3, 0, 0, 1066, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		29, "横扫伤害加成", 0, 10, 0, 4, 0, 0, 1071, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		30, "盾牌猛击伤害加成", 0, 10, 0, 4, 0, 0, 1072, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		31, "天罚之剑伤害加成", 0, 10, 0, 4, 0, 0, 1073, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		32, "祝福之盾伤害加成", 0, 10, 0, 4, 0, 0, 1074, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		33, "审判伤害加成", 0, 10, 0, 4, 0, 0, 1075, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		34, "天堂之拳伤害加成", 0, 10, 0, 4, 0, 0, 1076, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		35, "烈焰飞弹伤害加成", 0, 10, 0, 4, 0, 0, 1077, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		36, "陨石术伤害加成", 0, 10, 0, 4, 0, 0, 1078, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		37, "暴风雪伤害加成", 0, 10, 0, 4, 0, 0, 1079, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		38, "多头蛇伤害加成", 0, 10, 0, 4, 0, 0, 1080, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		39, "能量气旋伤害加成", 0, 10, 0, 4, 0, 0, 1081, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		40, "黑洞伤害加成", 0, 10, 0, 4, 0, 0, 1082, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		41, "锯齿箭伤害加成", 0, 10, 0, 4, 0, 0, 1083, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		42, "多重射击伤害加成", 0, 10, 0, 4, 0, 0, 1084, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		43, "集束箭伤害加成", 0, 10, 0, 4, 0, 0, 1085, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		44, "复仇之雨伤害加成", 0, 10, 0, 4, 0, 0, 1086, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		45, "飞轮刃伤害加成", 0, 10, 0, 4, 0, 0, 1087, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		46, "缠绕射击伤害加成", 0, 10, 0, 4, 0, 0, 1088, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		47, "锯齿箭伤害加成", 0, 10, 0, 4, 0, 0, 1089, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		48, "多重射击伤害加成", 0, 10, 0, 4, 0, 0, 1090, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		49, "集束箭伤害加成", 0, 10, 0, 4, 0, 0, 1091, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		50, "复仇之雨伤害加成", 0, 10, 0, 4, 0, 0, 1092, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		51, "飞轮刃伤害加成", 0, 10, 0, 4, 0, 0, 1093, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		52, "缠绕射击伤害加成", 0, 10, 0, 4, 0, 0, 1094, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		100, "物理抗性", 1, 10, 0, 5, 0, 0, 1015, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		101, "火焰抗性", 1, 10, 0, 5, 0, 0, 1016, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		102, "冰霜抗性", 1, 10, 0, 5, 0, 0, 1017, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		103, "闪电抗性", 1, 10, 0, 5, 0, 0, 1018, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		104, "毒素抗性", 1, 10, 0, 5, 0, 0, 1019, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		105, "神圣抗性", 1, 10, 0, 5, 0, 0, 1020, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		106, "物理伤害减免", 1, 10, 0, 6, 0, 0, 4022, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		107, "火焰伤害减免", 1, 10, 0, 6, 0, 0, 4023, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		108, "冰霜伤害减免", 1, 10, 0, 6, 0, 0, 4024, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		109, "闪电伤害减免", 1, 10, 0, 6, 0, 0, 4025, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		110, "毒素伤害减免", 1, 10, 0, 6, 0, 0, 4026, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		111, "神圣伤害减免", 1, 10, 0, 6, 0, 0, 4027, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		112, "近战伤害减免", 1, 10, 0, 7, 0, 0, 4028, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		113, "远程伤害减免", 1, 10, 0, 7, 0, 0, 4029, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		114, "法力秒回值", 1, 10, 0, 0, 0, 0, 1036, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		115, "移动速度加成", 1, 10, 0, 0, 0, 0, 1044, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		116, "状态抗性", 1, 10, 0, 0, 0, 0, 4046, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		117, "药水治疗效果", 1, 10, 0, 0, 0, 0, 1049, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		118, "荆棘", 1, 10, 0, 0, 0, 0, 1050, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		119, "受到精英怪伤害减免", 1, 10, 0, 0, 0, 0, 4053, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		120, "挂机金币加成", 1, 10, 0, 0, 0, 0, 1056, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		121, "挂机经验加成", 1, 10, 0, 0, 0, 0, 1057, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		122, "击杀怪物金币加成", 1, 10, 0, 0, 0, 0, 1058, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		123, "击杀怪物经验加成", 1, 10, 0, 0, 0, 0, 1059, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		124, "魔法物品掉率加成", 1, 10, 0, 0, 0, 0, 1060, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		125, "击中时造成昏迷", 1, 10, 0, 8, 0, 0, 1067, "5:10;11:20;21:30;40:40", 2, 
	},
	&EquipAffixes{
		126, "击中时造成冻结", 1, 10, 0, 8, 0, 0, 1068, "5:15;11:25;21:35;45:45", 2, 
	},
	&EquipAffixes{
		127, "击中时造成定身", 1, 10, 0, 8, 0, 0, 1069, "5:20;11:30;21:40;50:50", 2, 
	},
	&EquipAffixes{
		128, "击中时造成击退", 1, 10, 0, 8, 0, 0, 1070, "5:20;11:30;21:40;50:50", 2, 
	},
}

	func EquipAffixes_hot() {
		for _, val := range cnfEquipAffixes	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}