package main
import "fmt"

func jumDigit(n int, genap, ganjil *int) {
    if n == 0 {
        return
    }
    digit := n % 10
    if digit%2 == 0 {
        *genap++
    } else {
        *ganjil++
    }
    jumDigit(n/10, genap, ganjil)
}

func main() {
    var num int
    fmt.Scan(&num)
    
    var genap, ganjil int
    jumDigit(num, &genap, &ganjil)
    fmt.Printf("Genap: %d, Ganjil: %d", genap, ganjil)
}
