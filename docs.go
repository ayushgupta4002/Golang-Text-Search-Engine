package main

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// type docs struct {
// 	Documents []document `xml:"doc"`
// }

func loadDocs(dumpPath string) ([]document, error) {
	file, err := os.Open(dumpPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()
	xmlLoader := xml.NewDecoder(gzipReader)
	fmt.Println("Decoding XML")
	dumpData := struct {
		Documents []document `xml:"doc"`
	}{}

	err = xmlLoader.Decode(&dumpData)
	if err != nil {
		return nil, err
	}

	docs := dumpData.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
