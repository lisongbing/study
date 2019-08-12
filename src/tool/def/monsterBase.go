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

type MonsterBaseVersion struct {
	md5 string //配置表md5数据
}

type MonsterBaseClass struct {
	MonsterBaseVersion
}

var MonsterBaseinstance *MonsterBaseClass

func MonsterBaseInstance() *MonsterBaseClass {
	if MonsterBaseinstance == nil {
		MonsterBaseinstance = &MonsterBaseClass{}
	}
	return MonsterBaseinstance
}

type MonsterBase struct {
	Id int32 //唯一ID
	Name string //名字
	Des1 string //怪物等阶
	Des2 string //怪物描述
	Level int32 //基本等级
	Adaptive int32 //自适应
	CanDrop int32 //能否掉落
	DropID string //掉落id
	Exp int32 //基本经验
	Type int32 //类型
	Race int32 //种族
	SelectAble int32 //可否被选中
	BlockAble int32 //可否被阻挡
	RotateAble int32 //能否转向
	MoveAble int32 //能否移动
	Faction int32 //阵营
	Aggressivity int32 //攻击性
	WarnRange int32 //警戒范围
	RandomWalk string //随机移动范围像素
	Radius int32 //半径
	Scale float32 //缩放
	NormalAtk int32 //普通攻击
	AtkSpeedAjust float32 //攻击速度修正
	RandDelay string //随机延时
	BasePassiveSkill string //基本被动技能
	RandPassiveSkill string //随机被动技能库
	BaseActiveSkill string //基本主动技能
	RandActiveSkill string //随机主动技能库
	TargetChoose int32 //目标选择
	TargetChange int32 //目标切换
	AiScript string //AI脚本
	MoveSpeed int32 //移动速度
	Attribute string //所有属性
	InheritAttrAble int32 //可否继承属性
	InheritAttrMod string //继承属性修正
	ResType int32 //资源类型
	ModelID string //骨骼
	EffectID string //自带特效
	BornVfx string //出生特效
	DeathVfx string //死亡特效
	SpriteType int32 //类型
	TriggerPassive int32 //触发被动
	ReadyTime int32 //准备时间
	WorkTime int32 //陷阱持续毫秒
	WaitTime int32 //陷阱等待时间
	TrapOverTime int32 //陷阱超时
	MovePlus int32 //额外移动行为
}

func GetMonsterBaseByPk(id int32) (itm *MonsterBase, ok bool) {
	mtxMonsterBase.RLock()
	itm, ok = cnfMonsterBase[id]
	mtxMonsterBase.RUnlock()
	return
}

const (
	Adaptive_NoAda = 0 //非自适应
	Adaptive_Ada = 1 //自适应
)

const (
	CanDrop_CantDrop = 0 //不能掉落
	CanDrop_ = 1 //可以掉落
)

const (
	Race_NoRace = 0 //无种族
	Race_Demon = 1 // 恶魔
	Race_Undead = 2 //亡灵
)

const (
	SelectAble_Able =0 //能被选
	SelectAble_Unable = 1 //不能被选
)

const (
	Faction_Monster = 1 //怪物
	Faction_Player = 2 //玩家
	Faction_Red = 3  //战场红方
	Faction_Blue= 4 //战场蓝方
	Faction_NeutralHost = 5 //中立敌对
	Faction_NeutralFriend = 6 //中立友好
)

const (
	Aggressivity_None = 0 //无攻击性，永不攻击
	Aggressivity_Passive = 1 //被动，受攻击反击
	Aggressivity_Active= 2 //主动攻击
	Aggressivity_Always =3 //出生就一直攻击
)

const (
	TargetChoose_RandTarget = 0 //随机选择
	TargetChoose_NearestTarget = 1 //最近的目标
	TargetChoose_FarthestTarget = 2 //最远的目标
	TargetChoose_HalfTarget=3 // 50%几率最近或者随机选择
)

const (
	TargetChange_NoChange = 0 //不切换
	TargetChange_TenSecChange = 1 // 10秒切换1次
)

const (
	InheritAttrAble_Inherit = 0 //可以
	InheritAttrAble_NoInherit = 1 //不可以，如果不继承则使用attribute的数值
)

const (
	SpriteType_VisualSprite = 0 //特效精灵
	SpriteType_Henchman = 1 //仆从
	SpriteType_Trap =2 //陷阱，目标进入警戒范围就触发
)

