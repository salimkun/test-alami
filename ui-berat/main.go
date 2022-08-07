package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"text/template"

	"github.com/salimkun/sirclo-test/ui-berat/models"
	sc "github.com/salimkun/sirclo-test/ui-berat/repo"
)

func createFileTemp() {
	// Shopping chart
	// create file to temp database
	_, err := os.Create("weight-db.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
}

func main() {
	createFileTemp()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		var resp = &models.ListWeight{}
		var msgErr string
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var sc sc.WeighRepo = &sc.WeightData{}

		if r.Method == "POST" {
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
			}

			typeAction := r.FormValue("option-action")
			if typeAction == "edit" {
				dateWeight := r.FormValue("date-weight-update")
				maxWeight := r.FormValue("max-weight-update")
				minWeight := r.FormValue("min-weight-update")
				fmt.Println("ALALLALASKS ", dateWeight, "asns ", maxWeight)

				max, _ := strconv.Atoi(maxWeight)
				min, _ := strconv.Atoi(minWeight)

				err = sc.UpdateWeight(&models.Weight{
					Date: dateWeight,
					Max:  int32(max),
					Min:  int32(min),
				})
				if err != nil {
					fmt.Println("Error ", err.Error())
					msgErr = err.Error()
				}
			} else {
				dateWeight := r.FormValue("date-weight")
				maxWeight := r.FormValue("max-weight")
				minWeight := r.FormValue("min-weight")

				max, _ := strconv.Atoi(maxWeight)
				min, _ := strconv.Atoi(minWeight)

				err = sc.AddWeight(&models.Weight{
					Date: dateWeight,
					Max:  int32(max),
					Min:  int32(min),
				})
				if err != nil {
					fmt.Println("Error ", err.Error())
					msgErr = err.Error()
				}
			}
		} else if r.Method == "PUT" {
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
			}
			dateWeight := r.FormValue("date-weight")
			maxWeight := r.FormValue("max-weight")
			minWeight := r.FormValue("min-weight")

			max, _ := strconv.Atoi(maxWeight)
			min, _ := strconv.Atoi(minWeight)

			err = sc.AddWeight(&models.Weight{
				Date: dateWeight,
				Max:  int32(max),
				Min:  int32(min),
			})
			if err != nil {
				fmt.Println("Error ", err.Error())
				msgErr = err.Error()
			}
		} else if r.Method == "GET" {
			paramKey := r.URL.Query().Get("dateKey")
			if paramKey != "" {
				err = sc.DeleteWeight(paramKey)
				if err != nil {
					fmt.Println("Error ", err.Error())
					msgErr = err.Error()
				}
			}
		}

		//Request not POST
		resp = sc.ViewAllWeight(0, 0)
		resp.Error = msgErr
		err = tmpl.Execute(w, resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
