package generic

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"strings"
)

func (gSlice *MultiMapSlice) Filter(soft bool, gFilter Generic) (gSlice0 Slice) {
	gSlice1 := &MultiMapSlice{}

	gSlice1.tn = gSlice.tn
	gSlice1.Gs = filtera(soft, gFilter.(*MultiMapGeneric), gSlice.Gs)

	gSlice0 = gSlice1

	return
}

func filtera(soft bool, gFilter *MultiMapGeneric, gs []*MultiMapGeneric) ([]*MultiMapGeneric) {
	is := make([]int, len(gs))

	for i, _ := range is {
		is[i] = i
	}

	b, ok := gFilter.Bool0[fieldnames.Set]
	if ok && !b {
	    return []*MultiMapGeneric{}
	}

	is = filter(soft, gFilter, gs, is)

	gs0 := []*MultiMapGeneric{}
	for _, i := range is {
		gs0 = append(gs0, gs[i])
	}

	return gs0
}

func filter(soft bool, gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int) ([]int) {
	for fieldName, gFilter0 := range gFilter.Generic0 {
		switch gFilter0.Type().Name() {
		case typenames.StringFilter:
			is = filterStrings(soft, gFilter0, gs, is, fieldName)
		case typenames.Float64Filter:
			is = filterFloats(soft, gFilter0, gs, is, fieldName)
		case typenames.Int32Filter:
			is = filterInts(soft, gFilter0, gs, is, fieldName)
		case typenames.BoolFilter:
			is = filterBools(soft, gFilter0, gs, is, fieldName)
		case typenames.EnumFilter:
			is = filterStrings(soft, gFilter0, gs, is, fieldName)
		default:
			if gFilter0.Type().Flags().Is(typeflags.IsListFilter, true) {
				gFilter1, ok := gFilter0.Generic0[fieldnames.Some]
				if ok {
					is = filterList(soft, gFilter1, gs, is, fieldName, fieldnames.Some)

				    continue
				}

				gFilter1, ok = gFilter0.Generic0[fieldnames.None]
				if ok {
					is = filterList(soft, gFilter1, gs, is, fieldName, fieldnames.None)

					continue
				}

				gFilter1, ok = gFilter0.Generic0[fieldnames.Every]
				if ok {
					is = filterList(soft, gFilter1, gs, is, fieldName, fieldnames.Every)

					continue
				}
			}

			b, ok := gFilter0.Bool0[fieldnames.Set]
			if ok {
				is0 := []int{}

				if b {
					for _, i := range is {
						_, ok := gs[i].Generic0[fieldName]
						if ok {
							is0 = append(is0, i)
						} else {
						}
					}
				} else {
					for _, i := range is {
						_, ok := gs[i].Generic0[fieldName]
						if !ok {
							is0 = append(is0, i)
						} else {
						}
					}
				}

				is = is0
			}

			gs0 := []*MultiMapGeneric{}
			for _, g := range gs {
				g0, ok := g.Generic0[fieldName]
				if ok {
					gs0 = append(gs0, g0)
				} else {
					gs0 = append(gs0, &MultiMapGeneric{})
				}
			}

			is = filter(soft, gFilter0, gs0, is)
		}
	}

	return is
}

func filterList(soft bool, gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, fieldName string, kind string) ([]int) {
	is0 := []int{}

	switch kind {
	case fieldnames.None:
		for _, i := range is {
			gs0 := gs[i].GenericSlice0[fieldName].Gs

			gs1 := filtera(soft, gFilter, gs0)

			if len(gs1) == 0 {
				is0 = append(is0, i)
			}
		}
		break
	case fieldnames.Some:
		for _, i := range is {
			gs0 := gs[i].GenericSlice0[fieldName].Gs

			gs1 := filtera(soft, gFilter, gs0)

			if len(gs1) != 0 {
				is0 = append(is0, i)
			}
		}

		break
	case fieldnames.Every:
		for _, i := range is {
			gs0 := gs[i].GenericSlice0[fieldName].Gs

			gs1 := filtera(soft, gFilter, gs0)

			if len(gs1) == len(gs0) {
				is0 = append(is0, i)
			}
		}

		break
	}

	return is0
}

