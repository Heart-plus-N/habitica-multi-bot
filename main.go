package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	. "github.com/Heart-plus-N/habitica-multi-bot/bot"
	log "github.com/amoghe/distillog"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	. "gitlab.com/bfcarpio/gabit"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warningln("No .env file found")
	}
}

func main() {
	// Ensure we have a port
	portNum := os.Getenv("PORT")
	if portNum == "" {
		log.Infoln("Setting port to default :8080")
		portNum = "8080"
	}
	portStr := ":" + portNum

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			os.Getenv("HMB_USERNAME"),
			os.Getenv("HMB_PASSWORD"),
			os.Getenv("HMB_DB_HOST"),
			os.Getenv("HMB_DB_NAME"),
		)
	}
	connConfig, _ := pgx.ParseConfig(databaseUrl)
	connStr := stdlib.RegisterConnConfig(connConfig)
	db, dbErr := sql.Open("pgx", connStr)
	if dbErr != nil {
		panic(dbErr)
	}
	defer db.Close()

	habiticaUsername := os.Getenv("HMB_USERNAME")
	habiticaPassword := os.Getenv("HMB_PASSWORD")

	// Connect to habitica
	hapi := NewHabiticaAPI(nil, "", nil)
	_, authErr := hapi.Authenticate(habiticaUsername, habiticaPassword)
	if authErr != nil {
		panic(authErr)
	}
	log.Infoln("Logged into Habitica as: ", habiticaUsername)

	sc := SharedConfig{
		Api: hapi,
		Db:  db,
	}

	// Open routes
	r := chi.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bot up as of: %s", time.Now().String())
	})
	r.Post("/task", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorln(err)
		} else {
			go reporter.Notify(op.TaskEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/chat", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorln(err)
		} else {
			go reporter.Notify(op.GroupChatEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorln(err)
		} else {
			go reporter.Notify(op.UserEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/quest", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorln(err)
		} else {
			go reporter.Notify(op.QuestEvent, body)
		}

		w.WriteHeader(200)
	})

	log.Infoln("Listening on", portStr)
	serverErr := http.ListenAndServe(portStr, r)
	log.Errorln(serverErr)
}
