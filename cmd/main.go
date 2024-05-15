package main

import (
	"fmt"

	"test.db/internal/entry"
	"test.db/internal/repository"
	"test.db/internal/taskgorm"
)

func main() {
	envConfig, errEnvConfig := entry.InitializeEnv()
	if errEnvConfig != nil {
		fmt.Println("Error read config: ", errEnvConfig)
		return
	}

	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", envConfig.User, envConfig.Password, envConfig.Host, envConfig.Port, envConfig.DBName, envConfig.DBSslmode)

	db, err := repository.NewRepository(dbString)
	if err != nil {
		fmt.Println("Error create DB: ", err)
		return
	}

	fmt.Println("\nStart task 1 ................")
	err1 := db.Task1ConnectToDB()
	if err1 != nil {
		fmt.Println("Error task 1: ", err1)
		return
	}

	fmt.Println("\nStart task 8 ................")
	err8 := db.MigrateDB()
	if err8 != nil {
		fmt.Println("Error task 8:", err8)
	}

	fmt.Println("\nStart task 2 ................")
	err2 := db.Task2ExecuteQuery()
	if err2 != nil {
		fmt.Println("Error task 2: ", err2)
		return
	}

	fmt.Println("\nStart task 3 ................")
	err3 := db.Task3QueryWithParams("1")
	if err3 != nil {
		fmt.Println("Error task 3: ", err3)
		return
	}

	fmt.Println("\nStart task 4 ................")
	err4 := db.Task4InsertData("Roman", "Kostetskii", 1)
	if err4 != nil {
		fmt.Println("Error task 4: ", err4)
		return
	} else {
		fmt.Println("Successfully task 4")
	}

	fmt.Println("\nStart task 5 ................")
	err5 := db.Task5UpdateData("Enot", 5)
	if err5 != nil {
		fmt.Println("Error task 5: ", err5)
		return
	} else {
		fmt.Println("Successfully task 5")
	}

	fmt.Println("\nStart task 6 ................")
	err6 := db.Task6Transaction("developer420", 42)
	if err6 != nil {
		fmt.Println("Error task 6: ", err6)
	} else {
		fmt.Println("Successfully task 6")
	}

	err7 := taskgorm.Task7(envConfig)
	if err7 != nil {
		fmt.Println("Error task 7: ", err7)
	} else {
		fmt.Println("Successfully task 7")
	}

	defer db.DBClose()

}
