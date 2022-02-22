package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadAndReplaceFile(name, css, version string) string {
	abspath, _ := filepath.Abs("./")
	filename := abspath + fmt.Sprintf("/core/template/%s/%s.txt", version, "component")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	// fmt.Println("Contents of file:", string(file))
	// fmt.Println(name, css, version)

	strReplaceAll := ReplaceChar(string(file), "$name$", name)
	strReplaceAll = ReplaceChar(strReplaceAll, "$css$", strings.ToLower(css))
	// strReplaceAll = ReplaceChar(strReplaceAll, "$cname$", strings.ToLower(name))
	strReplaceAll = ReplaceChar(strReplaceAll, "$classname$", strings.ToLower(name))
	// fmt.Println(strReplaceAll)
	CreateFolder(strReplaceAll, name)
	return strReplaceAll
}

func ReplaceChar(orginStr string, fromChar string, toChar string) string {
	outStr := strings.ReplaceAll(orginStr, fromChar, toChar)
	return outStr
}

func CreateFolder(filedata, filename string) {
	s, _ := filepath.Abs("./")
	srcFoldPath := s + "/src"
	fmt.Println(srcFoldPath)
	if !Exists(srcFoldPath) {
		fmt.Println("请先在当前目录创建src目录")
		return
	}
	filePath := fmt.Sprintf("%s/%s.vue", srcFoldPath, filename)
	if Exists(filePath) {
		fmt.Printf("%s已存在", filePath)
		return
	}
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(filedata)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%s.vue创建成功", filename)

}

// 判断所给路径文件/文件夹是否存在1
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
