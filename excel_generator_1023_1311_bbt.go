// 代码生成时间: 2025-10-23 13:11:34
package main

import (
    "encoding/csv"
    "errors"
    "fmt"
    "os"
    "path/filepath"
    "time"

    "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// ExcelGenerator 结构体，用于封装生成Excel的相关属性和方法
type ExcelGenerator struct {
    Filepath string // Excel文件保存路径
    Filename string // Excel文件名称
}

// NewExcelGenerator 创建一个新的ExcelGenerator实例
func NewExcelGenerator(filepath, filename string) *ExcelGenerator {
    return &ExcelGenerator{
        Filepath: filepath,
        Filename: filename,
    }
}

// GenerateExcel 生成Excel文件
func (e *ExcelGenerator) GenerateExcel(data [][]string, sheetName string) error {
    // 创建Excel文件
    f := excelize.NewFile()
    // 创建一个sheet
    index := f.NewSheet(sheetName)
    if err := f.SetActiveSheet(index); err != nil {
        return err
    }

    // 设置header
    for i, header := range data[0] {
        f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+1), header)
    }

    // 设置数据行
    for i, row := range data[1:] {
        for j, cell := range row {
            f.SetCellValue(sheetName, fmt.Sprintf("%c%d", 'A'+j-1, i+2), cell)
        }
    }

    // 保存Excel文件
    if err := f.SaveAs(filepath + "/" + filename); err != nil {
        return err
    }

    return nil
}

// main函数，程序入口
func main() {
    // 示例数据
    data := [][]string{
        {""Sheet1"", ""Column A"", ""Column B"", ""Column C""},
        {""Row 1 Cell 1"", ""Row 1 Cell 2"", ""Row 1 Cell 3""},
        {""Row 2 Cell 1"", ""Row 2 Cell 2"", ""Row 2 Cell 3""},
        {""Row 3 Cell 1"", ""Row 3 Cell 2"", ""Row 3 Cell 3""},
    }

    // 文件生成路径和文件名
    filepath := "./"
    filename := fmt.Sprintf("example_%s.xlsx", time.Now().Format("2006-01-02_15-04-05"))

    // 创建Excel生成器实例
    generator := NewExcelGenerator(filepath, filename)

    // 生成Excel文件
    if err := generator.GenerateExcel(data, "Sheet1"); err != nil {
        fmt.Printf("Failed to generate Excel: %v
", err)
    } else {
        fmt.Printf("Excel file generated at: %s
", filepath+"/"+filename)
    }
}
