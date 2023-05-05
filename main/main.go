package main

import (
	//"flag"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	//"github.com/apex/gateway"

	"github.com/alexedwards/scs/v2"
	"github.com/bertoxic/aarc/config"
	"github.com/bertoxic/aarc/drivers"
	"github.com/bertoxic/aarc/handlers"
	"github.com/bertoxic/aarc/helpers"
	"github.com/bertoxic/aarc/models"
	"github.com/bertoxic/aarc/render"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// User represents a user account

var app config.AppConfig
var sessions *scs.SessionManager
const portNumber = ":8087"

func main() {
	gob.Register(helpers.PcProperties{})

	// dbHost := flag.String("dbhost","localhost","Database Host")
	// dbName := flag.String("dbname","","Database name")
	// dbUser := flag.String("dbuser","","Database user")
	// dbPass := flag.String("dbpass","","Database password")
	// dbPort := flag.String("dbport","5432","Database port")
	// dbSSL := flag.String("dbssl","disable","Database ssl settings (disable, prefer, require)")
	// port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
    // flag.Parse()
	// listener := gateway.ListenAndServe
	// portStr := "n/a"
	// if *port != -1 {
    //     portStr = fmt.Sprintf(":%d", *port)
    //     listener = http.ListenAndServe
    //     http.Handle("/", http.FileServer(http.Dir("./public")))
    //}
	app.InProduction = false
	app.UserCache= false  
	sessions = scs.New()
	sessions.Cookie.Persist = true
	sessions.Lifetime = 24 * time.Hour
	sessions.Cookie.Secure = app.InProduction
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	app.Sessions = sessions
    // create a connection to the database
	connectinString := "host=localhost port=5432 dbname=aarc user=postgres password=bert"
	// connectinString := "host=dpg-ch50v4cs3fvqdikfsu10-a	port=5432 dbname=practice_sngu user=bert password=oKSsFrN3yEuxcErXKvoAopsLkMh5rZsl"
	
	// connectinString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",*dbHost,*dbPort,*dbName,*dbUser,*dbPass,*dbSSL )
	db, err := drivers.ConnectSQL(connectinString)
	if err !=nil {
		//log.Fatal("cannot initiate dbase",err)
		log.Println("cannot initiate dbase")
	}
	defer db.SQL.Close()
	log.Println("connectioon established")

		if err !=nil {
		log.Fatalf(fmt.Sprintf("cound not ping:%v ",err))
	}

	//creating the postgresrepo
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println(err)
		return
	}
	repo :=handlers.NewRepo(db,&app)
	render.NewRenderer(&app)
	handlers.NewHandlers(repo)
	app.TemplateCache = tc
   // mailChan := make (chan models.MailData)
	//app.MailChan = mailChan
	// Assume user has provided these details during registration
	email := "user@example.com"
	password := "password123"

	// Create a new user account with a verification code
    verificationCode, expiryTime := helpers.GenerateVerificationCode()
	user := models.User{
		Email:        email,
		Password:     password,
		Verified:     false,
		Verification: verificationCode,
        VerificationExpiry: expiryTime,
	}

	// Send the verification email to the user
	
	listenForMail()
	// msg := models.MailData{
	// 	To:      "me@gmail.com",
	// 	From:    "ok@fmail.com",
	// 	Subject: "Hellllo people",
	// 	Content: "june dumps",
	// 	Template: "",
	// }
	log.Println("before listening email sent to", user.Email)
	
	
			//app.MailChan <- msg

	
	// Email sent successfully
	log.Println("Verification email sent to", user.Email)
	
	sqlStatement := `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	verification TEXT,
	verification_expiry TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`
D := db.SQL
// Execute the SQL statement to create the table
_, err = D.Exec(sqlStatement)
if err != nil {
		log.Println("failed to create database")
}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	
	// err =listener(   portStr, routes(&app))
	if err != nil {
		log.Fatal(err)
	}
	
	
}

// sendVerificationEmail sends a verification email to the user


// generateVerificationCode generates a random verification code for the user
