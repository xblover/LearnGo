package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// m := make([][]int, 1000)
	// for i := 0; i < 1000; i++ {
	// 	m[i] = make([]int, 1000)
	// 	for j := 0; j < 1000; j++ {
	// 		m[i][j] = 0
	// 	}
	// }
	// generateMatrixTXTFile(m)

	// testChannel()

	r := setupRouter()
	r.Run(":8080")

}
