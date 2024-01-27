package database

import (
	"fmt"
)

func GetLogs() (string, error) {

	defer db.Close()
	var log Log
	rows, err := db.Query("SELECT * FROM logs")
	if err != nil {
		return "", fmt.Errorf("logById %d", err)
	}

	for rows.Next(){
		if err := rows.Scan(&log.ID, &log.date, &log.module, &log.fonction, &log.arguments, &log.output); err != nil {
			return "", fmt.Errorf("scanning rows: %v", err)
		}
		fmt.Println(log)
	}
	return "success", nil
}