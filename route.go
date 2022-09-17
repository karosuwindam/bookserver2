package main

import (
	"bookserver/fileupload"
	"bookserver/textread"
	"bookserver/webserver"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	ROUTE = "./html"
)

type Htmldata struct {
	uploadpass *fileupload.UploadPass
}

//静的HTMLのページを返す
func viewhtml(w http.ResponseWriter, r *http.Request) {
	textdata := []string{".html", ".htm", ".css", ".js"}
	upath := r.URL.Path
	tmp := map[string]string{}
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	log.Println(r.Method + ":" + r.URL.Path)
	if upath == "/" {
		upath += "index.html"
	}
	if !textread.Exists(ROUTE + upath) {
		w.WriteHeader(404)
		log.Printf("ERROR request:%v\n", r.URL.Path)
		return
	} else {
		for _, data := range textdata {
			if len(upath) > len(data) {
				if upath[len(upath)-len(data):] == data {
					fmt.Fprint(w, textread.ConvertData(textread.ReadHtml(ROUTE+upath), tmp))
					return
				}
			}
		}
		buffer := textread.ReadOther(ROUTE + upath)
		// bodyに書き込み
		w.Write(buffer)
	}
	return
}

func urlAnalysis(url string) []string {
	tmp := []string{}
	for _, str := range strings.Split(url[1:], "/") {
		tmp = append(tmp, str)
	}
	return tmp
}

func (t *Htmldata) v1(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method + ":" + r.URL.Path)
	urldata := urlAnalysis(r.URL.Path)
	if len(urldata) > 1 {
		switch urldata[1] {
		case "":
			w.WriteHeader(400)
			fmt.Fprintf(w, "Err API request")
		case "upload":
			fileupload.FIleupload(t.uploadpass, w, r)
		default:
			fmt.Fprintf(w, "%s", r.URL.Path)
		}

	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Err API request")
	}
	return

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %v", r.URL.Path[1:])
}

func setupbaseRoute() (Htmldata, error) {
	var err error
	output := Htmldata{}
	uppass, err := fileupload.Setup()
	if err != nil {
		return output, err
	}
	output.uploadpass = uppass
	return output, nil
}

func (t *Htmldata) setupRoute(cfg *webserver.SetupServer) {
	cfg.Add("/v1/", t.v1)
	cfg.Add("/", viewhtml)
}
