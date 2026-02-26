package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Main(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "index.html")
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	file, header, err := request.FormFile("myFile")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	defer func() { _ = file.Close() }()

	bytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	result := service.ConvertBytes(bytes)
	if err = os.WriteFile(fmt.Sprintf("%s.%s", time.Now().UTC().String(), filepath.Ext(header.Filename)), result, 0755); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	_, _ = writer.Write(result)
}
