package controller

import (
	"fmt"
	"log"
	"net/http"

	"gowebapp/model"
	"gowebapp/shared/passhash"
	"gowebapp/shared/session"
	"gowebapp/shared/view"

	"github.com/gorilla/sessions"
	"github.com/josephspurrier/csrfbanana"
)

// loginAttempt increments the number of login attempts in sessions variable
func loginAttempt(sess *sessions.Session) {
	// Log the attempt
	if sess.Values["login_attempt"] == nil {
		sess.Values["login_attempt"] = 1
	} else {
		sess.Values["login_attempt"] = sess.Values["login_attempt"].(int) + 1
	}
}

// clearSessionVariables clears all the current session values
func clearSessionVariables(sess *sessions.Session) {
	// Clear out all stored values in the cookie
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}

func LoginGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "login"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	// Refill any form fields
	view.Repopulate([]string{"email"}, r.Form, v.Vars)
	v.Render(w)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Prevent brute force login attempts by not hitting database and pretending like it was invalid :-)
	if sess.Values["login_attempt"] != nil && sess.Values["login_attempt"].(int) >= 5 {
		log.Println("Brute force login prevented")
		sess.AddFlash(view.Flash{"Sorry, no brute force :-)", view.FlashNotice})
		sess.Save(r, w)
		LoginGET(w, r)
		return
	}

	// Validate with required fields
	if validate, missingField := view.Validate(r, []string{"email", "password"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		LoginGET(w, r)
		return
	}

	// Form values
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Get database result
	result, err := model.UserByEmail(email)
	fmt.Printf("model.UserByEmail=%v\nresult=%T=%+v\nerr=%T=%+v\n", email, result, result, err, err)
	// cls: force Login successfully
	clearSessionVariables(sess)
	sess.AddFlash(view.Flash{"Login successful!", view.FlashSuccess})
	sess.Values["email"] = email
	sess.Values["first_name"] = "Molly"
	sess.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
	return

	// Determine if user exists
	// if err == sql.ErrNoRows {
	if err != nil {
		loginAttempt(sess)
		sess.AddFlash(view.Flash{"Password is incorrect - Attempt: " + fmt.Sprintf("%v", sess.Values["login_attempt"]), view.FlashWarning})
		sess.Save(r, w)
	} else if err != nil {
		// Display error message
		log.Println(err)
		sess.AddFlash(view.Flash{"There was an error. Please try again later.", view.FlashError})
		sess.Save(r, w)
	} else if passhash.MatchString(result.Password, password) {
		// if result.Status_id != 1 {
		// 	// User inactive and display inactive message
		// 	sess.AddFlash(view.Flash{"Account is inactive so login is disabled.", view.FlashNotice})
		// 	sess.Save(r, w)
		// } else {
		// Login successfully
		clearSessionVariables(sess)
		sess.AddFlash(view.Flash{"Login successful!", view.FlashSuccess})
		// sess.Values["id"] = result.Id
		sess.Values["email"] = email
		sess.Values["first_name"] = result.First_name
		sess.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
		return
		// }
	} else {
		loginAttempt(sess)
		sess.AddFlash(view.Flash{"Password is incorrect - Attempt: " + fmt.Sprintf("%v", sess.Values["login_attempt"]), view.FlashWarning})
		sess.Save(r, w)
	}

	// Show the login page again
	LoginGET(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// If user is authenticated
	// if sess.Values["id"] != nil {
	if sess.Values["email"] != nil {
		clearSessionVariables(sess)
		sess.AddFlash(view.Flash{"Goodbye!", view.FlashNotice})
		sess.Save(r, w)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
