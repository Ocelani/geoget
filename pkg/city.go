package geoget

// City type is responsible for saving the city name and
// its routes to anothers cities.
type City struct {
	Name   string
	Routes map[string]*Route
}

// * PUBLIC * //

// NewCity constructor method of *City type.
func NewCity(location string) *City {
	return &City{
		Name:   location,
		Routes: map[string]*Route{},
	}
}

// SetRoutes collects the data of routes from this city to the others cities.
func (c *City) SetRoutes(cities []*City) (done bool) {
	ch := make(chan *Route)

	go func() {

		for _, to := range cities {
			go func(to *City) {

				if c.Name == to.Name {
					return
				}

				ch <- NewRoute(c, to) // Gets the route data and sends in channel
			}(to)
		}
	}()

	for {
		rt := <-ch                          // Wait until receive data in channel
		c.Routes[rt.To.Name] = rt           // Saves into this city refference
		if len(c.Routes) >= len(cities)-1 { // Returns when all data has been collected
			return true
		}
		rt.Log()
	}
}
