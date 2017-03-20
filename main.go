package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var fname, outFileName string
	flag.StringVar(&fname, "file", "", "file path")
	flag.StringVar(&outFileName, "out", "", "output file name")
	flag.Parse()

	parse(fname, outFileName)
}

func parse(path string, out string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatal("Error:", err)
	}

	var writer *bufio.Writer
	if out == "" {
		writer = bufio.NewWriter(os.Stdout)
	} else {
		o, err := os.Create(out)
		defer o.Close()
		if err != nil {
			log.Fatal("Error:", err)
		}
		writer = bufio.NewWriter(o)
	}
	defer writer.Flush()

	reader := csv.NewReader(f)
	isHeader := true
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error:", err)
		}
		newRecords := toMarkdownTable(records, isHeader)
		writer.Write(newRecords)
		isHeader = false
	}
}

func toMarkdownTable(records []string, isHeader bool) []byte {
	buf := make([]byte, 0, 10*len(records))
	for _, val := range records {
		// newline to '<br>'
		val = strings.Replace(val, "\n", "<br>", -1)
		buf = append(buf, "|"...)
		buf = append(buf, val...)
	}
	buf = append(buf, "|\n"...)

	if isHeader {
		for i := 0; i < len(records); i++ {
			buf = append(buf, "|:----"...)
		}
		buf = append(buf, "|\n"...)
	}

	return buf
}
