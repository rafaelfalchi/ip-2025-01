package main
import "fmt"

func main() {
    var ( 
        a[10] int
        b[5] int
        par[10] int
        impar[10] int
        soma, contpar, contimpar int
    )
    fmt.Println("Primeiro vetor")
    for i := 0; i < 10; i++{
        fmt.Printf("%d° termo: ", i+1)
        fmt.Scan(&a[i])
    }
    fmt.Println("Segundo vetor")
    for i := 0; i < 5; i++{
        fmt.Printf("%d° termo: ", i+1)
        fmt.Scan(&b[i])
        soma += b[i]
    }
    for i := 0; i < 10; i++{
        if a[i]%2 == 0{
            par[contpar] = a[i] + soma
            contpar ++
        }else{
            impar[contimpar] = a[i] + soma
            contimpar ++
        }
    }
    fmt.Printf("Primeiro vetor resultante: %d\n", par)
    fmt.Printf("Segundo vetor resultante: %d", impar)
}