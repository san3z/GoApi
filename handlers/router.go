package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage)

	//  маршрут к статическим файлам
	staticDir := http.Dir("./Pictures")
	fileServer := http.FileServer(staticDir)

	// любой запрос, начинающийся с "/images/", будет сопоставляться с этим маршрутом
	imageRoute := r.PathPrefix("/images/")

	// обработчик удаляет префикс "/images/" из URL запроса перед передачей запроса в fileServer
	imageRoute.Handler(http.StripPrefix("/images/", fileServer))

	// указывается ip адрес на котором запускается сайт, и логирует ошибку при ее появлении
	log.Println("Server is running")
	err := http.ListenAndServe("192.168.100.5:80", r)
	if err != nil {
		log.Fatal(err)
	}
}

// создается гипперсылка
func createHyperlink(url, displayText string) string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, url, displayText)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	url := "http://192.168.100.5:80/images/1.jpg"
	displayText := "images"
	hyperlink := createHyperlink(url, displayText)
	fmt.Println("Available pages:")
	fmt.Fprintln(w, hyperlink)
}

func Pictures(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "No pictures yet")
}
