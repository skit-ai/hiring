package main

import (
	"github.com/gocarina/gocsv"
	"live_coding/models"
	"os"
)

func main() {
	// Parse CSV into a grouped dataframe on the basis of the group ID
	in, groupedDf := parseCSV()
	defer closeSafely(in)

	// Map with
	var analyzedGroups []models.AnalyzedGroup

	// Iterating over records to analyse each group
	for groupId, _ := range groupedDf {
		analyzedGroup := models.AnalyzedGroup{
			GroupID: groupId,
		}
		// TODO: Insert your code
		// ...
		analyzedGroups = append(analyzedGroups, analyzedGroup)
	}

	// Write results to CSV
	file, _ := os.OpenFile("analysis.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	// Close file connection
	defer closeSafely(file)
	//write the csv file
	if err := gocsv.MarshalFile(&analyzedGroups, file); err != nil{
		panic(err)
	}
}

// Close connection to a file safely
func closeSafely(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}

// Parse the CSV and group the data on the basis of the group ID
func parseCSV() (*os.File, map[int][]*models.DataFrame) {
	in, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}

	var dfs []*models.DataFrame

	if err := gocsv.UnmarshalFile(in, &dfs); err != nil {
		panic(err)
	}

	var groupedDf = make(map[int][]*models.DataFrame)
	var exists bool
	var df *models.DataFrame

	for i := range dfs {
		df = dfs[i]
		// Initializing the map slice
		if _, exists = groupedDf[df.GroupID]; !exists {
			groupedDf[df.GroupID] = []*models.DataFrame{}
		}

		// Map group Ids to all the individual IDs
		groupedDf[df.GroupID] = append(groupedDf[df.GroupID], df)
	}
	return in, groupedDf
}
