package geoget

import (
	"github.com/jasonwinn/geocoder"
)

type Route struct {
	From     *City
	To       *City
	Time     string
	Distance float64
}

func NewRoute(from *City, to *City) *Route {
	r := &Route{From: from, To: to}
	r.setDirections()
	return r
}

func (r *Route) Log() {
	_zlog.Info().
		Str("from", r.From.Name).
		Str("to", r.To.Name).
		Float64("distance", r.Distance).
		Send()
}

func (r *Route) setDirections() {
	var (
		from = r.From.Name
		to   = r.To.Name
	)
	direct := geocoder.NewDirections(from, []string{to})
	direct.RouteType = "shortest"

	result, err := direct.Get()

	if err != nil {
		_zlog.Panic().
			Str("from", from).
			Str("to", to).
			Err(err).
			Send()
	}

	res := result.Route
	r.Time = res.FormattedTime
	r.Distance = res.Distance
}
