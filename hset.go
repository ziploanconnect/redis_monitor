package main

import "fmt"

type KeyVal struct {
	Key string
	Val string
}

//HSET
func saveHashKind(key string, values []string) {
	//prepare to iterate over cmd and see if hash contains values
	//save each value in key val and store it as json later
	var vals []KeyVal
	i := 0
	//all evens are key, odds are value
	var lastStruct KeyVal //save lastStruct for next loop item
	for _, k := range values {

		//see if even, means last key is obtained
		if i%2 != 0 {
			lastStruct.Val = k
			vals = append(vals, lastStruct)
		} else {
			lastStruct = KeyVal{
				Key: k,
			}
		}
		i++
		if len(values) == i {
			fmt.Println(key, vals)
			break
		}
	}
}
