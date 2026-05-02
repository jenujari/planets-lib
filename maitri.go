package baselib

import (
	"fmt"
)

// Relationship constants as defined in the Graha Maitri Chakra.
const (
	FRIEND  = "Friend"
	NEUTRAL = "Neutral"
	ENEMY   = "Enemy"
)

// grahaMaitriChakra codifies the Permanent Planetary Relationship Table (Narpatijayacharya).
// The map is structured as [base_planet][target_planet] -> Relationship.
var grahaMaitriChakra = map[string]map[string]string{
	SUN: {
		MOON:    FRIEND,
		MARS:    FRIEND,
		JUPITER: FRIEND,
		MERCURY: NEUTRAL,
		VENUS:   ENEMY,
		SATURN:  ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	MOON: {
		SUN:     FRIEND,
		MERCURY: FRIEND,
		MARS:    NEUTRAL,
		JUPITER: NEUTRAL,
		VENUS:   NEUTRAL,
		SATURN:  NEUTRAL,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	MARS: {
		SUN:     FRIEND,
		MOON:    FRIEND,
		JUPITER: FRIEND,
		VENUS:   NEUTRAL,
		SATURN:  NEUTRAL,
		MERCURY: ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	MERCURY: {
		SUN:     FRIEND,
		VENUS:   FRIEND,
		MARS:    NEUTRAL,
		JUPITER: NEUTRAL,
		SATURN:  NEUTRAL,
		MOON:    ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	JUPITER: {
		SUN:     FRIEND,
		MOON:    FRIEND,
		MARS:    FRIEND,
		SATURN:  NEUTRAL,
		MERCURY: ENEMY,
		VENUS:   ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	VENUS: {
		MERCURY: FRIEND,
		SATURN:  FRIEND,
		MARS:    NEUTRAL,
		JUPITER: NEUTRAL,
		SUN:     ENEMY,
		MOON:    ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	SATURN: {
		MERCURY: FRIEND,
		VENUS:   FRIEND,
		JUPITER: NEUTRAL,
		SUN:     ENEMY,
		MOON:    ENEMY,
		MARS:    ENEMY,
		RAHU:    ENEMY,
		// Note: Ketu is missing from the table in kb/graha_maitri_chakra.md for Saturn.
	},
	RAHU: {
		KETU:    FRIEND,
		SUN:     ENEMY,
		MOON:    ENEMY,
		MARS:    ENEMY,
		MERCURY: ENEMY,
		JUPITER: ENEMY,
		VENUS:   ENEMY,
		SATURN:  ENEMY,
	},
	KETU: {
		RAHU:    FRIEND,
		SUN:     ENEMY,
		MOON:    ENEMY,
		MARS:    ENEMY,
		MERCURY: ENEMY,
		JUPITER: ENEMY,
		VENUS:   ENEMY,
		SATURN:  ENEMY,
	},
}

// GetGrahaMaitri returns the permanent relationship of the basePlanet towards the targetPlanet.
// It uses the Narpatijayacharya Planetary Friendship Table logic.
func GetGrahaMaitri(basePlanet, targetPlanet string) (string, error) {
	baseMap, ok := grahaMaitriChakra[basePlanet]
	if !ok {
		return "", fmt.Errorf("base planet %s not covered in Graha Maitri Chakra", basePlanet)
	}

	relationship, ok := baseMap[targetPlanet]
	if !ok {
		if basePlanet == targetPlanet {
			return "Self", nil
		}
		return "", fmt.Errorf("relationship from %s to %s not defined in the table", basePlanet, targetPlanet)
	}

	return relationship, nil
}
