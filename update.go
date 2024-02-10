package spreads

import (
	"google.golang.org/api/sheets/v4"
)

func (p *Client) Update(row []interface{}) error {
	rangeKey := rangekey(p.sheetName, p.rangeKey)
	// 値の更新、該当行を上書きする
	if _, err := p.Sheets.Spreadsheets.Values.Update(p.spreadID, rangeKey, &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			row,
		},
	}).ValueInputOption("USER_ENTERED").Do(); err != nil {
		return SetError(err, "failed to update data")
	}

	return nil
}
