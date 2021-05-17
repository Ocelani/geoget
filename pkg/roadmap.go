package geoget

import "fmt"

type RoadMap struct {
	Cities   map[int]*City
	Routes   map[int]*Route
	Distance float64
}

func NewRoadMap(cities map[int]*City) *RoadMap {
	m := &RoadMap{
		Cities: cities,
		Routes: map[int]*Route{},
	}
	m.setRoutes()
	m.setDistance()
	return m
}

func (m *RoadMap) Log(rank int) {
	fmt.Println()
	_zlog.Info().
		Float64("distance", m.Distance).
		Msgf("%dº", rank)

	for _, r := range m.Routes {
		r.Log()
	}
}

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

func (m *RoadMap) setDistance() {
	for _, r := range m.Routes {
		m.Distance += r.Distance
	}
}
