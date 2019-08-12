package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/labstack/gommon/color"
	"github.com/tealeg/xlsx"
)

var (
	buf     *bytes.Buffer
	buf_cst *bytes.Buffer
	buf_int *bytes.Buffer
	cst_map = make(map[string]bool)
)

func main() {
	color.Enable()

	dir := ``
	if len(os.Args) > 1 {
		a1 := strings.Replace(os.Args[1], ` `, ``, -1)
		if len(a1) > 0 {
			dir = a1
		}
	}
	if dir == `` {
		dir = `./`
	} else if dir[len(dir)-1] != '/' && dir[len(dir)-1] != '\\' {
		dir += `/`
	}

	dst := dir + `def`
	err := os.Mkdir(dst, os.ModePerm)
	if err == nil || os.IsExist(err) {
		dst += `/`
		tmp := make([]byte, 1024*64)
		buf_int = bytes.NewBuffer(tmp[0:0])
		cmm := dst + `common.go`
		genRegexpAndLoad1(buf_int)

		tmp = make([]byte, 64*1024*1024)
		buf = bytes.NewBuffer(tmp[0:0])
		tmp = make([]byte, 1024)
		buf_cst = bytes.NewBuffer(tmp[0:0])
		err = walk(dir, dst)
		if err == nil {
			err = genRegexpAndLoad2(buf_int, cmm)
		}
	}

	if err != nil {
		color.Println(color.Red(err.Error()))
	} else {
		color.Println(color.Green(`Successful completed !`))
	}
	sig_ch := make(chan os.Signal)
	signal.Notify(sig_ch, syscall.SIGINT|syscall.SIGTERM)
	<-sig_ch
}

func walk(src, dst string) error {
	fs, err := filepath.Glob(src + `[^{0123456789}]*` )
	if err != nil {
		return errors.New(err.Error() + ` ` + src)
	}
	var info os.FileInfo
	for _, fn := range fs {
		ext := filepath.Ext(fn)
		df := dst + fn
		if ext == `` && fn != `def` {
			info, err = os.Stat(df)
			if err == nil && info.IsDir() {
				err = os.Mkdir(df, os.ModePerm)
				if err == nil || os.IsExist(err) {
					err = walk(src+fn+`/`, df+`/`)
					if err != nil {
						return err
					}
				} else {
					return errors.New(err.Error() + ` ` + df)
				}
			}
			continue
		}
		if ext != `.xlsx` {
			continue
		}
		if fn[0:2] == `~$` {
			continue
		}
		df = strings.Replace(df, `.xlsx`, `.go`, 1)
		err = transfer(src+fn, df)
		if err != nil {
			return err
		}
	}
	return nil
}

