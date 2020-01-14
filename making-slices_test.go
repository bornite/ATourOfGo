package main

func ExampleMain() {
    a := make([]int, 5)
    printSlice("a", a)

    b := make([]int, 0, 5)
    printSlice("b", b)

    c := b[:2]
    printSlice("c", c)

    d := c[2:5]
    printSlice("d", d)
    // Output:
    // a len=5 cap=5 [0 0 0 0 0]
    // b len=0 cap=5 []
    // c len=2 cap=5 [0 0]
    // d len=3 cap=3 [0 0 0]
}
