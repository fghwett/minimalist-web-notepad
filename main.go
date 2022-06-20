package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	port     = 9099
	savePath = "_tmp"
)

//go:embed static/*
var static embed.FS

func index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !strings.Contains(path, "/") {
		jump(w, r)
		return
	}

	path = strings.TrimPrefix(path, "/")
	b, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, path)
	if strings.Contains(path, "/") || !b || len(path) > 16 {
		jump(w, r)
		return
	}

	filePath := filepath.Join(savePath, path)

	if r.Method == http.MethodPost {
		if e := r.ParseForm(); e != nil {
			log.Printf("parse form error: %s\n", e)
			return
		}
		if !r.PostForm.Has("text") {
			return
		}
		text := r.PostFormValue("text")
		if text == "" {
			// 删除文件
			if _, e := os.Stat(filePath); os.IsNotExist(e) {
				log.Printf("delete file -> %s is not exist", filePath)
			} else if e != nil {
				log.Printf("delete file -> check %s stat error: %s\n", filePath, e)
			} else {
				if e = os.Remove(filePath); e != nil {
					log.Printf("delete file -> delete %s error: %s \n", filePath, e)
				}
			}
		} else {
			// 创建文件
			if e := ioutil.WriteFile(filePath, []byte(text), 0666); e != nil {
				log.Printf("write %s error: %s, text: %s\n", filePath, e, text)
			}
		}
		return
	} else if r.Method == http.MethodGet {
		if r.URL.Query().Has("text") {
			text := r.URL.Query().Get("text")
			if text == "" {
				// 删除文件
				if _, e := os.Stat(filePath); os.IsNotExist(e) {
					log.Printf("delete file -> %s is not exist", filePath)
				} else if e != nil {
					log.Printf("delete file -> check %s stat error: %s\n", filePath, e)
				} else {
					if e = os.Remove(filePath); e != nil {
						log.Printf("delete file -> delete %s error: %s \n", filePath, e)
					}
				}
			} else {
				// 创建文件
				if e := ioutil.WriteFile(filePath, []byte(text), 0666); e != nil {
					log.Printf("write %s error: %s, text: %s\n", filePath, e, text)
				}
			}
			_, _ = w.Write([]byte("ok"))
			return
		}

	}

	ua := r.Header.Get("user-agent")
	if r.URL.Query().Has("raw") || strings.HasPrefix(ua, "curl") || strings.HasPrefix(ua, "Wget") {
		if _, e := os.Stat(filePath); e != nil {
			http.NotFound(w, r)
		} else {
			w.Header().Set("Content-type", "text/plain")
			if c, e := ioutil.ReadFile(filePath); e != nil {
				log.Printf("read %s error: %s", filePath, e)
				http.Error(w, "read file error", http.StatusInternalServerError)
			} else {
				_, _ = w.Write(c)
			}
		}
		return
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content string
	if _, e := os.Stat(filePath); os.IsNotExist(e) {
		content = ""
	} else if e != nil {
		log.Printf("check %s stat error: %s\n", filePath, e)
		return
	} else {
		if c, e := ioutil.ReadFile(filePath); e != nil {
			log.Printf("read %s error: %s\n", filePath, e)
			return
		} else {
			content = string(c)
		}
	}

	tem, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Printf("read index error: %s \n", err)
		return
	}
	e := tem.Execute(w, struct {
		Title   string
		Content string
	}{
		Title:   path,
		Content: content,
	})
	if e != nil {
		log.Printf("write html error: %s\n", e)
	}
}

func jump(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/"+randStr(), http.StatusFound)
}

func randStr() string {
	words := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	str := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		index := rand.Intn(len(words))
		str += string(words[index])
	}
	return str
}

func main() {
	fs := http.FileServer(http.FS(static))

	http.Handle("/static/", fs)
	http.HandleFunc("/", index)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("run server error: ", err)
		return
	}
}
