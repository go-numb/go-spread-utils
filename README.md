# go-spread-utils

spreads is a package for easy manipulation of Google spreadsheets using the Google Sheets API. It provides functions for reading, updating, and modifying spreadsheet data in a structured way.

## Installation

To use spreads in your Go project, you can simply import it using:

```go
import "spreads"
```

Make sure to also install the necessary dependencies by running:

```sh
go get -u github.com/go-numb/go-spread-utils
```

## Usage

Here is an example of how to use spreads package in your Go code:

```go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"spreads"
)

const (
	CREDENTIALIFILE = "path/to/credential.json"

	SPREADID = "spreadID"
	SHEETID  = "sheetID"
	RANGEKEY = "A1:Z"
)


// struct for youres
type Row struct {
	ID   string `csv:"id"`
	Name string `csv:"name"`
	Age  string `csv:"age"`
}

func main() {
    // read credential.json
	f, err := os.Open(CREDENTIALIFILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cred, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

    // create client
    ctx := context.Background()
    client := spreads.New(ctx, cred)

    // Read
    // bind to struct, and more dataframes
    rows := []Row{}
    df, err := client.SetSpreadID(SPREADID).SetSheetName(SHEETID).SetRangeKey(RANGEKEY).Read(&rows)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("------------\n%+v\n------------\n", rows)

    // Update row number 1, with header
    records := df.Subset([]int{1}).Records()
    // string to interface row, and drop header
    recordsAny := spreads.ConvertStringToInterface(records)
    // drop header
    updateRow := recordsAny[1]
    if err := client.Update(updateRow); err != nil {
        log.Fatal(err)
    }

    // change spreadID, sheetName, rangeKey
    client.SetSpreadID("spreadID").SetSheetName("sheetName").SetRangeKey("A2")
}
```

In this example, the spreads package is used to read data from a Google spreadsheet into a struct, update a specific row in the spreadsheet, and change the spreadsheet's ID, sheet name, and range key.

## Documentation

For more detailed information on how to use spreads package, please refer to the [GoDoc documentation](https://pkg.go.dev/github.com/go-numb/go-spread-utils).