package main // import "github.com/alpe/ci-example-project"

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"runtime"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome! v0.0.6\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func HealthCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	httpPort := flag.String("port", "8080", "HTTP listen port")
	flag.Parse()

	glog.Info("GOMAXPROCS: ", runtime.GOMAXPROCS(-1))
	glog.Info("Port: ", *httpPort)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/healthcheck", HealthCheck)
	router.GET("/hello/:name", Hello)
	glog.Fatal(http.ListenAndServe(":8080", router))

}
