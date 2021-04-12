package controllers

import (
	"fmt"
	"net/http"
	"newspaper/views"
)


 func NewUsers() *Users{
	 return &Users{
		 NewView: views.NewView("main","views/users/new.html"),
	 }
 }

 type Users struct {
	 NewView *views.View
 }

 type SignupForm struct {
	 Email string `json:"email"`
	 Password string `json:"password"`
 }
 //GET/signup
 func (u *Users)New(w http.ResponseWriter, r *http.Request){
	 if err :=u.NewView.Render(w,nil);err!=nil{
		 fmt.Println(err)
	 }
 }
 //POST/signup
 func (u *Users)Create(w http.ResponseWriter, r *http.Request){
var form SignupForm
parseForm(r, &form)
fmt.Fprintln(w,form)
	
}