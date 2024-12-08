package main

type Point struct {
	X int
	Y int
}

type Antenna struct {
	P    Point
	Freq rune
}

func (p Point) DistanceTo(x Point) (int, int) {
	return x.X - p.X, x.Y - p.Y
}

func (p Point) IsAntinodeOf(a, b Antenna) bool {
	if p.IsHarmonicAntinodeOf(a, b) {
		ax, ay := p.DistanceTo(a.P)
		bx, by := p.DistanceTo(b.P)

		if (ax*2 == bx || bx*2 == ax) && (ay*2 == by || by*2 == ay) {
			return true
		}
	}

	return false
}

func (p Point) IsHarmonicAntinodeOf(a, b Antenna) bool {
	if a.Freq != b.Freq {
		return false
	}

	slope := float64(a.P.Y-b.P.Y) / float64(a.P.X-b.P.X)
	testSlope := float64(p.Y-a.P.Y) / float64(p.X-a.P.X)

	return testSlope == slope
}

func getAntennas(input []string) []Antenna {
	antennas := make([]Antenna, 0)
	for i, l := range input {
		for j, r := range l {
			if r != '.' {
				antennas = append(antennas, Antenna{
					P: Point{
						X: i,
						Y: j,
					},
					Freq: r,
				})
			}
		}
	}
	return antennas
}

func checkPoint(p Point, antennas []Antenna, harmonic bool) bool {
	for _, antenna := range antennas {
		for _, atest := range antennas {
			if atest != antenna {
				if harmonic && p.IsHarmonicAntinodeOf(antenna, atest) {
					return true
				} else if p.IsAntinodeOf(antenna, atest) {
					return true
				}
			}
		}
	}
	return false
}

func calcAntinodes(antennas []Antenna, mi, mj int, harmonic bool) []Point {
	antinodes := make([]Point, 0)

	for i := range mi {
		for j := range mj {
			p := Point{
				X: i,
				Y: j,
			}

			if checkPoint(p, antennas, harmonic) {
				antinodes = append(antinodes, p)
			}
		}
	}
	return antinodes
}

func Day8() (int, int) {
	input := getFileLines("./day8.txt")
	antennas := getAntennas(input)
	ilen := len(input)
	jlen := len(input[0])
	antinodes := calcAntinodes(antennas, ilen, jlen, false)
	harmonicAntinodes := calcAntinodes(antennas, ilen, jlen, true)

	return len(antinodes), len(harmonicAntinodes)
}
