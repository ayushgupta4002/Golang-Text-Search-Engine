package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()
	fmt.Println("dumpPath:", dumpPath)
	fmt.Println("query:", query)
	documents, err := loadDocs(dumpPath)
	start := time.Now()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(documents), time.Since(start))

	// indexing the document
	idx := make(index)
	idx.add(documents)

	// fmt.Println("docs:", documents)
}
