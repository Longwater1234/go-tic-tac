/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package player

type SymbolGame int

const (
	O SymbolGame = iota
	X
)

func (s SymbolGame) String() string {
	switch s {
	case O:
		return "O"
	case X:
		return "X"
	}
	return "unknown"
}
