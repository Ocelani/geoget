package geoget

type City struct {
	Name   string
	Routes map[string]*Route
}

func NewCity(location string) *City {
	return &City{
		Name:   location,
		Routes: map[string]*Route{},
	}
}

func (c *City) SetRoutes(cities []*City) (done bool) {
	ch := make(chan *Route)

	go func() {
		for _, to := range cities {
			go func(to *City) {
				if c.Name == to.Name {
					return
				}
				ch <- NewRoute(c, to)
			}(to)
		}
	}()

	for {
		rt := <-ch
		rt.Log()
		c.Routes[rt.To.Name] = rt

		if len(c.Routes) >= len(cities)-1 {
			return true
		}
	}
}
