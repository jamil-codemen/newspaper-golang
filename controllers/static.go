package controllers

import "newspaper/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("main", "views/static/home.html"),
		Contact: views.NewView("main", "views/static/contact.html"),
	}
}


type Static struct {
	Home *views.View
	Contact *views.View
}