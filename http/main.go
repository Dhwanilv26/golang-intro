package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {

	resp, err := http.Get("https://google.com")
	// .Get() returns a *http.Response type -> uske andar status, statusCode, proto, body, header sab hota hai
	// jo body hai, voh ReadCloser interface ko support karne vaala koi bhi type ka ho sakta hai,
	// ReadCloser interface ke andar 2 interfaces hai -> Reader,Closer
	// Reader interface ke andar Read method hai, jo byte slice accept karta hai, and no of bytes read and error return karega
	// this read method is from the net/http package and TCPConn type
	// imp point -> close connections after establishing using defer keyword

	// the read keyword needs bytes to transfer the read data, so we give the created byte slice in order to save memory because it will create a new byte slice, make keyword is used to dynamically allocate memory for the byte slice

	// use of reader interface is just to give a common read() contract to all different types of read cases like from file,terminal,http,websocket etc, just to simplify method names instead of using readfromfile(), readfromnetwork(), etc

	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	defer resp.Body.Close() // defer keyword stores everything in a stack, this will be executed at the last after the program finishes, to close the tcp connection

	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// io.Copy(os.Stdout, resp.Body) // first val -> implement the writer and 2nd value -> reader interface (dst writer, src reader) , copies from src to dest
	// os.stdout is a type of *file, file has the write method which satisfies the writer interface,
	// in go, everything is treated as a file, even terminals, so this works perfectly

}

func (logWriter) Write(bs []byte) (int, error) {
	// beauty of go interfaces, because write ko to hamne call bhi nai kiya, then also we can use it as per our needs!! GO >>>
	fmt.Println(string(bs))
	fmt.Println("bytes written", len(bs))
	return len(bs), nil
}
