package handler_func

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type studentInfo struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Address string `json:"address"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	student1 := studentInfo{Name: "Ishaq Ali", Id: "1530869042", Address: "Jigatola, Dhaka"}
	student2 := studentInfo{Name: "Sajjid Hossain", Id: "1530869043", Address: "Dhanmondi, Dhaka"}

	studentList := []studentInfo{student1, student2}

	for i := 0; i < 2; i++ {
		marshaledData, err := json.Marshal(studentList[i])

		if err != nil {
			log.Debug().Msg("marshal error")
			return
		}

		io.WriteString(w, string(marshaledData)+"\n")
	}

}
