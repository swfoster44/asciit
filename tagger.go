package asciit

const (
	numeric = iota
	alphal
	alphau
	spacing
	symbols
)

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

type codeRanges struct {
	ranges []*codeRange
	tag    int
}

func (c *codeRanges) tagger(e *entry) bool {
	for i := 0; i < len(c.ranges); i++ {
		r := c.ranges[i]
		if ok := r.InRange(e.code); ok {
			e.tag = c.tag
			return true
		}
	}

	return false
}

type Tagger struct {
	allRanges []*codeRanges
}

func (T *Tagger) tag(e *entry) bool {
	for i := 0; i < len(T.allRanges); i++ {
		cr := T.allRanges[i]
		if ok := cr.tagger(e); ok {
			return true
		}
	}

	return false
}

func NewTagger() *Tagger {
	all := make([]*codeRanges, 5, 5)
	all = append(all, numeric_ranges())
	all = append(all, alphau_ranges())
	all = append(all, alphal_ranges())
	all = append(all, symbols_ranges())
	all = append(all, spacing_ranges())
	return &Tagger{all}
}

func numeric_ranges() *codeRanges {
	r1 := codeRange{8, 13}
	r := make([]*codeRange, 1)
	r[0] = &r1
	cr := codeRanges{ranges: r, tag: numeric}
	return &cr
}

func alphal_ranges() *codeRanges {
	r1 := codeRange{97, 122}
	r := make([]*codeRange, 1)
	r[0] = &r1
	cr := codeRanges{ranges: r, tag: alphal}
	return &cr
}

func alphau_ranges() *codeRanges {
	r1 := codeRange{65, 90}
	r := make([]*codeRange, 1)
	r[0] = &r1
	cr := codeRanges{ranges: r, tag: alphau}
	return &cr
}

func spacing_ranges() *codeRanges {
	r1 := codeRange{8, 13}
	r2 := codeRange{i: 32}
	r := make([]*codeRange, 2)

	r[0] = &r1
	r[1] = &r2
	cr := codeRanges{ranges: r, tag: spacing}
	return &cr
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
	cr := codeRanges{ranges: r, tag: symbols}
	return &cr
}
