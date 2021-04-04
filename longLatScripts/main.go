package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const dir = "./datasets/"

var files = []string{dir + "hashtag_joebiden.csv", dir + "hashtag_donaldtrump.csv"}

type apiRes struct {
	results []results `json:"results"`
}

type results struct {
	state_name  string `json:"state_name"`
	county_name string `json:"county_name"`
}

func getCountyAndState(lat, long string) (county, state string) {
	strQuery := fmt.Sprintf("https://geo.fcc.gov/api/census/area?lat=%s&lon=%s&format=json", lat, long)
	res, err := http.Get(strQuery)
	if err != nil {
		log.Println("err making req: " + err.Error())
	}
	defer res.Body.Close()

	apiResponse := &apiRes{}

	bod := res.Body

	//err = json.NewDecoder(bod).Decode(apiResponse)
	//if err!= nil {
	//	log.Println("err decoding res: " + err.Error())
	//}

	format, _ := io.ReadAll(bod)
	fmt.Printf("%v", format)

	fmt.Printf("%v", apiResponse)
	return "", ""
}

// decoding a large json will actually be slower than converting to a string and parsing the string
func simpleParse() {

}

func main() {
	for _, file := range files {
		columnNamesToIndex := make(map[string]int)

		csvFile, err := os.Open(file)
		if err != nil {
			log.Fatal("couldn't open file: " + file)
			return
		}

		r := csv.NewReader(csvFile)

		// load up columnNamesToIndex
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for i, columnName := range record {
			columnNamesToIndex[columnName] = i
		}

		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			tweetID := record[columnNamesToIndex["tweet_id"]]
			lat := record[columnNamesToIndex["lat"]]
			long := record[columnNamesToIndex["long"]]

			if lat != "" && long != "" {
				getCountyAndState(lat, long)
			}
			fmt.Print(tweetID + lat + long)
		}
	}
}
