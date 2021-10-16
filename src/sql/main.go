package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)
 
const (
    host     = "192.168.2.111"
    port     = 5432
    user     = "postgres"
    password = "secretpassword"
    database = "gocam"
)
 
func main() {
    // Build the database connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable", host, port, user, password, database)
         
    // Connect to database using the sql package
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
     
    // Make sure connection is closed when done
    defer db.Close()
 
    // Verify you are connected and log result
    err = db.Ping()
    CheckError(err)
    fmt.Println(time.Now().Format(time.RFC850), ":: Connected to database")

	// Insert static data
	// insertData := `INSERT INTO "alarm_rec"("camera_id", "rec_path") VALUES('0001', '/opt/rec/12332423dsfsd232.snap')`
	// _, e := db.Exec(insertData)
	// CheckError(e)

	// Insert dynamic data
	// camID := 0002
	// recPath := "/opt/rec/546sdgv2314sdfgd232.snap"
	// insertVarData := `INSERT INTO "alarm_rec"("camera_id", "rec_path") VALUES($1, $2)`
	// _, e = db.Exec(insertVarData, camID, recPath)
	// CheckError(e)

	// Update a row
	// updateData := `UPDATE "alarm_rec" SET "rec_path"=$1 WHERE "camera_id"=$2`
	// _, e := db.Exec(updateData, "/opt/rec/false_positive/546sdgv2314sdfgd232.snap", 0001)
	// CheckError(e)

	// Delete a row
	// deleteData := `DELETE FROM "alarm_rec" WHERE "camera_id"=$1`
	// _, e := db.Exec(deleteData, 0001)
	// CheckError(e)

	selectData := `SELECT "rec_path" FROM "alarm_rec" WHERE "camera_id"=$1`
	rows, err := db.Query(selectData, 0002)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var path string
	
		err = rows.Scan(&path)
		CheckError(err)
	
		fmt.Println(path)
	}
	CheckError(err)
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}