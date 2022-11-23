package main

import "gorm.io/gorm"

type Employee struct{ 
	gorm.Model 
	EmpName string `json:"empname"`
	EmpRole string `json:"emprole"`
	EmpDepart string `json:"empdepart"`
	EmpSalary float64 `json:"salary"`
	Email string `json:"email"`
}


