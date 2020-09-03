package middleware

import (
	"database/sql" // package to encode and decode the json into struct and vice versa
	"fmt"          // models package where User schema is defined
	"log"          // used to access the request and response object of the api
	"os"           // used to read the environment variable
	"roster-api/models"

	// package used to covert string into int type
	// used to get the params from the route

	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// package used to read the .env file
	// postgres golang driver
)

/*
	create table in database
	CREATE TABLE weekRosters (
		Id SERIAL PRIMARY KEY,
		Year int,
		Week int,
		Mon json,
		Tue json,
		Wed json,
		Thu json,
		Fri json,
		Sat json,
		Sun json
	);
*/

func getAllRosters() ([]models.WeekRoster, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var rosters []models.WeekRoster

	// create the select sql query
	sqlStatement := `SELECT * FROM weekrosters`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var roster models.WeekRoster

		// unmarshal the row object to user
		err = rows.Scan(
			&roster.ID,
			&roster.Week,
			&roster.Year,
			&roster.Mon,
			&roster.Tue,
			&roster.Wed,
			&roster.Thu,
			&roster.Fri,
			&roster.Sat,
			&roster.Sun,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		rosters = append(rosters, roster)

	}

	// return empty user on error
	return rosters, err
}

// GetAllRoster send all Roster
func GetAllRoster(c *fiber.Ctx) {
	fmt.Println("get rosters ")

	rosters, err := getAllRosters()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	/* 	func (t *Test) MarshalJSON() ([]byte, error) {
		var array string
		if t.Array == nil {
			array = "null"
		} else {
			array = strings.Join(strings.Fields(fmt.Sprintf("%d", t.Array)), ",")
		}
		jsonResult := fmt.Sprintf(`{"Name":%q,"Array":%s}`, t.Name, array)
		return []byte(jsonResult), nil
	} */

	var location string = ""

	err := json.Unmarshal([]byte(location), &rosters[0].Mon)

	fmt.Println(err)
	// send all the users as response
	// json.NewEncoder(w).Encode(users)
	c.Send(rosters[0].Week)

}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
