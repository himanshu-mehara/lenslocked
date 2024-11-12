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
	// name := "new"
	// email := "newm@meh.io"

	// // query := fmt.Sprintf(`
	// // 	insert into users(name,email)
	// // 	values ('%s','%s');
	// // `,name,email)
	// // fmt.Printf("executing: %s\n",query)
	// // _,err = db.Exec(query)

	// row := db.QueryRow(`
	// insert into users(name,email)
	// values($1,$2)returning id;
	// `,name,email)
	// row.Err() 
	// var id int 
	// err  = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("user record inserted id ",id)

	// var name string 
	// var email string 
	// id := 1
	// row := db.QueryRow(`
	// select name,email
	// from users
	// where id=$1;`,id)
	// err = row.Scan(&name,&email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("error ,no rows!")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("user info: name=%s,email=%s\n",name,email)

	userID := 1 
	for i := 1; i <= 5; i++ {
		amount := i * 100 
		desc := fmt.Sprintf("fake order #%d",i)
		_,err := db.Exec(`
		insert into orders(user_id,amount,description)
		values($1,$2,$3)`,userID,amount,desc)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("created fale orders. ")
}
