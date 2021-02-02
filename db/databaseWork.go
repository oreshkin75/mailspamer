package db

import (
	"database/sql"
	"fmt"
)

type EmailInfo struct {
	Id         int
	Mail       string
	ReadStatus bool
	LinkStatus bool
	UniqueUrl  string
	UniquePictureUrl string
}

func GetEmailInfoFromDB(db *sql.DB) []EmailInfo {
	rows, err := db.Query("SELECT * FROM emailInfo")
	if err != nil {
		fmt.Println("db.Query error: ", err)
	}

	defer rows.Close()
	var infos []EmailInfo

	for rows.Next() {
		p := EmailInfo{}
		err := rows.Scan(&p.Id, &p.Mail, &p.ReadStatus, &p.LinkStatus, &p.UniqueUrl, &p.UniquePictureUrl)
		if err != nil {
			fmt.Println("rows.Scan error: ", err)
			continue
		}
		infos = append(infos, p)
	}

	return infos
}

func AddEmailInfoToDB(db *sql.DB, email, url, picurl string, readstat, linkstat bool) {
	_, err := db.Exec("insert into emailInfo (mail, readstatus, linkstatus, uniqueurl, uniquepictureurl) values ($1, $2, $3, $4, $5)",
		email, readstat, linkstat, url, picurl)
	if err != nil {
		fmt.Println("Database.Exec error: ", err)
	}
}

func SearchURLInDB(db *sql.DB, url string) EmailInfo {
	row := db.QueryRow("select * from emailInfo where uniqueurl = $1", url)
	info := EmailInfo{}
	err := row.Scan(&info.Id, &info.Mail, &info.ReadStatus, &info.LinkStatus, &info.UniqueUrl, &info.UniquePictureUrl)
	if err != nil {
		fmt.Println("row.Scan error: ", err)
	}
	return info
}

func SearchPICURLInDB(db *sql.DB, url string) EmailInfo {
	row := db.QueryRow("select * from emailInfo where uniquepictureurl = $1", url)
	info := EmailInfo{}
	err := row.Scan(&info.Id, &info.Mail, &info.ReadStatus, &info.LinkStatus, &info.UniqueUrl, &info.UniquePictureUrl)
	if err != nil {
		fmt.Println("row.Scan error (SearchPICURLInDB): ", err)
	}
	return info
}

func UpdateDB(db *sql.DB, id int) {
	_, err := db.Exec("update emailInfo set linkstatus = $1 where id = $2", true, id)
	if err != nil {
		fmt.Println("db.Exec (UpdateDB) error: ", err)
	}
}

func UpdateLinkPICURLDB(db *sql.DB, id int) {
	_, err := db.Exec("update emailInfo set readstatus = $1 where id = $2", true, id)
	if err != nil {
		fmt.Println("db.Exec (UpdateDB) error: ", err)
	}
}

// create database if not exists
func CreateDB (db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS emailInfo (id SERIAL, mail TEXT, readstatus BOOLEAN, linkstatus BOOLEAN, uniqueurl TEXT, uniquepictureurl TEXT);")
	if err != nil {
		fmt.Println("Database.Exec error (CreateDB ): ", err)
	}
}
