package points

import (
	"fmt"
	"strings"
)

type Coordinate2D [2]float64
type Points2D []Coordinate2D

func (c2d Coordinate2D) String() string {
	return fmt.Sprintf("(%.3f, %.3f)", c2d[0], c2d[1])
}

func (pts2d Points2D) String() string {
	var strPoints []string
	for _, pt := range pts2d {
		strPoints = append(strPoints, pt.String())
	}
	return "[" + strings.Join(strPoints, ", ") + "]"
}
