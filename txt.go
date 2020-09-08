package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func generateMatrixTXTFile(m [][]int) error {

	f, err := os.Create("text.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()

	q := make([]string, 1000)
	w := bufio.NewWriter(f)
	for _, v := range m {
		for i := 0; i < len(v); i++ {
			q[i] = strconv.Itoa(v[i])
		}
		fmt.Fprintln(w, strings.Join(q, "	"))
	}
	return w.Flush()
}
