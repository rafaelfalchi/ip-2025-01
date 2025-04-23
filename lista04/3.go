package main
import "fmt"

func main() {
    var vetor[10] int
    for i := 0; i < 10; i++{
        fmt.Scan(&vetor[i])
    }
    fmt.Println("Números pares:")
    soma := 0
    for i := 0; i < 10; i++{
        if vetor[i]%2 == 0{
            fmt.Printf("%d\n", vetor[i])
            soma += vetor[i]
        }
    }
    fmt.Printf("Soma dos pares: %d\n", soma)
    fmt.Println("Números ímpares:")
    quantidade := 0
    for i := 0; i < 10; i++{
        if vetor[i]%2 != 0{
            fmt.Printf("%d\n", vetor[i])
            quantidade ++
        }
    }
    fmt.Printf("Quantidade de números ímpares: %d", quantidade)
}