package main

import (
	"log"
	"net/http"

	"os"

	op "github.com/Heart-plus-N/habitica-multi-bot/observer_pattern"
	qq "github.com/Heart-plus-N/habitica-multi-bot/quest_queue"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/goware/httplog"
	"github.com/jinzhu/configor"
	. "gitlab.com/bfcarpio/gabit"
)

var Config = struct {
	Habitica struct {
		Username string `required:"true"`
		Password string `required:"true"`
	}
}{}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port, nil
}

func main() {
	// Load variable from config
	configor.Load(&Config, "config.toml")

	//
	hapi := NewHabiticaAPI(nil, "", nil)
	_, err := hapi.Authenticate(Config.Habitica.Username, Config.Habitica.Password)
	if err != nil {
		log.Fatalln("Could not log into Habitica")
		log.Fatalln(err)
	}
	log.Println("Logged into Habitica")

	reporter := op.Reporter{}

	quest_queue := qq.QuestQueue{Name: "QQ"}
	reporter.Subscribe(quest_queue)

	quest_queue_2 := qq.QuestQueue{Name: "QQ2"}
	reporter.Subscribe(quest_queue_2)

	r := chi.NewRouter()

	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})

	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/"))

	r.Post("/task", func(w http.ResponseWriter, r *http.Request) {
		body, err := hapi.ParseWebhookBody(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.TaskEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/chat", func(w http.ResponseWriter, r *http.Request) {
		body, err := hapi.ParseWebhookBody(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.GroupChatEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		body, err := hapi.ParseWebhookBody(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.UserEvent, body)
		}

		w.WriteHeader(200)
	})
	r.Post("/quest", func(w http.ResponseWriter, r *http.Request) {
		body, err := hapi.ParseWebhookBody(r.Body)
		if err != nil {
			log.Println(err)
		} else {
			go reporter.Notify(op.QuestEvent, body)
		}

		w.WriteHeader(200)
	})

	host, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on", host)
	err = http.ListenAndServe(host, r)
	log.Fatal(err)
}
