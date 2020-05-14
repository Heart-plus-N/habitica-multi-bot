package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"os"

	op "github.com/Heart-plus-N/habitica-multi-bot/observer_pattern"
	qq "github.com/Heart-plus-N/habitica-multi-bot/quest_queue"
	"github.com/spf13/viper"

	"github.com/go-chi/chi"
	. "gitlab.com/bfcarpio/gabit"
)

func main() {
	// Load ENV or config variables
	viper.SetDefault("PORT", ":8080")
	if os.Getenv("ENV") == "PROD" {
		log.Println("Loading environment variables")
		viper.AutomaticEnv()
	} else {
		log.Println("Loading config file")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		viper.ReadInConfig()
	}

	// Ensure we have a port
	port := ":" + viper.GetString("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Connect to habitica
	hapi := NewHabiticaAPI(nil, "", nil)
	_, err := hapi.Authenticate(viper.GetString("HABITICA_USERNAME"), viper.GetString("HABITICA_PASSWORD"))
	if err != nil {
		log.Fatalln("Could not log into Habitica")
		log.Fatalln(err)
	}
	log.Println("Logged into Habitica")

	// Set up oversevers
	reporter := op.Reporter{}

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
