package geoget

// * PUBLIC * //

// Permute calls f with each permutation of cities.
func Permute(cities []*City) []*RoadMap {
	roadCH := make(chan *RoadMap)
	defer close(roadCH)

	go permute(roadCH, sliceToMap(cities), 0)

	roadmap := []*RoadMap{}

	fac := factorial(len(cities))

	for len(roadmap) < fac {
		roadmap = append(roadmap, <-roadCH)
	}

	return roadmap
}

// * PRIVATE * //

// Permute the values.
func permute(roadCH chan<- *RoadMap, c map[int]*City, i int) {
	if i > len(c) {
		roadCH <- NewRoadMap(c)
		return
	}
	permute(roadCH, c, i+1)

	for j := i + 1; j < len(c); j++ {
		c[i], c[j] = c[j], c[i]
		permute(roadCH, c, i+1)
		c[i], c[j] = c[j], c[i]
	}
}

func sliceToMap(slice []*City) map[int]*City {
	m := map[int]*City{}
	for i, c := range slice {
		m[i] = c
	}
	return m
}

func factorial(n int) int {
	var fact = 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
