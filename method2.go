package main

import "fmt"

type Calc struct {atai1, atai2 int}

func (p Calc) Add() int {
    return p.atai1 + p.atai2
}

func (p Calc) Sub() int {
    return p.atai1 - p.atai2
}

func (p Calc) Multi() int {
    return p.atai1 * p.atai2
}

func (p Calc) Div() int {
    if p.atai2 == 0 {
        return 0
    }
    return p.atai1 / p.atai2
}

func main() {
    p := Calc{5, 2}
    fmt.Println(p.Add())
    fmt.Println(p.Sub())
    fmt.Println(p.Multi())
    fmt.Println(p.Div())

    q := Calc{5, 2}
    fmt.Println(q.Add())
}
