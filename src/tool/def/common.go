package def

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	re_num, _ = regexp.Compile(`[\|\+\-\*\/\^\&]`)
    errInvalidInt = errors.New("int格式不正确！")
	func_map map[string][]func(string) string
)
type Update struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Full string `json:"Full"`
	Md5 string `json:"md5"`
}
type HttpConfig struct {
	Http string `json:"http"`
}
func getInt(s string) (int32, error) {
	if len(s)>1 && s[0] == '-'{
		t, e := strconv.ParseInt(s, 10, 64)
		if e != nil {
			return 0,errInvalidInt
		}
		return int32(t),nil
	}
	arr := re_num.Split(s, -1)
	if len(arr) == 0 {
		return 0, errInvalidInt
	}
	var v, t int64
	var e error
	var n int32
	for idx, num := range arr {
		if num == "" {
			return 0, errInvalidInt
		}
		if idx == 0 {
			t, e = strconv.ParseInt(num, 10, 64)
			if e != nil {
				break
			}
			v = t
		} else {
			t, e = strconv.ParseInt(num, 10, 64)
			if e != nil {
				break
			}
			switch s[n] {
			case '|':
				v |= t
			case '+':
				v += t
			case '-':
				v -= t
			case '*':
				v *= t
			case '/':
				v /= t
			case '^':
				v ^= t
			case '&':
				v &= t
			}
			n += 1
		}
		n += int32(len(num))
	}
	return int32(v), e
}

func init() {
	func_map = map[string][]func(string) string{
		`charmRecipeData.xlsx`: { LoadCharmsRecipeData, SetCharmsRecipeDataVersion, GetCharmsRecipeDataVersion },
		`createrole.xlsx`: { LoadCreaterole, SetCreateroleVersion, GetCreateroleVersion },
		`decomposition.xlsx`: { LoadDecomposition, SetDecompositionVersion, GetDecompositionVersion },
		`diamondStore.xlsx`: { LoadDiamondStore, SetDiamondStoreVersion, GetDiamondStoreVersion },
		`dropData.xlsx`: { LoadDropData, SetDropDataVersion, GetDropDataVersion },
		`dropEquip.xlsx`: { LoadDropEquip, SetDropEquipVersion, GetDropEquipVersion },
		`dropGoods.xlsx`: { LoadDropGoods, SetDropGoodsVersion, GetDropGoodsVersion },
		`equipAffixes.xlsx`: { LoadEquipAffixes, SetEquipAffixesVersion, GetEquipAffixesVersion },
		`equipCompound.xlsx`: { LoadEquipCompound, SetEquipCompoundVersion, GetEquipCompoundVersion },
		`equipFusion.xlsx`: { LoadEquipFusion, SetEquipFusionVersion, GetEquipFusionVersion },
		`equipUpgrade.xlsx`: { LoadEquipUpgrade, SetEquipUpgradeVersion, GetEquipUpgradeVersion },
		`equipment.xlsx`: { LoadEquipment, SetEquipmentVersion, GetEquipmentVersion },
		`gemData.xlsx`: { LoadGemData, SetGemDataVersion, GetGemDataVersion },
		`gemRecipeData.xlsx`: { LoadGemRecipeData, SetGemRecipeDataVersion, GetGemRecipeDataVersion },
		`giftData.xlsx`: { LoadGiftData, SetGiftDataVersion, GetGiftDataVersion },
		`globalParam.xlsx`: { LoadGlobalParam, SetGlobalParamVersion, GetGlobalParamVersion },
		`innateAffixes.xlsx`: { LoadInnateAffixes, SetInnateAffixesVersion, GetInnateAffixesVersion },
		`item.xlsx`: { LoadItem, SetItemVersion, GetItemVersion },
		`loadTips.xlsx`: { LoadLoadTips, SetLoadTipsVersion, GetLoadTipsVersion },
		`mapres.xlsx`: { LoadMapRes, SetMapResVersion, GetMapResVersion },
		`monsterBase.xlsx`: { LoadMonsterBase, SetMonsterBaseVersion, GetMonsterBaseVersion },
		`player0Base.xlsx`: { LoadPlayerBase, SetPlayerBaseVersion, GetPlayerBaseVersion },
		`playerData.xlsx`: { LoadPlayerData, SetPlayerDataVersion, GetPlayerDataVersion },
		`runeRecipeData.xlsx`: { LoadRuneRecipeData, SetRuneRecipeDataVersion, GetRuneRecipeDataVersion },
		`svrsignInData.xlsx`: { LoadSignInData, SetSignInDataVersion, GetSignInDataVersion },
		`taskMain.xlsx`: { LoadTaskMain, SetTaskMainVersion, GetTaskMainVersion },
	}
}

