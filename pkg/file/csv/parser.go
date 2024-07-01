package csv

import (
	"encoding/csv"
	"log"
	"os"
)

type Parser struct {
}

//
//func (p *Parser) ParseDir(dir string) {
//	fmt.Println(fmt.Sprintf("Scan dir %s", dir))
//
//	files, err := ioutil.ReadDir(dir)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, file := range files {
//		records := p.ParseFile(dir + "/" + file.Name())
//
//		for _, record := range records[:1] {
//			fmt.Println(record)
//		}
//	}
//}

func (p *Parser) ParseFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
