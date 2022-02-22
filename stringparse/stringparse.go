// https://www.rapid7.com/blog/post/2016/08/04/build-a-simple-cli-tool-with-golang/

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse.")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}
