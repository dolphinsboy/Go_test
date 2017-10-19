package main
import (
	"net/http"
	"io"
	"log"
	"os"
	"io/ioutil"
	"html/template"
	"path"
	"strings"
)
const (
	UPLOAD_DIR  = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)
//templates := make(map[string]*template.Template)

func init() {
	files, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil{
		panic(err)
		return
	}

	var templateName, templatePath string

	for _, fileNo :=range files{
		templateName = fileNo.Name()
		if ext:=path.Ext(templateName); ext != ".html"{
			continue
		}

		templatePath = TEMPLATE_DIR + "/" + templateName
		t := template.Must(template.ParseFiles(templatePath))
		tmp1 := strings.Split(templateName, ".")[0]
		templates[tmp1] = t
	}
}

func checkError(w http.ResponseWriter, err error) bool{
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := readHtml(w,"upload", nil)

		//		t, err := template.ParseFiles("upload.html")
		if checkError(w, err){
			return
		}
//		t.Execute(w, nil)
		return
	}else if r.Method == "POST"{
		file, head, err := r.FormFile("image")

		if checkError(w, err){
			return
		}

		filename := head.Filename
		defer file.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if checkError(w, err){
			return
		}

		defer t.Close()
		_, err = io.Copy(t, file)

		if checkError(w, err){
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func isExists(path string) bool{
	_, err := os.Stat(path)
	if err == nil{
		return true
	}else {
		return false
	}
}

func listHandler(w http.ResponseWriter, r*http.Request) {
	files, err := ioutil.ReadDir(UPLOAD_DIR)
	if checkError(w, err){
		return
	}

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileno := range files{
		images = append(images, fileno.Name())
	}

	locals["images"] = images
	err =readHtml(w, "list", locals)
//	t, err := template.ParseFiles("list.html")
	if checkError(w, err){
		return
	}
//
//	t.Execute(w, locals)
}

func readHtml(w http.ResponseWriter, templateName string, locals map[string]interface{})(err error){
//	t, err := template.ParseFiles(templateName + ".html")
//	if err != nil{
//		return err
//	}
//	err = t.Execute(w, locals)
//	return err
	err = templates[templateName].Execute(w, locals)
	return
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists{
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		defer func() {
			if e, ok := recover().(error); ok{
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
			}
		}()
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))

	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		log.Fatal("ListenAndServe failed :", err.Error())
	}
}
