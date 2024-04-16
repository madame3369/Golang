package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "hwicheon", Email: "akeka0303@naver.com"}

	rd.JSON(w, http.StatusOK, user)
	// http코드로 작성할 경우 아래와 같이 작성함.
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// http코드로 작성할 경우 아래와 같이 작성함
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()

	rd.JSON(w, http.StatusOK, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "hwicheon", Email: "akeka0303@naver.com"}
	rd.HTML(w, http.StatusOK, "body", user)
	// http 코드로 짤 경우 아래와 같이 작성함
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "hwicheon")
}
func main() {
	rd = render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	//negroni 사용하기 전에 사용한 방식
	// mux.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", mux)
}
