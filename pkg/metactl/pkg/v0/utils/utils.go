package utils

import (
	"encoding/json"
	"fmt"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/types"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
	"reflect"
	"strings"
)

func Plural(s string) (string) {
	if strings.HasSuffix(s, "y") {
		s = strings.TrimSuffix(s, "y")
		s = s + "ie"
	}

	s = s + "s"

	return s
}

func ToRows(columns [][]string) (rows [][]string) {
	c := 0
	for {
		row := []string{}
		cSkipped := 0
		for i, _ := range columns {
			if len(columns[i]) <= c {
				row = append(row, "")

				cSkipped++

				continue
			}

			row = append(row, columns[i][c])
		}

		if cSkipped == len(columns) {
			break
		}

		c++

		rows = append(rows, row)
	}

	return
}

func GetTableString(header []string, rows [][]string, containNewLines []int) (string) {
	rows0 := [][]string{}

	for i, _ := range rows {
		lines := map[int][]string{}
		firstLines := map[int]string{}
		nMaxLines := 0

		for _, i0 := range containNewLines {
			lines[i0] = strings.Split(rows[i][i0], "\n")

			if len(lines[i0]) != 0 {
				firstLines[i0] = lines[i0][0]
			} else {
				firstLines[i0] = ""
			}

			if len(lines[i0]) > nMaxLines {
				nMaxLines = len(lines[i0])
			}
		}

		row := []string{}
		for i0, _ := range rows[i] {
			s, ok := firstLines[i0]
			if ok {
				row = append(row, s)

				continue
			}

			row = append(row, rows[i][i0])
		}

		rows0 = append(rows0, row)

		for i0 := 1; i0 < nMaxLines; i0++ {
			currentLines := map[int]string{}

			for _, i1 := range containNewLines {
				if i0 < len(lines[i1]) {
					currentLines[i1] = lines[i1][i0]
				}
			}

			row := []string{}

			for i1, _ := range rows[i] {
				s, ok := currentLines[i1]
				if ok {
					row = append(row, s)

					continue
				}

				row = append(row, "")
			}

			rows0 = append(rows0, row)
		}
	}

	tableString := &strings.Builder{}

	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header)

	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})

	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)

	table.SetCenterSeparator("")
	table.SetRowSeparator("")
	table.SetColumnSeparator("")

	table.SetAutoWrapText(false)

	table.AppendBulk(rows0)

	table.Render()

	return tableString.String()
}

func PrintReport(noColor bool, outputFormat string, returnData bool, messageReport types.MessageReport, o types.Output) (err error) {
	if returnData == true {
		err = printData(outputFormat, messageReport, o)
		if err != nil {
			return
		}
	} else {
		printText(noColor, messageReport, o)
	}

	return
}

func printData(outputFormat string, messageReport types.MessageReport, o types.Output) (err error) {
	dataReport := types.DataReport{
		MessageReport: messageReport,
		Data:          o.Data,
	}

	var b []byte
	switch outputFormat {
	case types.FORMAT_JSON:
		b, err = json.Marshal(dataReport)
		if err != nil {
			return
		}
	case types.FORMAT_YAML:
		b, err = yaml.Marshal(dataReport)
		if err != nil {
			return
		}
	default:
		panic("expected format to be either json or yaml")
	}

	print(string(b))

	return
}

func printText(noColor bool, messageReport types.MessageReport, o types.Output) {
	debug := "\u001b[37mdebug\u001b[0m"
	info := "\u001b[34minfo\u001b[0m"
	warning := "\u001b[33mwarning\u001b[0m"
	error0 := "\u001b[31merror\u001b[0m"

	if noColor {
		debug = "debug:"
		info = "info:"
		warning = "warning:"
		error0 = "error:"
	}

	for _, s := range messageReport.Debug {
		fmt.Printf("%v %v\n", debug, s)
	}

	for _, s := range messageReport.Info {
		fmt.Printf("%v %v\n", info, s)
	}

	print(o.Text)

	for _, s := range messageReport.Warning {
		fmt.Printf("%v %v\n", warning, s)
	}

	for _, s := range messageReport.Error {
		fmt.Printf("%v %v\n", error0, s)
	}
}

func StructsToNamedColumns(ss interface{}) (columns map[string][]string) {
	structColumns := []map[string]string{}

	v := reflect.ValueOf(ss)

	for i := 0; i < v.Len(); i++ {
		structColumns = append(structColumns, StructToNamedColumns("", v.Index(i).Interface()))
	}

	columnNames := map[string]bool{}
	for _, c := range structColumns {
		for k, _ := range c {
			columnNames[k] = true
		}
	}

	columns = map[string][]string{}
	for _, c := range structColumns {
		for k, _ := range columnNames {
			v := c[k]
			columns[k] = append(columns[k], v)
		}
	}

	return
}

func StructToNamedColumns(prefix string, s interface{}) (columns map[string]string) {
	if prefix != "" {
		prefix += "."
	}

	columns = map[string]string{}

	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		fName := prefix + t.Field(i).Name

		if f.IsNil() {
			continue
		}

		switch f.Elem().Kind() {
		case reflect.Struct:
			columns0 := StructToNamedColumns(fName, f.Elem().Interface())
			for k, v := range columns0 {
				columns[k] = v
			}
		case reflect.Int32:
			columns[fName] = fmt.Sprintf("%v", f.Elem().Int())
		case reflect.Float64:
			columns[fName] = fmt.Sprintf("%v", f.Elem().Float())
		case reflect.Bool:
			columns[fName] = fmt.Sprintf("%v", f.Elem().Bool())
		case reflect.String:
			columns[fName] = f.Elem().String()
		}
	}

	return
}

func GetTableStringFromStructs(ss interface{}) (string) {
	namedColumns := StructsToNamedColumns(ss)

	header := []string{}
	columns := [][]string{}
	for k, v := range namedColumns {
		header = append(header, k)
		columns = append(columns, v)
	}

	rows := ToRows(columns)

	return GetTableString(header, rows, nil)
}

func ConcatTaskSets(taskSets...[]types.RenderTask) (set []types.RenderTask) {
	for _, set0 := range taskSets {
		set = append(set, set0...)
	}

	return
}