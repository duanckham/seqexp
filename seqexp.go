package seqexp

// Seqexp ...
type Seqexp struct {
	i          int
	expression string
}

// Match ...
func (s *Seqexp) Match(seq *Seq) bool {
	return s.isMonotonic(seq)
}

func (s *Seqexp) isMonotonic(seq *Seq) bool {
	inc, dec := true, true
	for i := 0; i < seq.Length()-1; i++ {
		if seq.GT(i, i+1) {
			inc = false
		}
		if seq.LT(i, i+1) {
			dec = false
		}
	}
	return inc || dec
}
