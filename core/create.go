package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 创建vue Component组件
func CreateComponent(name, css, version string) string {
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
	CreateFolder(strReplaceAll, name, "vue", "components")
	return strReplaceAll
}

//  创建vuex store
func CreateStore(name string) {
	abspath, _ := filepath.Abs("./")
	filename := abspath + fmt.Sprintf("/core/template/%s.txt", "store")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	CreateFolder(string(file), name, "js", "store")
}

func ReplaceChar(orginStr string, fromChar string, toChar string) string {
	outStr := strings.ReplaceAll(orginStr, fromChar, toChar)
	return outStr
}

func CreateFolder(filedata, filename string, fileExtention string, targetSource string) {
	s, _ := filepath.Abs("./")
	srcFoldPath := s + "/src"
	targetSourceFold := s + "/src/" + targetSource
	fmt.Println(srcFoldPath)
	// 判断Src目录存在否
	if !Exists(srcFoldPath) {
		fmt.Println("请先在当前目录创建src目录")
		return
	}
	// 判断生成存放目录存在否
	if !Exists(targetSourceFold) {
		fmt.Printf("请先在src目录创建%s目录", targetSource)
		return
	}

	// 判断写入的文件存在否
	filePath := fmt.Sprintf("%s/%s.%s", srcFoldPath, filename, fileExtention)
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
	fmt.Printf("%s.%s创建成功", filename, fileExtention)

}

// 判断所给路径文件/文件夹是否存在1
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
