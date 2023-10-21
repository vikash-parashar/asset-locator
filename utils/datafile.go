package utils

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jung-kurt/gofpdf"
)

// TODO:
func GeneratePDFFile(data interface{}) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Your PDF Content Here")
	// Add your data to the PDF

	return pdf.OutputFileAndClose("output.pdf")
}

// TODO:

// GenerateExcelFile generates an Excel file from the provided data and returns the Excel file object.
func GenerateExcelFile(data interface{}) (*excelize.File, error) {
	// Create a new Excel file
	xlsx := excelize.NewFile()

	// Create a new sheet
	sheetName := "Sheet1"
	xlsx.NewSheet(sheetName)

	// Set headers (assuming you have headers for your data)
	headers := []string{"Header1", "Header2", "Header3"}
	for col, header := range headers {
		xlsx.SetCellValue(sheetName, excelize.ToAlphaString(col+1)+"1", header)
	}

	// Add data rows
	for row, item := range data {
		// You'll need to adjust the column indexes and data fields based on your data structure
		xlsx.SetCellValue(sheetName, "A"+strconv.Itoa(row+2), item.Field1)
		xlsx.SetCellValue(sheetName, "B"+strconv.Itoa(row+2), item.Field2)
		xlsx.SetCellValue(sheetName, "C"+strconv.Itoa(row+2), item.Field3)
		// Add more fields as needed
	}

	return xlsx, nil
}
