package expansion

import (
	"fmt"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
)

func Step(verbosityLevel int, rn *graph.RootNode, name string, f func()) {
	var missingBefore []string
	var unusedBefore []string
	var idsBefore []string

	if verbosityLevel != 0 {
		missingBefore = graph.ToStrings(graph.GetMissing(rn).Type.Ids...)
		unusedBefore = graph.ToStrings(graph.GetUnused(rn, graph.TYPE).Type.Ids...)
		idsBefore = graph.ToStrings(rn.Types.GetIds()...)
	}

	f()

	if verbosityLevel == 0 {
		return
	}

	missingAfter := graph.ToStrings(graph.GetMissing(rn).Type.Ids...)
	unusedAfter := graph.ToStrings(graph.GetUnused(rn, graph.TYPE).Type.Ids...)
	idsAfter := graph.ToStrings(rn.Types.GetIds()...)

	idsAdded, _ := change(idsBefore, idsAfter)
	unusedAdded, unusedRemoved := change(unusedBefore, unusedAfter)
	_, unusedUnchanged := change(unusedAfter, unusedAdded)
	missingAdded, missingRemoved := change(missingBefore, missingAfter)
	_, missingUnchanged := change(missingAfter, missingAdded)

	missing := map[string][]string{}
	rn.Types.Each(func(n *graph.TypeNode) {
		missing[graph.ToString(n.Id())] = graph.ToStrings(n.Edges.Types.Resolver.Misses()...)
	})

	//missingUnchanged = expandMissingNames(missing, missingUnchanged)
	//missingAdded = expandMissingNames(missing, missingAdded)

	data := [][]string{
		idsAdded,
		unusedUnchanged,
		unusedAdded,
		unusedRemoved,
		missingUnchanged,
		missingAdded,
		missingRemoved,
	}

	sortColumns(data)

	data = toRows(data)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		fmt.Sprintf("types (%v -> %v)", len(idsBefore), len(idsAfter)),
		fmt.Sprintf("unused (%v -> %v)", len(unusedBefore), len(unusedAfter)),
		"",
		"",
		fmt.Sprintf("missing (%v -> %v)", len(missingBefore), len(missingAfter)),
		"",
		"",
	})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgRedColor},
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgRedColor},
	)

	table.SetColMinWidth(0, 10)
	table.SetColMinWidth(1, 10)
	table.SetColMinWidth(2, 10)
	table.SetColMinWidth(3, 10)
	table.SetColMinWidth(4, 10)
	table.SetColMinWidth(5, 10)
	table.SetColMinWidth(6, 10)

	table.AppendBulk(data)

	fmt.Printf("\n\nstep: %v\n", name)
	table.Render()
}

func change(before, after []string) (added, removed []string) {
	f := func(a, b []string) (diff []string) {
		m := map[string]bool{}

		for _, v := range a {
			m[v] = true
		}

		for _, v := range b {
			_, ok := m[v]
			if !ok {
				diff = append(diff, v)
			}
		}

		return
	}

	added = f(before, after)
	removed = f(after, before)

	return
}

func toRows(columns [][]string) (rows [][]string) {
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

func sortColumns(columns [][]string) () {
	for _, c := range columns {
		sort.Strings(c)
	}
}

func expandMissingNames(missing map[string][]string, names []string) (expanded []string) {
	for _, n := range names {
		expanded = append(expanded, fmt.Sprintf("-> %v %v", n, missing[n]))
	}

	return
}