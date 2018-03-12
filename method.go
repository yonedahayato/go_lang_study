package main

import "fmt"

type Calc struct {atai1, atai2 int}

func Add(q Calc) int {
    return q.atai1 + q.atai2
}

func (p Calc) Add() int {
    return p.atai1 + p.atai2
}

func main() {
    q := Calc{8, 6}
    fmt.Println(Add(q))

    p := Calc{3, 2}
    fmt.Println(p.Add())
}
