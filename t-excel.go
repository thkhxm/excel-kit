package main

import (
	"flag"
	"fmt"
	"github.com/thkhxm/excel-kit/excel"
	"os"
	"strings"
)

// ***************************************************
// @Link  https://github.com/thkhxm/tgf
// @Link  https://gitee.com/timgame/tgf
// @QQ群 7400585
// author tim.huang<thkhxm@gmail.com>
// @Description
// 2025/1/16
func main() {

	// Define command-line flags
	//path := flag.String("path", "./", "Path to the excel files")
	//jsonPath := flag.String("jsonPath", "./generated/conf/js", "Path to the json files")
	//goPath := flag.String("goPath", "./generated/conf", "Path to the go files")
	//tsPath := flag.String("tsPath", "./generated/conf/ts", "Path to the ts files")
	path := flag.String("path", "./", "Path to the excel files")
	jsonPath := flag.String("jsonPath", "./generated/conf/js", "Path to the json files")
	goPath := flag.String("goPath", "", "Path to the go files")
	tsPath := flag.String("tsPath", "", "Path to the ts files")
	// Parse command-line flags
	flag.Parse()

	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "help" {
		flag.Usage()
		fmt.Println("\nSupported command-line flags:")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("  -%s: %s\n", f.Name, f.Usage)
		})
		return
	}

	// Validate the provided paths
	if _, err := os.Stat(*path); os.IsNotExist(err) {
		fmt.Printf("Path does not exist: %s\n", *path)
		return
	}
	//if _, err := os.Stat(*jsonPath); os.IsNotExist(err) {
	//	fmt.Printf("Json path does not exist: %s\n", *jsonPath)
	//	return
	//}
	//if _, err := os.Stat(*goPath); os.IsNotExist(err) {
	//	fmt.Printf("Go path does not exist: %s\n", *goPath)
	//	return
	//}
	//if _, err := os.Stat(*tsPath); os.IsNotExist(err) {
	//	fmt.Printf("TS path does not exist: %s\n", *tsPath)
	//	return
	//}

	// Set paths for excel package
	excel.SetPath(*path)
	excel.SetToJsonPath(*jsonPath)
	excel.SetToGoPath(*goPath)
	excel.SetToTsPath(*tsPath)

	// Start export
	excel.Export()
}