func transfer(sf, df string) error {
	f, err := xlsx.OpenFile(sf)
	if err != nil {
		return errors.New(err.Error() + ` ` + sf)
	}
	if len(f.Sheets) == 0 || len(f.Sheets[0].Rows) < 4 {
		return errors.New(fmt.Sprintf(`sheet中没有内容，或内容格式不正确，文件：%s`, sf))
	}

	buf.Reset()
	buf.WriteString("package def\r\n\r\n")
	buf.WriteString("import (\r\n")
	buf.WriteString("\t\"fmt\"\r\n")
	buf.WriteString("\t\"os\"\r\n")
	buf.WriteString("\t\"strconv\"\r\n")
	buf.WriteString("\t\"strings\"\r\n")
	buf.WriteString("\t\"sync\"\r\n")
	buf.WriteString("\r\n\t\"github.com/tealeg/xlsx\"")
	buf.WriteString("\r\n\t\"reflect\"")
	buf.WriteString("\r\n\t\"os/exec\"")
	buf.WriteString("\r\n\t\"path/filepath\"")
	buf.WriteString("\r\n)\r\n\r\n")
	buf_cst.Reset()
	lth, val := len(f.Sheets), ``
	uniqs := make([]int32, lth)
	names := make([]string, lth)
	keys := make([][]string, lth)
	types := make([][]string, lth)
	shref := make([]int64, len(f.Sheets[0].Rows[0].Cells)) //sheet ref
	flt_hot := &bytes.Buffer{}
	for sdx, sheet := range f.Sheets {
		if sheet.Name == "注释页" {
			continue
		}
		if len(sheet.Rows) < 4 {
			return errors.New(fmt.Sprintf(`sheet中没有内容，或内容格式不正确，文件：%s 第 %d 张表`, sf, sdx+1))
		}
		names[sdx] = sheet.Name

		ok, err := regexp.Match(`[a-z][a-z0-9\-]*`, []byte(names[sdx]))
		if err != nil {
			return err
		}
		if !ok {
			return errors.New(fmt.Sprintf(
				`sheet名称必须是小写字母、数字与横杠的组合，并且开头必须是字母，参考格式：eg, eg-name, eg-123，在文件%s第%d张表`,
				sf, sdx+1,
			))
		}
		names[sdx] = camelStyleName(names[sdx])

		//添加version
		buf.WriteString("type ")
		buf.WriteString(names[sdx])
		buf.WriteString("Version struct {\r\n")
		buf.WriteString("\tmd5 string //配置表md5数据")
		buf.WriteString("\r\n}\r\n\r\n")

		buf.WriteString("type ")
		buf.WriteString(names[sdx])
		buf.WriteString("Class struct {\r\n\t")
		buf.WriteString(names[sdx])
		buf.WriteString("Version\r\n")
		//buf.WriteString("\tlId")
		//buf.WriteString(names[sdx])
		//buf.WriteString(" map[int32][]*")
		//buf.WriteString(names[sdx])
		buf.WriteString("}\r\n\r\n")

		//添加单例访问
		buf.WriteString("var ")
		buf.WriteString(names[sdx])
		buf.WriteString("instance *")
		buf.WriteString(names[sdx])
		buf.WriteString("Class\r\n\n")

		buf.WriteString("func ")
		buf.WriteString(names[sdx])
		buf.WriteString("Instance() *")
		buf.WriteString(names[sdx])
		buf.WriteString("Class {\r\n\t")
		buf.WriteString("if ")
		buf.WriteString(names[sdx])
		buf.WriteString("instance == nil {\r\n\t\t")
		buf.WriteString(names[sdx])
		buf.WriteString("instance = &")
		buf.WriteString(names[sdx])
		buf.WriteString("Class{}\r\n")
		buf.WriteString("\t}\r\n")
		buf.WriteString("\treturn ")
		buf.WriteString(names[sdx])
		buf.WriteString("instance")
		buf.WriteString("\r\n}\r\n\n")

		//buf.WriteString("var g")
		//buf.WriteString(names[sdx])
		//buf.WriteString("Class ")
		//buf.WriteString(names[sdx])
		//buf.WriteString("Class = ")
		//buf.WriteString(names[sdx])
		//buf.WriteString("Class {\r\n")
		//buf.WriteString("\tlId")
		//buf.WriteString(names[sdx])
		//buf.WriteString(": make(map[int32][]*")
		//buf.WriteString(names[sdx])
		//buf.WriteString("),\r\n")
		//buf.WriteString("}\r\n\r\n")

		buf.WriteString("type ")
		buf.WriteString(names[sdx])
		buf.WriteString(" struct {\r\n")
		keys[sdx] = make([]string, len(sheet.Rows[1].Cells))
		types[sdx] = make([]string, len(sheet.Rows[1].Cells))
		var flt_keys, flt_types []string
		c3lth := len(sheet.Rows[3].Cells)
		for idx, cell := range sheet.Rows[1].Cells {
			keys[sdx][idx] = strings.Trim(cell.String(), " _\t\r\n")
			if "" == keys[sdx][idx] {
				continue
			}
			ok, _ = regexp.Match(`[a-z][a-z0-9\-]*`, []byte(keys[sdx][idx]))
			if !ok {
				return errors.New(fmt.Sprintf(
					`字段名称必须是小写字母、数字与横杠的组合，并且开头必须是字母，参考格式：eg, eg-name, eg-123，在文件%s第%d张表第%d列`,
					sf, sdx+1, idx+1,
				))
			}
			buf.WriteString("\t")
			keys[sdx][idx] = camelStyleName(keys[sdx][idx])
			buf.WriteString(keys[sdx][idx])

			types[sdx][idx] = strings.Trim(sheet.Rows[2].Cells[idx].String(), " \t\r\n")

			uType := `unique`
			uTypes := strings.Split(types[sdx][idx], ",")
			if len(uTypes) == 2 {
				uType = uTypes[0]
			} else {
				uType = types[sdx][idx]
			}
			switch uType {
			case `unique`:
				if uniqs[sdx] > 0 {
					return errors.New(fmt.Sprintf(`unique字段只能出现一次，在文件%s第%d张表第%d列`, sf, sdx+1, idx+1))
				}
				uniqs[sdx] = int32(idx + 1)
				buf.WriteString(` int32`)
			case `const`:
				if _, ok := cst_map[keys[sdx][idx]]; !ok {
					buf_cst.WriteString("const (\r\n")
					cst_arr := bytes.Split([]byte(sheet.Rows[3].Cells[idx].String()), []byte{'\n'})
					ok := len(cst_arr) > 0
					for _, cst_bts := range cst_arr {
						cst_bts = bytes.Trim(cst_bts, " \t\r")
						if len(cst_bts) == 0 {
							continue
						}
						if bytes.HasPrefix(cst_bts, []byte(`filter`)) {
							flt_keys = append(flt_keys, keys[sdx][idx])
							flt_types = append(flt_types, `int32`)
							continue
						}
						ok, _ = regexp.Match(`[a-zA-Z]\w*\s*=\s*(?:\d+|\d+<<\d+|\d+(?:[\|\+\-\^]\d+)+)(?:\s*\/\/.+)?`, cst_bts)
						if !ok {
							break
						}
						buf_cst.WriteByte('\t')
						buf_cst.WriteString(keys[sdx][idx])
						buf_cst.WriteString(`_`)
						//buf_cst.Write(cst_bts)
						buf_cst.Write(camelStyleConst(cst_bts))
						buf_cst.WriteString("\r\n")
					}
					if !ok {
						return errors.New(fmt.Sprintf(`const字段需要预定义，请按格式写在第四行中，在文件%s第%d张表第%d列`, sf, sdx+1, idx+1))
					}
					buf_cst.WriteString(")\r\n\r\n")
					cst_map[keys[sdx][idx]] = true
				}
				buf.WriteString(` int32`)
			case `int`:
				buf.WriteString(` int32`)
				if idx < c3lth {
					val = strings.Trim(sheet.Rows[3].Cells[idx].String(), " \t\r\n")
					if val == `filter` {
						flt_keys = append(flt_keys, keys[sdx][idx])
						flt_types = append(flt_types, `int32`)
					}
				}
			case `float`:
				buf.WriteString(` float32`)
				if idx < c3lth {
					val = strings.Trim(sheet.Rows[3].Cells[idx].String(), " \t\r\n")
					if val == `filter` {
						flt_keys = append(flt_keys, keys[sdx][idx])
						flt_types = append(flt_types, `float32`)
					}
				}
			case `string`:
				buf.WriteString(` string`)
				if idx < c3lth {
					val = strings.Trim(sheet.Rows[3].Cells[idx].String(), " \t\r\n")
					if val == `filter` {
						flt_keys = append(flt_keys, keys[sdx][idx])
						flt_types = append(flt_types, `string`)
					}
				}
			case `json`:
				buf.WriteString(` ` + uTypes[1])
			case `bool`:
				buf.WriteString(` bool`)
			case `sheet`:
				if sdx > 0 {
					return sheetFieldError(sf, int32(sdx), int32(idx), keys[sdx][idx])
				}
				buf.WriteString(` *`)
				buf.WriteString(keys[sdx][idx])
				shref[idx], err = strconv.ParseInt(sheet.Rows[3].Cells[idx].String(), 10, 64)
				if err != nil || shref[idx] < 2 || int32(shref[idx]) > int32(lth) {
					return errors.New(fmt.Sprintf(`sheet字段需要在第四行指明索引的表序号，在文件%s第%d列`, sf, idx+1))
				}
			case `referer`:
				if sdx > 0 {
					return sheetFieldError(sf, int32(sdx), int32(idx), keys[sdx][idx])
				}
				buf.WriteString(` *`)
				buf.WriteString(keys[sdx][idx])
			default:
				return errors.New(fmt.Sprintf(
					`类型必须是：bool、int、const、unique、float、string、json、sheet、referer其中之一，在文件%s第%d张表第%d行，类型为："%s"`,
					sf, sdx+1, idx+1, uType,
				))
			}

			cmm := strings.Trim(sheet.Rows[0].Cells[idx].String(), " \t\r\n")
			if cmm == `` {
				cmm = strings.Trim(sheet.Rows[3].Cells[idx].String(), " \t\r\n")
			}
			buf.WriteString(` //`)
			buf.WriteString(cmm)
			buf.WriteString("\r\n")
		}
		buf.WriteString("}\r\n\r\n")

		//如果是json，把第四行添加进来
		for idx, _ := range sheet.Rows[1].Cells {
			types[sdx][idx] = strings.Trim(sheet.Rows[2].Cells[idx].String(), " \t\r\n")
			uTypes := strings.Split(types[sdx][idx], ",")
			if uTypes[0] == `json` {
				buf.WriteString(sheet.Rows[3].Cells[idx].String())
				buf.WriteString("\r\n\r\n")
			}
		}

		for fdx, key := range flt_keys {
			filterData(buf, names[sdx], key, flt_types[fdx])
		}
		if len(flt_keys) > 0 {
			initFilter1(flt_hot, names[sdx], flt_keys)
		}
	}
	if flt_hot.Len() > 0 {
		initFilter2(names[0], buf, flt_hot)
	}
	if uniqs[0] > 0 {
		primaryReferer(buf, names[0])
	}
	if buf_cst.Len() > 0 {
		buf.Write(buf_cst.Bytes())
	}
	//versioin get set
	//buf.WriteString("func Set")
	//buf.WriteString(names[0])
	//buf.WriteString("Version (md5 string) string {\r\n")
	//buf.WriteString("/tg")
	//buf.WriteString(names[0])
	//buf.WriteString("Class.md5 = md5\r\n")
	//buf.WriteString("return ``")
	//buf.WriteString("\r\n}\r\n")
	//
	//buf.WriteString("func Get")
	//buf.WriteString(names[0])
	//buf.WriteString("Version (md5 string) string {\r\n")
	//buf.WriteString("/treturn g")
	//buf.WriteString(names[0])
	//buf.WriteString("Class.md5\r\n")
	//buf.WriteString("}\r\n")

	buf.WriteString("func Set")
	buf.WriteString(names[0])
	buf.WriteString("Version(md5 string) string {\r\n\t")
	buf.WriteString(names[0])
	buf.WriteString("Instance().md5 = md5\r\n\t")
	buf.WriteString("return ``\r\n\t}\r\n")

	buf.WriteString("func Get")
	buf.WriteString(names[0])
	buf.WriteString("Version(md5 string) string {\r\n\t")
	buf.WriteString("return ")
	buf.WriteString(names[0])
	buf.WriteString("Instance().md5\r\n")
	buf.WriteString("}\r\n")

	buf_int.WriteString("\t\t`")
	buf_int.WriteString(filepath.Base(sf))
	buf_int.WriteString("`: { Load")
	buf_int.WriteString(names[0])
	buf_int.WriteString(", Set")
	buf_int.WriteString(names[0])
	buf_int.WriteString("Version")
	buf_int.WriteString(", Get")
	buf_int.WriteString(names[0])
	buf_int.WriteString("Version }")
	buf_int.WriteString(",\r\n")

	buf.WriteString("func Get")
	buf.WriteString(names[0])
	buf.WriteString("()")

	table := f.Sheets[0]
	if uniqs[0] > 0 {
		buf.WriteString(" map[int32]*")
	} else {
		buf.WriteString(" []*")
	}
	buf.WriteString(names[0])
	buf.WriteString("{\r\n")
	buf.WriteString("\tmtx")
	buf.WriteString(names[0])
	buf.WriteString(".RLock()\r\n")
	buf.WriteString("\tcnf := cnf")
	buf.WriteString(names[0])
	buf.WriteString("\r\n\tmtx")
	buf.WriteString(names[0])
	buf.WriteString(".RUnlock()\r\n")
	buf.WriteString("\treturn cnf\r\n}\r\n\r\n")

	for sdx, sheet := range f.Sheets {
		if sheet.Name == "注释页" {
			continue
		}
		for idx, _ := range sheet.Rows[1].Cells {
			if "" == keys[sdx][idx] {
				continue
			}
			if idx > 0 {
				buf.WriteString("func (this *")
				buf.WriteString(names[0])
				buf.WriteString(") get")
				buf.WriteString(keys[sdx][idx])
				buf.WriteString("() ")
				strs := strings.Split(types[0][idx], ",")
				if `int` == types[0][idx] || `const` == types[0][idx] || `unique` == types[0][idx] {
					buf.WriteString("int32")
				} else if `float` == types[0][idx] {
					buf.WriteString("float32")
				} else if len(strs) == 2 && `json` == strs[0] {
					buf.WriteString(strs[1])
				} else {
					buf.WriteString(types[0][idx])
				}
				buf.WriteString(" {\r\n")
				buf.WriteString("\treturn this.")
				buf.WriteString(keys[sdx][idx])
				buf.WriteString(" \r\n}\r\n\r\n")
			}
		}
	}

	buf.WriteString("func Load")
	buf.WriteString(names[0])
	buf.WriteString("(file string) string {\r\n")
	buf.WriteString("\tvar clen = []int32{")
	for idx, k2s := range keys {
		if idx > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(strconv.FormatInt(int64(len(k2s)), 10))
	}
	buf.WriteString("}\r\n")
	buf.WriteString("\tsf := `")
	buf.WriteString(filepath.Base(sf))
	buf.WriteString("`")
	buf.WriteString(`
	fi, _ := exec.LookPath(os.Args[0])
	pa, _ := filepath.Abs(fi)
	rst := filepath.Dir(pa)
`)

	buf.WriteString("\t" + `sf = rst + "/" +"res"+"/" + filepath.Base(sf)`)
	//	buf.WriteString("`")
	buf.WriteString(`
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
    var shref = []int64{`)
	for idx, val := range shref {
		if idx > 0 {
			buf.WriteString(", ")
		}
		if val > 0 {
			val -= 1
		}
		buf.WriteString(strconv.FormatInt(val, 10))
	}
	buf.WriteString("}\r\n\t_ = shref\r\n")

	if uniqs[0] > 0 {
		buf.WriteString("\tcnf := make(map[int32]*")
		buf.WriteString(names[0])
		buf.WriteString(")\r\n")
	} else {
		buf.WriteString("var cnf []*")
		buf.WriteString(names[0])
		buf.WriteString("\r\n")
	}
	buf.WriteString("\tfor rdx, row := range f.Sheets[0].Rows[4:] {\r\n")
	buf.WriteString("\t\tif int32(len(row.Cells)) == 0 || strings.Replace(row.Cells[0].String(), \" \\t\\r\\n\", ``, -1) == `` {\r\n")
	buf.WriteString("\t\t\tcontinue\r\n")
	buf.WriteString("\t\t}\r\n")
	buf.WriteString("\t\tif int32(len(row.Cells)) != clen[0] {\r\n")
	buf.WriteString("\t\t\treturn fmt.Sprintf(`列数错误，在文件%s的第%d行，期望列数：\"%d\"，实际列数：\"%d\"`, sf, rdx+5, clen[0], len(row.Cells))\r\n")
	buf.WriteString("\t\t}\r\n")
	buf.WriteString("\t\titm := &")
	buf.WriteString(names[0])
	buf.WriteString("{}\r\n")
	for idx, field := range keys[0] {
		col := strconv.FormatInt(int64(idx), 10)
		buf.WriteString("\t\tval = strings.Replace(row.Cells[")
		buf.WriteString(col)
		buf.WriteString("].String(), \" \\t\\r\\n\", ``, -1)\r\n")
		switch types[0][idx] {
		case `bool`:
			buf.WriteString("\t\titm.")
			buf.WriteString(field)
			buf.WriteString(", err = strconv.ParseBool(val)\r\n")
			buf.WriteString("\t\tif err != nil {\r\n")
			buf.WriteString("\t\t\treturn fmt.Sprintf(`bool解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, ")
			buf.WriteString(col)
			buf.WriteString(", val)\r\n")
			buf.WriteString("\t\t}\r\n")
		case `unique`:
			fallthrough
		case `const`:
			fallthrough
		case `int`:
			buf.WriteString("\t\titm.")
			buf.WriteString(field)
			buf.WriteString(", err = getInt(val)\r\n")
			buf.WriteString("\t\tif err != nil {\r\n")
			buf.WriteString("\t\t\treturn fmt.Sprintf(`int解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, ")
			buf.WriteString(col)
			buf.WriteString(", val)\r\n")
			buf.WriteString("\t\t}\r\n")
		case `float`:
			buf.WriteString("\t\tf64, err = strconv.ParseFloat(val, 32)\r\n")
			buf.WriteString("\t\tif err != nil {\r\n")
			buf.WriteString("\t\t\treturn fmt.Sprintf(`float解析失败，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, ")
			buf.WriteString(col)
			buf.WriteString(", val)\r\n")
			buf.WriteString("\t\t}\r\n")
			buf.WriteString("\t\titm.")
			buf.WriteString(field)
			buf.WriteString(" = float32(f64)\r\n")
		case `string`:
			buf.WriteString("\t\titm.")
			buf.WriteString(field)
			buf.WriteString(" = val\r\n")
		case `referer`:
			buf.WriteString("\t\t\tr64, err = strconv.Atoi(val)\r\n")
			buf.WriteString("\t\t\tif err != nil {\r\n")
			buf.WriteString("\t\t\t\treturn fmt.Sprintf(`无法解析referer引用，在文件%s的第%d行第%d列，值为：\"%s\"`, sf, rdx+5, ")
			buf.WriteString(col)
			buf.WriteString("+1, val)\r\n\t\t\t}\r\n")
			buf.WriteString("\t\titm.")
			buf.WriteString(field)
			buf.WriteString(", ok = cnf")
			buf.WriteString(field)
			buf.WriteString("(r64)\r\n")
			buf.WriteString("\t\tif !ok {\r\n")
			buf.WriteString("\t\t\treturn fmt.Sprintf(`referer引用了不存在的主键，在文件%s的第%d行第%d列，值为：%s`, sf, rdx+5, ")
			buf.WriteString(col)
			buf.WriteString(", val)\r\n")
			buf.WriteString("\t\t}\r\n")
		case `sheet`:
			buf.WriteString("\t\tif val == `-` {\r\n")
			buf.WriteString("\t\t\titm.")
			buf.WriteString(field)
			buf.WriteString(" = nil\r\n\t\t} else {\r\n")
			buf.WriteString("\t\t\titm2 := &")
			buf.WriteString(field)
			buf.WriteString("{}\r\n")

			buf.WriteString("\t\t\tr64, err = strconv.Atoi(val)\r\n")
			buf.WriteString("\t\t\tif err != nil || r64 < 5 || r64 > rlen[shref[")
			buf.WriteString(col)
			buf.WriteString("]] {\r\n")
			buf.WriteString("\t\t\t\treturn fmt.Sprintf(`无法解析sheet引用或引用错误，在文件%s的第%d行第%d列，值为：\"%s\"`, sf, rdx+5, ")
			buf.WriteString(col)
			buf.WriteString("+1, val)\r\n\t\t\t}\r\n")
			buf.WriteString("\t\t\tr64 -= 1\r\n")

			buf.WriteString("\t\t\tif len(f.Sheets[shref[")
			buf.WriteString(col)
			buf.WriteString("]].Rows[r64].Cells) < clen[shref[")
			buf.WriteString(col)
			buf.WriteString("]] {\r\n")
			buf.WriteString("\t\t\t\treturn fmt.Sprintf(`应用表列数不匹配，在文件%s第%d张表第%d行`, sf, ")
			sh2 := strconv.FormatInt(shref[idx], 10)
			buf.WriteString(sh2)
			buf.WriteString(", r64+1)\r\n")
			buf.WriteString("\t\t\t}\r\n")
			for cdx, fname := range keys[shref[idx]-1] {
				col2 := strconv.FormatInt(int64(cdx), 10)
				buf.WriteString("\t\t\tval = strings.Replace(f.Sheets[shref[")
				buf.WriteString(col)
				buf.WriteString("]].Rows[r64].Cells[")
				buf.WriteString(col2)
				buf.WriteString("].String(), \" \\t\\r\\n\", ``, -1)\r\n")
				switch types[shref[idx]-1][cdx] {
				case `bool`:
					buf.WriteString("\t\t\titm2.")
					buf.WriteString(fname)
					buf.WriteString(", err = strconv.ParseBool(val)\r\n")
					buf.WriteString("\t\t\tif err != nil {\r\n")
					buf.WriteString("\t\t\t\treturn fmt.Sprintf(`bool解析失败，在文件%s的第%d张表第%d行第%d列，值为：%s`, sf, rdx+5, ")
					buf.WriteString(sh2)
					buf.WriteString(", r64, ")
					buf.WriteString(col2)
					buf.WriteString(", val)\r\n")
					buf.WriteString("\t\t\t}\r\n")
				case `int`:
					buf.WriteString("\t\t\titm2.")
					buf.WriteString(fname)
					buf.WriteString(", err = getInt(val)\r\n")
					buf.WriteString("\t\t\tif err != nil {\r\n")
					buf.WriteString("\t\t\t\treturn fmt.Sprintf(`int解析失败，在文件%s的第%d张表第%d行第%d列，值为：%s`, sf, rdx+5, ")
					buf.WriteString(sh2)
					buf.WriteString(", r64, ")
					buf.WriteString(col2)
					buf.WriteString(", val)\r\n")
					buf.WriteString("\t\t\t}\r\n")
				case `float`:
					buf.WriteString("\t\t\tf64, err = strconv.ParseFloat(val, 32)\r\n")
					buf.WriteString("\t\t\tif err != nil {\r\n")
					buf.WriteString("\t\t\t\treturn fmt.Sprintf(`float解析失败，在文件%s的第%d张表第%d行第%d列，值为：%s`, sf, rdx+5, ")
					buf.WriteString(sh2)
					buf.WriteString(", r64, ")
					buf.WriteString(col2)
					buf.WriteString(", val)\r\n")
					buf.WriteString("\t\t\t}\r\n")
					buf.WriteString("\t\t\titm2.")
					buf.WriteString(fname)
					buf.WriteString(" = float32(f64)\r\n")
				case `string`:
					buf.WriteString("\t\t\titm2.")
					buf.WriteString(fname)
					buf.WriteString(" = val\r\n")
				}
			}

			buf.WriteString("\t\t\titm.")
			buf.WriteString(field)
			buf.WriteString(" = itm2\r\n")
			buf.WriteString("\t\t}\r\n")
		}
	}
	if uniqs[0] > 0 {
		buf.WriteString("\t\tcnf[itm.")
		buf.WriteString(keys[0][uniqs[0]-1])
		buf.WriteString("] = itm\r\n")
	} else {
		buf.WriteString("\t\tcnf = append(cnf, itm)\r\n")
	}
	buf.WriteString("\t}\r\n")
	buf.WriteString("\tmtx")
	buf.WriteString(names[0])
	buf.WriteString(".Lock()\r\n")
	buf.WriteString("\tcnf")
	buf.WriteString(names[0])
	buf.WriteString(" = cnf\r\n")
	if flt_hot.Len() > 0 {
		buf.Write(flt_hot.Bytes())
	}
	buf.WriteString("\tmtx")
	buf.WriteString(names[0])
	buf.WriteString(".Unlock()\r\n")
	buf.WriteString("\treturn ``\r\n}\r\n\r\n")

	buf.WriteString("var _ = strconv.ErrRange\r\n")
	buf.WriteString("var mtx")
	buf.WriteString(names[0])
	buf.WriteString(" = new(sync.RWMutex)\r\n")
	buf.WriteString("var cnf")
	buf.WriteString(names[0])
	buf.WriteString(" = ")
	if uniqs[0] > 0 {
		buf.WriteString("map[int32]*")
	} else {
		buf.WriteString("[]*")
	}
	buf.WriteString(names[0])
	buf.WriteString("{\r\n")

	umap := make(map[int32]map[string]bool)
	var i64 int64
	var f64 float64
	var bl8 bool
	for rdx, row := range table.Rows[4:] {
		if len(row.Cells) == 0 || strings.Replace(row.Cells[0].String(), " \t\r\n", ``, -1) == `` {
			continue
		}
		if uniqs[0] == 0 {
			buf.WriteString("\t&")
			buf.WriteString(names[0])
			buf.WriteString("{\r\n\t\t")
		} else {
			val = strings.Replace(row.Cells[uniqs[0]-1].String(), " \t\r\n", ``, -1)
			buf.WriteString("\t")
			buf.WriteString(val)
			buf.WriteString(": &")
			buf.WriteString(names[0])
			buf.WriteString("{\r\n\t\t")
		}
		for idx, cell := range row.Cells {
			if "" == keys[0][idx] {
				continue
			}
			val = valInit(strings.Replace(cell.String(), " \t\r\n", ``, -1), types[0][idx])
			uTypes := strings.Split(types[0][idx], ",")
			switch uTypes[0] {
			case `bool`:
				bl8, err = strconv.ParseBool(val)
				if err != nil {
					return errors.New(fmt.Sprintf(`bool类型解析错误，在文件%s的第%d行第%d列，值为："%s"`, sf, rdx+5, idx+1, val))
				}
				buf.WriteString(strconv.FormatBool(bl8))
				buf.WriteString(`, `)
			case `unique`:
				fmap, ok := umap[int32(idx)]
				if !ok {
					fmap = make(map[string]bool)
					umap[int32(idx)] = fmap
				}
				_, ok = fmap[val]
				if ok {
					return errors.New(fmt.Sprintf(`唯一类型字段有重复值，在文件%s的第%d行第%d列，值为："%s"`, sf, rdx+5, idx+1, val))
				}
				i64, err = strconv.ParseInt(val, 10, 64)
				if err != nil || i64 < 0 {
					return errors.New(fmt.Sprintf(`unique类型的值必须为大于 等于 0的整数，在文件%s的第%d行第%d列，值为："%s"`, sf, rdx+5, idx+1, val))
				}
				val = strconv.FormatInt(i64, 10)
				fmap[val] = true
				buf.WriteString(val)
				buf.WriteString(`, `)
			case `const`:
				fallthrough
			case `int`:
				i64, err = strconv.ParseInt(val, 10, 64)
				if err != nil {
					ok, _ := regexp.Match(`^\d(?:[\|\+\-\*\/\^\&]\d+)*$`, []byte(val))
					if ok {
						buf.WriteString(val)
					} else if val == "" {
						buf.WriteString(`0`)
					} else {
						return errors.New(fmt.Sprintf(`int、const类型字段的值必须为一个整数，在文件%s的第%d行第%d列，值为："%s"`, sf, rdx+5, idx+1, val))
					}
				} else {
					buf.WriteString(strconv.FormatInt(i64, 10))
				}
				buf.WriteString(`, `)
			case `float`:
				f64, err = strconv.ParseFloat(val, 64)
				if err != nil {
					return errors.New(fmt.Sprintf(`float字段类型的值必须为一个浮点数，在文件%s的第%d行第%d列，值为："%s"`, sf, rdx+5, idx+1, val))
				}
				buf.WriteString(strconv.FormatFloat(f64, 'f', -1, 64))
				buf.WriteString(`, `)
			case `string`:
				buf.WriteString(`"`)
				buf.WriteString(val)
				buf.WriteString(`", `)
			case `json`:
				//todo 写每行的值

				buf.WriteString(uTypes[1])
				buf.WriteString("{source:")
				buf.WriteString("`")
				if val != `` {
					buf.WriteString(val)
				} else {
					buf.WriteString("")
				}
				buf.WriteString("`}")
				buf.WriteString(", ")

			case `referer`:
				i64, err = strconv.ParseInt(val, 10, 64)
				if err != nil {
					return errors.New(fmt.Sprintf(`referer字段类型的值必须为一个整数，在文件%s的第%d行第%d列，值为："%s"`, sf, rdx+5, idx+1, val))
				}
				buf.WriteString("cnf")
				buf.WriteString(keys[0][idx])
				buf.WriteByte('[')
				buf.WriteString(strconv.FormatInt(i64, 10))
				buf.WriteString("],")
			case `sheet`:
				if val == `-` {
					buf.WriteString(`nil, `)
					break
				}
				r64, err := strconv.Atoi(val)
				if err != nil || r64 > len(f.Sheets[shref[idx]-1].Rows) {
					return errors.New(fmt.Sprintf(`无法解析sheet引用或引用错误，在文件%s的第%d行第%d列，值为："%s"`, sf, rdx+5, idx+1, val))
				}
				tmp_row := f.Sheets[shref[idx]-1].Rows[r64-1]
				buf.WriteString("\r\n\t\t&")
				buf.WriteString(keys[0][idx])
				buf.WriteString(`{`)
				for ydx, ycell := range tmp_row.Cells {
					if ydx > 0 {
						buf.WriteString(`, `)
					}
					val = valInit(strings.Replace(ycell.String(), " \t\r\n", ``, -1), types[shref[idx]-1][ydx])
					switch types[shref[idx]-1][ydx] {
					case `bool`:
						bl8, err = strconv.ParseBool(val)
						if err != nil {
							return errors.New(fmt.Sprintf(
								`bool类型解析错误，在文件%s的第%d张表第%d行第%d列，值为："%s"`,
								sf, shref[idx], r64, ydx+1, val,
							))
						}
						buf.WriteString(strconv.FormatBool(bl8))
					case `int`:
						i64, err = strconv.ParseInt(val, 10, 64)
						if err != nil {
							ok, _ := regexp.Match(`^\d(?:[\|\+\-\*\/\^\&]\d+)*$`, []byte(val))
							if ok {
								buf.WriteString(val)
							} else {
								return errors.New(fmt.Sprintf(
									`int类型必须为整数，在文件%s的第%d张表第%d行第%d列，值为："%s"`,
									sf, shref[idx], r64, ydx+1, val,
								))
							}
						} else {
							buf.WriteString(strconv.FormatInt(i64, 10))
						}
					case `float`:
						f64, err = strconv.ParseFloat(val, 64)
						if err != nil {
							return errors.New(fmt.Sprintf(
								`float类型解析错误，在文件%s的第%d张表第%d行第%d列，值为："%s"`,
								sf, shref[idx], r64, ydx+1, val,
							))
						}
						buf.WriteString(strconv.FormatFloat(f64, 'f', -1, 64))
					case `string`:
						buf.WriteString(`"`)
						buf.WriteString(val)
						buf.WriteString(`"`)
					default:
						return errors.New(fmt.Sprintf(
							`被引用的sheet表中只能使用bool、int、const、json、float、string类型，在文件%s的第%d张表第%d行第%d列，值为："%s"`,
							sf, shref[idx], r64, ydx+1, val,
						))
					}
				}
				buf.WriteString("},")
				if idx+1 < len(row.Cells) {
					buf.WriteString("\r\n\t\t")
				}
			}
		}
		buf.WriteString("\r\n\t},\r\n")
	}

	buf.WriteString("}\r\n")
	//添加json解析
	buf.WriteString(`
	func `)
	buf.WriteString(names[0])
	buf.WriteString(`_hot() {`)
	for idx, _ := range keys[0] {
		strs := strings.Split(types[0][idx], ",")
		if len(strs) == 2 && strs[0] == `json` {
			buf.WriteString("\r\n ")
			buf.WriteString(strings.ToLower(strs[1]))
			buf.WriteString(` := make([]*`)
			strs[1] = strs[1][:len(strs[1])-6]
			buf.WriteString(strs[1])
			buf.WriteString(`, 0)`)
		}
	}
	buf.WriteString(`
		for _, val := range cnf`)
	buf.WriteString(names[0])
	buf.WriteString(`	{
			v := reflect.ValueOf(*val)
			t := reflect.TypeOf(*val)
			fmt.Printf("%s",t.Name())
			for k := 0; k < v.NumField(); k++ {
				if strings.Contains(v.Field(k).Type().String(), "_Json") {`)
	for idx, _ := range keys[0] {
		strs := strings.Split(types[0][idx], ",")
		if len(strs) == 2 && strs[0] == `json` {
			buf.WriteString("\r\n")
			buf.WriteString(`if t.Field(k).Name == `)
			buf.WriteString(`"` + keys[0][idx] + `"`)
			buf.WriteString(`	 {
			if err := json.Unmarshal([]byte(val.`)
			buf.WriteString(keys[0][idx])
			buf.WriteString(`.source), &`)
			buf.WriteString(strings.ToLower(strs[1]))
			buf.WriteString(`); err == nil {
						val.`)
			buf.WriteString(keys[0][idx])
			buf.WriteString(`.item = `)
			buf.WriteString(strings.ToLower(strs[1]))
			buf.WriteString(`
			continue
					}
			`)
			buf.WriteString(`}`)
		}

	}
	buf.WriteString(`}
			}
		}
	}`)
	err = ioutil.WriteFile(df, buf.Bytes(), os.ModePerm)
	if err != nil {
		return errors.New(err.Error() + ` ` + sf)
	}
	return nil
}

