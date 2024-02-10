package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	spreads "github.com/go-numb/go-spread-utils"
)

const (
	CREDENTIALIFILE = "path/to/credential.json"

	SPREADID = "spreadID"
	SHEETID  = "sheetID"
	RANGEKEY = "A1:Z"
)

type Row struct {
	ID   string `csv:"id"`
	Name string `csv:"name"`
	Age  string `csv:"age"`
}

func main() {
	f, err := os.Open(CREDENTIALIFILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cred, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

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
