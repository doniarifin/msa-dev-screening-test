package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Structs to model the API response
type ApiResponse struct {
	Code int             `json:"code"`
	Data ApiResponseData `json:"data"`
}

type ApiResponseData struct {
	Rows []GraduateData `json:"rows"`
}

type GraduateData struct {
	VaultID                  string `json:"vault_id"`
	Year                     string `json:"year"`
	University               string `json:"university"`
	School                   string `json:"school"`
	Degree                   string `json:"degree"`
	EmploymentRateOverall    string `json:"employment_rate_overall"`
	EmploymentRateFtPerm     string `json:"employment_rate_ft_perm"`
	BasicMonthlyMean         string `json:"basic_monthly_mean"`
	BasicMonthlyMedian       string `json:"basic_monthly_median"`
	GrossMonthlyMean         string `json:"gross_monthly_mean"`
	GrossMonthlyMedian       string `json:"gross_monthly_median"`
	GrossMonthly25Percentile string `json:"gross_mthly_25_percentile"`
	GrossMonthly75Percentile string `json:"gross_mthly_75_percentile"`
}

func main() {
	concurrentLimit := flag.Int("concurrent_limit", 2, "Limit of concurrent goroutines")
	outputDir := flag.String("output", "./csv", "Output directory for CSV files")
	flag.Parse()

	apiURL := "https://api-production.data.gov.sg/v2/internal/api/datasets/d_3c55210de27fcccda2ed0c63fdd2b352/rows?limit=1000"

	// fetch data from API
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// decode json response
	var apiResp ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		log.Fatalf("erorr: %v", err)
	}

	// map data by year
	yearlyData := make(map[string][]GraduateData)

	// process data and organize by year
	for _, data := range apiResp.Data.Rows {
		yearlyData[data.Year] = append(yearlyData[data.Year], data)
	}

	concurrency := make(chan struct{}, *concurrentLimit)
	var wg sync.WaitGroup

	// process each year's data concurrently
	for year, data := range yearlyData {
		wg.Add(1)
		concurrency <- struct{}{}
		go func(year string, data []GraduateData) {
			defer func() {
				<-concurrency
				wg.Done()
			}()
			err := writeCSV(*outputDir, year, data)
			if err != nil {
				log.Printf("Error writing csv for %s: %v", year, err)
			}
		}(year, data)
	}

	wg.Wait()
	fmt.Println("CSV files have been created successfully.")
}

// function to create csv file
func writeCSV(outputDir, year string, data []GraduateData) error {
	// create directory if it doesn't exist
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory: %v", err)
	}

	// construct filename
	fileName := filepath.Join(outputDir, fmt.Sprintf("%s.csv", year))

	// create csv file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file %s: %v", fileName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"Vault ID",
		"Year",
		"University",
		"School",
		"Degree",
		"Employment Rate Overall",
		"Employment Rate FT Perm",
		"Basic Monthly Mean",
		"Basic Monthly Median",
		"Gross Monthly Mean",
		"Gross Monthly Median",
		"Gross Monthly 25th Percentile",
		"Gross Monthly 75th Percentile",
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error csv header: %v", err)
	}

	// write each row of data to csv
	for _, d := range data {
		row := []string{
			d.VaultID,
			d.Year,
			d.University,
			d.School,
			d.Degree,
			d.EmploymentRateOverall,
			d.EmploymentRateFtPerm,
			d.BasicMonthlyMean,
			d.BasicMonthlyMedian,
			d.GrossMonthlyMean,
			d.GrossMonthlyMedian,
			d.GrossMonthly25Percentile,
			d.GrossMonthly75Percentile,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error csv row: %v", err)
		}
	}

	return nil
}
