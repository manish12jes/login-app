package component

import(
	"html/template"
	"path"

)


func TemplateRenderer(filepath string) (*template.Template, error){
	fpHeader := path.Join("view", "admin/header.html")
	fpScript := path.Join("view", "admin/script.html")
	fpSidebar := path.Join("view", "admin/sidebar.html")
	fpTopNavbar := path.Join("view", "admin/navbar.html")
	fpContent := path.Join("view", "admin/" + filepath)
	templ, err := template.ParseFiles(fpContent, fpHeader, fpScript, fpSidebar, fpTopNavbar)
	return templ, err
}