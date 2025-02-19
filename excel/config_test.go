package excel_test

import (
	"github.com/thkhxm/excel-kit/excel"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2025/1/16
//***************************************************

func ExampleExport() {
	//设置excel文件路径
	excel.SetPath("./")
	//设置excel导出的go文件路径
	excel.SetToGoPath("./generated/conf")
	//设置excel导出的json文件路径
	excel.SetToJsonPath("./generated/conf/js")
	//开始导出excel
	excel.Export()
}
