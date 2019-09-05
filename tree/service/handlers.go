package service

import (
	"io"
	"io/ioutil"
	"strconv"
	"tree/db"
	"tree/logger"
	"tree/models"
	"tree/utils"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Tree service")
}

func FindNodeChilds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	nId := ps.ByName("node_id")
	nodeId, err := strconv.Atoi(nId)
	if err != nil {
		logger.Error(err)
	}
	NodeChilds, err := db.FindNodeChilds(ctx, nodeId)
	if err != nil {
		logger.Error(err)
		utils.WriteError(&w, err)
	}
	if NodeChilds == nil {
		fmt.Fprintf(w, "[]")
		return
	}
	err = json.NewEncoder(w).Encode(NodeChilds)
	if err != nil {
		logger.Error(err)
		utils.WriteError(&w, err)
	}
}

func UpdateParent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	nId := ps.ByName("node_id")
	nodeId, err := strconv.Atoi(nId)
	if err != nil {
		logger.Error(err)
	}
	pId := ps.ByName("parent_id")
	parentId, err := strconv.Atoi(pId)
	if err != nil {
		logger.Error(err)
	}
	err = db.UpdateParent(ctx, nodeId, parentId)
	if err != nil {
		logger.Error(err)
		utils.WriteError(&w, err)
	}
	fmt.Fprintf(w, "Updated successfully")
}

func StoreNode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx := r.Context()

	var node models.Node

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		logger.Error(err)
	}
	if err := r.Body.Close(); err != nil {
		logger.Error(err)
	}

	// Save JSON to the node struct
	if err := json.Unmarshal(body, &node); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		err = json.NewEncoder(w).Encode(err)
		logger.Error(err)
		return
	}

	// storing at the database
	err = db.StoreNode(ctx, node)
	if err != nil {
		logger.Error(err)
		utils.WriteError(&w, err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
