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

type GiftDataVersion struct {
	md5 string //配置表md5数据
}

type GiftDataClass struct {
	GiftDataVersion
}

var GiftDatainstance *GiftDataClass

func GiftDataInstance() *GiftDataClass {
	if GiftDatainstance == nil {
		GiftDatainstance = &GiftDataClass{}
	}
	return GiftDatainstance
}

type GiftData struct {
	Id int32 //id
	GiftLv int32 //礼包等级
	AwardMethod int32 //奖励类型
	Award string //奖励
	DropID string //掉落id
}

func GetGiftDataByPk(id int32) (itm *GiftData, ok bool) {
	mtxGiftData.RLock()
	itm, ok = cnfGiftData[id]
	mtxGiftData.RUnlock()
	return
}

const (
	AwardMethod_Default = 0 //填什么给什么
	AwardMethod_ClassChoice = 1 //按职业限制奖励,根据道具的职业限制，自动筛选出当前主角职业可用的物品
)

func SetGiftDataVersion(md5 string) string {
	GiftDataInstance().md5 = md5
	return ``
	}
func GetGiftDataVersion(md5 string) string {
	return GiftDataInstance().md5
}
func GetGiftData() map[int32]*GiftData{
	mtxGiftData.RLock()
	cnf := cnfGiftData
	mtxGiftData.RUnlock()
	return cnf
}

func (this *GiftData) getGiftLv() int32 {
	return this.GiftLv 
}

func (this *GiftData) getAwardMethod() int32 {
	return this.AwardMethod 
}

func (this *GiftData) getAward() string {
	return this.Award 
}

func (this *GiftData) getDropID() string {
	return this.DropID 
}

func LoadGiftData(file string) string {
	var clen = []int32{5}
	sf := `giftData.xlsx`
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
	cnf := make(map[int32]*GiftData)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &GiftData{}
		val = strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1)
		itm.Id, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 0, val)
		}
		val = strings.Replace(row.Cells[1].String(), " \t\r\n", ``, -1)
		itm.GiftLv, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 1, val)
		}
		val = strings.Replace(row.Cells[2].String(), " \t\r\n", ``, -1)
		itm.AwardMethod, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 2, val)
		}
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.Award = val
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.DropID = val
		cnf[itm.Id] = itm
	}
	mtxGiftData.Lock()
	cnfGiftData = cnf
	mtxGiftData.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxGiftData = new(sync.RWMutex)
var cnfGiftData = map[int32]*GiftData{
	201: &GiftData{
		201, 10, 0, "1001:1;41107:1", "10000;10001", 
	},
}

	func GiftData_hot() {
		for _, val := range cnfGiftData	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}