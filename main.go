package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type empData struct {
	ID               int64
	Name             string
	Age              int32
	Balanced         int64
	PreviousBalanced int64
	AverageBalanced  int64
	FreeTransfer     int64
	ThreadNo1        string
	ThreadNo2a       string
	ThreadNo2b       string
	ThreadNo3        string
}

var (
	bonus100 = 100
)

func main() {
	// read file from input
	f, _ := os.Open("Before-Eod.csv")
	defer f.Close()

	// worker run function
	concuRSwWP(f)
}

// with Worker pools
func concuRSwWP(f *os.File) {
	fcsv := csv.NewReader(f)
	rs := make([]*empData, 0)
	numWps := 5
	jobs := make(chan []string, numWps)
	res := make(chan *empData)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *empData, workerID string) {
		for {
			select {
			case job, ok := <-jobs: // you must check for readable state of the channel.
				if !ok {
					return
				}
				results <- parseStruct(job, workerID)
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		workerID := w
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, res, strconv.Itoa(workerID))
		}()
	}

	go func() {
		for {
			rStr, err := fcsv.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		rs = append(rs, r)
	}

	// create file for output
	newFile, err := os.Create("After-Eod.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// write to output file
	csvwriter := csv.NewWriter(newFile)

	// write header
	err = csvwriter.Write([]string{
		"id;Nama;Age;Balanced;No 2b Thread-No;No 3 Thread-No;Previous Balanced;Average Balanced;No 1 Thread-No;Free Transfer;No 2a Thread-No"})
	if err != nil {
		log.Println("Cannot write header to CSV file:", err)
	}

	for _, empRow := range rs {
		strData := fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s", strconv.Itoa(int(empRow.ID)),
			empRow.Name, strconv.Itoa(int(empRow.Age)), strconv.Itoa(int(empRow.Balanced)),
			empRow.ThreadNo2b, empRow.ThreadNo3,
			strconv.Itoa(int(empRow.PreviousBalanced)),
			strconv.Itoa(int(empRow.AverageBalanced)),
			empRow.ThreadNo1, strconv.Itoa(int(empRow.FreeTransfer)),
			empRow.ThreadNo2a)
		err = csvwriter.Write([]string{strData})
		if err != nil {
			log.Println("Cannot write to CSV file:", err)
		}
	}
	csvwriter.Flush()
	newFile.Close()
}

func parseStruct(data []string, workerID string) *empData {
	tempData := strings.Split(data[0], ";")
	id, _ := strconv.Atoi(tempData[0])
	age, _ := strconv.Atoi(tempData[2])
	balanced, _ := strconv.Atoi(tempData[3])
	pbalanced, _ := strconv.Atoi(tempData[4])
	abalanced, _ := strconv.Atoi(tempData[5])
	ftransfer, _ := strconv.Atoi(tempData[6])
	emp := &empData{
		ID:               int64(id),
		Name:             tempData[1],
		Age:              int32(age),
		Balanced:         int64(balanced),
		PreviousBalanced: int64(pbalanced),
		AverageBalanced:  int64(abalanced),
		FreeTransfer:     int64(ftransfer),
	}

	// count Average
	emp.AverageBalanced = (emp.Balanced + emp.PreviousBalanced) / 2
	emp.ThreadNo1 = workerID

	// free transfer
	if emp.Balanced >= 100 || emp.Balanced <= 150 {
		emp.FreeTransfer = 5
		emp.ThreadNo2a = workerID
	}

	//add balance bonus 25
	if emp.Balanced > 150 {
		emp.Balanced = emp.Balanced + 25
		emp.ThreadNo2b = workerID
	}

	// add balance bonus 10 for first 1000 user
	if bonus100 != 0 {
		bonus100 = bonus100 - 10
		emp.Balanced = emp.Balanced + 10
		emp.ThreadNo3 = workerID
	}

	return emp
}
