package main
import (
"fmt"
"net/http"
"os"
"io/ioutil")

func main () {
for _ , url := range os.Args[1:] {
resp ,err := http.Get(url)
if err != nil {
fmt.Fprintf(os.Stderr , "fetch error :%v\n",err )
os.Exit(1)

}
b ,err :=  ioutil.ReadAll(resp.Body)
resp.Body.Close()
if err != nil {
fmt.Fprintf(os.Stderr , "error in fetchib %s %v \n",url , err)
os.Exit(1)
}
fmt.Printf("%s",b)
}
}