func SetMonsterBaseVersion(md5 string) string {
	MonsterBaseInstance().md5 = md5
	return ``
	}
func GetMonsterBaseVersion(md5 string) string {
	return MonsterBaseInstance().md5
}
func GetMonsterBase() map[int32]*MonsterBase{
	mtxMonsterBase.RLock()
	cnf := cnfMonsterBase
	mtxMonsterBase.RUnlock()
	return cnf
}

func (this *MonsterBase) getName() string {
	return this.Name 
}

func (this *MonsterBase) getDes1() string {
	return this.Des1 
}

func (this *MonsterBase) getDes2() string {
	return this.Des2 
}

func (this *MonsterBase) getLevel() int32 {
	return this.Level 
}

func (this *MonsterBase) getAdaptive() int32 {
	return this.Adaptive 
}

func (this *MonsterBase) getCanDrop() int32 {
	return this.CanDrop 
}

func (this *MonsterBase) getDropID() string {
	return this.DropID 
}

func (this *MonsterBase) getExp() int32 {
	return this.Exp 
}

func (this *MonsterBase) getType() int32 {
	return this.Type 
}

func (this *MonsterBase) getRace() int32 {
	return this.Race 
}

func (this *MonsterBase) getSelectAble() int32 {
	return this.SelectAble 
}

func (this *MonsterBase) getBlockAble() int32 {
	return this.BlockAble 
}

func (this *MonsterBase) getRotateAble() int32 {
	return this.RotateAble 
}

func (this *MonsterBase) getMoveAble() int32 {
	return this.MoveAble 
}

func (this *MonsterBase) getFaction() int32 {
	return this.Faction 
}

func (this *MonsterBase) getAggressivity() int32 {
	return this.Aggressivity 
}

func (this *MonsterBase) getWarnRange() int32 {
	return this.WarnRange 
}

func (this *MonsterBase) getRandomWalk() string {
	return this.RandomWalk 
}

func (this *MonsterBase) getRadius() int32 {
	return this.Radius 
}

func (this *MonsterBase) getScale() float32 {
	return this.Scale 
}

func (this *MonsterBase) getNormalAtk() int32 {
	return this.NormalAtk 
}

func (this *MonsterBase) getAtkSpeedAjust() float32 {
	return this.AtkSpeedAjust 
}

func (this *MonsterBase) getRandDelay() string {
	return this.RandDelay 
}

func (this *MonsterBase) getBasePassiveSkill() string {
	return this.BasePassiveSkill 
}

func (this *MonsterBase) getRandPassiveSkill() string {
	return this.RandPassiveSkill 
}

func (this *MonsterBase) getBaseActiveSkill() string {
	return this.BaseActiveSkill 
}

func (this *MonsterBase) getRandActiveSkill() string {
	return this.RandActiveSkill 
}

func (this *MonsterBase) getTargetChoose() int32 {
	return this.TargetChoose 
}

func (this *MonsterBase) getTargetChange() int32 {
	return this.TargetChange 
}

func (this *MonsterBase) getAiScript() string {
	return this.AiScript 
}

func (this *MonsterBase) getMoveSpeed() int32 {
	return this.MoveSpeed 
}

func (this *MonsterBase) getAttribute() string {
	return this.Attribute 
}

func (this *MonsterBase) getInheritAttrAble() int32 {
	return this.InheritAttrAble 
}

func (this *MonsterBase) getInheritAttrMod() string {
	return this.InheritAttrMod 
}

func (this *MonsterBase) getResType() int32 {
	return this.ResType 
}

func (this *MonsterBase) getModelID() string {
	return this.ModelID 
}

func (this *MonsterBase) getEffectID() string {
	return this.EffectID 
}

func (this *MonsterBase) getBornVfx() string {
	return this.BornVfx 
}

func (this *MonsterBase) getDeathVfx() string {
	return this.DeathVfx 
}

func (this *MonsterBase) getSpriteType() int32 {
	return this.SpriteType 
}

func (this *MonsterBase) getTriggerPassive() int32 {
	return this.TriggerPassive 
}

func (this *MonsterBase) getReadyTime() int32 {
	return this.ReadyTime 
}

func (this *MonsterBase) getWorkTime() int32 {
	return this.WorkTime 
}

