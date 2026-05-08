package main

import (
	"fmt"
	"math"

	baselib "github.com/jenujari/planets-lib"
	"github.com/jenujari/planets-lib/bal"
)

func main() {
	// Example 1: Calculate tithy (lunar day)
	moon := 84 + 12.0/60.0 // 84°12'
	sun := 227 + 46.0/60.0 // 227°46'
	tithy := baselib.CalcTithy(moon, sun)
	fmt.Printf("Example: CalcTithy\n  Moon: %.6f°, Sun: %.6f° -> Tithy: %d\n\n", moon, sun, tithy)

	// Example 2: Map a longitude to a zodiac sign
	longitude := 359.999
	sign := baselib.GetSignFrmDegree(longitude)
	fmt.Printf("Example: GetSignFrmDegree\n  Longitude: %.6f° -> Sign: %s\n\n", longitude, sign)

	// Example 3: Get nakshatra and pada from a longitude
	nak := baselib.GetNakshatraPadaFromDegree(longitude)
	fmt.Printf("Example: GetNakshatraPadaFromDegree\n  Longitude: %.6f° -> Nakshatra: %s, Pada: %d\n\n", longitude, nak.Name, nak.Pada)

	// Example 4: PlanetCord usage and derived values
	// Note: longitude can be unnormalized (e.g., >360), CalculateDerivedValues will use normalized angle for lookups.
	p := &baselib.PlanetCord{
		Name:      baselib.MARS,
		Longitude: 721.5,  // will normalize to 1.5°
		Latitude:  -1.234, // example latitude
		SpeedLong: -0.512, // retrograde example
	}

	p.CalculateDerivedValues()

	fmt.Printf("Example: PlanetCord.CalculateDerivedValues\n")
	fmt.Printf("  Raw Longitude: %.6f°\n", p.Longitude)
	fmt.Printf("  Normalized Sign: %s\n", p.Sign)
	fmt.Printf("  Nakshatra: %s (Pada %d)\n", p.Nakshatra.Name, p.Nakshatra.Pada)
	fmt.Printf("  IsRetro: %v\n", p.IsRetro)
	fmt.Printf("  SpeedCategory: %s\n", p.SpeedCategory)
	fmt.Printf("  Vedha: %s\n", p.Vedha)
	fmt.Printf("  Longitude DMS: %s\n\n", p.LongitudeDMS.String())

	// Example 6: Calculate Uchch Bala (exaltation strength)
	plName := baselib.JUPITER
	plLong := 95.0 // Exaltation point for Jupiter
	uchhBal := bal.UchhBal(plLong, plName)
	fmt.Printf("Example: UchhBal\n  Planet: %s, Longitude: %.2f° -> Uchch Bala: %.2f\n\n", plName, plLong, uchhBal)

	// Example 7: Calculate Navansh Rashi
	navRashiNum, navRashiName := baselib.CalcNavanshRashi(plLong)
	fmt.Printf("Example: CalcNavanshRashi\n  Longitude: %.2f° -> Navansh Rashi: %s (%d)\n\n", plLong, navRashiName, navRashiNum)

	// Example 8: Calculate Udaybal (combustion/rising strength)
	sunLong := 10.0
	marsLong := 40.0
	marsSpeed := 0.5
	udayBal := bal.UdayBal(sunLong, marsLong, marsSpeed, baselib.MARS)
	fmt.Printf("Example: UdayBal\n  Sun: %.2f°, Mars: %.2f°, Speed: %.2f -> Udaybal: %.2f\n\n", sunLong, marsLong, marsSpeed, udayBal)

	// Example 9: Calculate Vakra Bal (retrograde strength)
	retroMarsSpeed := -0.4
	vakraBal := bal.VakraBal(retroMarsSpeed, baselib.MARS)
	fmt.Printf("Example: VakraBal\n  Planet: %s, Speed: %.2f -> Vakra Bal: %.2f\n\n", baselib.MARS, retroMarsSpeed, vakraBal)

	// Example 10: Get Graha Maitri (Planetary Relationship)
	base := baselib.SUN
	target := baselib.SATURN
	rel, _ := baselib.GetGrahaMaitri(base, target)
	fmt.Printf("Example: GetGrahaMaitri\n  Relationship from %s to %s is %s\n\n", base, target, rel)

	// Example 11: Calculate Kshetra Bal (positional strength)
	ksBal, _ := bal.KshetraBal(127.5, baselib.SUN) // 7.5° into Leo
	fmt.Printf("Example: KshetraBal\n  Sun at 127.5° (Leo) -> Kshetra Bal: %.2f\n\n", ksBal)

	// Example 12: Calculate Navansh Bal (Navamsha positional strength)
	navBal, _ := bal.NavanshBal(1.666666, baselib.MARS) // Aries Navamsha Midpoint
	fmt.Printf("Example: NavanshBal\n  Mars in Aries Navamsha Midpoint -> Navansh Bal: %.2f%%\n\n", navBal)

	// Example 5: Handling invalid inputs (NaN / Inf)
	invalid := math.NaN()
	fmt.Printf("Example: Invalid inputs\n")
	fmt.Printf("  CalcTithy(NaN, %.6f) -> %d (0 indicates invalid input)\n", sun, baselib.CalcTithy(invalid, sun))
	fmt.Printf("  GetSignFrmDegree(NaN) -> %q (empty string indicates invalid input)\n", baselib.GetSignFrmDegree(invalid))
	fmt.Printf("  GetNakshatraPadaFromDegree(NaN) -> %+v (zero-value indicates invalid input)\n", baselib.GetNakshatraPadaFromDegree(invalid))
}
