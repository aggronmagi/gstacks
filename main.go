package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/aggronmagi/gstacks/stacks/ast"
	"github.com/aggronmagi/gstacks/stacks/lexer"
	"github.com/aggronmagi/gstacks/stacks/parser"
)

var cfg = struct {
	startLimit             int
	blockLimit             int
	sameStartAndBlockLimit int
	sameStartAndBlockTop   int
}{
	startLimit:             5,
	blockLimit:             5,
	sameStartAndBlockLimit: 5,
	sameStartAndBlockTop:   1,
}

func init() {
	flag.IntVar(&cfg.startLimit, "start_limit", cfg.startLimit, "same grountine start ,limit count")
	flag.IntVar(&cfg.blockLimit, "block_limit", cfg.blockLimit, "same grountine block code ,limit count")
	flag.IntVar(&cfg.sameStartAndBlockLimit, "same_limit", cfg.sameStartAndBlockLimit, "same grountine start and block,limit count")
	flag.IntVar(&cfg.sameStartAndBlockTop, "same_top", cfg.sameStartAndBlockTop, "same grountine start and block,show top count")
}

func main() {
	flag.Parse()
	f := os.Args[1]
	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(data)
	val, err := parser.NewParser().Parse(l)
	if err != nil {
		panic(err)
	}

	gs, ok := val.([]*ast.Grountine)
	if !ok {
		fmt.Printf("%#v", val)
		panic("invalid result")
	}

	fmt.Println(len(gs))

	// 相同起始的协程统计
	calcGrountines("same-start", cfg.startLimit, 0, func(m map[string][]*ast.Grountine) {
		for _, v := range gs {
			startStack := v.Stacks[len(v.Stacks)-1]
			m[startStack.File] = append(m[startStack.File], v)
		}
	})
	//
	calcGrountines("same-block", cfg.blockLimit, 0, func(m map[string][]*ast.Grountine) {
		for _, v := range gs {
			startStack := v.Stacks[0]
			m[startStack.File] = append(m[startStack.File], v)
		}
	})
	//
	calcGrountines("same-start-and-block", cfg.sameStartAndBlockLimit, cfg.sameStartAndBlockTop, func(m map[string][]*ast.Grountine) {
		for _, v := range gs {
			s := v.Stacks[0].File + " -> " + v.Stacks[len(v.Stacks)-1].File
			m[s] = append(m[s], v)
		}
	})
}

func calcGrountines(name string, limit, printTop int, deal func(map[string][]*ast.Grountine)) {
	same := make(map[string][]*ast.Grountine)
	deal(same)
	fmt.Println(name)
	rank := sortGrountines(same)
	for _, v := range rank {
		if v.count < limit {
			continue
		}
		fmt.Printf("%5d %s\n", v.count, v.name)
	}
	for k := 0; k < len(rank); k++ {
		if k+1 >= printTop {
			break
		}
		fmt.Println()
		g := rank[k].g
		fmt.Println(g.Status.Status, g.Status.Minutes)
		for _, s := range g.Stacks {
			fmt.Println(s.Func)
			fmt.Println("\t", s.File)
		}
	}
}

func sortGrountines(in map[string][]*ast.Grountine) (res []*struct {
	name  string
	count int
	g     *ast.Grountine
}) {
	for k, v := range in {
		res = append(res, &struct {
			name  string
			count int
			g     *ast.Grountine
		}{
			name:  k,
			count: len(v),
			g:     v[0],
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].count >= res[j].count
	})
	return
}
