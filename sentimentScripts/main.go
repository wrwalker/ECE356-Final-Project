package main

import (
	"encoding/csv"
	"fmt"
	"github.com/cdipaolo/sentiment"
	"io"
	"log"
	"os"
	"time"
)


const dir = "../datasets/"

var files = []string{dir + "hashtag_joebiden.csv", dir + "hashtag_donaldtrump.csv"}

const totalRecords = 1747800.0

var recordsProcessed = 0;
var misses = 0;

var nameList = []string{"b", "t"}

func getSentiment(tweetContent string)(retSentimentAsInt int, retSentimentAsFloat float64, ok bool) {
	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	analysis := model.SentimentAnalysis(tweetContent, sentiment.English)

	// summing method isn't as accurate
	sum := float64(0)
	//for _, word := range analysis.Words{
	//	sum += float64(word.Score)
	//}
	return int(analysis.Score), sum/float64(len(analysis.Words)), true
}

func main() {

	timeStart := time.Now()
	for dex, file := range files {
		log.Println(fmt.Sprintf("starting file: %s", file))
		columnNamesToIndex := make(map[string]int)

		// write file headers
		path := dir+"new_sentiment.csv"
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
		f.WriteString(fmt.Sprintf("tweetID,trump_or_biden,sentiment_score\n"))

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
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			tweetID := record[columnNamesToIndex["tweet_id"]]
			tweetContent := record[columnNamesToIndex["tweet"]]

			sentimentInt, _, ok := getSentiment(tweetContent)
			if ok {
				//isCorrect := "correct"
				//if sentimentInt ==1 && sentimentFloat >= 0.5{
				//	isCorrect = "correct"
				//} else {
				//	isCorrect = "wrong"
				//}
				//log.Println(fmt.Sprintf("%s %s,%s,%d,%.2f\n",isCorrect, tweetID,nameList[dex],sentimentInt, sentimentFloat))

				f.WriteString(fmt.Sprintf("%s,%s,%d\n",tweetID,nameList[dex],sentimentInt))
			}

			recordsProcessed +=1
			timePerRecord := float64(time.Since(timeStart).Milliseconds())/float64(recordsProcessed)
			eta:= timePerRecord * float64(totalRecords-recordsProcessed) / float64(time.Hour.Milliseconds())
			if recordsProcessed % 100 == 0{
				log.Println(fmt.Sprintf("%d/%d rows handled(%.2f%%) and on file:%s ETA: %.2f hours \n", recordsProcessed, int64(totalRecords), float64(recordsProcessed)/totalRecords*float64(100), file, eta))
			}
		}
	}
}