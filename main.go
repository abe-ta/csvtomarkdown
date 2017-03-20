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

	var (
		fname       string
		outFileName string
		headerFlg   bool
	)

	flag.StringVar(&fname, "file", "", "CSV file path")
	flag.StringVar(&fname, "f", "", "CSV file path")
	flag.StringVar(&outFileName, "out", "", "Output file name")
	flag.StringVar(&outFileName, "o", "", "Output file name")
	flag.BoolVar(&headerFlg, "header", false, "Use first line as headers")
	flag.Parse()

	parse(fname, outFileName, headerFlg)
}

func parse(path string, out string, headerFlg bool) {
	// csv
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatal("Error:", err)
	}

	// writer
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

	// read csv
	reader := csv.NewReader(f)
	isHeader := true
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error:", err)
		}
		newRecords := toMarkdownTable(records, isHeader, headerFlg)
		writer.Write(newRecords)
		isHeader = false
	}
}

// TODO: move it to lib
func toMarkdownTable(records []string, isHeader bool, headerFlg bool) []byte {
	buf := make([]byte, 0, 10*len(records))
	for _, val := range records {
		// newline to '<br>'
		val = strings.Replace(val, "\n", "<br>", -1)
		buf = append(buf, "|"...)
		buf = append(buf, val...)
	}
	buf = append(buf, "|\n"...)

	if isHeader && headerFlg {
		for i := 0; i < len(records); i++ {
			buf = append(buf, "|:----"...)
		}
		buf = append(buf, "|\n"...)
	}

	return buf
}
