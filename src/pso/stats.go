package pso

import (
	"fmt"
	"strings"
)

const separator = ";"

type SwarmStats struct{ strings.Builder }

func newStats() *SwarmStats {
	return &SwarmStats{
		Builder: strings.Builder{},
	}
}

func (ss SwarmStats) ToCsv() string {
	return ss.csvHeaders() + ss.Builder.String()
}

func (ss SwarmStats) addIterData(mksp float64, x []float64, seq []int) {
	ss.WriteString(fmt.Sprintf("%.3f%s %s%s %s\n", mksp, separator, encodeFloat(x), separator, encodeInt(seq)))
}

func (ss SwarmStats) csvHeaders() string {
	return strings.Join([]string{"makespan", "best_pos", "best_seq"}, separator) + "\n"
}

func encodeFloat(list []float64) string {
	encodingSeparator := ", "
	var sb strings.Builder
	sb.WriteString("(")
	for i, v := range list {
		if i > 0 {
			sb.WriteString(encodingSeparator)
		}
		sb.WriteString(fmt.Sprintf("%.3f", v))
	}
	sb.WriteString(")")
	return sb.String()
}

func encodeInt(list []int) string {
	encodingSeparator := ", "
	var sb strings.Builder
	sb.WriteString("(")
	for i, v := range list {
		if i > 0 {
			sb.WriteString(encodingSeparator)
		}
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	sb.WriteString(")")
	return sb.String()
}
