package main

import (
	"encoding/json"
	"net/http"
	"presenter"
)

var systemInformationPresenter = presenter.NewSystemInformationPresenter()

func GetCpuCount(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.CpuCount()
	writeGetRequest(res, err, writer, r)
}

func GetGpuCount(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.GpuCount()
	writeGetRequest(res, err, writer, r)
}

func GetMemorySize(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.MemorySize()
	writeGetRequest(res, err, writer, r)
}

func GetDiskPartition(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.DiskPartition()
	writeGetRequest(res, err, writer, r)
}

func writeGetRequest(object interface{}, err error, writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if request.Method == "GET" {
		if err != nil {
			writer.WriteHeader(http.StatusNoContent)
			json.NewEncoder(writer).Encode(err)
		} else {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(object)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/cpu/count", GetCpuCount)
	http.HandleFunc("/gpu/count", GetGpuCount)
	http.HandleFunc("/memory/size", GetMemorySize)
	http.HandleFunc("/disk/partition", GetDiskPartition)
	http.ListenAndServe(":8080", nil)
}
