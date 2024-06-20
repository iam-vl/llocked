package controllers

import (
	"fmt"
	"net/http"

	"github.com/iam-vl/llocked/models"
)

type Users struct {
	Templates struct {
		// New views.Template
		New    TemplateExecuter
		SignIn TemplateExecuter
	}
	UserService *models.UserService
	// Coupons struct {
	// 	New Couponer
	// }
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	fmt.Println(data.Email)
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	user, err := u.UserService.Create(email, pwd)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "user created: %+v", user)
}

// func (u Users) Coupon(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("processing")
// 	// fmt.Fprint(w, "processing coupon...")
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "unable to parse form submission", http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Fprintf(w, "<p>Coupon: %s</p>", r.PostForm.Get("coupon"))
// 	fmt.Fprintf(w, "<p>PName: %s</p>", r.PostForm.Get("name"))
// 	// u.Coupons.New.Execute(w, nil)
// }

// func (u Users) NewCoupon(w http.ResponseWriter, r *http.Request) {
// 	// err
// }
