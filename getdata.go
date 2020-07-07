package excel_utils

// @description   通过列名对数据进行筛选
func (df *DataFrame) ValueLoc(cols []string) DataFrame {
	var dataframe DataFrame
	rows := df.Data
	colIndex := make([]int, len(cols))

	// 获取每个col的所在序列号
	for index, row := range rows {
		if index == 0 {
			num := 0
			for _, col := range cols {
				for key, colCell := range row {
					if colCell == col {
						colIndex[num] = key + 1
						num++
					}
				}
			}
		}
	}

	//	对存在的量进行重新矫正，以解决初始变量长度问题
	res_len := 0
	for _, coli := range colIndex {
		if coli-1 >= 0 {
			res_len++
		}
	}

	// 获取数据
	res_data := make([][]string, len(rows)-1)
	res_index := 0
	for index, row := range rows {
		if index != 0 {
			data := make([]string, res_len)
			for i, colindex := range colIndex {
				for key, colCell := range row {
					if key == colindex-1 {
						data[i] = colCell
					}
				}
			}
			res_data[res_index] = data
			res_index++
		}
	}
	dataframe.Data = res_data
	dataframe.ColLen = len(rows) - 1
	dataframe.RowLen = res_len
	return dataframe
}

// @description   通过列序号对数据进行筛选
func (df *DataFrame) ValueIndex(cols []int) DataFrame {
	var dataframe DataFrame
	rows := df.Data
	colIndex := cols
	// 获取数据
	res_data := make([][]string, len(rows)-1)
	res_index := 0
	for index, row := range rows {
		if index != 0 {
			data := make([]string, len(cols))
			for i, colindex := range colIndex {
				for key, colCell := range row {
					if key == colindex {
						data[i] = colCell
					}
				}
			}
			res_data[res_index] = data
			res_index++
		}
	}
	dataframe.Data = res_data
	dataframe.ColLen = len(rows) - 1
	dataframe.RowLen = len(cols)
	return dataframe
}

// @description   获取指定范围行数据
func (df *DataFrame) RowRange(start, end int) DataFrame {
	var dataframe DataFrame
	rows := df.Data
	// 获取数据
	res_data := make([][]string, end-start)
	res_index := 0
	for index, row := range rows {
		if index >= start && index < end {
			res_data[res_index] = row
			res_index++
		}
	}
	dataframe.Data = res_data
	dataframe.ColLen = len(res_data)
	dataframe.RowLen = len(rows[0])
	return dataframe
}

// @description   获取指定行列单元格数据
func (df *DataFrame) Value(rowNum, colNum int) string {
	rows := df.Data
	// 获取数据
	for irow, row := range rows {
		if irow == rowNum {
			for iloc, loc := range row {
				if iloc == colNum {
					return loc
				}
			}
		}
	}
	return ""
}
