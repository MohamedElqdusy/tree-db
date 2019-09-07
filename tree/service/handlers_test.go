package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"tree/db"
	"tree/models"

	"github.com/julienschmidt/httprouter"
)

func TestFindNodeChilds(t *testing.T) {

	mockStore := db.InitMockStore()
	nodes := []models.Node{
		{
			ID:      0,
			Content: "node_0",
			Root:    0,
			Parent:  0,
			Height:  0,
			Path:    "0,",
		},
		{
			ID:      1,
			Content: "node_1",
			Root:    0,
			Parent:  0,
			Height:  1,
			Path:    "0,1,",
		},
		{
			ID:      2,
			Content: "node_2",
			Root:    0,
			Parent:  1,
			Height:  2,
			Path:    "0,1,2,",
		},
		{
			ID:      3,
			Content: "node_3",
			Root:    0,
			Parent:  1,
			Height:  2,
			Path:    "0,1,3,",
		},
		{
			ID:      4,
			Content: "node_4",
			Root:    0,
			Parent:  2,
			Height:  3,
			Path:    "0,1,2,4,",
		},
		{
			ID:      5,
			Content: "node_5",
			Root:    0,
			Parent:  2,
			Height:  3,
			Path:    "0,1,2,5,",
		},
		{
			ID:      6,
			Content: "node_6",
			Root:    0,
			Parent:  4,
			Height:  4,
			Path:    "0,1,2,4,6,",
		},
		{
			ID:      7,
			Content: "node_7",
			Root:    0,
			Parent:  5,
			Height:  4,
			Path:    "0,1,2,5,7,",
		},
	}
	mockStore.On("FindNodeChilds").Return(nodes, nil).Once()

	req, err := http.NewRequest("GET", "/nodechilds/0/", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	router.Handle("GET", "/nodechilds/:node_id/", FindNodeChilds)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := nodes

	b := []models.Node{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// the expectations that we defined in the `On` method are asserted here
	mockStore.AssertExpectations(t)
}

func TestUpdateParent(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("UpdateParent").Return(nil).Once()

	req, err := http.NewRequest("PUT", "/update/2/3", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := httprouter.New()
	router.Handle("PUT", "/update/:node_id/:parent_id", UpdateParent)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func TestStoreNode(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("StoreNode").Return(nil).Once()
	node := models.Node{
		ID:      0,
		Content: "node_0",
		Root:    0,
		Parent:  0,
		Height:  0,
		Path:    "0,",
	}

	nodeJson, err := json.Marshal(node)

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/node", bytes.NewBuffer(nodeJson))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := httprouter.New()
	router.Handle("POST", "/node", StoreNode)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}
