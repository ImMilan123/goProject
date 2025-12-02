package main
import ("fmt"
	"net/http"
	"io/ioutil"
)

func main() {
    resp, err := http.Get("https://donkeyontheedge.com/mahir/")
    if err != nil {
    }
    defer resp.Body.Close()
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)
    if err != nil {
    }
}