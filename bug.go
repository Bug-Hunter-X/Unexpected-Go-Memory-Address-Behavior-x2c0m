```Go
func main() {
    var i int
    fmt.Println(i) // Output: 0
    fmt.Println(&i) // Output: 0xc0000140a8

    var s string
    fmt.Println(s) // Output: 
    fmt.Println(&s) //Output: 0xc0000140b0

    var b bool
    fmt.Println(b) //Output: false
    fmt.Println(&b) //Output: 0xc0000140b8
}
```