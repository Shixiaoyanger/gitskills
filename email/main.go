package main

import ( //"github.com/gin-gonic/gin"
	//"log"
	//"net/http"
	//"email/models" //"database/sql"
	"email/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	routers.Router()
}
