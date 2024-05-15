package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"

	"test.db/internal/models"
	"test.db/internal/utils"
)

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(config string) (*Repository, error) {
	db, err := sqlx.Connect("postgres", config)
	if err != nil {
		return nil, err
	}

	return &Repository{DB: db}, nil
}
func (r *Repository) MigrateDB() error {

	driver, err := pgx.WithInstance(r.DB.DB, &pgx.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"pgx", driver,
	)
	if err != nil {
		return fmt.Errorf("Could not create migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil /*&& !errors.Is(err, migrate.ErrNoChange)*/ {
		return err
	}

	return nil
}

func (r *Repository) Close() error {
	if r.DB != nil {
		return r.DB.Close()
	}
	return nil
}

func (r *Repository) Task1ConnectToDB() error {
	err := r.DB.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to the database")
	return nil
}

func (r *Repository) Task2ExecuteQuery() error {
	rows, err := r.DB.Query("SELECT * FROM departments")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var departmentid int
		var departmentname string
		err = rows.Scan(&departmentid, &departmentname)
		if err != nil {
			return err
		}
		fmt.Println(departmentid, departmentname)
	}
	return nil
}

func (r *Repository) Task3QueryWithParams(paramValue string) error {
	var d []models.Employees
	query := "SELECT employeeid, firstname, lastname, departmentid FROM employees WHERE departmentid = $1"
	err := r.DB.Select(&d, query, paramValue)
	if err != nil {
		return err
	}
	utils.PrintStructure(d)
	return nil
}

func (r *Repository) Task4InsertData(firstname, lastname string, departmentid int) error {
	_, err := r.DB.Exec("INSERT INTO employees (firstname, lastname, departmentid) VALUES ($1, $2, $3)", firstname, lastname, departmentid)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Task5UpdateData(firstname string, departmentid int) error {
	_, err := r.DB.Exec("UPDATE employees SET firstname = $1 WHERE employeeid = $2", firstname, departmentid)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Task6Transaction(departmentname string, idDep int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		err = fmt.Errorf("Произошла ошибка открытия транзакции\nError: %s", err.Error())
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var existingDepartmentID int
	err = tx.QueryRow("SELECT departmentid FROM employees WHERE departmentid = $1", idDep).Scan(&existingDepartmentID)
	if err != nil {
		err = fmt.Errorf("Такого departmentid = %d для замены нет\nError: %s", idDep, err.Error())
		return err
	}

	var departmentid int
	err = tx.QueryRow("INSERT INTO departments (departmentname) VALUES ($1) RETURNING departmentid", departmentname).Scan(&departmentid)
	if err != nil {
		err = fmt.Errorf("Такое имя departmentname = %s уже есть\nError: %s", departmentname, err.Error())
		return err
	}

	_, err = tx.Exec("UPDATE employees SET departmentid = $1 WHERE departmentid = $2", departmentid, idDep)
	if err != nil {
		err = fmt.Errorf("Произошла ошибка обновления departmentid\nError: %s", err.Error())
		return err
	}

	return nil
}

func (r *Repository) DBClose() {
	fmt.Println("\nStart close to the database................")
	if r.DB != nil {
		errCloseDB := r.DB.Close()
		if errCloseDB != nil {
			fmt.Println("Error close db: ", errCloseDB)
			return
		} else {
			fmt.Println("Successfully close to the database")
		}
	} else {
		fmt.Println("Error close db ")
		return
	}
}
