package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const (
	start  = 32
	finish = 126

	groups = 4
)

func main() {
	fmt.Println("ascii shall receive...")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Dec", "ASCII", "Hex", "Oct", "Bin"})

	row := make([]string, 5)

	for i := start; i <= finish; i++ {
		row[0] = fmt.Sprint(i)
		row[1] = string(rune(i))

		// Hex
		src := []byte(row[1])
		hx := strings.ToUpper(hex.EncodeToString(src))
		row[2] = hx

		// Oct
		oct := strconv.FormatInt(int64(i), 8)
		row[3] = fmt.Sprintf("%03s", oct)

		// Bin
		bin := fmt.Sprintf(strconv.FormatInt(int64(i), 2))
		row[4] = fmt.Sprintf("%08s", bin)

		// Doesn't seem right that alignment should have to be set for every row but it doesn't work anywhere else
		table.SetColumnAlignment([]int{
			tablewriter.ALIGN_RIGHT,
			tablewriter.ALIGN_CENTER,
			tablewriter.ALIGN_RIGHT,
			tablewriter.ALIGN_RIGHT,
			tablewriter.ALIGN_RIGHT},
		)

		table.Append(row)
	}

	table.Render()

	os.Exit(0)
}
