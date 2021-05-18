package geoget

import "fmt"

// RoadMap object saves the city objects
// and the corresponding routes it will need
// to take. It preserves the sequences order,
// and calculates the total distance of the routes.
type RoadMap struct {
	Cities   map[int]*City
	Routes   map[int]*Route
	Distance float64
}

// * PUBLIC * //

// NewRoadMap constructor method of *RoadMap type.
func NewRoadMap(cities map[int]*City) *RoadMap {
	m := &RoadMap{
		Cities: cities,
		Routes: map[int]*Route{},
	}
	m.setRoutes()
	m.setDistance()
	return m
}

// Log this *RoadMap info.
func (m *RoadMap) Log(rank int) {
	fmt.Println()
	zlog.Info().
		Float64("distance", m.Distance).
		Msgf("%dº", rank)

	for _, r := range m.Routes {
		r.Log()
	}
}

// * PRIVATE * //

// setRoutes gets each route of this *RoadMap and
// saves into this *RoadMap pointer object.
func (m *RoadMap) setRoutes() {
	for i := range m.Cities {
		var to string
		from := m.Cities[i].Routes

		if i+1 >= len(m.Cities) {
			to = m.Cities[0].Name
		} else {
			to = m.Cities[i+1].Name
		}
		m.Routes[i] = from[to]
	}
}

// setDistance saves the total distance
// of this *RoadMap object, based on its mapped *Routes.
func (m *RoadMap) setDistance() {
	for _, r := range m.Routes {
		m.Distance += r.Distance
	}
}
