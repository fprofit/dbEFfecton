package taskgorm

import (
	"fmt"
	"time"

	"test.db/internal/entry"
	"test.db/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Task7(envConfig entry.EnvConfig) error {
	fmt.Println("\nStart task 7 GORM")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Moscow", envConfig.Host, envConfig.User, envConfig.Password, envConfig.DBName, envConfig.Port, envConfig.DBSslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	dbInstance, err := db.DB()
	if err != nil {
		return err
	}
	dbInstance.SetConnMaxIdleTime(time.Second * 5)
	defer dbInstance.Close()

	if errGet := GetDepartments(db); errGet != nil {
		return errGet
	}
	departmentID, errAdd := AddDepartments(db)
	if errAdd != nil {
		return errAdd
	} else {
		errEdit := EditNameDepartments(db, departmentID)
		if errEdit != nil {
			return errEdit
		}
	}
	return nil
}

func GetDepartments(db *gorm.DB) error {
	fmt.Println("\nGet Departments")
	var departments []models.Departments
	if err := db.Find(&departments).Error; err != nil {
		fmt.Println("Error Get Departments: ", err)
	} else {
		for _, v := range departments {
			fmt.Printf("%d. %s\n", v.Departmentid, v.Departmentname)
		}
	}
	fmt.Println("Successfully Get Departments")
	return nil
}

func AddDepartments(db *gorm.DB) (int, error) {
	fmt.Println("\nAdd Department")
	department := &models.Departments{
		Departmentname: "GORM",
	}
	result := db.Create(department)
	if result.Error != nil {
		fmt.Println("Error Add Departments: ", result.Error)
		return 0, result.Error
	}
	departmentID := department.Departmentid
	fmt.Println("Departmentid: ", departmentID)
	return departmentID, nil
}

func EditNameDepartments(db *gorm.DB, id int) error {
	fmt.Println("\nEdit Department id = ", id)
	if err := db.Model(&models.Departments{}).Where("departmentid = ?", id).Update("departmentname", "GORM UPDATE").Error; err != nil {
		fmt.Println("Error Add Departments: ", err)
		return err
	}

	fmt.Println("Department name updated successfully")
	return nil
}
