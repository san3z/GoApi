package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	users "GoodApi/database"

	"github.com/gorilla/mux"
)

func NewRouter() {
	r := mux.NewRouter()

	// маршрут к статическим файлам
	staticDir := http.Dir("./assets")
	fileServer := http.FileServer(staticDir)

	// любой запрос, начинающийся с "/images/", будет сопоставляться с этим маршрутом
	imageRoute := r.PathPrefix("/images/")

	// обработчик удаляет префикс "/images/" из URL запроса перед передачей запроса в fileServer
	imageRoute.Handler(http.StripPrefix("/images/", fileServer))

	r.HandleFunc("/", HomePage)
	r.HandleFunc("/DataBase", DataBase).Methods("GET")

	ip, ipErr := getLocalIP()
	if ipErr != nil {
		log.Fatal(ipErr)
	}
	port := getFreePort()
	serverAddress := fmt.Sprintf("%s:%d", ip, port)
	log.Println("Server is running", serverAddress)
	err := http.ListenAndServe(serverAddress, r)
	if err != nil {
		log.Fatal(err)
	}
}

// создается гипперсылка
func createHyperlink(url, displayText string) string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, url, displayText)
}

// Генерирует список гиперссылок
func generateHyperlinks(links []struct{ URL, DisplayText string }) string {
	var linksHTML string
	for _, link := range links {
		linksHTML += createHyperlink(link.URL, link.DisplayText) + "<br>"
	}
	return linksHTML
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	serverAddress := r.Host
	links := []struct{ URL, DisplayText string }{
		{fmt.Sprintf("http://%s/images/1.jpg", serverAddress), "Image 1"},
		{fmt.Sprintf("http://%s/DataBase", serverAddress), "UsrDB"},
	}

	// Генерация HTML-строки с гиперссылками
	hyperlinks := generateHyperlinks(links)
	fmt.Fprintln(w, hyperlinks)
}

func DataBase(w http.ResponseWriter, r *http.Request) {
	usersData := users.Getusers()
	userDataJSON, err := json.Marshal(usersData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userDataJSON)
}
