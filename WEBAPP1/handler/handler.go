package handler

// Importing libraries
import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/ibilalkayy/WEBAPP1/database"
	"golang.org/x/crypto/bcrypt"
)

// Declaring variables
type Signupdata struct {
	Name, Email, Password, About string
}

type Templatedata struct {
	Signupsuccess, Loginfailure string
}

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

// Attaching this function to every template
func makeTemplate(path string) *template.Template {
	files := []string{path, "templates/footer.html", "templates/base.html"}
	return template.Must(template.ParseFiles(files...))
}

var (
	homeTmpl      = makeTemplate("templates/home.html")
	aboutTmpl     = makeTemplate("templates/about.html")
	aboutFormTmpl = makeTemplate("templates/aboutform.html")
	signupTmpl    = makeTemplate("templates/signup.html")
	loginTmpl     = makeTemplate("templates/login.html")
	logoutTmpl    = makeTemplate("templates/logout.html")
	pageErrorTmpl = makeTemplate("templates/pageerror.html")
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		if err := pageErrorTmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := homeTmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

// Converting the password into hash
func hashAndSalt(pass []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// Comparing password with hash values
func comparePasswords(hashedPass string, plainPass []byte) bool {
	byteHash := []byte(hashedPass)

	if err := bcrypt.CompareHashAndPassword(byteHash, plainPass); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Signup page
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if err := signupTmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm-password")
	hashPassword := hashAndSalt([]byte(password))
	if comparePasswords(hashPassword, []byte(confirmPassword)) {
		stuff := Signupdata{name, email, hashPassword, ""}
		database.Insertdata(stuff)
		sm := Templatedata{Signupsuccess: "Your account is successfully created"}
		if err := loginTmpl.Execute(w, sm); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		fm := struct{ Failure string }{Failure: "Both passwords are not matched"}
		if err := signupTmpl.Execute(w, fm); err != nil {
			log.Fatal(err)
		}
	}
}

// Login page
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if err := loginTmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	match := database.Findaccount(email, password)
	if match == true {
		session, _ := store.Get(r, "session-name")
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/about", http.StatusFound)
	} else {
		fm := Templatedata{Loginfailure: "Enter the correct email or password"}
		if err := loginTmpl.Execute(w, fm); err != nil {
			log.Fatal(err)
		}
	}
}

// About authentication
func About(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var authenticated interface{} = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			if err := loginTmpl.Execute(w, nil); err != nil {
				panic(err)
			}
			return
		}
		showAbout(w, r)
	} else {
		if err := loginTmpl.Execute(w, nil); err != nil {
			panic(err)
		}
		return
	}
}

// About page
func showAbout(w http.ResponseWriter, r *http.Request) {
	d := struct{ Aboutdata string }{Aboutdata: database.Account.About}
	if r.Method == "GET" {
		if database.Account.About == "" {
			if err := aboutFormTmpl.Execute(w, nil); err != nil {
				log.Fatal(err)
			}
			return
		} else {
			if err := aboutTmpl.Execute(w, d); err != nil {
				log.Fatal(err)
			}
			return
		}
	} else if r.Method == "POST" {
		content := r.FormValue("content")
		update := database.Updatedata("about", content)
		if update == true {
			d := struct{ Aboutdata string }{Aboutdata: content}
			if err := aboutTmpl.Execute(w, d); err != nil {
				log.Fatal(err)
			}
			return
		} else {
			f := struct{ Aboutfailure string }{Aboutfailure: "Your data is not updated"}
			if err := aboutFormTmpl.Execute(w, f); err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}

// Logout page
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	if err := logoutTmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
