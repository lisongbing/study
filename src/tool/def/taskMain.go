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

type TaskMainVersion struct {
	md5 string //配置表md5数据
}

type TaskMainClass struct {
	TaskMainVersion
}

var TaskMaininstance *TaskMainClass

func TaskMainInstance() *TaskMainClass {
	if TaskMaininstance == nil {
		TaskMaininstance = &TaskMainClass{}
	}
	return TaskMaininstance
}

type TaskMain struct {
	Id int32 //任务ID
	Name string //任务名称
	Desc string //任务描述
	Type int32 //任务分类
	Number int32 //任务序号
	Activation int32 //任务激活条件
	Conditions string //激活条件参数
	Style int32 //任务完成类型
	Complete string //完成任务参数
	Trace string //任务追踪配置
	PositionID string //目标位置信息
	TaskReward int32 //任务奖励ID
	NextTask string //完成任务后添加任务
	Gonpc string //进行中任务NPC信息
	GodialogueRole string //进行中任务NPC对白
	Awardnpc string //领奖任务NPC信息
	AwarddialogueRole string //领奖任务NPC对白
}

func GetTaskMainByPk(id int32) (itm *TaskMain, ok bool) {
	mtxTaskMain.RLock()
	itm, ok = cnfTaskMain[id]
	mtxTaskMain.RUnlock()
	return
}

const (
	Activation_Actlevel=0//达到普通等级
	Activation_Actlegend=1//达到巅峰等级
	Activation_Acttalk=2 // 对话(地图id+XY坐标）
	Activation_Acttask=3 //由其他任务完成后激活
	Activation_Actkill=4// 杀怪数量
	Activation_Actcopyid=7// 到达特定副本ID
	Activation_Actmaptime=8// 计时副本ID+时间限制
	Activation_Actitem=9//获得道具数量
	Activation_Actequip=10//获得装备
	Activation_Actupskill=15// 升级技能
	Activation_Actunion=16// 加入公会
)

const (
	Style_Getlevel=0//达到普通等级
	Style_Getlegend=1//巅峰等级
	Style_Gettalk=2//对话(地图id+XY坐标）
	Style_Getcollect=3//采集
	Style_Getkill=4//杀怪数量
	Style_Getdrop=5//杀怪掉落
	Style_Getflush=6//主动刷怪杀怪
	Style_Getcopyid=7//传送到特定副本ID
	Style_Getmaptime=8//计时副本ID+时间限制
	Style_Getitem=9//获得道具数量
	Style_Getequip=10//获得装备
	Style_Getrec=11//任务回收道具or装备数量
	Style_Getupequip=12//升级魂装
	Style_Getingem=13//镶嵌宝石
	Style_Getinamu=14//镶嵌护符
	Style_Getupskill=15//升级技能
	Style_Getunion=16//加入公会
)

func SetTaskMainVersion(md5 string) string {
	TaskMainInstance().md5 = md5
	return ``
	}
func GetTaskMainVersion(md5 string) string {
	return TaskMainInstance().md5
}
func GetTaskMain() map[int32]*TaskMain{
	mtxTaskMain.RLock()
	cnf := cnfTaskMain
	mtxTaskMain.RUnlock()
	return cnf
}

func (this *TaskMain) getName() string {
	return this.Name 
}

func (this *TaskMain) getDesc() string {
	return this.Desc 
}

func (this *TaskMain) getType() int32 {
	return this.Type 
}

func (this *TaskMain) getNumber() int32 {
	return this.Number 
}

func (this *TaskMain) getActivation() int32 {
	return this.Activation 
}

func (this *TaskMain) getConditions() string {
	return this.Conditions 
}

func (this *TaskMain) getStyle() int32 {
	return this.Style 
}

func (this *TaskMain) getComplete() string {
	return this.Complete 
}

func (this *TaskMain) getTrace() string {
	return this.Trace 
}

func (this *TaskMain) getPositionID() string {
	return this.PositionID 
}

func (this *TaskMain) getTaskReward() int32 {
	return this.TaskReward 
}

func (this *TaskMain) getNextTask() string {
	return this.NextTask 
}

func (this *TaskMain) getGonpc() string {
	return this.Gonpc 
}

func (this *TaskMain) getGodialogueRole() string {
	return this.GodialogueRole 
}

func (this *TaskMain) getAwardnpc() string {
	return this.Awardnpc 
}

func (this *TaskMain) getAwarddialogueRole() string {
	return this.AwarddialogueRole 
}

