package main

import "fmt"

func main() {

	// We can store multiple items in a map as well

	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"realname":"Clark Kent",
			"city":"Metropolis",
		},

		"Batman": map[string]string{
			"realname":"Bruce Wayne",
			"city":"Gotham City",
		},
	}

	// We can output data where the key matches Superman

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
	}

}
