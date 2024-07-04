// This file contains types that are used in this server (global)
package types

type Tree struct {
	X      int
	Y      int
	Height int
}

type Estate struct {
	ID     string
	Length int
	Width  int
}

type Stats struct {
	Count  int
	Max    int
	Min    int
	Median int
}
