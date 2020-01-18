package main

func ExampleMain() {
    var s []int
    printSlice(s)

    // append works on nil slices.
    s = append(s, 0)
    printSlice(s)

    // The slice grows as needed.
    s = append(s, 1)
    printSlice(s)

    // We can add more than one element at a time.
    s = append(s, 2, 3, 4)
    printSlice(s)

    // Output:
    // len=0 cap=0 []
    // len=1 cap=1 [0]
    // len=2 cap=2 [0 1]
    // len=5 cap=6 [0 1 2 3 4]
}
