package baseLib

import (
	"fmt"
	"math"
)

type DMS struct {
	isNegative bool
	D          int
	M          int
	S          float32
}

func NewDMS(degree float64) DMS {
	dms := DMS{}
	dms.ParseFromDegree(degree)
	return dms
}

func (d *DMS) ParseFromDegree(degree float64) {
	d.isNegative = false
	if degree < 0 {
		d.isNegative = true
	}
	degree = math.Abs(degree)
	d.D = int(degree)
	d.M = int((degree - float64(d.D)) * 60)
	d.S = float32((degree - float64(d.D) - float64(d.M)/60) * 3600)
}

func (dms DMS) ToDegree() float64 {
	degree := float64(dms.D) + float64(dms.M)/60 + float64(dms.S)/3600
	if dms.isNegative {
		return -degree
	}
	return degree
}

func (d DMS) String() string {
	Degree := int(math.Abs(float64(d.D)))
	minutes := int(math.Abs(float64(d.M)))
	second := float32(math.Abs(float64(d.S)))
	if d.isNegative {
		return fmt.Sprintf("-%d°%d'%.2f\"", Degree, minutes, second)
	} else {
		return fmt.Sprintf("%d°%d'%.2f\"", Degree, minutes, second)
	}
}

func RoundTo2Decimals[T float64 | float32](val T) T {
	return T(math.Round(float64(val)*100) / 100)
}
