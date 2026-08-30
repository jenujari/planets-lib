# Domain glossary

Inferred from README, wiki, and source. Terms marked *(inferred)* were not previously recorded.

## Core quantities

- **Longitude** *(inferred)*: ecliptic position in degrees. Callers may pass any finite value; the longitude module wraps it into `[0, 360)` and rejects NaN / ±Inf.
- **Tithy**: lunar day, 1..30. Each tithy is a 12° slice of the moon−sun separation. 1..15 waxing (Sukla), 16..30 waning (Krishna). Invalid inputs return 0.
- **Sign (Rashi)**: one of 12 zodiac signs, 30° each, starting at 0° Aries.
- **Nakshatra**: lunar mansion. 28 names including Abhijit, each divided into 4 **padas**.
- **Pada**: quarter of a nakshatra. Irregular widths around Uttara Ashadha, Abhijit, and Shravana.
- **Navamsha (Navansh Rashi)**: D9 sign. Each sign is split into nine 3°20' segments; the starting sign depends on the element's fire/earth/air/water group.
- **Vargottama**: the planet's sign and navamsha sign are the same.

## Planet state

- **PlanetCord**: raw Swiss-Ephemeris-style coordinates plus derived fields (sign, nakshatra, DMS, retro, speed category, vedha, lordship, vargottama).
- **Graha Maitri**: permanent friendship of one planet toward another (Friend, Neutral, Enemy, Self). Table from Narpatijayacharya. Outer planets (Uranus, Neptune, Pluto) are not in the table.
- **Speed category**: traditional gati — kutil, ati-vakra, vakra, ati-mand, mand, madhyam, sama, sheeghra, ati-sheeghra.
- **Vedha**: obstruction direction (left, right, front, none) from speed category, plus a nakshatra **Vedha target**.

## Strength (Bal)

- **Uchh Bal**: exaltation strength from distance to the planet's exaltation degree.
- **Uday Bal**: rising / combustion strength from distance to the Sun and retrograde state.
- **Vakra Bal**: retrograde strength from speed vs max retro speed.
- **Kshetra Bal**: positional strength in the occupied sign, weighted by Graha Maitri to the sign lord, peaked at 15°.
- **Navansh Bal**: same idea in the navamsha segment, peaked at 100'.
