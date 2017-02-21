package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// curl -w "total %{time_total}\n" -X POST 'http://localhost:3000/' -F "file=@/Users/hoge/test.txty"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v\n", handler.Header)

		f, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/html; charset=utf8")
		w.Write([]byte("OK!\n"))
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Print("Starting server")
	log.Fatal(http.ListenAndServe(":3100", nil))
}
