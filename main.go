package main

import (
	geoget "geoget/pkg"
)

func main() {
	geoget.Run([]string{
		"Belo Horizonte, MG, Brazil",
		"Manaus, AM, Brazil",
		"Curitiba, PR, Brazil",
		"Rio de Janeiro, RJ, Brazil",
	})
}
