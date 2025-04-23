package main
import "fmt"

func main() {
    var vetor[10] int
    cont := 0
    for i := 0; i < 10; i++{
        fmt.Scan(&vetor[i])
    }
    for i := 0; i < 10; i++{
        if vetor[i] > 50{
            fmt.Printf("%d: %d° elemento\n", vetor[i], i+1)
            cont = 1
        }
    }
    if cont == 0{
        fmt.Print("Nenhum elemento é superior a 50")
    }
}