package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	// "github.com/backend-guchitter-app/config"
	"github.com/backend-guchitter-app/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationFilePath = "file://./db/migrations/"
)

func main() {

	// extract command line args
	flag.Parse()
	args := flag.Args()
	fmt.Println("args:", args)

	// ↓↓↓これもうまく行かない
	// isTest := flag.CommandLine.Lookup("test")

	isTest := false
	// ↓↓↓何故か一回も呼ばれない...
	// flag.Visit(func(f *flag.Flag) {
	// 	fmt.Println("f.Name:", f.Name)
	// 	if f.Name == "test" {
	// 		isTest = true
	// 	}
	// })
	for _, arg := range args {
		if arg == "--test" {
			isTest = true
		}
	}

	var m *migrate.Migrate
	if isTest {
		// migrate for test db
		m = newMigrate(true)
	} else {
		m = newMigrate(false)
	}
	v, dirty, versionErr := m.Version()
	if versionErr != nil {
		fmt.Println(errors.Wrap(versionErr, "error at m.Version()"))
	}
	fmt.Println("version:", v)
	if dirty {
		fmt.Printf("version %d is dirty\n", v)
		m.Force(int(v))
	}

	// execute migration by command line args
	var err error
	if len(args) > 0 && args[0] == "up" {
		err = m.Up()
	} else if len(args) > 0 && args[0] == "down" {
		err = m.Down()
	} else {
		fmt.Println("command line args[0] must be 'up' or 'down'")
		panic("")
	}
	if err.Error() == "no change" {
		fmt.Println("no change")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("migration finished")
		v, _, _ := m.Version()
		fmt.Println("version:", v)
	}
}

func newMigrate(isTest bool) *migrate.Migrate {
	dsn := config.GetDsn(isTest)
	fmt.Println("dsn:", dsn)

	db, openErr := sql.Open("mysql", dsn)
	if openErr != nil {
		fmt.Println(errors.Wrap(openErr, "error at sql.Open()"))
	}

	// error at mysql.WithInstance(): Error 1045: Access denied for user 'nakajimahidenari'@'172.27.0.1' (using password: YES)
	// →dsnがおかしい。環境変数「USER」が正しく読み込まれていない...?
	driver, instanceErr := mysql.WithInstance(db, &mysql.Config{})
	if instanceErr != nil {
		fmt.Println(errors.Wrap(instanceErr, "error at mysql.WithInstance()"))
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationFilePath,
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println(errors.Wrap(err, "error at migrate.NewWithDatabaseInstance()"))
	}

	return m
}
