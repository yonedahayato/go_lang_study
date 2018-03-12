package main

import "fmt"

type Calc struct{atai1, atai2 int}
type Sums []Calc

func (p Calc) Add() int {
    return p.atai1 + p.atai2
}

func (s Sums) Adds() int {
    ans := 0
    for _, s := range s {
        ans += s.Add()
    }
    return ans
}

func main() {
    sums := Sums{
        {1, 3},
        {2, 4},
        {3, 5}.
    }
}
