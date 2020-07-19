package main

import "fmt"
import "strings"
import "sync"
import "time"
import "encoding/json"
import "io/ioutil"
import "net/http"

type Task struct {
	ID					string 		`json:"id"`
	Title				string 		`json:"title"`
	CreatedAt			time.Time	`json:"createdAt"`
	CompletedAt			time.Time	`json:"completedAt"`
	Status				string		`json:"status"`
}

type taskHandlers struct {
	sync.Mutex
	store map[string]Task
}	

func (h *taskHandlers) get(w http.ResponseWriter, r *http.Request) {
	tasks := make([]Task, len(h.store))

	h.Lock()
	i := 0
	for _, task := range h.store {
		tasks[i] = task
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(tasks)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *taskHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	contentType := r.Header.Get("content-type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json, but got '%s", contentType)))
		return
	}

	var task Task
	err = json.Unmarshal(bodyBytes, &task)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	task.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	task.CreatedAt = time.Now()

	if task.Status == ""{
		task.Status = "New"
	}

	h.Lock()
	h.store[task.ID] = task
	defer h.Unlock()
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *taskHandlers) getTask(w http.ResponseWriter, r *http.Request) {
	url_parts := strings.Split(r.URL.String(), "/")
	if len(url_parts) != 3{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	task, ok := h.store[url_parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(task)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *taskHandlers) updateTask(w http.ResponseWriter, r *http.Request) {
	url_parts := strings.Split(r.URL.String(), "/")
	if len(url_parts) != 3{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	_, ok := h.store[url_parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	contentType := r.Header.Get("content-type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json, but got '%s", contentType)))
		return
	}

	var task Task
	err = json.Unmarshal(bodyBytes, &task)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if task.ID != url_parts[2]{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID mismatch"))
		return
	}

	h.Lock()
	h.store[task.ID] = task
	defer h.Unlock()
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}



func (h *taskHandlers) tasks(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *taskHandlers) task(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		h.getTask(w, r)
		return
	case "PUT":
		h.updateTask(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func newTaskHandlers() *taskHandlers {
	return &taskHandlers{
		store: map[string]Task{},
	}
}

func main() {
	taskHandlers := newTaskHandlers()
	http.HandleFunc("/tasks", taskHandlers.tasks)
	http.HandleFunc("/tasks/", taskHandlers.task)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}