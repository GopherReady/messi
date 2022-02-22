/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/GopherReady/messi/core"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "messi add -c <name> 生成文件模板",
	Long: ` 快速创建Vue文件 For example:
	messi add -c  <name>  生成name.vue
	messi add -s  <name>  生成name.js
	messi add -r  <name>  生成name.js	
`,
	Run: func(cmd *cobra.Command, args []string) {
		vueVersion = "v" + vueVersion
		fmt.Println(css)
		if _, ok := cssFileType[css]; !ok {
			fmt.Println("请输出正确的css类型")
			return
		}
		result := fmt.Sprintf("正在添加组件名:%s.vue, CSS文件类型:%s Vue版本:Vue%s 编辑语言:%s", componentsName, css, vueVersion, jsType)
		fmt.Println(result)
		core.ReadAndReplaceFile(componentsName, css, vueVersion)
	},
}

//  css文件类型
var cssFileType = map[string]bool{
	"ccs":  true,
	"scss": true,
	"sass": true,
	"less": true,
}

var (
	componentsName string
	css            string
	vueVersion     string
	jsType         string
)

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&componentsName, "components", "v", "defaultName", "生成components.vue file")
	addCmd.Flags().StringVarP(&css, "css", "c", "scss", "vue css类型")
	addCmd.Flags().StringVarP(&vueVersion, "vue", "t", "3", "vue版本类型")
	addCmd.Flags().StringVarP(&jsType, "jsType", "j", "ts", "vue版本类型")
}
