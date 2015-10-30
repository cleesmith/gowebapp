package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"

	"gowebapp/route"
	"gowebapp/shared/database"
	"gowebapp/shared/email"
	"gowebapp/shared/jsonconfig"
	"gowebapp/shared/recaptcha"
	"gowebapp/shared/server"
	"gowebapp/shared/session"
	"gowebapp/shared/view"
	"gowebapp/shared/view/plugin"
)

// *****************************************************************************
// Application Logic
// *****************************************************************************

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores: no longer needed with Go 1.5+
	// runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("GOMAXPROCS=%v\n", runtime.GOMAXPROCS(-1))
}

func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// Connect to database
	database.Connect(config.Database)
	defer database.DB.Close()

	// Configure the Google reCAPTCHA prior to loading view plugins
	recaptcha.Configure(config.Recaptcha)

	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	view.LoadPlugins(plugin.TemplateFuncMap(config.View))

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Database       `json:"Database"`
	Email     email.SMTPInfo          `json:"Email"`
	Recaptcha recaptcha.RecaptchaInfo `json:"Recaptcha"`
	Server    server.Server           `json:"Server"`
	Session   session.Session         `json:"Session"`
	Template  view.Template           `json:"Template"`
	View      view.View               `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// Create an Admin user is needed
func CreateAdmin() {
	// 1. if Admin user already exists then return
	// 2. if not then create an Admin user with a known password
	//    let Admin change these after login:
	//      email: admin@admin.admin
	//      password: admin

	// first_name := r.FormValue("first_name")
	// last_name := r.FormValue("last_name")
	// email := r.FormValue("email")
	// password, errp := passhash.HashString(r.FormValue("password"))
	// // If password hashing failed
	// if errp != nil {
	// 	log.Println(errp)
	// 	sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
	// 	sess.Save(r, w)
	// 	http.Redirect(w, r, "/register", http.StatusFound)
	// 	return
	// }

	// _, err := model.UserIdByEmail(email)
	// if err == sql.ErrNoRows { // If success (no user exists with that email)
	// 	ex := model.UserCreate(first_name, last_name, email, password)
	// 	// Will only error if there is a problem with the query
	// 	if ex != nil {
	// 		log.Println(ex)
	// 		sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
	// 		sess.Save(r, w)
	// 	} else {
	// 		sess.AddFlash(view.Flash{"Account created successfully for: " + email, view.FlashSuccess})
	// 		sess.Save(r, w)
	// 		http.Redirect(w, r, "/login", http.StatusFound)
	// 		return
	// 	}
	// } else if err != nil { // Catch all other errors
	// 	log.Println(err)
	// 	sess.AddFlash(view.Flash{"An error occurred on the server. Please try again later.", view.FlashError})
	// 	sess.Save(r, w)
	// } else { // Else the user already exists
	// 	sess.AddFlash(view.Flash{"Account already exists for: " + email, view.FlashError})
	// 	sess.Save(r, w)
	// }
}
