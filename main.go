package main

import (
	geoget "geoget/pkg"
	"runtime/debug"
)

func main() {
	debug.SetPanicOnFault(true)
	geoget.Run([]string{
		"Belo Horizonte, MG, Brazil",
		"Manaus, AM, Brazil",
		"Curitiba, PR, Brazil",
		"Rio de Janeiro, RJ, Brazil",
	})
}

// func printResult(query string, result *geocoder.GeocodingResult) {
// 	var (
// 		v   = result.Results[0].Locations[0]
// 		log = StartLogger()
// 	)
// 	log.Info().Str("Query", query).Send()
// 	log.Info().
// 		Str("\n1."+v.AdminArea1Type, v.AdminArea1).
// 		Str("\n3."+v.AdminArea3Type, v.AdminArea3).
// 		Str("\n4."+v.AdminArea5Type, v.AdminArea5).
// 		Float64("\n5.Lat", v.LatLng.Lat).
// 		Float64("\n6.Lng", v.LatLng.Lng).
// 		Send()
// }

// func run() {
// 	geocoder.SetAPIKey("cWcTA6ytM8gG4xymPxXJ7Biyravc22Gr")
// 	zlog := StartLogger()

// 	// Recebe um array de cidades, a primeira é a origem
// 	places := []string{
// 		"Belo Horizonte, Brazil",
// 		"Rio de Janeiro, Brazil",
// 		"Curitiba, Brazil",
// 		"Manaus, Brazil",
// 	}

// 	cities := geoget.ValidateLocations(places)

// 	t := [][]Route{}

// 	for _, p := range Perm(places) {
// 		ee := []Route{}
// 		for i := range p {
// 			if i+1 >= len(p) {
// 				continue
// 			}
// 			from := p[i]
// 			to := p[i+1]
// 			v := routes[from][to]
// 			ee = append(ee, v)
// 		}
// 		t = append(t, ee)
// 	}

// 	type Distance struct {
// 		Routes   []Route
// 		Distance float64
// 	}
// 	var distances = []Distance{}

// 	for _, rs := range t {
// 		res := Distance{Routes: rs}
// 		for _, r := range rs {
// 			res.Distance += r.Distance
// 		}
// 		distances = append(distances, res)
// 	}

// 	sort.Slice(distances, func(i, j int) bool {
// 		return distances[i].Distance > distances[j].Distance
// 	})

// 	for i, d := range (distances)[:5] {
// 		for _, r := range d.Routes {
// 			zlog.Info().
// 				Str("from", r.From).
// 				Str("to", r.To).
// 				Float64("distance", r.Distance).
// 				Str("time", r.Time).
// 				Send()
// 		}
// 		zlog.Info().Float64("total_distance", d.Distance).Msgf("%dº opção", i+1)
// 	}
// }

// type Trip struct {
// 	Routes   []*Route
// 	Distance float64
// }

// // Perm calls f with each permutation of a.
// func Perm(a []string) (xx []map[int]string) {
// 	ch := make(chan map[int]string)
// 	go perm(a, 0, ch)

// 	d := time.Second
// 	t := time.NewTimer(d)

// 	for {
// 		select {
// 		case x := <-ch:
// 			xx = append(xx, x)
// 			if !t.Stop() {
// 				<-t.C
// 			}
// 			t.Reset(d)
// 		case <-t.C:
// 			return xx
// 		}
// 	}
// }

// // Permute the values at index i to len(a)-1.
// func perm(a []string, i int, ch chan<- map[int]string) {
// 	if i > len(a) {
// 		ch <- mapTrip(a)
// 		return
// 	}

// 	// fmt.Println(a, i+1)
// 	perm(a, i+1, ch)
// 	// fmt.Println(a, i+1)

// 	for j := i + 1; j < len(a); j++ {
// 		a[i], a[j] = a[j], a[i]
// 		perm(a, i+1, ch)
// 		a[i], a[j] = a[j], a[i]
// 	}
// }

// func mapTrip(a []string) map[int]string {
// 	m := map[int]string{}
// 	for n, v := range a {
// 		m[n] = v
// 	}
// 	return m
// }