func valInit(v, t string) string {
	if "" != v {
		return v
	}
	switch t {
	case `bool`:
		return "f"
	case `int`:
		return "0"
	case `float`:
		return "0.0"
	default:
		break
	}
	return v
}

func sheetFieldError(sf string, sdx, idx int32, val string) error {
	return errors.New(fmt.Sprintf("被引用的sheet表中出现不支持的字段类型，在文件%s第%d张表第%d列，值为：\"%s\"", sf, sdx+1, idx+1, val))
}

func camelStyleName(name string) string {
	fns := bytes.Split([]byte(name), []byte{'_'})
	for _, tmp := range fns {
		if len(tmp) > 0 && tmp[0] > 96 && tmp[0] < 123 {
			tmp[0] -= 32
		}
	}
	return string(bytes.Join(fns, []byte{}))
}
func camelStyleConst(data []byte) []byte {
	if len(data) < 1 {
		return data
	}
	pos := bytes.Index(data, []byte{' '})
	tmp := data
	if pos > 0 {
		tmp = data[0:pos]
	}
	fns := bytes.Split(tmp, []byte{'_'})
	rst := make([]byte, len(data))
	rst = rst[0:0]
	for _, tmp := range fns {
		if len(tmp) > 0 && tmp[0] > 96 && tmp[0] < 123 {
			tmp[0] -= 32
			rst = append(rst, tmp...)
		}
	}
	if pos > 0 {
		rst = append(rst, data[pos:]...)
	}
	return rst
}

