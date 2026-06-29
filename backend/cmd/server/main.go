package main
import (
	"log"
	"github.com/harshh-2/slink/internal/config"
	"github.com/harshh-2/slink/internal/db"
)
func main(){
	cfg,err := config.Load()
	if err!=nil{
		log.Fatal(err)
	}
	pool,err := db.NewPostgresPool(cfg)
	if err!=nil{
		log.Fatal(err)
	}
	defer pool.Close()
	log.Println("Connection to DB Succesfull")
}