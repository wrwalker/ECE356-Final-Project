package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
)

const dir = "./datasets/"

var files = []string{dir + "hashtag_joebiden.csv", dir + "hashtag_donaldtrump.csv"}

const waitTime = time.Millisecond * 20
const totalRecords = 1747800.0

var recordsProcessed = 0;
var misses = 0;

//time to completion ~=
//1747800*(100/1000)/60/60 ~=5.5h

type results struct {
	state_name  string
	county_name string
}

func getCountyAndState(lat, long string) (county, state string, ok bool) {
	strQuery := fmt.Sprintf("https://geo.fcc.gov/api/census/area?lat=%s&lon=%s&format=json", lat, long)
	res, err := http.Get(strQuery)
	if err != nil {
		log.Println("err making req: " + err.Error())
	}
	defer res.Body.Close()

	apiResponse := map[string][]interface{}{}

	bod := res.Body

	err = json.NewDecoder(bod).Decode(&apiResponse)
	if err!= nil {
		//log.Println("err decoding res: " + err.Error() + "\n")
	}

	if len(apiResponse["results"]) <1{
		misses +=1
		return "","",false
	}

	respMap := apiResponse["results"][0]

	respStruct := map[string]interface{}{}
	v := reflect.ValueOf(respMap)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)

			respStruct[key.Interface().(string)] = strct.Interface()
		}
	}


	//fmt.Printf("%v %v", respStruct["county_name"].(string), respStruct["state_name"].(string))
	return respStruct["county_name"].(string), respStruct["state_name"].(string), true
}

// decoding a large json wirespMapll actually be slower than converting to a string and parsing the string
func simpleParse() {

}

func main() {
	for _, file := range files {
		log.Println(fmt.Sprintf("starting file: %s", file))
		columnNamesToIndex := make(map[string]int)

		// write file headers
		path := dir+"new_lat_lon.csv"
		err := os.Remove(path)
		if err != nil {
			//log.Fatal("couldn't delete file")
			//return
		}

		f, err := os.Create(path)
		if err != nil {
			log.Fatal("couldn't write file")
			return
		}
		defer f.Close()
		f.WriteString(fmt.Sprintf("lat,lon,tweetID,county_name,state_name\n"))

		// open file
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
			time.Sleep(waitTime)
			// Read each record from csv
			_, err := r.Read()
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
				county, state, ok := getCountyAndState(lat, long)
				if ok {
					//fmt.Printf("%s %s %s\n", county, state, tweetID)
					f.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s\n",lat,long,tweetID,county,state))
				}
			}
			recordsProcessed +=1
			if recordsProcessed % 100 == 0{
				log.Println(fmt.Sprintf("%d/%v rows handled(%v%%) with %d misses and on file:%s\n", recordsProcessed, totalRecords, float64(recordsProcessed)/totalRecords*float64(100), misses, file))
			}
		}
	}
}
