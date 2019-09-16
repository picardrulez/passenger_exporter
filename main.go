package main

import (
	"net/http"

	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var NAME string = "passenger-exporter"
var VERSION string = "v1.5.4"
var PORT string = "8000"
var EXECCOMMAND string = "passenger-status"
var EXECOPTIONS string = "--show=xml"

func main() {
	testFlag := flag.Bool("t", false, "check if test")
	versionFlag := flag.Bool("v", false, "check version")
	flag.Parse()
	if *versionFlag {
		fmt.Println(NAME + " " + VERSION)
		return
	}
	if *testFlag {
		log.Info("running in test mode")
		EXECCOMMAND = "xml-shitter"
		EXECOPTIONS = ""
	}
	recordMetrics()
	//start http server
	http.Handle("/metrics", promhttp.Handler())
	log.Info(NAME + " " + VERSION + " started on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
