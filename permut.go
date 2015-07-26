package main
import (
  "bufio"
  "fmt"
  "io"
  "strconv"
  "strings"
  "os"
  "math"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func Factorial(n int) int {
  result :=1
  for i:=1 ; i<=n; i++ {
    result *= i
  }

  return result
}

func FindPermut(n int, k int) []int {
  // find k'th element of n-size permutation

  // there are n! permutations
  // each number is on first position (n-1)! times
  totalP := Factorial(n)
// 19 -> 20
// 4! -> each can be 6 times
// 20/6 -> 3.x -> 4
// 6 / 3 ->  2
 result := make([]int, n)
 result[0] = int(math.Ceil(float64(k+1.0)/ ( float64(totalP) / float64(n))))
 result[1] = totalP
  return result
}

func main() {
    argsWithoutProg := os.Args[1:]
    var input []int
    var err error

    if(len(argsWithoutProg) == 0){
      input, err = ReadInts(os.Stdin)
    } else {
      input, err = ReadInts(strings.NewReader(argsWithoutProg[0]));
    }

    fmt.Println(input, err)
}
