package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"my-crud-project/models"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(employee *models.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), employee)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeById(employeeId string) (*models.Employee, error) {
	var employee models.Employee

	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "employee_id", Value: employeeId}}).Decode(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepo) FindAllEmployees() ([]models.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var employees []models.Employee
	err = results.All(context.Background(), &employees)
	if err != nil {
		return nil, fmt.Errorf("error while decoding employees: %s", err.Error())
	}
	return employees, nil
}

func (r *EmployeeRepo) UpdateEmployeeById(employeeId string, employee *models.Employee) (int64, error) {
	results, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "employee_id", Value: employeeId}}, bson.D{{Key: "$set", Value: employee}})
	if err != nil {
		return 0, err
	}
	return results.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeById(employeeId string) (int64, error) {
	results, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "employee_id", Value: employeeId}})
	if err != nil {
		return 0, err
	}
	return results.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployees() (int64, error) {
	results, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return results.DeletedCount, nil
}
 