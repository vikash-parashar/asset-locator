package utils

import (
	"reflect"

	"github.com/tealeg/xlsx"
)

// Function to generate an Excel file for any device detail type
func GenerateExcelFile(data interface{}, filename string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("DeviceDetails")
	if err != nil {
		return err
	}

	// Get the type and value of the data
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	// Add headers
	headerRow := sheet.AddRow()
	for i := 0; i < dataType.NumField(); i++ {
		headerRow.AddCell().SetString(dataType.Field(i).Tag.Get("json"))
	}

	// Add data
	dataSlice := dataValue.Interface()
	sliceValue := reflect.ValueOf(dataSlice)
	for i := 0; i < sliceValue.Len(); i++ {
		dataRow := sheet.AddRow()
		for j := 0; j < dataType.NumField(); j++ {
			field := sliceValue.Index(i).Field(j)
			dataRow.AddCell().SetValue(field.Interface())
		}
	}

	// Save the file
	err = file.Save(filename)
	if err != nil {
		return err
	}

	return nil
}
