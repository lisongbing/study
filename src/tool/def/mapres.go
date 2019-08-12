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

type MapResVersion struct {
	md5 string //配置表md5数据
}

type MapResClass struct {
	MapResVersion
}

var MapResinstance *MapResClass

func MapResInstance() *MapResClass {
	if MapResinstance == nil {
		MapResinstance = &MapResClass{}
	}
	return MapResinstance
}

type MapRes struct {
	Id int32 //id
	ResID int32 //资源ID
	Number int32 //资源编号
	Describe string //说明
	Res string //资源
	BornPos string //出生点
	Music string //背景音乐
	MapID int32 //地图阻挡ID
	Width int32 //宽
	Height int32 //高
	GroupMonster string //怪物组刷新编号
	Elite string //本地图精英怪
	Boss string //本地图BOSS怪
}

func GetMapResByPk(id int32) (itm *MapRes, ok bool) {
	mtxMapRes.RLock()
	itm, ok = cnfMapRes[id]
	mtxMapRes.RUnlock()
	return
}

func SetMapResVersion(md5 string) string {
	MapResInstance().md5 = md5
	return ``
	}
func GetMapResVersion(md5 string) string {
	return MapResInstance().md5
}
func GetMapRes() map[int32]*MapRes{
	mtxMapRes.RLock()
	cnf := cnfMapRes
	mtxMapRes.RUnlock()
	return cnf
}

func (this *MapRes) getResID() int32 {
	return this.ResID 
}

func (this *MapRes) getNumber() int32 {
	return this.Number 
}

func (this *MapRes) getDescribe() string {
	return this.Describe 
}

func (this *MapRes) getRes() string {
	return this.Res 
}

func (this *MapRes) getBornPos() string {
	return this.BornPos 
}

func (this *MapRes) getMusic() string {
	return this.Music 
}

func (this *MapRes) getMapID() int32 {
	return this.MapID 
}

func (this *MapRes) getWidth() int32 {
	return this.Width 
}

func (this *MapRes) getHeight() int32 {
	return this.Height 
}

func (this *MapRes) getGroupMonster() string {
	return this.GroupMonster 
}

func (this *MapRes) getElite() string {
	return this.Elite 
}

func (this *MapRes) getBoss() string {
	return this.Boss 
}

func LoadMapRes(file string) string {
	var clen = []int32{13}
	sf := `mapres.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*MapRes)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &MapRes{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.ResID, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.Number, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Describe = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.Res = val
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.BornPos = val
		val = strings.Replace(row.Cells[6].String(), " \t\r\n", ``, -1)
		itm.Music = val
		val = strings.Replace(row.Cells[7].String(), " \t\r\n", ``, -1)
		itm.MapID, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 7, val)
		}
		val = strings.Replace(row.Cells[8].String(), " \t\r\n", ``, -1)
		itm.Width, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 8, val)
		}
		val = strings.Replace(row.Cells[9].String(), " \t\r\n", ``, -1)
		itm.Height, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 9, val)
		}
		val = strings.Replace(row.Cells[10].String(), " \t\r\n", ``, -1)
		itm.GroupMonster = val
		val = strings.Replace(row.Cells[11].String(), " \t\r\n", ``, -1)
		itm.Elite = val
		val = strings.Replace(row.Cells[12].String(), " \t\r\n", ``, -1)
		itm.Boss = val
		cnf[itm.Id] = itm
	}
	mtxMapRes.Lock()
	cnfMapRes = cnf
	mtxMapRes.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxMapRes = new(sync.RWMutex)
var cnfMapRes = map[int32]*MapRes{
	101: &MapRes{
		101, 1, 1, "剧情1", "map11_tmx", "1010:1712", "BG_act1_1", 1001, 2048, 2048, "monster1", "1013", "1013", 
	},
	201: &MapRes{
		201, 2, 1, "剧情2", "map12_tmx", "200:1176", "BG_act1_1", 1001, 1536, 1536, "monster1", "1016", "1016", 
	},
	301: &MapRes{
		301, 3, 1, "剧情3", "map13_tmx", "1388:1742", "BG_act1_1", 1001, 2560, 2048, "monster1", "1005", "1005", 
	},
	401: &MapRes{
		401, 4, 1, "剧情4", "map14_tmx", "361:1975", "BG_act1_1", 1001, 1536, 2048, "monster1", "1015", "1015", 
	},
	501: &MapRes{
		501, 5, 1, "剧情5", "map15_tmx", "146:1934", "BG_act1_1", 1001, 1536, 2048, "monster1", "1004", "1004", 
	},
	601: &MapRes{
		601, 6, 1, "剧情6", "map16_tmx", "1333:639", "BG_act1_1", 1001, 1536, 2048, "monster1", "1008", "1008", 
	},
	701: &MapRes{
		701, 7, 1, "剧情7", "map17_tmx", "237:1363", "BG_act1_1", 1001, 2560, 1536, "monster1", "1019", "1019", 
	},
}

	func MapRes_hot() {
		for _, val := range cnfMapRes	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}