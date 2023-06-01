// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 2
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 3
		case r == 37: // ['%','%']
			return 3
		case r == 40: // ['(','(']
			return 1
		case r == 41: // [')',')']
			return 1
		case r == 42: // ['*','*']
			return 3
		case r == 43: // ['+','+']
			return 4
		case r == 44: // [',',',']
			return 5
		case r == 45: // ['-','-']
			return 6
		case r == 46: // ['.','.']
			return 7
		case r == 47: // ['/','/']
			return 3
		case r == 48: // ['0','0']
			return 8
		case 49 <= r && r <= 57: // ['1','9']
			return 9
		case r == 58: // [':',':']
			return 3
		case r == 64: // ['@','@']
			return 3
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case r == 91: // ['[','[']
			return 10
		case r == 93: // [']',']']
			return 11
		case r == 95: // ['_','_']
			return 12
		case r == 97: // ['a','a']
			return 3
		case r == 98: // ['b','b']
			return 13
		case r == 99: // ['c','c']
			return 14
		case 100 <= r && r <= 102: // ['d','f']
			return 3
		case r == 103: // ['g','g']
			return 15
		case 104 <= r && r <= 108: // ['h','l']
			return 3
		case r == 109: // ['m','m']
			return 16
		case 110 <= r && r <= 122: // ['n','z']
			return 3
		case r == 123: // ['{','{']
			return 1
		case r == 125: // ['}','}']
			return 1
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case r == 48: // ['0','0']
			return 19
		case 49 <= r && r <= 57: // ['1','9']
			return 20
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 21
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 22
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 9
		case r == 120: // ['x','x']
			return 23
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 9
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case r == 58: // [':',':']
			return 24
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 120: // ['a','x']
			return 17
		case r == 121: // ['y','y']
			return 25
		case r == 122: // ['z','z']
			return 17
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 113: // ['a','q']
			return 17
		case r == 114: // ['r','r']
			return 26
		case 115 <= r && r <= 122: // ['s','z']
			return 17
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 110: // ['a','n']
			return 17
		case r == 111: // ['o','o']
			return 27
		case 112 <= r && r <= 122: // ['p','z']
			return 17
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 104: // ['a','h']
			return 17
		case r == 105: // ['i','i']
			return 28
		case 106 <= r && r <= 122: // ['j','z']
			return 17
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 20
		case r == 120: // ['x','x']
			return 29
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 20
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 21
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 30
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 31
		case r == 63: // ['?','?']
			return 32
		case 65 <= r && r <= 70: // ['A','F']
			return 32
		case 97 <= r && r <= 102: // ['a','f']
			return 32
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 100: // ['a','d']
			return 17
		case r == 101: // ['e','e']
			return 33
		case 102 <= r && r <= 122: // ['f','z']
			return 17
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 113: // ['a','q']
			return 17
		case r == 114: // ['r','r']
			return 34
		case 115 <= r && r <= 122: // ['s','z']
			return 17
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 109: // ['a','m']
			return 17
		case r == 110: // ['n','n']
			return 35
		case 111 <= r && r <= 122: // ['o','z']
			return 17
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case r == 63: // ['?','?']
			return 37
		case 65 <= r && r <= 70: // ['A','F']
			return 37
		case 97 <= r && r <= 102: // ['a','f']
			return 37
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 31
		case r == 63: // ['?','?']
			return 32
		case 65 <= r && r <= 70: // ['A','F']
			return 32
		case 97 <= r && r <= 102: // ['a','f']
			return 32
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 31
		case r == 63: // ['?','?']
			return 32
		case 65 <= r && r <= 70: // ['A','F']
			return 32
		case 97 <= r && r <= 102: // ['a','f']
			return 32
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case r == 97: // ['a','a']
			return 38
		case 98 <= r && r <= 122: // ['b','z']
			return 17
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 110: // ['a','n']
			return 17
		case r == 111: // ['o','o']
			return 39
		case 112 <= r && r <= 122: // ['p','z']
			return 17
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 116: // ['a','t']
			return 17
		case r == 117: // ['u','u']
			return 40
		case 118 <= r && r <= 122: // ['v','z']
			return 17
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case r == 63: // ['?','?']
			return 37
		case 65 <= r && r <= 70: // ['A','F']
			return 37
		case 97 <= r && r <= 102: // ['a','f']
			return 37
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case r == 63: // ['?','?']
			return 37
		case 65 <= r && r <= 70: // ['A','F']
			return 37
		case 97 <= r && r <= 102: // ['a','f']
			return 37
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 115: // ['a','s']
			return 17
		case r == 116: // ['t','t']
			return 41
		case 117 <= r && r <= 122: // ['u','z']
			return 17
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 116: // ['a','t']
			return 17
		case r == 117: // ['u','u']
			return 42
		case 118 <= r && r <= 122: // ['v','z']
			return 17
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 115: // ['a','s']
			return 17
		case r == 116: // ['t','t']
			return 43
		case 117 <= r && r <= 122: // ['u','z']
			return 17
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 100: // ['a','d']
			return 17
		case r == 101: // ['e','e']
			return 44
		case 102 <= r && r <= 122: // ['f','z']
			return 17
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 115: // ['a','s']
			return 17
		case r == 116: // ['t','t']
			return 45
		case 117 <= r && r <= 122: // ['u','z']
			return 17
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 100: // ['a','d']
			return 17
		case r == 101: // ['e','e']
			return 46
		case 102 <= r && r <= 122: // ['f','z']
			return 17
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 99: // ['a','c']
			return 17
		case r == 100: // ['d','d']
			return 47
		case 101 <= r && r <= 122: // ['e','z']
			return 17
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 104: // ['a','h']
			return 17
		case r == 105: // ['i','i']
			return 48
		case 106 <= r && r <= 122: // ['j','z']
			return 17
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 114: // ['a','r']
			return 17
		case r == 115: // ['s','s']
			return 49
		case 116 <= r && r <= 122: // ['t','z']
			return 17
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 109: // ['a','m']
			return 17
		case r == 110: // ['n','n']
			return 50
		case 111 <= r && r <= 122: // ['o','z']
			return 17
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 100: // ['a','d']
			return 17
		case r == 101: // ['e','e']
			return 51
		case 102 <= r && r <= 122: // ['f','z']
			return 17
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 17
		case r == 37: // ['%','%']
			return 17
		case r == 42: // ['*','*']
			return 17
		case r == 45: // ['-','-']
			return 17
		case r == 46: // ['.','.']
			return 17
		case r == 47: // ['/','/']
			return 17
		case 48 <= r && r <= 57: // ['0','9']
			return 18
		case r == 58: // [':',':']
			return 17
		case r == 64: // ['@','@']
			return 17
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 95: // ['_','_']
			return 17
		case 97 <= r && r <= 122: // ['a','z']
			return 17
		}
		return NoState
	},
}
