package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/joho/godotenv"

	bot "github.com/Heart-plus-N/habitica-multi-bot/bot"
	op "github.com/Heart-plus-N/habitica-multi-bot/observer_pattern"
	qq "github.com/Heart-plus-N/habitica-multi-bot/quest_queue"

	"github.com/go-chi/chi"
	. "gitlab.com/bfcarpio/gabit"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	// Ensure we have a port
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		log.Println("Setting port to default :8080")
		port = ":8080"
	}

	sc := op.SharedConfig{
		HabiticaUsername: os.Getenv("HMB_USERNAME"),
		HabiticaPassword: os.Getenv("HMB_PASSWORD"),
	}

	// Connect to habitica
	hapi := NewHabiticaAPI(nil, "", nil)
	_, err := hapi.Authenticate(sc.HabiticaUsername, sc.HabiticaPassword)
	if err != nil {
		log.Fatalln("Could not log into Habitica")
		log.Fatalln(err)
	}
	log.Println("Logged into Habitica")

	// Set up oversevers
	reporter := op.NewReporter(sc)

	bot_utils := bot.Bot{Name: "Bot Utils"}
	reporter.Subscribe(bot_utils)

	quest_queue := qq.QuestQueue{Name: "QQ"}
	reporter.Subscribe(quest_queue)

	quest_queue_2 := qq.QuestQueue{Name: "QQ2"}
	reporter.Subscribe(quest_queue_2)

	// Open routes
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bot up as of: %s", time.Now().String())
	})
	r.Post("/task", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.TaskEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/chat", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.GroupChatEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.UserEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/quest", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.QuestEvent, body)
		}

		w.WriteHeader(200)
	})

	log.Println("Listening on", port)
	err = http.ListenAndServe(port, r)
	log.Fatal(err)
}
