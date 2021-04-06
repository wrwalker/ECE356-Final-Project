package main

import (
	"encoding/csv"
	"fmt"
	"github.com/cdipaolo/sentiment"
	"io"
	"log"
	"os"
	"sync"
	"time"
)


const dir = "../datasets/"

var files = []string{dir + "hashtag_joebiden.csv", dir + "hashtag_donaldtrump.csv"}

const totalRecords = 1747800.0
const maxGoRoutines = 50

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

type tweet struct {
	sentimentScore int
	tweetID string
	person string
}

func main() {
	tweetChan := make(chan *tweet)
	exitChan := make(chan int)

	timeStart := time.Now()

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

	wg := sync.WaitGroup{}
	sem := make(chan struct{}, maxGoRoutines)

	go func() {
		counter := -1;
		for dex, file := range files {
			log.Println(fmt.Sprintf("starting file: %s", file))
			columnNamesToIndex := make(map[string]int)

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
					wg.Wait()
					if dex == 1{
						exitChan <- 1;
						return
					}
					break
				}

				counter +=1

				//TODO comment out if you want to do > 50 line csv output
				//if counter >= 50{
				//	wg.Wait()
				//	exitChan <-1
				//	return
				//}

				sem <- struct{}{}
				go func(tweetID, tweetContent string) {
					wg.Add(1)
					defer wg.Done()
					defer func() {
						<- sem
					}()
					sentimentInt, _, _ := getSentiment(tweetContent)

					//isCorrect := "correct"
					//if sentimentInt ==1 && sentimentFloat >= 0.5{
					//	isCorrect = "correct"
					//} else {
					//	isCorrect = "wrong"
					//}
					//log.Println(fmt.Sprintf("%s %s,%s,%d,%.2f\n",isCorrect, tweetID,nameList[dex],sentimentInt, sentimentFloat))

					tc := &tweet{
						sentimentScore:sentimentInt,
						tweetID:tweetID,
						person: nameList[dex],
					}

					tweetChan <-tc
					return
				}(record[columnNamesToIndex["tweet_id"]], record[columnNamesToIndex["tweet"]])
			}
		}
		return
	}()

	for {
		select {
		case incomingTweet := <-tweetChan:
			f.WriteString(fmt.Sprintf("%s,%s,%d\n",incomingTweet.tweetID,incomingTweet.person,incomingTweet.sentimentScore))

			recordsProcessed +=1
			timePerRecord := float64(time.Since(timeStart).Milliseconds())/float64(recordsProcessed)
			eta:= timePerRecord * float64(totalRecords-recordsProcessed) / float64(time.Hour.Milliseconds())
			if recordsProcessed % 10 == 0{
				log.Println(fmt.Sprintf("%d/%d rows handled(%.2f%%) ETA: %.2f hours \n", recordsProcessed, int64(totalRecords), float64(recordsProcessed)/totalRecords*float64(100), eta))
			}
		case <- exitChan:
			close(exitChan)
			close(tweetChan)
			return
		}
	}
}