func filterFloats(soft bool, gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	if soft {
		return filterFloatsSoft(gFilter, gs, is, f)
	}

	return filterFloatsHard(gFilter, gs, is, f)
}

func filterFloatsSoft(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	fb, ok := gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
		} else {
			for _, i := range is {
				_, ok := gs[i].Float640[f]
				if ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	ff, ok := gFilter.Float640[fieldnames.Is]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if !ok || v == ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Not]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if !ok || v != ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ffl, ok := gFilter.Float64Slice0[fieldnames.In]
	if ok && len(ffl) != 0 {
		is0 := []int{}

		for _, i := range is {
			for _, v0 := range ffl {
				v, ok := gs[i].Float640[f]
				if !ok || v == v0 {
					is0 = append(is0, i)

					break
				}
			}
		}

		is = is0
	}

	ffl, ok = gFilter.Float64Slice0[fieldnames.NotIn]
	if ok && len(ffl) != 0 {
		is0 := []int{}

	OUTER:
		for _, i := range is {
			for _, v0 := range ffl {
				v, ok := gs[i].Float640[f]
				if !ok || v == v0 {
					continue OUTER
				}
			}

			is0 = append(is0, i)
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Lt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if !ok || v < ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Lte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if !ok || v <= ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Gt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if !ok || v > ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Gte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if !ok || v >= ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	return is
}

func filterFloatsHard(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	fb, ok := gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
			for _, i := range is {
				_, ok := gs[i].Float640[f]
				if ok {
					is0 = append(is0, i)
				}
			}
		} else {
			for _, i := range is {
				_, ok := gs[i].Float640[f]
				if !ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	ff, ok := gFilter.Float640[fieldnames.Is]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if ok && v == ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Not]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if !ok || v != ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ffl, ok := gFilter.Float64Slice0[fieldnames.In]
	if ok && len(ffl) != 0 {
		is0 := []int{}

		for _, i := range is {
			for _, v0 := range ffl {
				v, ok := gs[i].Float640[f]
				if ok && v == v0 {
					is0 = append(is0, i)

					break
				}
			}
		}

		is = is0
	}

	ffl, ok = gFilter.Float64Slice0[fieldnames.NotIn]
	if ok && len(ffl) != 0 {
		is0 := []int{}

	OUTER:
		for _, i := range is {
			for _, v0 := range ffl {
				v, ok := gs[i].Float640[f]
				if ok && v == v0 {
					continue OUTER
				}
			}

			is0 = append(is0, i)
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Lt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if ok && v < ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Lte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if ok && v <= ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Gt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if ok && v > ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	ff, ok = gFilter.Float640[fieldnames.Gte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Float640[f]
			if ok && v >= ff {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	return is
}

func filterInts(soft bool, gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	if soft {
		return filterIntsSoft(gFilter, gs, is, f)
	}

	return filterIntsHard(gFilter, gs, is, f)
}

func filterIntsSoft(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	fb, ok := gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
		} else {
			for _, i := range is {
				_, ok := gs[i].Int320[f]
				if !ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	fi, ok := gFilter.Int320[fieldnames.Is]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if !ok || v == fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Not]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if !ok || v != fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fil, ok := gFilter.Int32Slice0[fieldnames.In]
	if ok && len(fil) != 0 {
		is0 := []int{}

		for _, i := range is {
			for _, v0 := range fil {
				v, ok := gs[i].Int320[f]
				if !ok || v == v0 {
					is0 = append(is0, i)

					break
				}
			}
		}

		is = is0
	}

	fil, ok = gFilter.Int32Slice0[fieldnames.NotIn]
	if ok && len(fil) != 0 {
		is0 := []int{}

	OUTER:
		for _, i := range is {
			for _, v0 := range fil {
				v, ok := gs[i].Int320[f]
				if !ok || v == v0 {
					continue OUTER
				}
			}

			is0 = append(is0, i)
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Lt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if !ok || v < fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Lte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if !ok || v <= fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Gt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if !ok || v > fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Gte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if !ok || v >= fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	return is
}

func filterIntsHard(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	fb, ok := gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
			for _, i := range is {
				_, ok := gs[i].Int320[f]
				if ok {
					is0 = append(is0, i)
				}
			}
		} else {
			for _, i := range is {
				_, ok := gs[i].Int320[f]
				if !ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	fi, ok := gFilter.Int320[fieldnames.Is]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if ok && v == fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Not]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if !ok || v != fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fil, ok := gFilter.Int32Slice0[fieldnames.In]
	if ok && len(fil) != 0 {
		is0 := []int{}

		for _, i := range is {
			for _, v0 := range fil {
				v, ok := gs[i].Int320[f]
				if ok && v == v0 {
					is0 = append(is0, i)

					break
				}
			}
		}

		is = is0
	}

	fil, ok = gFilter.Int32Slice0[fieldnames.NotIn]
	if ok && len(fil) != 0 {
		is0 := []int{}

	OUTER:
		for _, i := range is {
			for _, v0 := range fil {
				v, ok := gs[i].Int320[f]
				if ok && v == v0 {
					continue OUTER
				}
			}

			is0 = append(is0, i)
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Lt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if ok && v < fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Lte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if ok && v <= fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Gt]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if ok && v > fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fi, ok = gFilter.Int320[fieldnames.Gte]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Int320[f]
			if ok && v >= fi {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	return is
}

func filterStrings(soft bool, gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	if soft {
		return filterStringsSoft(gFilter, gs, is, f)
	}

	return filterStringsHard(gFilter, gs, is, f)
}

func filterStringsSoft(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	caseSensitive := false

	fb, ok := gFilter.Bool0[fieldnames.CaseSensitive]
	if ok && fb {
		caseSensitive = true
	}

	fb, ok = gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
		} else {
			for _, i := range is {
				_, ok := gs[i].String0[f]
				if !ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	fs, ok := gFilter.String0[fieldnames.Is]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || v == fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.ToLower(v) == fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.Not]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || v != fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.ToLower(v) != fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fsl, ok := gFilter.StringSlice0[fieldnames.In]
	if ok && len(fsl) != 0 {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if !ok || v == v0 {
						is0 = append(is0, i)

						break
					}
				}
			}

			is = is0
		} else {
			is0 := []int{}

			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if !ok || strings.ToLower(v) == strings.ToLower(v0) {
						is0 = append(is0, i)

						break
					}
				}
			}

			is = is0
		}
	}

	fsl, ok = gFilter.StringSlice0[fieldnames.NotIn]
	if ok && len(fsl) != 0 {
		if caseSensitive {
			is0 := []int{}

		OUTERA:
			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if !ok || v == v0 {
						continue OUTERA
					}
				}

				is0 = append(is0, i)
			}

			is = is0
		} else {
			is0 := []int{}

		OUTERB:
			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if !ok || strings.ToLower(v) == strings.ToLower(v0) {
						continue OUTERB
					}
				}

				is0 = append(is0, i)
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.Contains]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.Contains(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.Contains(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.NotContains]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.Contains(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.Contains(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.StartsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.HasPrefix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.HasPrefix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.NotStartsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasPrefix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasPrefix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.EndsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.HasSuffix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.HasSuffix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.NotEndsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasSuffix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasSuffix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	return is
}

func filterStringsHard(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	caseSensitive := false

	fb, ok := gFilter.Bool0[fieldnames.CaseSensitive]
	if ok && fb {
		caseSensitive = true
	}

	fb, ok = gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
			for _, i := range is {
				_, ok := gs[i].String0[f]
				if ok {
					is0 = append(is0, i)
				}
			}
		} else {
			for _, i := range is {
				_, ok := gs[i].String0[f]
				if !ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	fs, ok := gFilter.String0[fieldnames.Is]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && v == fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && strings.ToLower(v) == fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		}

	}

	fs, ok = gFilter.String0[fieldnames.Not]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || v != fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || strings.ToLower(v) != fs {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fsl, ok := gFilter.StringSlice0[fieldnames.In]
	if ok && len(fsl) != 0 {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if ok && v == v0 {
						is0 = append(is0, i)

						break
					}
				}
			}

			is = is0
		} else {
			is0 := []int{}

			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if ok && strings.ToLower(v) == strings.ToLower(v0) {
						is0 = append(is0, i)

						break
					}
				}
			}

			is = is0
		}
	}

	fsl, ok = gFilter.StringSlice0[fieldnames.NotIn]
	if ok && len(fsl) != 0 {
		if caseSensitive {
			is0 := []int{}

		OUTERA:
			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if ok && v == v0 {
						continue OUTERA
					}
				}

				is0 = append(is0, i)
			}

			is = is0
		} else {
			is0 := []int{}

		OUTERB:
			for _, i := range is {
				for _, v0 := range fsl {
					v, ok := gs[i].String0[f]
					if ok && strings.ToLower(v) == strings.ToLower(v0) {
						continue OUTERB
					}
				}

				is0 = append(is0, i)
			}

			is = is0
		}

	}

	fs, ok = gFilter.String0[fieldnames.Contains]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && strings.Contains(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && strings.Contains(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.NotContains]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.Contains(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.Contains(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}

	}

	fs, ok = gFilter.String0[fieldnames.StartsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && strings.HasPrefix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && strings.HasPrefix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.NotStartsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasPrefix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasPrefix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.EndsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && strings.HasSuffix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if ok && strings.HasSuffix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	fs, ok = gFilter.String0[fieldnames.NotEndsWith]
	if ok {
		if caseSensitive {
			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasSuffix(v, fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		} else {
			fs = strings.ToLower(fs)

			is0 := []int{}

			for _, i := range is {
				v, ok := gs[i].String0[f]
				if !ok || !strings.HasSuffix(strings.ToLower(v), fs) {
					is0 = append(is0, i)
				}
			}

			is = is0
		}
	}

	return is
}

func filterBools(soft bool, gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	if soft {
		return filterBoolsSoft(gFilter, gs, is, f)
	}

	return filterBoolsHard(gFilter, gs, is, f)
}

func filterBoolsSoft(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	fb, ok := gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
		} else {
			for _, i := range is {
				_, ok := gs[i].Bool0[f]
				if !ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	fb, ok = gFilter.Bool0[fieldnames.Is]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Bool0[f]
			if !ok || v == fb {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fb, ok = gFilter.Bool0[fieldnames.Not]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Bool0[f]
			if !ok || v != fb {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	return is
}

func filterBoolsHard(gFilter *MultiMapGeneric, gs []*MultiMapGeneric, is []int, f string) ([]int) {
	fb, ok := gFilter.Bool0[fieldnames.Set]
	if ok {
		is0 := []int{}

		if fb == true {
			for _, i := range is {
				_, ok := gs[i].Bool0[f]
				if ok {
					is0 = append(is0, i)
				}
			}
		} else {
			for _, i := range is {
				_, ok := gs[i].Bool0[f]
				if !ok {
					is0 = append(is0, i)
				}
			}
		}

		is = is0
	}

	fb, ok = gFilter.Bool0[fieldnames.Is]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Bool0[f]
			if ok && v == fb {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	fb, ok = gFilter.Bool0[fieldnames.Not]
	if ok {
		is0 := []int{}

		for _, i := range is {
			v, ok := gs[i].Bool0[f]
			if !ok || v != fb {
				is0 = append(is0, i)
			}
		}

		is = is0
	}

	return is
}
