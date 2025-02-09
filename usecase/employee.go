package usecase

import (
	"net/http"

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

}
func (s *EmployeeService) GetEmployeeById(w http.ResponseWriter, r *http.Request) {

}
func (s *EmployeeService) GetAllEmployees(w http.ResponseWriter, r *http.Request) {

}
func (s *EmployeeService) UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {

}
func (s *EmployeeService) DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {

}
func (s *EmployeeService) DeleteAllEmployees(w http.ResponseWriter, r *http.Request) {

}