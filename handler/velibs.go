package handler

import (
	"VelibFinder/dao"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// handleGetVelib returns the list of Velib stations capacities near Splio HQ
func (s *server) handleGetVelib() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		vs, err := dao.GetVelibsStation(); if err != nil {
			resp.Error = "Internal Server Error"
			resp.Message = "Error getting Velib stations near Splio"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
		}


		vs, err = dao.GetAvailableVelibsForStation(vs); if err != nil {
			resp.Error = "Internal Server Error"
			resp.Message = "Error reading Velib API results"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
		}

		resp.Data = append(resp.Data, vs)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		resp.Meta.Query = fmt.Sprintln("List of Velib station capacities near Splio HQ")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

