package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if(err != nil) {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(3 * time.Minute)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancelFunc()

	//fail fast saying
	if err := db.PingContext(ctx); err != nil { //(context) basically sends data here it's trying to connect with the db 
		return nil, fmt.Errorf("ping.Context: %w", err)
	}
	
	return db, nil
}