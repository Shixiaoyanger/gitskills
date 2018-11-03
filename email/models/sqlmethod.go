package models

import ( //"github.com/gin-gonic/gin"
	//"fmt"
	//	"log"
	//"net/http"
	//"email/routers"

	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", "root:123456@/email")

func Getinfo(username string) (string, string, bool) {
	var usrname, passwd string
	err := db.QueryRow("select username, password from userinfo where username = ? ", username).Scan(&usrname, &passwd)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("There is no row")
		} else {
			log.Fatal(err)
		}

		return "", "", false
	}
	return usrname, passwd, true

}
func Insert(username, password string) {
	_, err := db.Exec("insert into userinfo(username,password) values(?,?)", username, password)
	if err != nil {
		log.Fatal(err)
	}
	return
}

/*
var db = &sql.DB{}
func init(){
	//user@unix(/path/to/socket)/dbname?charset=utf8
	//user:password@tcp(localhost:5555)/dbname?charset=utf8
	//user:password@/dbname
	//user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	//db,err := sql.Open("mysql","root:123456@tcp(120.0.0.1:3306)/?charset=utf8")
	db,_ := sql.Open("mysql","root:123456@/test")

}
*/

/*

func Insert(){
	//checkErr(err)
	name := "4"
	rows,err := db.Query("select id,code from userinfo where id = ?",name)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next(){
		var id int
		var code string
		if err := rows.Scan(&id,&code);err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s is %d\n",code,id)
	}
}
func Update(){

}
func Delete(){

}
*/
