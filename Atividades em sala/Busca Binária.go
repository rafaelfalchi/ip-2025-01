package main
import "fmt"

func main() {
    var(
        vetor[1000] int
        x, e, d, m int = 0, 0, 999, 0
    )
    for i := 0; i < 1000; i++{
        vetor[i] = (i+1)
    }
    fmt.Scan(&x)
    for e <= d{
        m = (e+d)/2
        if vetor[m] == x{
            fmt.Print(m)
            break
        }else if x > vetor[m]{
            e = m+1
        }else{
            d = m-1
        }
        return -1
    }
}