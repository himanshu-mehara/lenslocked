package main

import (
	"database/sql"
	_"github.com/jackc/pgx/v4/stdlib"
	"fmt"
)

// type User struct {
// 	Name string
// 	Bio  string
// }

type PostgresConfig struct {
	Host string
	Port string 
	User string
	Password string
	Database string
	SSLMode string
}

func (cfg PostgresConfig) String() string{
	return fmt.Sprintf( "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",cfg.Host,cfg.Port,cfg.User,cfg.Password,cfg.Database,cfg.SSLMode)

}
func main() {
	cfg := PostgresConfig{
		Host: "localhost",
		Port: "5433",
		User: "baloo",
		Password: "junglebook",
		Database:"lenslocked",
		SSLMode: "disable",
	}
	
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	//create a table
	_,err = db.Exec(`
	create table if not exists users(
	id serial primary key,
	name text , 
	email text unique not null
	);

	create table if not exists orders(
	id serial primary key,
	user_id int not null,
	amount int ,
	description text
	);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("tables created")

	// insert some data 
	name := "himanshumehara"
	email := "hm@meh.io"
	_,err = db.Exec(`
	insert into users(name,email)
	values($1,$2);
	`,name,email)
	if err != nil {
		panic(err)
	}
	fmt.Println("user record inserted")
}
