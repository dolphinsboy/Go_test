package main
import (
	"net/http"
	"io"
	"log"
	"os"
)
const (
	UPLOAD_DIR  = "./uploads"
)

func checkError(w http.ResponseWriter, err error) bool{
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//设置返回的格式为Html，否则输出到浏览器上就是字符串
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" " +
		" enctype=\"multipart/form-data\">" +
		"Choose an image to upload:<input name=\"image\" type=\"file\" />" +
		"<input type=\"submit\" value=\"Upload\" />" + "</form>")
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		log.Fatal("ListenAndServe failed :", err.Error())
	}
}
