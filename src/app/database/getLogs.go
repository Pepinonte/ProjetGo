package database

import (
	"fmt"
)

// func GetLogs() (string, error) {
// 	var log Log
// 	rows := db.QueryRow("SELECT * FROM logs")
// 	if err := rows.Scan(&log.ID, &log.date, &log.module, &log.fonction, &log.arguments, &log.output); err != nil {
// 		if err == sql.ErrNoRows {
// 			return "", fmt.Errorf("logById: no such log")
// 		}
// 		return "", fmt.Errorf("logById %d:", err)
// 	}

// 	for rows.Next(){
// 		var alb Album
// 		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
// 			return nil, fmt.Errorf("scanning rows: %v", err)
// 		}
// 		albums = append(albums, alb)
// 	}
// 	return "success", nil
// }


func GetLogs() (string, error) {
	var log Log
	rows, err := db.Query("SELECT * FROM logs")
	if err != nil {
		return "", fmt.Errorf("logById %d:", err)
	}

	for rows.Next(){
		if err := rows.Scan(&log.ID, &log.date, &log.module, &log.fonction, &log.arguments, &log.output); err != nil {
			return "", fmt.Errorf("scanning rows: %v", err)
		}
		fmt.Println(log)
	}
	return "success", nil
}