func (this *MonsterBase) getWaitTime() int32 {
	return this.WaitTime 
}

func (this *MonsterBase) getTrapOverTime() int32 {
	return this.TrapOverTime 
}

func (this *MonsterBase) getMovePlus() int32 {
	return this.MovePlus 
}

func LoadMonsterBase(file string) string {
	var clen = []int32{47}
	sf := `monsterBase.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*MonsterBase)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &MonsterBase{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Name = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Des1 = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Des2 = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Level, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.Adaptive, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.CanDrop, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 6, val)
		}
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.DropID = val
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.Exp, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.Type, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 9, val)
		}
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.Race, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 10, val)
		}
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		itm.SelectAble, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 11, val)
		}
		val = strings.Replace(row.Cells[12].String(), " \t\r\n", ``, -1)
		itm.BlockAble, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 12, val)
		}
		val = strings.Replace(row.Cells[13].String(), " \t\r\n", ``, -1)
		itm.RotateAble, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 13, val)
		}
		val = strings.Replace(row.Cells[14].String(), " \t\r\n", ``, -1)
		itm.MoveAble, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 14, val)
		}
		val = strings.Replace(row.Cells[15].String(), " \t\r\n", ``, -1)
		itm.Faction, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 15, val)
		}
		val = strings.Replace(row.Cells[16].String(), " \t\r\n", ``, -1)
		itm.Aggressivity, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 16, val)
		}
		val = strings.Replace(row.Cells[17].String(), " \t\r\n", ``, -1)
		itm.WarnRange, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 17, val)
		}
		val = strings.Replace(row.Cells[18].String(), " \t\r\n", ``, -1)
		itm.RandomWalk = val
		val = strings.Replace(row.Cells[19].String(), " \t\r\n", ``, -1)
		itm.Radius, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 19, val)
		}
		val = strings.Replace(row.Cells[20].String(), " \t\r\n", ``, -1)
		f64, err = strconv.ParseFloat(val, 32)
		if err != nil {
			return fmt.Sprintf(`float解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 20, val)
		}
		itm.Scale = float32(f64)
		val = strings.Replace(row.Cells[21].String(), " \t\r\n", ``, -1)
		itm.NormalAtk, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 21, val)
		}
		val = strings.Replace(row.Cells[22].String(), " \t\r\n", ``, -1)
		f64, err = strconv.ParseFloat(val, 32)
		if err != nil {
			return fmt.Sprintf(`float解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 22, val)
		}
		itm.AtkSpeedAjust = float32(f64)
		val = strings.Replace(row.Cells[23].String(), " \t\r\n", ``, -1)
		itm.RandDelay = val
		val = strings.Replace(row.Cells[24].String(), " \t\r\n", ``, -1)
		itm.BasePassiveSkill = val
		val = strings.Replace(row.Cells[25].String(), " \t\r\n", ``, -1)
		itm.RandPassiveSkill = val
		val = strings.Replace(row.Cells[26].String(), " \t\r\n", ``, -1)
		itm.BaseActiveSkill = val
		val = strings.Replace(row.Cells[27].String(), " \t\r\n", ``, -1)
		itm.RandActiveSkill = val
		val = strings.Replace(row.Cells[28].String(), " \t\r\n", ``, -1)
		itm.TargetChoose, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 28, val)
		}
		val = strings.Replace(row.Cells[29].String(), " \t\r\n", ``, -1)
		itm.TargetChange, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 29, val)
		}
		val = strings.Replace(row.Cells[30].String(), " \t\r\n", ``, -1)
		itm.AiScript = val
		val = strings.Replace(row.Cells[31].String(), " \t\r\n", ``, -1)
		itm.MoveSpeed, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 31, val)
		}
		val = strings.Replace(row.Cells[32].String(), " \t\r\n", ``, -1)
		itm.Attribute = val
		val = strings.Replace(row.Cells[33].String(), " \t\r\n", ``, -1)
		itm.InheritAttrAble, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 33, val)
		}
		val = strings.Replace(row.Cells[34].String(), " \t\r\n", ``, -1)
		itm.InheritAttrMod = val
		val = strings.Replace(row.Cells[35].String(), " \t\r\n", ``, -1)
		itm.ResType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 35, val)
		}
		val = strings.Replace(row.Cells[36].String(), " \t\r\n", ``, -1)
		itm.ModelID = val
		val = strings.Replace(row.Cells[37].String(), " \t\r\n", ``, -1)
		itm.EffectID = val
		val = strings.Replace(row.Cells[38].String(), " \t\r\n", ``, -1)
		itm.BornVfx = val
		val = strings.Replace(row.Cells[39].String(), " \t\r\n", ``, -1)
		itm.DeathVfx = val
		val = strings.Replace(row.Cells[40].String(), " \t\r\n", ``, -1)
		itm.SpriteType, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 40, val)
		}
		val = strings.Replace(row.Cells[41].String(), " \t\r\n", ``, -1)
		itm.TriggerPassive, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 41, val)
		}
		val = strings.Replace(row.Cells[42].String(), " \t\r\n", ``, -1)
		itm.ReadyTime, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 42, val)
		}
		val = strings.Replace(row.Cells[43].String(), " \t\r\n", ``, -1)
		itm.WorkTime, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 43, val)
		}
		val = strings.Replace(row.Cells[44].String(), " \t\r\n", ``, -1)
		itm.WaitTime, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 44, val)
		}
		val = strings.Replace(row.Cells[45].String(), " \t\r\n", ``, -1)
		itm.TrapOverTime, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 45, val)
		}
		val = strings.Replace(row.Cells[46].String(), " \t\r\n", ``, -1)
		itm.MovePlus, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 46, val)
		}
		cnf[itm.Id] = itm
	}
	mtxMonsterBase.Lock()
	cnfMonsterBase = cnf
	mtxMonsterBase.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxMonsterBase = new(sync.RWMutex)
var cnfMonsterBase = map[int32]*MonsterBase{
	1: &MonsterBase{
		1, "小怪甲", "小怪近战", "", 10, 0, 0, "10000;10001", 0, 1, 0, 0, 0, 0, 0, 0, 0, 800, "100;200", 100, 1.5, 2, 1.1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s01", "", "", "", 0, 0, 0, 5000, 0, 0, 0, 
	},
	2: &MonsterBase{
		2, "小怪乙", "精英远程", "", 20, 0, 0, "10000;10001", 0, 1, 0, 0, 0, 0, 0, 0, 1, 800, "200;400", 150, 1, 2, 1.2, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:4500;2:120;3:150", 0, "", 0, "mon_001_s02", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	3: &MonsterBase{
		3, "小怪丙", "boss近战", "1章boss", 30, 0, 1, "10000;10001", 0, 1, 0, 0, 0, 0, 0, 0, 2, 800, "100;200", 300, 1, 3, 1.3, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3600;2:50;3:150", 0, "", 0, "mon_001_s03", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	4: &MonsterBase{
		4, "小怪丁", "boss远程", "2章boss", 50, 0, 1, "10000;10001", 0, 1, 0, 0, 0, 0, 0, 0, 0, 800, "200;400", 200, 1.1, 3, 0.8, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:6000;2:80;3:150", 0, "", 0, "mon_010_b01", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	5: &MonsterBase{
		5, "暴风雪", "精灵", "", 1, 0, 0, "", 0, 4, 0, 1, 1, 1, 1, 0, 3, 10000, "", 0, 1, 4, 1, "2000;4000", "", "", "", "", 1, 0, "", 0, "", 0, "1005:100;", 1, "1_ts01_u_000", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	501: &MonsterBase{
		501, "暗月部落萨满", "普通", "1-6", 30, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s01", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	502: &MonsterBase{
		502, "暗月部落战士", "普通", "1-6", 30, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s02", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	503: &MonsterBase{
		503, "刺脊魔", "普通", "1-4", 20, 0, 0, "", 0, 1, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s03", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	504: &MonsterBase{
		504, "行尸", "普通", "1-1", 5, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s04", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	505: &MonsterBase{
		505, "黑暗狂暴者", "精英", "1-5", 25, 0, 0, "", 0, 2, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s05", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	506: &MonsterBase{
		506, "憎恶", "精英", "1-3", 15, 0, 0, "", 0, 2, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s06", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	507: &MonsterBase{
		507, "活死人", "普通", "1-1", 5, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s07", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	508: &MonsterBase{
		508, "厉鬼", "普通", "1-7", 35, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s08", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	509: &MonsterBase{
		509, "骨灰", "精英", "1-6", 30, 0, 0, "", 0, 2, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s09", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	510: &MonsterBase{
		510, "骷髅", "普通", "1-3", 15, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s10", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	511: &MonsterBase{
		511, "骷髅持盾者", "普通", "1-4", 20, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s11", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	512: &MonsterBase{
		512, "骷髅弓手", "普通", "1-2", 10, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s12", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	513: &MonsterBase{
		513, "尸虫", "普通", "1-3", 15, 0, 0, "", 0, 1, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s13", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	514: &MonsterBase{
		514, "尸母", "精英", "1-1", 5, 0, 0, "", 0, 2, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s14", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	515: &MonsterBase{
		515, "食尸鬼", "普通", "1-2", 10, 0, 0, "", 0, 1, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s15", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	516: &MonsterBase{
		516, "废土游魔", "精英", "1-4", 20, 0, 0, "", 0, 2, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s16", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	517: &MonsterBase{
		517, "无棺尸魔", "精英", "1-2", 10, 0, 0, "", 0, 2, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s17", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	518: &MonsterBase{
		518, "凶暴蛮牛怪", "普通", "1-5", 25, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s18", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	519: &MonsterBase{
		519, "月亮族战士", "普通", "1-5", 25, 0, 0, "", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s19", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	520: &MonsterBase{
		520, "骷髅王", "首领", "1-7", 35, 0, 0, "", 0, 3, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s20", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1000: &MonsterBase{
		1000, "暗月部落萨满", "普通", "1-6", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s01", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1001: &MonsterBase{
		1001, "暗月部落战士", "普通", "1-6", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s02", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1002: &MonsterBase{
		1002, "刺脊魔", "普通", "1-4", 1, 0, 0, "10000;10001", 0, 1, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s03", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1003: &MonsterBase{
		1003, "行尸", "普通", "1-1", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s04", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1004: &MonsterBase{
		1004, "黑暗狂暴者", "精英", "1-5", 1, 0, 1, "10000;10001", 0, 2, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s05", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1005: &MonsterBase{
		1005, "憎恶", "精英", "1-3", 1, 0, 1, "10000;10001", 0, 2, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s06", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1006: &MonsterBase{
		1006, "活死人", "普通", "1-1", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s07", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1007: &MonsterBase{
		1007, "厉鬼", "普通", "1-7", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s08", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1008: &MonsterBase{
		1008, "骨灰", "精英", "1-6", 1, 0, 1, "10000;10001", 0, 2, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s09", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1009: &MonsterBase{
		1009, "骷髅", "普通", "1-3", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s10", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1010: &MonsterBase{
		1010, "骷髅持盾者", "普通", "1-4", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s11", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1011: &MonsterBase{
		1011, "骷髅弓手", "普通", "1-2", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s12", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1012: &MonsterBase{
		1012, "尸虫", "普通", "1-3", 1, 0, 0, "10000;10001", 0, 1, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s13", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1013: &MonsterBase{
		1013, "尸母", "精英", "1-1", 1, 0, 1, "10000;10001", 0, 2, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s14", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1014: &MonsterBase{
		1014, "食尸鬼", "普通", "1-2", 1, 0, 0, "10000;10001", 0, 1, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s15", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1015: &MonsterBase{
		1015, "废土游魔", "精英", "1-4", 1, 0, 1, "10000;10001", 0, 2, 1, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s16", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1016: &MonsterBase{
		1016, "无棺尸魔", "精英", "1-2", 1, 0, 1, "10000;10001", 0, 2, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s17", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1017: &MonsterBase{
		1017, "凶暴蛮牛怪", "普通", "1-5", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s18", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1018: &MonsterBase{
		1018, "月亮族战士", "普通", "1-5", 1, 0, 0, "10000;10001", 0, 1, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s19", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
	1019: &MonsterBase{
		1019, "骷髅王", "首领", "1-7", 1, 0, 1, "10000;10001", 0, 3, 2, 0, 0, 0, 0, 1, 1, 800, "100;200", 300, 1, 2, 1, "2000;4000", "", "", "", "", 0, 0, "", 0, "1:3000;2:100;3:160", 0, "", 0, "mon_001_s20", "", "", "", 0, 0, 0, 0, 0, 0, 0, 
	},
}

	func MonsterBase_hot() {
		for _, val := range cnfMonsterBase	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}