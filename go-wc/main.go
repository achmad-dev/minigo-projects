package main

import (
	"fmt"

	"github.com/achmad-dev/minigo-projects/go-wc/wc"
	initLog "github.com/achmad-dev/minigo-projects/internal/log"
)

func main() {
	log := initLog.InitLog()
	filepath := "./file/test.txt"
	wc := wc.NewWc(log)
	w, _ := wc.W(filepath)
	l, _ := wc.L(filepath)
	c, _ := wc.C(filepath)
	m, _ := wc.M(filepath)
	fmt.Println("w: ", w, "l: ", l, "c:", c, "m: ", m)
}
