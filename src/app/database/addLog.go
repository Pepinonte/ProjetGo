package database

import "fmt"

func AddLog(date string, module string, fonction string, arguments string, output string) (int64, error) {

	var log Log = Log{
		date: date,
		module: module,
		fonction: fonction,
		arguments: arguments,
		output: output,
	} 

	result, err := db.Exec("INSERT INTO logs (date, module, fonction, arguments, output) VALUES (?, ?, ?, ?, ?)", log.date, log.module, log.fonction, log.arguments, log.output)
	if err != nil {
		return 0, fmt.Errorf("addLog: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addLog: %v", err)
	}
	return id, nil
}



		// id, err := addLog(Log{
		// 	date:  "2020-01-01",
		// 	module: "module1",
		// 	fonction:  "fonction1",
		// 	arguments:  "arguments1",
		// 	output:  "output1",
		// })
