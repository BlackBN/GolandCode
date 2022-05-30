package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Payment struct {
	ID     string `json:"id"`
	Serial string `json:"serial"`
}

func main() {
	DB, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置最大闲置连接数
	DB.SetMaxIdleConns(10)

	if err = DB.Ping(); err != nil {
		fmt.Println("open database fail")
		panic(err)
	}
	fmt.Println("connection success")
	rows, err := DB.Query("select * from payment")
	if err != nil {
		panic(err)
	}
	paymentList := make([]*Payment, 0)

	for rows.Next() {
		payment := &Payment{}
		err := rows.Scan(&(payment.ID), &(payment.Serial))
		if err != nil {
			panic(err)
		}
		paymentList = append(paymentList, payment)
	}
	for _, pay := range paymentList {
		fmt.Println(pay.ID + "  " + pay.Serial)
	}

}
