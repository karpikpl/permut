package main
import (
  "bufio"
  "fmt"
  "io"
  "strconv"
  "strings"
  "os"
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
