package main

import (
	"fmt"
	//"strconv"
)

// type TZ int

// func (d Day) String() string {
// 	weekStr := []string{"MO", "TU", "WE", "TH", "FI", "SA", "SU"}
// 	return weekStr[d]
// }

// const (
// 	MO = iota
// 	TU
// 	WE
// 	TH
// 	FI
// 	SA
// 	SU
// )

// func main() {
// 	fmt.Println(TU)
// 	var i Day
// 	i = WE
// 	fmt.Println(i)
// }

type TZ int

const (
	UTC = iota
	PBC
)

func (t TZ) String() string {
	ZoneMap := map[TZ]string{UTC: "Universal Greenwich time", PBC: "CHINA"}
	return ZoneMap[t]
}

func main() {
	var t TZ
	t = 0
	fmt.Println(t)
}
