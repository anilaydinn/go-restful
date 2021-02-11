package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {

	statement, driverError := dbDriver.Prepare(train)

	if driverError != nil {
		log.Panicln(driverError)
	}

	//Create train table
	_, statementError := statement.Exec()

	if statementError != nil {
		log.Println("Table already exists!")
	}

	statement, driverError = dbDriver.Prepare(station)

	if driverError != nil {
		log.Panicln(driverError)
	}

	//Create station table
	_, statementError = statement.Exec()

	if statementError != nil {
		log.Println("Table already exists!")
	}

	statement, driverError = dbDriver.Prepare(schedule)

	if driverError != nil {
		log.Panicln(driverError)
	}

	//Create schedule table
	_, statementError = statement.Exec()

	if statementError != nil {
		log.Println("Table already exists!")
	}

	log.Println("All tables created/initialized successfully!")
}
