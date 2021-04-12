package views

import (
	"fmt"
	"html/template"
	"net/http"
)

func NewView (layout string,files ...string)*View {
	files = append(files, "views/layouts/main.html","views/layouts/navbar.html","views/layouts/footer.html")
	t,err:= template.ParseFiles(files...)
	if err!= nil {
		fmt.Println(err)
	}
	return &View{
		Template:t,
		Layout:layout,
		
	}
}

type View struct {
	Template *template.Template
	Layout string
}


func(v *View)ServeHTTP(w http.ResponseWriter, r *http.Request){
	if err := v.Render(w,nil);err !=nil{
		panic(err)
	}
}

func (v *View) Render (w http.ResponseWriter, data interface{})error{
	w.Header().Set("Content-Type","text/html")
	return v.Template.ExecuteTemplate(w,v.Layout,data)
}