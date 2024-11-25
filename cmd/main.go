// package main

// import (
// 	"fmt"
// 	"webdev/models"

// 	_ "github.com/jackc/pgx/v4/stdlib"
// )

// // type User struct {
// // 	Name string
// // 	Bio  string
// // }

// // type PostgresConfig struct {
// // 	Host string
// // 	Port string
// // 	User string
// // 	Password string
// // 	Database string
// // 	SSLMode string
// // }

// // func (cfg PostgresConfig) String() string{
// // 	return fmt.Sprintf( "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",cfg.Host,cfg.Port,cfg.User,cfg.Password,cfg.Database,cfg.SSLMode)

// // }
// func main() {
// 	cfg := models.DefaultPostgresConfig()
// 	db, err := models.Open(cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("connected")

// 	us := models.UserService{
// 		DB: db,
// 	}
// 	user, err := us.Create("bob4@bob.com", "bob123")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(user)
// }

//create a table
// _,err = db.Exec(`
// create table if not exists users(
// id serial primary key,
// name text ,
// email text unique not null
// );

// create table if not exists orders(
// id serial primary key,
// user_id int not null,
// amount int ,
// description text
// );
// `)
// if err != nil {
// 	panic(err)
// }
// fmt.Println("tables created")

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

// userID := 1
// for i := 1; i <= 5; i++ {
// 	amount := i * 100
// 	desc := fmt.Sprintf("fake order #%d",i)
// 	_,err := db.Exec(`
// 	insert into orders(user_id,amount,description)
// 	values($1,$2,$3)`,userID,amount,desc)
// 	if err != nil {
// 		panic(err)
// 	}
// }
// fmt.Println("created fale orders. ")

// 	type Order struct {
// 		ID int
// 		UserID  int
// 		Amount int
// 		Description string
// 	}

// 	var orders []Order
// 	userID := 1
// 	rows,err := db.Query(`
// 	select id ,amount , description
// 	from orders
// 	where user_id=$1`,userID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var order Order
// 		order.UserID = userID
// 		err := rows.Scan(&order.ID,&order.Amount,&order.Description)
// 		if err != nil {
// 			panic(err)
// 		}
// 		orders = append(orders, order)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("orders:",orders)
// 	// check for an error
// }

package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	Host     = "sandbox.smtp.mailtrap.io"
	Port     = 2525
	Username = "0a25f7a7d9236c"
	Password = "09efec939b67ef"
)

func main() {
	from := "test@lenslocked.com"
	to := "mehrahimanshu1708@gmail.com"
	subject := "this is a test email "
	plaintext := "this is the body of the email"
	html := `<h1> hello there buddy </h1> <p> this is the email </p> 
	<p> hope you enjoy it </p>`
	msg := mail.NewMessage()

	msg.SetHeader("to", to)
	msg.SetHeader("from", from)
	msg.SetHeader("subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)

	msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(Host, Port, Username, Password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("message sent ")

}
