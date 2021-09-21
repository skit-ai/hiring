package pkg1

import "live_coding/models"

func count(df []*models.DataFrame) int {
	return len(df)
}
