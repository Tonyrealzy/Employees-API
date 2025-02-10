package usecase

import (
	"encoding/json"
	"log"
	"my-crud-project/models"
	"my-crud-project/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data omitempty"`
	Error string      `json:"error omitempty"`
}

func (s *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error decoding request body: ", err)
		res.Error = err.Error()
		return
	}
	employee.EmployeeID = uuid.NewString()
	repo := repository.EmployeeRepo{MongoCollection: s.MongoCollection}

	insertID, err := repo.InsertEmployee(&employee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error inserting employee", err)
		res.Error = err.Error()
		return
	}
	res.Data = employee.EmployeeID
	w.WriteHeader(http.StatusOK)
	log.Println("created employee with ID: ", insertID, employee)
}

func (s *EmployeeService) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	employeeID := mux.Vars(r)["id"]
	log.Println("fetched employeeID: ", employeeID)

	repo := repository.EmployeeRepo{MongoCollection: s.MongoCollection}
	employee, err := repo.FindEmployeeById(employeeID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error fetching employee", err)
		res.Error = err.Error()
		return
	}
	res.Data = employee
	w.WriteHeader(http.StatusOK)
}

func (s *EmployeeService) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: s.MongoCollection}
	employee, err := repo.FindAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error fetching employee", err)
		res.Error = err.Error()
		return
	}
	res.Data = employee
	w.WriteHeader(http.StatusOK)
}
func (s *EmployeeService) UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
}
func (s *EmployeeService) DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {

}
func (s *EmployeeService) DeleteAllEmployees(w http.ResponseWriter, r *http.Request) {

}
