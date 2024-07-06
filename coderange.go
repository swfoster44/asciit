package asciit

import "slices"

type codeRange struct {
	i int
	j int
}

func (c *codeRange) InRange(x int) bool {
	if c.j == 0 {
		return x == c.i
	}
	return x >= c.i && x <= c.j
}

// -----------------------------------------------------------------

type codeRanges struct {
	crange []*codeRange
}

type Ranges struct {
	allRanges []*codeRanges
	data      []string
}

func (R *Ranges) Len() int {
	return len(R.allRanges)
}
func (R *Ranges) All() []*codeRanges {
	return R.allRanges
}

func (R *Ranges) load(rows []string) {
	for i := range R.allRanges {
		crs := R.allRanges[i]

		for j := range crs.crange {
			cr := crs.crange[j]

			// if cr.j == 0, that range only contains a single entry
			if cr.j != 0 {
				R.data = slices.Concat(R.data, rows[cr.i:cr.j+1])
			} else if cr.j == 0 {
				R.data = slices.Concat(R.data, []string{rows[cr.i]})
			}
		}
	}
}

func NewRanges() *Ranges {
	all := make([]*codeRanges, 5, 5)
	all[0] = numeric_ranges()
	all[1] = alphau_ranges()
	all[2] = alphal_ranges()
	all[3] = symbols_ranges()
	all[4] = spacing_ranges()
	return &Ranges{allRanges: all}

}

//-----------------------------------------------------------------

func numeric_ranges() *codeRanges {
	r1 := codeRange{48, 57}
	r := make([]*codeRange, 1)
	r[0] = &r1
	crs := codeRanges{crange: r}
	return &crs
}

func alphal_ranges() *codeRanges {
	r1 := codeRange{97, 122}
	r := make([]*codeRange, 1)
	r[0] = &r1
	crs := codeRanges{crange: r}
	return &crs
}

func alphau_ranges() *codeRanges {
	r1 := codeRange{65, 90}
	r := make([]*codeRange, 1)
	r[0] = &r1
	crs := codeRanges{crange: r}
	return &crs
}

func spacing_ranges() *codeRanges {
	r1 := codeRange{8, 13}
	r2 := codeRange{i: 32}
	r := make([]*codeRange, 2)

	r[0] = &r1
	r[1] = &r2
	crs := codeRanges{crange: r}
	return &crs
}

func symbols_ranges() *codeRanges {
	r1 := codeRange{33, 47}
	r2 := codeRange{58, 64}
	r3 := codeRange{91, 96}
	r4 := codeRange{123, 126}

	r := make([]*codeRange, 4)
	r[0] = &r1
	r[1] = &r2
	r[2] = &r3
	r[3] = &r4
	crs := codeRanges{crange: r}
	return &crs
}
