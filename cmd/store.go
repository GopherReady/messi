/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/GopherReady/messi/core"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "快速生成vue store.ts文件",
	Long:  `快速生成vue store.ts文件或者store.js文件，快速开发`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CreateStore(storeName, jsCode)
	},
}

var storeName string
var jsCode string

func init() {
	rootCmd.AddCommand(storeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	storeCmd.Flags().StringVarP(&storeName, "name", "s", "defaultName", "生成<defaultName>.ts")
	addCmd.PersistentFlags().StringVarP(&jsCode, "jsCode", "j", "ts", "store编程语言")
}
