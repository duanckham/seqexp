package seqexp

import (
	"math"
	"strconv"
)

// Seq ...
type Seq []interface{}

func (s Seq) Length() int {
	return len(s)
}

// Value for get value.
func (s Seq) Value(i int) float64 {
	switch i := s[i].(type) {
	case float64:
		return float64(i)
	case float32:
		return float64(i)
	case int64:
		return float64(i)
	case int32:
		return float64(i)
	case int16:
		return float64(i)
	case int8:
		return float64(i)
	case uint64:
		return float64(i)
	case uint32:
		return float64(i)
	case uint16:
		return float64(i)
	case uint8:
		return float64(i)
	case int:
		return float64(i)
	case uint:
		return float64(i)
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return math.NaN()
		}
		return f
	default:
		return math.NaN()
	}
}

// GT i greater than j.
func (s Seq) GT(i, j int) bool {
	return s.Value(i) > s.Value(j)
}

// LT i less than j.
func (s Seq) LT(i, j int) bool {
	return s.Value(i) < s.Value(j)
}
