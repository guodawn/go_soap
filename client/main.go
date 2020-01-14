package main

import (
    "fmt"

    soap "github.com/guodawn/go-soap/pkg/myservice"
)

func main() {
    c := soap.NewCalculatorSoap("", false, nil)
    r, err := c.Add(&soap.Add{
        IntA: 2,
        IntB: 3,
    })

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(r.AddResult)
}

