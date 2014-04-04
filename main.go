package main

import (
    //"./stoic" -> same error as philosopher below
    "github.com/javouhey/epictetus/philosopher"
    "github.com/javouhey/epictetus/stoic"
    stoic2 "github.com/javouhey/epictetus/stoic"
    "fmt"
    _ "fmt"
    "reflect"
    //. "./philosopher"  // equivalent to python's from a import *
    // during go get, error => local import "./philosopher" in non-local package
)

func main() {
    fmt.Println(philosopher.Salutation("Panaetius"))
    fmt.Println(philosopher.Salutation(philosopher.Name("Musonius")))

    v := "Zeno"
    fmt.Println(philosopher.Salutation(philosopher.Name(v)))
    fmt.Println(philosopher.Salutation(philosopher.Name(65)), "\n")

    h := stoic.Philosopher(0)
    fmt.Printf("%s %s\n", "stoic.Quote(5)", stoic.Quote(5))
    fmt.Printf("%s %s\n", "stoic2.Quote(stoic2.Epictetus)", stoic2.Quote(stoic2.Epictetus))
    fmt.Printf("%s %s\n", "stoic.Quote(stoic.Seneca)", stoic.Quote(stoic.Seneca))
    fmt.Printf("%s %s\n", "stoic.Quote(h)", stoic.Quote(h))
    //fmt.Printf("%s %s\n", "stoic.Quote(stoic2.gavin)", stoic2.Quote(stoic2.gavin))
    var _ = reflect.TypeOf(1)

    // compile error - overflow
    //fmt.Printf(stoic.Quote(1234567891234567891234567891238912356789))

    // compile error - truncation
    //fmt.Printf(stoic.Quote(1234.56))

    fmt.Println("Change after added as submodule to seneca (fun)")
}
