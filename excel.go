package excel_utils

import (
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type DataFrame struct {
	Data   [][]string //excel表二维数据
	ColLen int        //列数
	RowLen int        //行数
}

// @description   对excel表格进行数据读取
func ReadExcel(excel_path, sheet_name string) (DataFrame, error) {

	var dataframe DataFrame

	xlsx, err := excelize.OpenFile(excel_path)
	if err != nil {
		os.Exit(1)
		return dataframe, err
	}

	rows := xlsx.GetRows(sheet_name)

	dataframe.Data = rows
	dataframe.RowLen = len(rows[0])
	dataframe.ColLen = len(rows)
	return dataframe, nil
}
