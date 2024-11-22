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
	flag.StringVar(&query, "q", "what are best cat breeds", "search query")
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

	matchIds := idx.search(query)
	// fmt.Println("Matched IDs:", matchIds)

	for _, id := range matchIds {
		doc := documents[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}
