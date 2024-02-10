package spreads

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
)

func (p *Client) Read(binder any) (dataframe.DataFrame, error) {
	df := dataframe.DataFrame{}

	rangeKey := rangekey(p.sheetName, p.rangeKey)
	res, err := p.Sheets.Spreadsheets.Values.Get(p.spreadID, rangeKey).Do()
	if err != nil {
		return df, SetError(err, fmt.Sprintf("failed to read data, spread ID: %s, range key: %s", p.spreadID, rangeKey))
	}

	records := ConvertInterfaceToString(res.Values)
	if err := Bind(records, binder); err != nil {
		return df, SetError(err, "failed to bind data")
	}

	return dataframe.LoadRecords(records), nil
}
