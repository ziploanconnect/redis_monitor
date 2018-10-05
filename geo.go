package main

import "fmt"

type GeoVal struct {
	Long   string
	Lat    string
	Member string
}

//GETADD
func saveGeoKind(key string, values []string) {
	//prepare to iterate over cmd and see if hash contains values
	//save each value in key val and store it as json later
	var vals []GeoVal
	i := 3
	//all evens are key, odds are value
	var lastStruct GeoVal //save lastStruct for next loop item
	var first bool        //hacks to check if lat is saved
	for _, k := range values {

		if i%3 == 0 {
			lastStruct = GeoVal{
				Long: k,
			}
			first = true
		} else {

			if first {
				lastStruct.Lat = k
				first = false
			} else {
				lastStruct.Member = k
				vals = append(vals, lastStruct)
			}

		}
		i++
		if len(values) == i-3 {
			fmt.Println(key, vals)
			break
		}
	}

}
