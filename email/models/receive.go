package models

import ( //"github.com/gin-gonic/gin"
	"email/utils/go-pop3" //"log"
	"encoding/base64"     //"net/http"
	//"email/models" //"database/sql"
	//"email/routers"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

import "golang.org/x/text/encoding/simplifiedchinese"

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func main() {

	address := "pop.163.com:110"
	user := "13156211236@163.com"
	pass := "shi13156211236"
	if err := pop3.ReceiveMail(address, user, pass,
		func(number int, uid, data string, err error) (bool, error) {
			log.Printf("%d, %s\n", number, uid)

			r := strings.NewReader(data)
			m, err := mail.ReadMessage(r)
			if err != nil {
				log.Fatal(err)
			}

			header := m.Header
			fmt.Println("Date:", header.Get("Date"))
			//fmt.Println("From:", header.Get("From"))
			//fmt.Println("To:", ConvertByte2String([]byte(header.Get("To")), GB18030))
			//fmt.Println("Subject:", header.Get("Subject"))

			body, err := ioutil.ReadAll(m.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("body%s", body)

			decodeBytes, err := base64.StdEncoding.DecodeString("")
			if err != nil {
				log.Fatalln(err)
			}
			print(string(decodeBytes))
			//ls := ConvertByte2String(body, GB18030)
			//clearprint(ls)

			// implement your own logic here

			return false, nil
		}); err != nil {
		log.Fatalf("%v\n", err)
	}

	//routers.Router()
}

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
