package baselib

import (
	"fmt"
	"math"
)

type DMS struct {
	IsNegative bool    `json:"is_negative"`
	D          int     `json:"d"`
	M          int     `json:"m"`
	S          float32 `json:"s"`
}

func NewDMS(degree float64) DMS {
	dms := DMS{}
	dms.ParseFromDegree(degree)
	return dms
}

func (d *DMS) ParseFromDegree(degree float64) {
	// Handle NaN and infinities defensively
	if math.IsNaN(degree) || math.IsInf(degree, 0) {
		d.IsNegative = false
		d.D = 0
		d.M = 0
		d.S = 0
		return
	}

	// Preserve original behavior: treat -0 as non-negative
	d.IsNegative = degree < 0
	deg := math.Abs(degree)

	// Compute degrees, minutes, seconds robustly to avoid rounding issues
	d.D = int(math.Floor(deg))
	fractionalMinutes := (deg - float64(d.D)) * 60
	d.M = int(math.Floor(fractionalMinutes))
	seconds := (fractionalMinutes - float64(d.M)) * 60

	// Round seconds to 2 decimals to keep a consistent String() output
	seconds = RoundTo2Decimals[float64](seconds)

	// Handle cases where rounding pushes seconds or minutes to the next unit
	if seconds >= 60 {
		seconds -= 60
		d.M += 1
	}
	if d.M >= 60 {
		d.M -= 60
		d.D += 1
	}

	// Clamp to non-negative values (sign is stored in IsNegative)
	seconds = max(0, seconds)

	d.S = float32(seconds)
}

func (dms DMS) ToDegree() float64 {
	degree := float64(dms.D) + float64(dms.M)/60 + float64(dms.S)/3600
	if dms.IsNegative {
		return -degree
	}
	return degree
}

func (d DMS) String() string {
	deg := int(math.Abs(float64(d.D)))
	minutes := int(math.Abs(float64(d.M)))
	second := float32(math.Abs(float64(d.S)))
	if d.IsNegative {
		return fmt.Sprintf("-%d°%d'%.2f\"", deg, minutes, second)
	} else {
		return fmt.Sprintf("%d°%d'%.2f\"", deg, minutes, second)
	}
}

func RoundTo2Decimals[T float64 | float32](val T) T {
	// Convert once to float64 for rounding and then convert back to T
	return T(math.Round(float64(val)*100) / 100)
}
