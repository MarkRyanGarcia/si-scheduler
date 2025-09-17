package drive

import (
	"fmt"
	"os"
)

// Reads the Google Sheets data and returns all of the form responses
func ReadSheetData() ([][]string, error) {
	srv, err := Setup()
	if err != nil {
		return nil, err
	}
	// Just testing shit out for now
	spreadsheetId := os.Getenv("TEST_SHEET_ID")
	if spreadsheetId == "" {
		return nil, fmt.Errorf("TEST_SHEET_ID environment variable not set")
	}
	readRange := "'Form Responses 1'!A2:BV"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}
	res := make([][]string, len(resp.Values))
	for i, row := range resp.Values {
		res[i] = make([]string, len(row))
		for j, cell := range row {
			if cell != nil {
				res[i][j] = fmt.Sprintf("%v", cell)
			} else {
				res[i][j] = ""
			}
		}
	}
	return res, nil
}
