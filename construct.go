package main

import "fmt"

type Language struct {
    Name string
    LangType string
}

func NewLanguage(name string, langType string) *Language {
    l := new(Language)
    l.Name = name
    l.LangType = langType

    return l
}

func main() {
    l := NewLanguage("Go", "Static")
    fmt.Println("名前" + l.Name)
    fmt.Println("言語" + l.LangType)
}
