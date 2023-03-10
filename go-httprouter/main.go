package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// membuat struct middlerware
type LogMiddleware struct {
	http.Handler
}

// membuat midleware, ini akan dieksekusi sebelum masuk ke request
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ini midleware")
	middleware.Handler.ServeHTTP(writer, request)
}

func main() {
	router := httprouter.New()

	// how to create get in httprouter
	router.GET("/product/:id/size/:size", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		size, _ := strconv.Atoi(params.ByName("size"))
		text := fmt.Sprintf("haloo ini produk dengan id %s dan ukuran %d", params.ByName("id"), size)
		fmt.Fprintf(writer, text)
	})

	// biasanya pada catch all parameter *filtepath, digunakan untuk url ata path
	router.GET("/image/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := fmt.Sprintf("haloo ini poto dengan image %s", params.ByName("image"))
		fmt.Fprintf(writer, image)
	})

	// panic handler
	router.GET("/panic", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		panic("panic gaaa? panicc woyyy")
	})

	// ketika request panic seperti diatas, dsini membuat panic hadler
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic :", error)
	}

	// ketika request tidak ketemu urlnya maka akan dikembalikan not found, kita dapat untuk custom sebagai berikut
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "tidak ditemukann yaaa")
	})

	// ketika kita mempunyai http method akan tetapi saat memanggil salah http method maka akan terjadi method not allowed
	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "ini Post")
	})

	// ini dapat kita lakukan mengubah error method not allowed
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "salah method woy")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "localhost:3000", nil)
	recorder := httptest.NewRecorder()
	middleware.ServeHTTP(recorder, request)

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
