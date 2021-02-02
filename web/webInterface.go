package web

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"mailspamer/db"
	"mailspamer/mail"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var Database *sql.DB
var ExternalServer string
var ServerPort string

//send emails page
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	//ClientInfo(r)
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("r.ParseForm error: ", err)
		}

		email := r.FormValue("emails")
		emails := strings.Fields(email)

		//many e-mails enter to db
		for _, email := range emails {
			if mail.CheckMail(email) {
				path := createUrl(email)
				picturePath := createUrlToPic(email)
				mail.SendMail(email, path, picturePath)
				db.AddEmailInfoToDB(Database, email, path, picturePath,false, false)
				if err != nil {
					fmt.Println("http.Get error: ", err)
				}
			}
		}

		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "public/html/create.html")
	}
}

//main page (with db)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//ClientInfo(r)
	// press Send e-mails button
	if r.Method == "POST" {
		http.Redirect(w, r, "/create", 301)
	} else {
		if r.URL.RequestURI() == "/" {
			infos := db.GetEmailInfoFromDB(Database)
			sort.Slice(infos, func(i, j int) bool { return infos[i].Id < infos[j].Id })
			//ClientInfo(r)
			tmpl, _ := template.ParseFiles("public/html/index.html")
			tmpl.Execute(w, infos)
		} else if strings.Contains(r.URL.RequestURI(), "email=") {
			info := db.SearchURLInDB(Database, ExternalServer+r.URL.RequestURI())
			if info.Mail != "" {
				fmt.Println("Link status changed:", info.Id, info.Mail)
				db.UpdateDB(Database, info.Id)
			}
		} else if strings.Contains(r.URL.RequestURI(), "picture=") {
			info := db.SearchPICURLInDB(Database, ExternalServer+r.URL.RequestURI())
			if info.Mail != "" {
				fmt.Println("Mail checked:", info.Id, info.Mail)
				db.UpdateLinkPICURLDB(Database, info.Id)
			}
		}
	}
}

func WebServer() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/create", CreateHandler)

	fmt.Println("Server is listening on port:", ServerPort)
	http.ListenAndServe(":" + ServerPort, nil)
}

// print web clients info
func ClientInfo(req *http.Request) {
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {

		fmt.Println("userip: is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Println("userip: is not IP:port", req.RemoteAddr)
		return
	}
	fmt.Printf("Client IP %s PORT: %s\n", ip, port)

}

// create unique URL with email plus current time
func createUrl(email string) string {
	t := time.Now()
	uniqueUrl := randomString(email)[1:8] + t.Format(time.RFC3339)
	baseMailByte := []byte(uniqueUrl)
	baseMailStr := base64.StdEncoding.EncodeToString(baseMailByte)
	params := "email=" + url.QueryEscape(baseMailStr)
	path := fmt.Sprintf(ExternalServer + "/get?%s", params)
	return path
}

func createUrlToPic(email string) string {
	t := time.Now()
	uniqueUrl := randomString(email)[1:3] + t.Format(time.RFC3339) + "pic"
	baseMailByte := []byte(uniqueUrl)
	baseMailStr := base64.StdEncoding.EncodeToString(baseMailByte)
	params := "picture=" + url.QueryEscape(baseMailStr)
	path := fmt.Sprintf(ExternalServer + "/get?%s", params)
	return path
}

//generate random string for unique url
func randomString(str string) string {
	chars := []rune(str)
	var b strings.Builder
	for i := 0; i < len(str); i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str = b.String()
	return str
}
