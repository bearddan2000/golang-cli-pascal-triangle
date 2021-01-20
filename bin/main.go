//go 1.10.4

package main
import "fmt"

 func pascalTriangle(n int) {
     for i := 0; i < n; i++ {
      var c = 1
      for k:=0; k<i+1; k++ {
        fmt.Printf("%d, ", c)
        c = c * (i-k)/(k+1)
      }
      fmt.Println()
  }
 }

func main() {
    N := 10
	fmt.Printf("[INPUT] %d\n", N)
    fmt.Println("[OUTPUT]")
    pascalTriangle(N)
}
