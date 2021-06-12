  package main
 
   import (
       "fmt"
       "io/ioutil"
       "log"
       "net/http"
       "os/exec"
       test "github.com/wittymindstech/go_demo"
   )

var output string 
 
  func helloWorld(w http.ResponseWriter, r *http.Request) {
      if r.URL.Path != "/" {
              http.NotFound(w, r)
              return
      }

      switch r.Method {
      case "GET":
              // hard disk  status on mzone
              out, err := exec.Command("df -h").Output()
              output := string(out[:]) 
              for k, v := range r.URL.Query() {
                     fmt.Printf("%s: %s\n", k, v)
                    // out, err := exec.Command("ls").Output()
                     if err != nil {
                        fmt.Printf("%s", err)
                     }

      fmt.Println("Command Successfully Executed")
      fmt.Println(output)
              }
             w.Write([]byte(output))
              w.Write([]byte("Received a GET request !!! \n"))
     case "POST":
              reqBody, err := ioutil.ReadAll(r.Body)
              if err != nil {
                       log.Fatal(err)
              }

              fmt.Printf("%s\n", reqBody)
              w.Write([]byte("Received a POST request\n"))
     default:
              w.WriteHeader(http.StatusNotImplemented)
              w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
     }

  }

  func main() {
      demo :=  test.add_demo(4,5)
      http.HandleFunc("/", helloWorld)
      http.ListenAndServe(":8000", nil)
  }
