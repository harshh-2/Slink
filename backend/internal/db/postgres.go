package db
import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/harshh-2/slink/internal/config"
)

func NewPostgresPool(cfg *config.Config) (*pgxpool.Pool,error){
	ctx := context.Background()
	pool,err := pgxpool.New(ctx,cfg.DatabaseURL)
	if err!=nil{
		return nil,err
	}
	if err := pool.Ping(ctx); err != nil {
    pool.Close()
    return nil, err
	}
	return pool,nil
}