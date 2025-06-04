package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	rayv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	"github.com/ray-project/kuberay/ray-operator/controllers/ray/utils"
)

func SetupJobServer(port string) {
	http.HandleFunc(utils.JobPath, ServeJobInfo)

	fmt.Printf("Starting HTTP server on port %s\n", port)
	serverAddr := ":" + port
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		panic(err)
	}
}

func ServeJobInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	jobID := strings.TrimPrefix(path, utils.JobPath)
	jobID = strings.TrimSuffix(jobID, "/")
	if jobID == "" {
		http.Error(w, "Job ID not provided", http.StatusBadRequest)
		return
	}

	jobInfo := utils.RayJobInfo{
		JobId:     jobID,
		JobStatus: rayv1.JobStatusSucceeded,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(jobInfo); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
