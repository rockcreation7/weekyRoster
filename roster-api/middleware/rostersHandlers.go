package middleware

import (
	// package to encode and decode the json into struct and vice versa
	"database/sql"
	"fmt" // models package where User schema is defined
	"log" // used to access the request and response object of the api
	"os"

	// used to read the environment variable
	"roster-api/models"

	// package used to covert string into int type
	// used to get the params from the route

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"

	// db driver
	_ "github.com/lib/pq"
)

// GetAllRoster ...
func GetAllRoster(c *fiber.Ctx) {
	rosters, err := getAllRosters()
	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	fmt.Println(err)
	c.JSON(rosters)
}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// CreateRoster ...
func CreateRoster(c *fiber.Ctx) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	u := new(models.DayRoster)
	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}

	fmt.Println(u)
	return
	// create an empty user of type models.User
	var Roster models.DayRoster

	Roster.Date = "20200903"
	Roster.UpperStaff = "david"
	Roster.UpperTime = "david"
	Roster.LowerStaff = "david"
	Roster.LowerTime = " "
	Roster.CustomMessage = " "

	// call insert user function and pass the user
	insertID, err := insertRoster(Roster)

	if err != nil {
		c.JSON(err)
		return
	}

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Roster created successfully",
	}

	// send the response
	c.JSON(res)
}

func insertRoster(Roster models.DayRoster) (int64, error) {

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO Rosters (Date, UpperStaff, UpperTime, LowerStaff, LowerTime, CustomMessage) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, Roster.Date, Roster.UpperStaff, Roster.UpperTime, Roster.LowerStaff, Roster.LowerTime, Roster.CustomMessage).Scan(&id)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
		fmt.Printf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, err
}

func getAllRosters() ([]models.DayRoster, error) {
	// create the postgres db connection
	db := createConnection()
	defer db.Close()

	var rosters []models.DayRoster
	sqlStatement := `SELECT * FROM Rosters`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var roster models.DayRoster
		err = rows.Scan(
			&roster.ID,
			&roster.Date,
			&roster.UpperStaff,
			&roster.UpperTime,
			&roster.LowerStaff,
			&roster.LowerTime,
			&roster.CustomMessage,
		)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		rosters = append(rosters, roster)
	}

	return rosters, err
}

// GetAllRoster send all Roster

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
