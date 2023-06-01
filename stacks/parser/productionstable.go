// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/aggronmagi/gstacks/stacks/ast"

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Start	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Start : Goroutines	<< ast.CheckGoroutines(X[0]) >>`,
		Id:         "Start",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.CheckGoroutines(X[0])
		},
	},
	ProdTabEntry{
		String: `Goroutines : empty	<< ast.NewGoroutines() >>`,
		Id:         "Goroutines",
		NTType:     2,
		Index:      2,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewGoroutines()
		},
	},
	ProdTabEntry{
		String: `Goroutines : Goroutines Goroutine	<< ast.AppendGoroutines(X[0], X[1]) >>`,
		Id:         "Goroutines",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.AppendGoroutines(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Goroutines : Goroutines tok_enter Goroutine	<< ast.AppendGoroutines(X[0], X[2]) >>`,
		Id:         "Goroutines",
		NTType:     2,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.AppendGoroutines(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Goroutine : "goroutine" tok_int "[" GStatus "]:" tok_enter GStacks	<< ast.NewGoroutine(X[1],X[3],X[6]) >>`,
		Id:         "Goroutine",
		NTType:     3,
		Index:      5,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewGoroutine(X[1],X[3],X[6])
		},
	},
	ProdTabEntry{
		String: `GStatus : tok_identifier	<< ast.NewStatus(0, X[0]) >>`,
		Id:         "GStatus",
		NTType:     4,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewStatus(0, X[0])
		},
	},
	ProdTabEntry{
		String: `GStatus : tok_identifier "," tok_int "minutes"	<< ast.NewStatus(X[2], X[0]) >>`,
		Id:         "GStatus",
		NTType:     4,
		Index:      7,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewStatus(X[2], X[0])
		},
	},
	ProdTabEntry{
		String: `GStatus : tok_identifier tok_identifier	<< ast.NewStatus(0, X[0], X[1]) >>`,
		Id:         "GStatus",
		NTType:     4,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewStatus(0, X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `GStatus : tok_identifier tok_identifier "," tok_int "minutes"	<< ast.NewStatus(X[3], X[0], X[1]) >>`,
		Id:         "GStatus",
		NTType:     4,
		Index:      9,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewStatus(X[3], X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `GStacks : empty	<< ast.NewStacks() >>`,
		Id:         "GStacks",
		NTType:     5,
		Index:      10,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewStacks()
		},
	},
	ProdTabEntry{
		String: `GStacks : GStacks GStack	<< ast.AppendStack(X[0], X[1]) >>`,
		Id:         "GStacks",
		NTType:     5,
		Index:      11,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.AppendStack(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `GStack : GFuncCall tok_identifier tok_hex tok_enter	<< ast.NewStack(X[0], X[1]) >>`,
		Id:         "GStack",
		NTType:     6,
		Index:      12,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewStack(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `GStack : GFuncCall tok_identifier tok_enter	<< ast.NewStack(X[0], X[1]) >>`,
		Id:         "GStack",
		NTType:     6,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewStack(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `GFuncCall : tok_identifier GARGS tok_enter	<< ast.NewCall(X[0]) >>`,
		Id:         "GFuncCall",
		NTType:     7,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCall(X[0])
		},
	},
	ProdTabEntry{
		String: `GFuncCall : tok_identifier tok_identifier tok_identifier GARGS tok_enter	<< ast.NewCall(X[0],X[1],X[2]) >>`,
		Id:         "GFuncCall",
		NTType:     7,
		Index:      15,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCall(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `GFuncCall : "created" "by" tok_identifier tok_identifier tok_identifier GARGS tok_enter	<< ast.NewCall(X[0],X[1],X[2],X[3],X[4]) >>`,
		Id:         "GFuncCall",
		NTType:     7,
		Index:      16,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCall(X[0],X[1],X[2],X[3],X[4])
		},
	},
	ProdTabEntry{
		String: `GFuncCall : "created" "by" tok_identifier GARGS tok_enter	<< ast.NewCall(X[0],X[1], X[2]) >>`,
		Id:         "GFuncCall",
		NTType:     7,
		Index:      17,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCall(X[0],X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `GARGS : empty	<<  >>`,
		Id:         "GARGS",
		NTType:     8,
		Index:      18,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `GARGS : GARGS tok_hex	<<  >>`,
		Id:         "GARGS",
		NTType:     8,
		Index:      19,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `GARGS : GARGS "," tok_hex	<<  >>`,
		Id:         "GARGS",
		NTType:     8,
		Index:      20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `GARGS : GARGS "..."	<<  >>`,
		Id:         "GARGS",
		NTType:     8,
		Index:      21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `GARGS : GARGS "," "..."	<<  >>`,
		Id:         "GARGS",
		NTType:     8,
		Index:      22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `GARGS : GARGS "_"	<<  >>`,
		Id:         "GARGS",
		NTType:     8,
		Index:      23,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
}