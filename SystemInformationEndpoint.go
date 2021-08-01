package main

import (
	"encoding/json"
	"net/http"
	"presenter"
)

func exposeSystemInformationEndpoints() {
	http.HandleFunc("/cpu/count", cpuCountRequestHandler)
	http.HandleFunc("/gpu/count", gpuCountRequestHandler)
	http.HandleFunc("/memory/size", memorySizeRequestHandler)
	http.HandleFunc("/disk/partition", diskSizeRequestHandler)
}

var systemInformationPresenter = presenter.NewSystemInformationPresenter()

func cpuCountRequestHandler(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.CpuCount()
	writeGetRequest(res, err, writer, r)
}

func gpuCountRequestHandler(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.GpuCount()
	writeGetRequest(res, err, writer, r)
}

func memorySizeRequestHandler(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.MemorySize()
	writeGetRequest(res, err, writer, r)
}

func diskSizeRequestHandler(writer http.ResponseWriter, r *http.Request) {
	res, err := systemInformationPresenter.DiskPartition()
	writeGetRequest(res, err, writer, r)
}

func writeGetRequest(object interface{}, err error, writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if request.Method == http.MethodGet {
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte(err.Error()))
		} else {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(object)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(`{"message": "not found"}`))
	}
}
