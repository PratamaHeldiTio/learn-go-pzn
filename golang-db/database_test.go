package golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	res, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES($1, $2);", "tama", "Tama")

	if err != nil {
		panic(err)
	}

	a, err := res.RowsAffected()
	fmt.Println(a)
}

func TestSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "SELECT * FROM customer;"
	rows, err := db.QueryContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name, email string
		var balance sql.NullInt32 // ketika data dari db bisa null
		var rating float32
		var birdDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birdDate, &married)

		if err != nil {
			panic(err)
		}

		fmt.Println(id, name, email, balance, rating, createdAt, birdDate, married)
	}
	defer rows.Close()
}

func TestSQLInject(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin' ; --" //-- komtar pad postgressxX
	pasword := "ASA"

	sqlQuery := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + pasword + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("login sukses")
	} else {
		fmt.Println("gagal")
	}
	defer rows.Close()
}

func TestQueryParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin" //-- komtar pad postgressxX
	pasword := "admin"

	query := "SELECT username FROM users WHERE username = $1 AND password = $2 LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, pasword)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("login sukses")
	} else {
		fmt.Println("gagal")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	query := `INSERT INTO comments("name", "comment") VALUES($1, $2) RETURNING id`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}

	var id int32
	if err := stmt.QueryRow("heldi", "hahaha").Scan(&id); err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func TestPrepare(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(name, comment) VALUES($1, $2) RETURNING id"

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		comment := "lorem" + strconv.Itoa(i)
		name := "heldi" + strconv.Itoa(i)

		var id int32
		if err := stmt.QueryRow(name, comment).Scan(&id); err != nil {
			panic(err)
		}

		fmt.Println(id)

	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	kondisi := false
	for i := 11; i <= 20; i++ {
		comment := "lorem" + strconv.Itoa(i)
		name := "heldi" + strconv.Itoa(i)

		query := "SELECT id from comments where id = $1"

		res, err := tx.ExecContext(ctx, query, i)
		if err != nil {
			panic(err)
		}

		afect, _ := res.RowsAffected()

		if afect > 0 {
			fmt.Println("sebelum rollbsck")
			kondisi = true
			return
		}

		query = "INSERT INTO comments(id,name, comment) VALUES($1, $2, $3) RETURNING id"
		var id int32
		_, err = tx.ExecContext(ctx, query, i, name, comment)
		if err != nil {
			panic(err)
		}

		fmt.Println(id)
	}

	fmt.Println("sebelum komit")
	if kondisi == false {
		err = tx.Commit()
		if err != nil {
			panic(err)
		}
	} else {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
	}
}
