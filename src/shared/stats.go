package shared

import (
	"fmt"
	"strings"
)

const separator = ";"

type SwarmStats struct{ strings.Builder }

func NewStats() *SwarmStats {
	return &SwarmStats{
		Builder: strings.Builder{},
	}
}

func (ss *SwarmStats) ToCsv() string {
	return ss.csvHeaders() + ss.String()
}

func (ss *SwarmStats) AddIterData(mksp float64, x []float64, seq []int) {
	fmt.Fprintf(ss, "%.3f%s %s%s %s\n", mksp, separator, encode(x), separator, encode(seq))
}

func (ss *SwarmStats) csvHeaders() string {
	return strings.Join([]string{"makespan", "best_pos", "best_seq"}, separator) + "\n"
}

func encode[T int | float64](list []T) string {
	if len(list) == 0 {
		return ""
	}
	encodingSeparator := ", "
	var sb strings.Builder
	sb.WriteString("(")

	var str string
	switch any(list[0]).(type) {
	case float64:
		str = "%.3f"
	case int:
		str = "%d"
	}

	for i, v := range list {
		if i > 0 {
			sb.WriteString(encodingSeparator)
		}
		sb.WriteString(fmt.Sprintf(str, v))
	}
	sb.WriteString(")")
	return sb.String()
}
