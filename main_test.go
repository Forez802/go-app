package main

import "testing"

func TestSuma(t *testing.T) {
    if Suma(2, 3) != 5 {
        t.Fail()
    }
}
