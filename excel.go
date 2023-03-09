package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func WriteToCsv(csvData CsvData) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "Product page")
	f.SetCellValue("Sheet1", "B1", "Prompt")

	f.SetCellValue("Sheet1", "C1", "GPT3 Summary")
	f.SetCellValue("Sheet1", "D1", "Hindi")
	f.SetCellValue("Sheet1", "E1", "Image")

	//f.SetCellValue("Sheet1", "D1", "Hindi Summary")

	f.SetCellValue("Sheet1", "A"+strconv.Itoa(2), csvData.ProductLink)
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(2), csvData.Prompt)

	f.SetCellValue("Sheet1", "C"+strconv.Itoa(2), csvData.MetaDescription)

	f.SetCellValue("Sheet1", "D"+strconv.Itoa(2), csvData.HindiTranslate)
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(2), csvData.Image)

	f.SetActiveSheet(index)
	if err := f.SaveAs("micro_health.xlsx"); err != nil {
		return err
	}

	fmt.Println(csvData)

	return nil
}
