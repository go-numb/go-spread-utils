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

	if len(records) <= 1 {
		return df, SetError(fmt.Errorf("no data"), "failed to read data")
	}
	for i := 0; i < len(records); i++ {
		if len(records[i]) != len(records[0]) {
			return df, SetError(fmt.Errorf("invalid data, row has not required columns"), "failed to read data")
		}
	}

	return dataframe.LoadRecords(records), nil
}
