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

type PlayerBaseVersion struct {
	md5 string //配置表md5数据
}

type PlayerBaseClass struct {
	PlayerBaseVersion
}

var PlayerBaseinstance *PlayerBaseClass

func PlayerBaseInstance() *PlayerBaseClass {
	if PlayerBaseinstance == nil {
		PlayerBaseinstance = &PlayerBaseClass{}
	}
	return PlayerBaseinstance
}

type PlayerBase struct {
	Id int32 //唯一id
	Name string //名字
	Class int32 //职业
	AttrAjust string //属性修正
	Skill string //职业技能
	PassiveSkill string //被动技能
	Radius int32 //半径
	ModelID string //骨骼
	Face string //半身像
	CreateEffect string //创建特效
	Standby string //待机动作
	StandbyEffect string //待机特效
	DisplayAction string //展示动作
	DisplayEffect string //展示特效
	SkillAction string //技能动作
	SkillEffect string //技能特效
	RoleIcons string //角色图标
	FeaturePicture string //特性图片
}

func GetPlayerBaseByPk(id int32) (itm *PlayerBase, ok bool) {
	mtxPlayerBase.RLock()
	itm, ok = cnfPlayerBase[id]
	mtxPlayerBase.RUnlock()
	return
}

func SetPlayerBaseVersion(md5 string) string {
	PlayerBaseInstance().md5 = md5
	return ``
	}
func GetPlayerBaseVersion(md5 string) string {
	return PlayerBaseInstance().md5
}
func GetPlayerBase() map[int32]*PlayerBase{
	mtxPlayerBase.RLock()
	cnf := cnfPlayerBase
	mtxPlayerBase.RUnlock()
	return cnf
}

func (this *PlayerBase) getName() string {
	return this.Name 
}

func (this *PlayerBase) getClass() int32 {
	return this.Class 
}

func (this *PlayerBase) getAttrAjust() string {
	return this.AttrAjust 
}

func (this *PlayerBase) getSkill() string {
	return this.Skill 
}

func (this *PlayerBase) getPassiveSkill() string {
	return this.PassiveSkill 
}

func (this *PlayerBase) getRadius() int32 {
	return this.Radius 
}

func (this *PlayerBase) getModelID() string {
	return this.ModelID 
}

func (this *PlayerBase) getFace() string {
	return this.Face 
}

func (this *PlayerBase) getCreateEffect() string {
	return this.CreateEffect 
}

func (this *PlayerBase) getStandby() string {
	return this.Standby 
}

func (this *PlayerBase) getStandbyEffect() string {
	return this.StandbyEffect 
}

func (this *PlayerBase) getDisplayAction() string {
	return this.DisplayAction 
}

func (this *PlayerBase) getDisplayEffect() string {
	return this.DisplayEffect 
}

func (this *PlayerBase) getSkillAction() string {
	return this.SkillAction 
}

func (this *PlayerBase) getSkillEffect() string {
	return this.SkillEffect 
}

func (this *PlayerBase) getRoleIcons() string {
	return this.RoleIcons 
}

func (this *PlayerBase) getFeaturePicture() string {
	return this.FeaturePicture 
}

func LoadPlayerBase(file string) string {
	var clen = []int32{18}
	sf := `player0Base.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*PlayerBase)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &PlayerBase{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Name = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Class, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.AttrAjust = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Skill = val
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.PassiveSkill = val
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.Radius, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.ModelID = val
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.Face = val
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.CreateEffect = val
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.Standby = val
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		itm.StandbyEffect = val
		val = strings.Replace(row.Cells[12].String(), " \t\r\n", ``, -1)
		itm.DisplayAction = val
		val = strings.Replace(row.Cells[13].String(), " \t\r\n", ``, -1)
		itm.DisplayEffect = val
		val = strings.Replace(row.Cells[14].String(), " \t\r\n", ``, -1)
		itm.SkillAction = val
		val = strings.Replace(row.Cells[15].String(), " \t\r\n", ``, -1)
		itm.SkillEffect = val
		val = strings.Replace(row.Cells[16].String(), " \t\r\n", ``, -1)
		itm.RoleIcons = val
		val = strings.Replace(row.Cells[17].String(), " \t\r\n", ``, -1)
		itm.FeaturePicture = val
		cnf[itm.Id] = itm
	}
	mtxPlayerBase.Lock()
	cnfPlayerBase = cnf
	mtxPlayerBase.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxPlayerBase = new(sync.RWMutex)
var cnfPlayerBase = map[int32]*PlayerBase{
	1: &PlayerBase{
		1, "圣教军", 1, "2:15000;5:8000", "1000;1001;1002;1003;1004;1005", "1100;1101;1102;1103;1104;1105;1106;1107", 100, "1_eq01_1", "2", "0", "0", "0", "0", "0", "0", "0", "0", "0", 
	},
	2: &PlayerBase{
		2, "猎魔人", 2, "2:12000;5:8000", "", "", 100, "2_eq01_1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", 
	},
	3: &PlayerBase{
		3, "法师", 3, "2:11000;5:8000", "3000;3001;3003;3005;3007;3009", "", 100, "3_eq01_1", "3", "0", "0", "0", "0", "0", "0", "0", "0", "0", 
	},
}

	func PlayerBase_hot() {
		for _, val := range cnfPlayerBase	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}