func LoadTaskMain(file string) string {
	var clen = []int32{17}
	sf := `taskMain.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*TaskMain)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &TaskMain{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.Name = val
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Desc = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Type, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Number, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.Activation, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.Conditions = val
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.Style, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.Complete = val
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.Trace = val
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.PositionID = val
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		itm.TaskReward, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 11, val)
		}
		val = strings.Replace(row.Cells[12].String(), " \t\r\n", ``, -1)
		itm.NextTask = val
		val = strings.Replace(row.Cells[13].String(), " \t\r\n", ``, -1)
		itm.Gonpc = val
		val = strings.Replace(row.Cells[14].String(), " \t\r\n", ``, -1)
		itm.GodialogueRole = val
		val = strings.Replace(row.Cells[15].String(), " \t\r\n", ``, -1)
		itm.Awardnpc = val
		val = strings.Replace(row.Cells[16].String(), " \t\r\n", ``, -1)
		itm.AwarddialogueRole = val
		cnf[itm.Id] = itm
	}
	mtxTaskMain.Lock()
	cnfTaskMain = cnf
	mtxTaskMain.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxTaskMain = new(sync.RWMutex)
var cnfTaskMain = map[int32]*TaskMain{
	10001: &TaskMain{
		10001, "10102", "10107", 1, 1, 0, "0", 2, "map:11001;x:154;y:1314", "10107", "map:11001;x:154;y:1314", 90005, "10002", "11001:6", "0:10103;1:10104;1:10105;0:10106", "11001:1", "1:10108;0:10109;1:10110;0:10111", 
	},
	10002: &TaskMain{
		10002, "10112", "10118", 1, 2, 3, "0", 4, "1003:30", "10118", "map:11001;x:500;y:1068", 90005, "10003", "11001:1", "1:10113;0:10114;1:10115;0:10116;0:10117", "11001:6", "1:10119;0:10120", 
	},
	10003: &TaskMain{
		10003, "10121", "10123", 1, 3, 3, "0", 2, "map:11001;x:1506;y:1140", "10123", "map:11001;x:1506;y:1140", 90005, "10004", "11001:1", "1:10122", "11001:2", "1:10124;0:10125;0:10126", 
	},
	10004: &TaskMain{
		10004, "10127", "10131", 1, 4, 3, "0", 4, "1006:30", "10131", "map:11001;x:1094;y:738", 90005, "10005", "11001:2", "1:10128;0:10129;1:10130", "11001:2", "0:10132;0:10133", 
	},
	10005: &TaskMain{
		10005, "10134", "10138", 1, 5, 3, "0", 4, "1003:60", "10138", "map:11001;x:872;y:1520", 90005, "10006", "11001:2", "0:10135;0:10136;0:10137", "11001:2", "0:10139;0:10140", 
	},
	10006: &TaskMain{
		10006, "10141", "10145", 1, 6, 3, "0", 4, "1006:60", "10145", "map:11001;x:1338;y:1714", 90005, "10007", "11001:2", "0:10142;0:10143;0:10144", "11001:2", "0:10146;1:10147;0:10148", 
	},
	10007: &TaskMain{
		10007, "10149", "10154", 1, 7, 3, "0", 6, "1013:1", "10154", "map:11001;x:1026;y:1200", 90005, "10008", "11001:2", "0:10150;0:10151;0:10152;0:10153", "11001:3", "0:10155;0:10156;1:10157", 
	},
	10008: &TaskMain{
		10008, "10159", "10163", 1, 8, 3, "0", 4, "1011:30", "10163", "map:11002;x:211;y:1196", 90005, "10009", "11002:3", "0:10160;0:10161;0:10162", "11002:3", "0:10164;0:10165", 
	},
	10009: &TaskMain{
		10009, "10166", "10169", 1, 9, 3, "0", 4, "1014:30", "10169", "map:11002;x:1015;y:783", 90005, "10010", "11002:3", "0:10167;0:10168", "11002:7", "0:10170", 
	},
	10010: &TaskMain{
		10010, "10171", "10175", 1, 10, 3, "0", 4, "1011:60", "10175", "map:11002;x:835;y:363", 90005, "10011", "11002:7", "0:10172;0:10173;1:10174", "11002:7", "0:10176;0:10177", 
	},
	10011: &TaskMain{
		10011, "10178", "10181", 1, 11, 3, "0", 4, "1014:60", "10181", "map:11002;x:404;y:465", 90005, "10012", "11002:7", "0:10179;0:10180", "11002:7", "0:10182", 
	},
	10012: &TaskMain{
		10012, "10183", "10187", 1, 12, 3, "0", 6, "1016:1", "10187", "map:11002;x:944;y:467", 90005, "10013", "11002:7", "0:10184;0:10185;1:10186", "11002:8", "0:10188;1:10189;0:10190", 
	},
	10013: &TaskMain{
		10013, "10192", "10195", 1, 13, 3, "0", 4, "1012:30", "10195", "map:11003;x:838;y:1586", 90005, "10014", "11003:3", "1:10193;0:10194", "11003:3", "0:10196;0:10197", 
	},
	10014: &TaskMain{
		10014, "10198", "10202", 1, 14, 3, "0", 4, "1009:30", "10202", "map:11003;x:1170;y:1262", 90005, "10015", "11003:3", "0:10199;0:10200;1:10201", "11003:2", "1:10203;0:10204", 
	},
	10015: &TaskMain{
		10015, "10205", "10208", 1, 15, 3, "0", 4, "1012:60", "10208", "map:11003;x:838;y:1586", 90005, "10016", "11003:2", "1:10206;0:10207", "11003:2", "0:10209;0:10210", 
	},
	10016: &TaskMain{
		10016, "10218", "10215", 1, 16, 3, "0", 4, "1009:60", "10215", "map:11003;x:1170;y:1262", 90005, "10017", "11003:2", "1:10212;0:10213;0:10214", "11003:2", "0:10216;0:10217", 
	},
	10017: &TaskMain{
		10017, "10211", "10220", 1, 17, 3, "0", 6, "1005:1", "10220", "map:11003;x:1170;y:1262", 90005, "10018", "11003:7", "0:10219", "11003:7", "0:10221;1:10222;1:10223", 
	},
	10018: &TaskMain{
		10018, "10225", "10228", 1, 18, 3, "0", 2, "map:11004;x:341;y:1871", "10228", "map:11004;x:341;y:1871", 90005, "10019", "11004:2", "0:10226;1:10227", "11004:3", "0:10229;0:10230", 
	},
	10019: &TaskMain{
		10019, "10231", "10234", 1, 19, 3, "0", 4, "1002:30", "10234", "map:11004;x:912;y:1470", 90005, "10020", "11004:2", "0:10232;0:10233", "11004:3", "1:10235", 
	},
	10020: &TaskMain{
		10020, "10236", "10239", 1, 20, 3, "0", 4, "1010:30", "10239", "map:11004;x:329;y:776", 90005, "10021", "11004:3", "0:10237;0:10238", "11004:3", "0:10240", 
	},
	10021: &TaskMain{
		10021, "10241", "10245", 1, 21, 3, "0", 4, "1002:60", "10245", "map:11004;x:912;y:1470", 90005, "10022", "11004:3", "0:10242;0:10243;0:10244", "11004:7", "0:10246;0:10247;0:10248;0:10249;0:10250;0:10251", 
	},
	10022: &TaskMain{
		10022, "10252", "10255", 1, 22, 3, "0", 4, "1010:60", "10255", "map:11004;x:329;y:776", 90005, "10023", "11004:7", "0:10253;0:10254", "11004:7", "0:10256;0:10257", 
	},
	10023: &TaskMain{
		10023, "10258", "10260", 1, 23, 3, "0", 6, "1015:1", "10260", "map:11004;x:669;y:1148", 90005, "10024", "11004:7", "0:10259", "11004:7", "0:10261;0:10262", 
	},
	10024: &TaskMain{
		10024, "10264", "10267", 1, 24, 3, "0", 2, "map:11005;x:118;y:1890", "10267", "map:11005;x:118;y:1890", 90005, "10025", "11005:3", "0:10265;0:10266", "11005:2", "1:10268;0:10269", 
	},
	10025: &TaskMain{
		10025, "10270", "10272", 1, 25, 3, "0", 4, "1018:30", "10272", "map:11005;x:178;y:1490", 90005, "10026", "11005:3", "0:10271", "11005:3", "0:10273;1:10274", 
	},
	10026: &TaskMain{
		10026, "10275", "10278", 1, 26, 3, "0", 4, "1017:30", "10278", "map:11005;x:733;y:760", 90005, "10027", "11005:3", "0:10276;0:10277", "11005:3", "0:10279", 
	},
	10027: &TaskMain{
		10027, "10280", "10282", 1, 27, 3, "0", 4, "1018:60", "10282", "map:11005;x:178;y:1490", 90005, "10028", "11005:3", "0:10281", "11005:2", "0:10283", 
	},
	10028: &TaskMain{
		10028, "10284", "10287", 1, 28, 3, "0", 4, "1017:60", "10287", "map:11005;x:733;y:760", 90005, "10029", "11005:3", "0:10285;0:10286", "11005:3", "0:10288;0:10289;1:10290", 
	},
	10029: &TaskMain{
		10029, "10291", "10293", 1, 29, 3, "0", 6, "1004:1", "10293", "map:11005;x:641;y:921", 90005, "10030", "11005:3", "0:10292", "11005:3", "0:10294;0:10295;1:10296", 
	},
	10030: &TaskMain{
		10030, "10298", "10303", 1, 30, 3, "0", 2, "map:11006;x:1300;y:541", "10303", "map:11006;x:1300;y:541", 90005, "10031", "11006:3", "0:10299;1:10300;0:10301;0:10302", "11006:3", "0:10304;0:10305", 
	},
	10031: &TaskMain{
		10031, "10306", "10309", 1, 31, 3, "0", 4, "1001:30", "10309", "map:11006;x:665;y:615", 90005, "10032", "11006:3", "0:10307;1:10308", "11006:2", "0:10310;0:10311", 
	},
	10032: &TaskMain{
		10032, "10312", "10316", 1, 32, 3, "0", 4, "1000:30", "10316", "map:11006;x:145;y:1551", 90005, "10033", "11006:2", "0:10313;0:10314;0:10315", "11006:3", "0:10317;0:10318", 
	},
	10033: &TaskMain{
		10033, "10319", "10323", 1, 33, 3, "0", 4, "1001:60", "10323", "map:11006;x:665;y:615", 90005, "10034", "11006:2", "0:10320;0:10321;0:10322", "11006:3", "0:10324;0:10325", 
	},
	10034: &TaskMain{
		10034, "10326", "10331", 1, 34, 3, "0", 4, "1000:60", "10331", "map:11006;x:145;y:1551", 90005, "10035", "11006:3", "0:10327;0:10328;0:10329;0:10330", "11006:3", "0:10332;0:10333", 
	},
	10035: &TaskMain{
		10035, "10334", "10338", 1, 35, 3, "0", 6, "1008:1", "10338", "map:11006;x:308;y:1908", 90005, "10036", "11006:3", "0:10335;0:10336;0:10337", "11006:2", "0:10339;0:10340", 
	},
	10036: &TaskMain{
		10036, "10342", "10346", 1, 36, 3, "0", 2, "map:11007;x:439;y:1300", "10346", "map:11007;x:439;y:1300", 90005, "10037", "11007:3", "0:10343;0:10344;0:10345", "11007:2", "0:10347;0:10348", 
	},
	10037: &TaskMain{
		10037, "10349", "10352", 1, 37, 3, "0", 4, "1007:30", "10352", "map:11007;x:1799;y:1205", 90005, "10038", "11007:3", "0:10350;0:10351", "11007:2", "0:10353;1:10354", 
	},
	10038: &TaskMain{
		10038, "10355", "10358", 1, 38, 3, "0", 4, "1010:30", "10358", "map:11007;x:1199;y:913", 90005, "10039", "11007:2", "0:10356;0:10357", "11007:2", "0:10359;0:10360;0:10361", 
	},
	10039: &TaskMain{
		10039, "10362", "10365", 1, 39, 3, "0", 4, "1007:60", "10365", "map:11007;x:1799;y:1205", 90005, "10040", "11007:2", "0:10363;0:10364", "11007:2", "0:10366;0:10367", 
	},
	10040: &TaskMain{
		10040, "10368", "10371", 1, 40, 3, "0", 4, "1010:60", "10371", "map:11007;x:1199;y:913", 90005, "10041", "11007:3", "0:10369;0:10370", "11007:2", "0:10372;0:10373;0:10374", 
	},
	10041: &TaskMain{
		10041, "10375", "10380", 1, 41, 3, "0", 6, "1019:1", "10380", "map:11007;x:2088;y:323", 90005, "0", "11007:2", "0:10376;0:10377;0:10378;1:10379", "11007:3", "1:10381;0:10382;0:10383", 
	},
}

	func TaskMain_hot() {
		for _, val := range cnfTaskMain	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}