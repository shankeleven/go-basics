package main

import "fmt"
import "net/http"


/*


http.Handler Interface: The fundamental interface in Go's HTTP ecosystem.
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

http.HandlerFunc: A function type that implements the Handler interface.
type HandlerFunc func(ResponseWriter, *Request)

http.Request: Represents an HTTP request received by a server.
type Request struct {
    Method string
    URL *url.URL
    Header Header
    Body io.ReadCloser
    // ... many other fields
}

http.ResponseWriter: Interface used to construct HTTP responses.
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}

*/




/*
from what i understand
ListenAndServe accepts connections , tcp connections on the specified ports , reads the bytes as per http and makes the req, responseWriter
→ builds Request + ResponseWriter
→ calls a mux
→ mux selects a handler
→ handler writes to the response
→ bytes go back to the client
*/



/*
As soon as the first byte of the response is written , the header is sent and cannot be edited
as a http response on wire would look like this
HTTP/1.1 200 OK\r\n
Content-Type: text/plain\r\n
Content-Length: 5\r\n
\r\n
hello


and http.ResponseWriter is not a response object.
It is a controlled stream writer with rules:
It buffers headers until the first body byte
It tracks whether headers are already sent
It enforces HTTP correctness
If you didn’t call WriteHeader
Go assumes 200 OK
It sends headers immediately
Then streams the body


*/


/*
Why Go doesn’t “just buffer everything”
Because buffering everything would:
Blow memory on large responses
Prevent streaming (files, video, SSE)
Delay first byte (bad latency)

If required we can explicitly ask for the buffer

buf := bytes.NewBuffer(nil)
fmt.Fprint(buf, "hello")

w.Header().Set("X-Test", "yes")
w.WriteHeader(200)
w.Write(buf.Bytes())



this could be useful in cases , where for example, if something not computed successfully we would need to send a different status code
and headers depend on the body
Content-Length
ETag
Digest


Streaming can’t do this unless:
You precompute
Or you buffer

but else we just do Transfer-Encoding: chunked when streaming and after http1.1 , all works fine



*/


func homelander (w http.ResponseWriter, r *http.Request){
	w.Write([]byte(fmt.Sprintf("this is the home we talked about:%d:",5)))
}

func methodhandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		fmt.Println("getting contacts");
	}
}

func main(){

	fmt.Println("Staring...")
	mux:= http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hello hello")
		w.Write([]byte("Pranaam"))  // both do the same work and are equivalent
	})

	mux.HandleFunc("/writer", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"write write" );
		w.Write([]byte(" wrong"))
	})
	// basic routing with serverMux
	mux.HandleFunc("/home", homelander);

	// routing using methods
	mux.HandleFunc("func", methodhandler);




	server:= &http.Server{
		Addr: ":8080",
		Handler: mux,

		/*
These exits because real servers have policies, not just handlers.

http.Server is where behavioral constraints live:
Timeouts
Resource limits
TLS configuration
Connection lifecycle rules
Graceful shutdown hooks


read and write timeouts helps save the server from clients who never finish sending headers and handlers that accidentally block forever
header size limits This caps memory usage per request.
Without it:
A client can send massive headers
You allocate memory before even hitting your handler


Once you own the *http.Server, you can do this:
server.Shutdown(ctx)


Which means:
Stop accepting new connections
Let in-flight requests finish
Close cleanly

You cannot do this cleanly with http.ListenAndServe alone, because you never held the server reference.
This matters for:
Deployments
Kubernetes
Systemd
Zero-downtime restarts
		*/
	}

	fmt.Println("Starting a server on localhost:8080")
	err:= server.ListenAndServe();
	if err!= nil{
		fmt.Print("error: ",err)
	}



}