func LoadAll() []string {
    var earr []string
    for _, f := range func_map {
        err := f[0]("")
        if err != "" {
            earr = append(earr, err)
        }
    }
    return earr
}

func LoadFile(file string) string {
    name := filepath.Base(file)
    f, ok := func_map[name]
    if ok {
        return f[0](file)
    }
    return "这是一个新的配置文件，无法热加载！"
}
func LoadDiff() {
	dir,_ := getCurrentPath()

	c := LoadHttpConfig(dir)

	updatelist, _ := GetFiles(dir)
	for _, Update := range updatelist {

		url := c.Http + Update.Path + Update.Name
		res, gerr := http.Get(url)
		if gerr != nil {
			panic(gerr)
		}

		//文件是否存在

		b, _ := PathExists(dir + "/res/" + Update.Name)
		if false == b {
			fmt.Println("文件不存在 ", Update.Name)
			continue
		}

		//判断md5
		md5 := GetMd5(dir + "/res/" + Update.Name)
		if Update.Md5 == md5 {
			fmt.Println("md5一致 ", Update.Name)
			continue
		}

		//删除文件
		derr := DelFile(dir + "/res/" + Update.Name)
		if nil != derr {
			fmt.Println("删除失败 ", Update.Name)
			panic(derr)
		}

		f, cerr := os.Create(dir + "/res/" + Update.Name)
		if cerr != nil {
			fmt.Println("创建新文件失败 ", Update.Name)
			panic(cerr)
		}
		defer f.Close()

		io.Copy(f, res.Body)

		fmt.Printf("开始更新 %s", Update.Name)
		LoadFile(dir + "/res/" + Update.Name)
		SetMd5(dir + "/res/"+Update.Name, Update.Md5)
	}

}
func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New("")
	}
	return string(path[0 : i+1]), nil
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DelFile(file string) error {
	derr := os.Remove(file) //删除文件test.txt
	if derr != nil {
		fmt.Printf("file remove Error:%s", derr)
		return derr
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
		return nil
	}
}
func LoadHttpConfig(dir string) *HttpConfig {
	b := filepath.Join(dir,"/res/http.json")
	bb, err := ioutil.ReadFile(b)
	if err != nil {
		panic(err)
	}

	var config *HttpConfig
	json.Unmarshal(bb, &config)

	return config

}

func GetFiles(dir string) ([]*Update, error) {
	updatelist := make([]*Update, 0)

	b, err := ioutil.ReadFile(dir + "res\\updatelist.json")
	if err != nil {
		return nil, err
	}

	json.Unmarshal(b, &updatelist)

	return updatelist, nil
}

//func GetMd5(path string) string {
//
//	fi, err := os.Open(path)
//	defer fi.Close()
//	if err != nil {
//		panic(err)
//	}
//	buff, _ := ioutil.ReadAll(fi)
//	return MD5(buff)
//}
//
//func MD5(b []byte) string {
//	vCrypto := md5.New()
//	vCrypto.Write(b)
//	return hex.EncodeToString(vCrypto.Sum(nil))
//}
func SetMd5(file string, md5 string) {
	name := filepath.Base(file)
	f, ok := func_map[name]
	if ok {
		f[1](md5)
	}
}

func GetMd5(file string) string {
	name := filepath.Base(file)
	f, ok := func_map[name]
	if ok {
		return f[2]("")
	}
	return ""
}
