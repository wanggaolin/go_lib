package w

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
)

var Xlsx *xlsx

func init() {
	Xlsx = &xlsx{}
}

type Args_Xlsx_get_data struct {
	File_path   string   //文件路径
	Menu        []string //对应菜单栏，可为空
	Sheet_index int      //读取sheet索引，默认为0
	X           int      //正文从第几行开始读(除了菜单行)，默认为0
	Y           int      //正文从第几列开始读，默认为0
}

type Args_Xlsx_write_data struct {
	File_path  string //创建文件路径
	Sheet_Name string //创建表名
	Data       [][]Args_Xlsx_line_data
}

type Args_Xlsx_line_data struct {
	Content    string // 内容
	Color      string // 文字颜色,十六进制
	Background string // 背景填充色,十六进制
	Horizontal string // 居中样式，center
	Width      int64  // 列宽
}

func (x *xlsx) max_cell(row [][]string) (n int) {
	for _, item := range row {
		i := len(item)
		if i > n {
			n = i
		}
	}
	return n
}

func (x *xlsx) get_default_menu(menu []string, n int) (m []string) {
	for index, item := range Make_range(0, n+3) {
		if len(menu) > index {
			m = append(m, menu[index])
		} else {
			m = append(m, fmt.Sprintf("%v", item))
		}
	}
	return m
}

func (x *xlsx) to_column(number int) string {
	var result string
	for number > 0 {
		number--
		result = string(rune('A'+number%26)) + result
		number /= 26
	}
	return result
}

func (x *xlsx) to_String(n interface{}) string {
	return fmt.Sprintf("%v", n)
}

func (x *xlsx) get_sheet_name(sheel_list []string, n int) (name string, err error) {
	for index, item := range sheel_list {
		if index == n {
			return item, nil
		}
	}
	return name, errors.New("invalid sheet index")
}

func (x *xlsx) get_cell_style(arg Args_Xlsx_line_data) (style *excelize.Style, is_set bool) {
	style = &excelize.Style{}
	if arg.Color != "" {
		style.Font = &excelize.Font{
			Color: arg.Color,
		}
		is_set = true
	}
	if arg.Background != "" {
		style.Fill = excelize.Fill{
			Type:    "pattern",
			Pattern: 2,
			Color:   []string{arg.Background},
		}
		is_set = true
	}
	if arg.Horizontal != "" {
		style.Alignment = &excelize.Alignment{
			Horizontal: "center",
		}
		is_set = true
	}
	return style, is_set
}

func (x *xlsx) Read(arg Args_Xlsx_get_data) (mapData []map[string]string, err error) {
	f, err := excelize.OpenFile(arg.File_path)
	if err != nil {
		return mapData, err
	}
	defer f.Close()
	sheet_name, err := x.get_sheet_name(f.GetSheetList(), arg.Sheet_index)
	if err != nil {
		return mapData, err
	}

	rows, err := f.GetRows(sheet_name)
	if err != nil {
		return mapData, err
	}

	sheet_menu := x.get_default_menu(arg.Menu, x.max_cell(rows)-arg.Y)
	for line_index, line_item := range rows {
		if line_index <= (arg.X - 1) {
			continue
		}
		fmt.Println(line_item)
		row_data := map[string]string{}
		for cell_index, cell_item := range line_item {
			if cell_index <= (arg.Y - 1) {
				continue
			}
			row_data[sheet_menu[cell_index-arg.Y]] = cell_item
		}
		mapData = append(mapData, row_data)
	}
	return mapData, err
}

func (x *xlsx) Write(arg Args_Xlsx_write_data) (err error) {
	if arg.Sheet_Name == "" {
		arg.Sheet_Name = "Sheet1"
	}
	f := excelize.NewFile()
	defer f.Close()
	_, err = f.NewSheet(arg.Sheet_Name)
	if err != nil {
		return err
	}
	row_width := map[string]int64{}
	for row_index, row_item := range arg.Data {
		placeY := x.to_String(row_index + 1)
		for cell_index, cell_item := range row_item {
			placeX := x.to_column(cell_index + 1)
			if err = f.SetCellValue(arg.Sheet_Name, placeX+placeY, cell_item.Content); err != nil {
				return err
			}
			// set content style
			style, is_set := x.get_cell_style(cell_item)
			if is_set {
				rowStyle, err := f.NewStyle(style)
				if err != nil {
					return err
				}
				if err = f.SetCellStyle(arg.Sheet_Name, placeX+placeY, placeX+placeY, rowStyle); err != nil {
					return err
				}
			}
			// set content width
			row_width[placeX] = cell_item.Width
		}
	}

	for key, value := range row_width {
		if value > 0 {
			if err = f.SetColWidth(arg.Sheet_Name, key, key, float64(value)); err != nil {
				return err
			}
		}
	}

	if err = f.SaveAs(arg.File_path); err != nil {
		return err
	}
	return err
}