func genRegexpAndLoad1(buf *bytes.Buffer) {
	buf.WriteString(`package def

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

var (`)
	buf.WriteString("\r\n\tre_num, _ = regexp.Compile(`[\\|\\+\\-\\*\\/\\^\\&]`)")
	buf.WriteString(`
    errInvalidInt = errors.New("int格式不正确！")
	func_map map[string][]func(string) string
)`)

	buf.WriteString("\ntype Update struct {\n")
	buf.WriteString("\tName string `json:")
	buf.WriteString(`"name"`)
	buf.WriteString("`\n")
	buf.WriteString("\tPath string `json:")
	buf.WriteString(`"path"`)
	buf.WriteString("`\n")
	buf.WriteString("\tFull string `json:")
	buf.WriteString(`"Full"`)
	buf.WriteString("`\n")
	buf.WriteString("\tMd5 string `json:")
	buf.WriteString(`"md5"`)
	buf.WriteString("`\n}\n")

	buf.WriteString("type HttpConfig struct {\n")
	buf.WriteString("\tHttp string `json:")
	buf.WriteString(`"http"`)
	buf.WriteString("`\n}")

	buf.WriteString(`
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
}`)
	buf.WriteString("\r\n\r\n")
	buf.WriteString("func init() {\r\n")
	buf.WriteString("\tfunc_map = map[string][]func(string) string{\r\n")
}

