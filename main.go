package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	dataSource := GetParams()
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		fmt.Println("False")
		fmt.Println(err)
		return
	}
	_, err = db.Query("SELECT 1")
	if err != nil {
		fmt.Println("False")
		fmt.Println(err)
		return
	}
	if err != nil {
		return
	}
	db.Close()
	fmt.Println("True")
}

func GetParams() string {
	res := make(map[string]string)
	res["-h"] = "localhost"
	res["-P"] = "5432"
	res["-u"] = "postgres"
	res["-p"] = ""
	for i := 1; i < len(os.Args); i += 2 {
		switch os.Args[i] {
		case "-h":
			if i+1 >= len(os.Args) {
				fmt.Println("Miss param -h")
				fmt.Println("example usage: psql-test-conn -h localhost -P 5432 -u postgres -p 123")
				os.Exit(2)
			}
			res[os.Args[i]] = os.Args[i+1]
		case "-p":
			if i+1 >= len(os.Args) {
				fmt.Println("Miss param -p")
				fmt.Println("example usage: psql-test-conn -h localhost -P 5432 -u postgres -p 123")
				os.Exit(2)
			}
			res[os.Args[i]] = os.Args[i+1]
		case "-P":
			if i+1 >= len(os.Args) {
				fmt.Println("Miss param -P")
				fmt.Println("example usage: psql-test-conn -h localhost -P 5432 -u postgres -p 123")
				os.Exit(2)
			}
			res[os.Args[i]] = os.Args[i+1]
		case "-u":
			if i+1 >= len(os.Args) {
				fmt.Println("Miss param -u")
				fmt.Println("example usage: psql-test-conn -h localhost -P 5432 -u postgres -p 123")
				os.Exit(2)
			}
			res[os.Args[i]] = os.Args[i+1]
		default:
			fmt.Println("UnKnow param", os.Args[i])
			fmt.Println("example usage: psql-test-conn -h localhost -P 5432 -u postgres -p 123")
		}
	}
	resStr := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable TimeZone=Asia/Shanghai", res["-h"], res["-u"], res["-p"], res["-P"])
	return resStr
}
