package excel_utils

// @description   获取指定行添加数据
func (df *DataFrame) Add(data []string, rowNum int) DataFrame {
	var dataframe DataFrame
	rows := df.Data
	// 获取数据
	res_data := make([][]string, len(rows)+1)
	res_data[rowNum] = data
	for irow, row := range rows {
		if irow < rowNum {
			res_data[irow] = row
		} else {
			res_data[irow+1] = row
		}
	}
	dataframe.Data = res_data
	dataframe.ColLen = len(res_data)
	dataframe.RowLen = len(res_data[0])
	return dataframe
}
