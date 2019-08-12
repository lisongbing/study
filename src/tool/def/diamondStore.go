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

type DiamondStoreVersion struct {
	md5 string //配置表md5数据
}

type DiamondStoreClass struct {
	DiamondStoreVersion
}

var DiamondStoreinstance *DiamondStoreClass

func DiamondStoreInstance() *DiamondStoreClass {
	if DiamondStoreinstance == nil {
		DiamondStoreinstance = &DiamondStoreClass{}
	}
	return DiamondStoreinstance
}

type DiamondStore struct {
	Id int32 //
	Type int32 //
	Note string //注释
	ItemID int32 //道具id
	BuyCount int32 //每日可购买数量
	Price int32 //钻石价格
}

func GetDiamondStoreByPk(id int32) (itm *DiamondStore, ok bool) {
	mtxDiamondStore.RLock()
	itm, ok = cnfDiamondStore[id]
	mtxDiamondStore.RUnlock()
	return
}

func SetDiamondStoreVersion(md5 string) string {
	DiamondStoreInstance().md5 = md5
	return ``
	}
func GetDiamondStoreVersion(md5 string) string {
	return DiamondStoreInstance().md5
}
func GetDiamondStore() map[int32]*DiamondStore{
	mtxDiamondStore.RLock()
	cnf := cnfDiamondStore
	mtxDiamondStore.RUnlock()
	return cnf
}

func (this *DiamondStore) getType() int32 {
	return this.Type 
}

func (this *DiamondStore) getNote() string {
	return this.Note 
}

func (this *DiamondStore) getItemID() int32 {
	return this.ItemID 
}

func (this *DiamondStore) getBuyCount() int32 {
	return this.BuyCount 
}

func (this *DiamondStore) getPrice() int32 {
	return this.Price 
}

func LoadDiamondStore(file string) string {
	var clen = []int32{6}
	sf := `diamondStore.xlsx`
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
    var shref = []int64{0, 0, 0, 0, 0, 0}
	_ = shref
	cnf := make(map[int32]*DiamondStore)
	for rdx, row := range f.Sheets[0].Rows[4:] {
		if int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if int32(len(row.Cells)) != clen[0] {
			return fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数："%d"，实际列数："%d"`, sf, rdx+5, clen[0], len(row.Cells))
		}
		itm := &DiamondStore{}
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
		itm.Note = val
		val = strings.Replace(row.Cells[3].String(), " \t\r\n", ``, -1)
		itm.ItemID, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 3, val)
		}
		val = strings.Replace(row.Cells[4].String(), " \t\r\n", ``, -1)
		itm.BuyCount, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 4, val)
		}
		val = strings.Replace(row.Cells[5].String(), " \t\r\n", ``, -1)
		itm.Price, err = getInt(val)
		if err != nil {
			return fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, 5, val)
		}
		cnf[itm.Id] = itm
	}
	mtxDiamondStore.Lock()
	cnfDiamondStore = cnf
	mtxDiamondStore.Unlock()
	return ``
}

var _ = strconv.ErrRange
var mtxDiamondStore = new(sync.RWMutex)
var cnfDiamondStore = map[int32]*DiamondStore{
	1: &DiamondStore{
		1, 1, "白宝石1级", 1001, -1, 2, 
	},
	2: &DiamondStore{
		2, 1, "白宝石2级", 1002, -1, 4, 
	},
	3: &DiamondStore{
		3, 1, "白宝石3级", 1003, 10, 6, 
	},
	4: &DiamondStore{
		4, 1, "白宝石4级", 1004, 10, 8, 
	},
	5: &DiamondStore{
		5, 1, "白宝石5级", 1005, 10, 10, 
	},
	6: &DiamondStore{
		6, 1, "白宝石6级", 1006, 10, 12, 
	},
	7: &DiamondStore{
		7, 1, "白宝石7级", 1007, 10, 14, 
	},
	8: &DiamondStore{
		8, 1, "白宝石8级", 1008, 10, 16, 
	},
	9: &DiamondStore{
		9, 1, "白宝石9级", 1009, 10, 18, 
	},
	10: &DiamondStore{
		10, 1, "白宝石10级", 1010, 10, 20, 
	},
	100: &DiamondStore{
		100, 2, "铸魂材料", 101, 20, 20, 
	},
	101: &DiamondStore{
		101, 2, "灵魂石", 102, -1, 100, 
	},
	200: &DiamondStore{
		200, 3, "李奥瑞克的黄金护颈", 40005, 15, 200, 
	},
	201: &DiamondStore{
		201, 3, "骷髅王肩铠", 40006, -1, 20, 
	},
	300: &DiamondStore{
		300, 4, "贵族钻石", 40003, 100, 10, 
	},
	301: &DiamondStore{
		301, 4, "李奥瑞克的红宝石", 40004, -1, 20, 
	},
}

	func DiamondStore_hot() {
		for _, val := range cnfDiamondStore	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {}
			}
		}
	}