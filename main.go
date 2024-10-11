// You can edit this code!
// Click here and start typing.
package main

import "jovens/artes"

func main() {
	artes.DesenhaTriangulo("#", 4)

	for _, c := range []string{"#", "*", ".", "O"} {
		artes.DesenhaLosango(c, 3)
	}
}
