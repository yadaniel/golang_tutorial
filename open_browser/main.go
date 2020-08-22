package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func myHandler0(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go => %v", time.Now())
}

func myHandler1(w http.ResponseWriter, r *http.Request) {
	for {
		fmt.Fprint(w, "show some data")
		time.Sleep(1000)
	}
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	cnt := 0
	for {
		cnt += 1
		fmt.Fprintf(w, "\r%v", cnt)
		// time.Sleep(100_000_000)
		time.Sleep(100 * time.Millisecond)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

// <script type="text/javascript">print(new Date().toLocaleString());</script>
// opens printer dialog from browser

// <meta http-equiv="refresh" content="3" />
// reload page every 3 seconds

var page string = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>The HTML5 </title>
  <meta http-equiv="refresh" content="3" />
  <meta name="description" content="HTML5">
  <meta name="author" content="Test">   
</head>    
<body style="background-color:grey;">
  This is html page <br>
  <script type="text/javascript"> document.write(new Date().toLocaleString()); </script>
  <script type="text/javascript"> console.log(new Date().toLocaleString()); </script>
</body>
</html>`

func myHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", page)
	time.Sleep(time.Second)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func main() {
	http.HandleFunc("/", myHandler0)
	http.HandleFunc("/data", myHandler1)
	http.HandleFunc("/count", myHandler2)
	http.HandleFunc("/page", myHandler3)

	go func() {
		<-time.After(100 * time.Millisecond)
		err := exec.Command("explorer", "http://127.0.0.1:8080").Run()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Println("running at port localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