func genRegexpAndLoad2(buf *bytes.Buffer, df string) error {
	buf.WriteString("\t}\r\n}\r\n")
	buf.WriteString(`
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
`)
	return ioutil.WriteFile(df, buf.Bytes(), os.ModePerm)
}

func filterData(buf *bytes.Buffer, snm, cnm, ctp string) {
	vname := cnm + snm

	buf.WriteString("var f")
	buf.WriteString(vname)
	buf.WriteString(" = make(map[")
	buf.WriteString(ctp)
	buf.WriteString("][]*")
	buf.WriteString(snm)
	buf.WriteString(")\r\n\r\n")

	buf.WriteString("func pre")
	buf.WriteString(vname)
	buf.WriteString("() {\r\n")
	buf.WriteString("\tfor _, val := range cnf")
	buf.WriteString(snm)
	buf.WriteString("\t{\r\n")
	buf.WriteString("\t\tarr, ok := f")
	buf.WriteString(vname)
	buf.WriteString("[val.")
	buf.WriteString(cnm)
	buf.WriteString("]\r\n")
	buf.WriteString("\t\tif !ok {\r\n")
	buf.WriteString("\t\t\tarr = make([]*")
	buf.WriteString(snm)
	buf.WriteString(", 8)\r\n")
	buf.WriteString("\t\t\tarr = arr[0:0]\r\n")
	buf.WriteString("\t\t\tf")
	buf.WriteString(vname)
	buf.WriteString("[val.")
	buf.WriteString(cnm)
	buf.WriteString("] = arr\r\n")
	buf.WriteString("\t\t}\r\n")
	buf.WriteString("\t\tarr = append(arr, val)\r\n")
	buf.WriteString("\t}\r\n")
	buf.WriteString("}\r\n")

	buf.WriteString("func Filter")
	buf.WriteString(snm)
	buf.WriteString(`By`)
	buf.WriteString(cnm)
	buf.WriteString(`(id `)
	buf.WriteString(ctp)
	buf.WriteString(`) []*`)
	buf.WriteString(snm)
	buf.WriteString(" {\r\n")
	buf.WriteString("\tmtx")
	buf.WriteString(snm)
	buf.WriteString(".RLock()\r\n")
	buf.WriteString("\tarr, _ := f")
	buf.WriteString(vname)
	buf.WriteString("[id]\r\n")
	buf.WriteString("\tmtx")
	buf.WriteString(snm)
	buf.WriteString(".RUnlock()\r\n")
	buf.WriteString("\treturn arr\r\n")
	buf.WriteString("}\r\n\r\n")
}
func initFilter1(buf *bytes.Buffer, snm string, cnm []string) {
	for _, val := range cnm {
		buf.WriteString("\tpre")
		buf.WriteString(val)
		buf.WriteString(snm)
		buf.WriteString("()\r\n")
	}
}
func initFilter2(name string, buf, hot *bytes.Buffer) {
	buf.WriteString("func init() {\r\n")
	buf.Write(hot.Bytes())
	buf.WriteString(name)
	buf.WriteString("_hot()\n")
	buf.WriteString("}\r\n\r\n")
}
func primaryReferer(buf *bytes.Buffer, snm string) {
	buf.WriteString("func Get")
	buf.WriteString(snm)
	buf.WriteString("ByPk(id int32) (itm *")
	buf.WriteString(snm)
	buf.WriteString(", ok bool) {\r\n")
	buf.WriteString("\tmtx")
	buf.WriteString(snm)
	buf.WriteString(".RLock()\r\n")
	buf.WriteString("\titm, ok = cnf")
	buf.WriteString(snm)
	buf.WriteString("[id]\r\n")
	buf.WriteString("\tmtx")
	buf.WriteString(snm)
	buf.WriteString(".RUnlock()\r\n")
	buf.WriteString("\treturn\r\n")
	buf.WriteString("}\r\n\r\n")
}
