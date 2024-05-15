package models

type Employees struct {
	EmployeeID   int    `db:"employeeid"`
	Firstname    string `db:"firstname"`
	Lastname     string `db:"lastname"`
	Departmentid int    `db:"departmentid"`
}

type Departments struct {
	Departmentid   int    `db:"departmentid" gorm:"primaryKey;autoIncrement"`
	Departmentname string `db:"departmentname"`
}
