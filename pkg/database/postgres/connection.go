package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"golang-utility/pkg/errors"
	"golang-utility/pkg/errors/common"
	"golang-utility/pkg/logging"
	"os"
	"time"

	"github.com/jackc/pgconn"
	pgxPool "github.com/jackc/pgx/v4/pgxpool"
)

var logger = logging.GetLogger("PostgresClient")

type PostgresConfiguration struct {
	UserName     string
	Password     string
	DatabaseName string
	InstanceName string
	Host         string
	Port         string
}

func InitConnectionPool(config PostgresConfiguration) *pgxPool.Pool {
	localEnvVar := os.Getenv("LOCAL")

	logger.Info("env var LOCAL: " + localEnvVar)

	if localEnvVar == "true" {
		logger.Info("Initializing TCP Connection pool (IP address)")
		connPool := initTCPConnectionPool(config)
		configureConnectionPool(connPool)
		return connPool
	} else {
		logger.Info("Initializing Socket Connection pool")
		connPool := initSocketConnectionPool(config)
		configureConnectionPool(connPool)
		return connPool
	}
}

func initSocketConnectionPool(config PostgresConfiguration) *pgxPool.Pool {
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", config.UserName, config.Password, config.DatabaseName, socketDir, config.InstanceName)

	dbPool, err := pgxPool.Connect(context.Background(), dbURI)
	if err != nil {
		logger.Error("initSocketConnectionPool: unable to connect ", errors.ToError(err))
		panic(&PostgresConnectionError{
			BaseError: errors.NewBaseError("Error initializing Postgres Socket Connection pool",
				errors.ToError(err), true),
		})
	}

	return dbPool
}

func initTCPConnectionPool(config PostgresConfiguration) *pgxPool.Pool {
	localEnvVar := os.Getenv("LOCAL")
	var sslMode string
	if localEnvVar == "true" {
		sslMode = "disable"
	} else {
		sslMode = "enable"
	}

	dbURI := fmt.Sprintf("sslmode=%s host=%s user=%s password=%s port=%s database=%s", sslMode, config.Host, config.UserName, config.Password, config.Port, config.DatabaseName)

	dbPool, err := pgxPool.Connect(context.Background(), dbURI)
	if err != nil {
		panic(&PostgresConnectionError{
			BaseError: errors.NewBaseError(
				"Error initializing Postgres TCP Connection pool",
				errors.ToError(err),
				true),
		})
	}

	configureConnectionPool(dbPool)

	return dbPool
}

func configureConnectionPool(pool *pgxPool.Pool) {
	pool.Config().MaxConns = 10
	pool.Config().MinConns = 5
	pool.Config().MaxConnLifetime = 1800 * time.Second
}

func Exec(ctx context.Context, db DBorTx, query string, args ...any) (pgconn.CommandTag, errors.Error) {
	result, err := db.Exec(ctx, query, args...)

	if err != nil {
		logger.Error("Failed to execute query: ", errors.ToError(err))
		return nil, &QueryExecError{
			BaseError: errors.NewBaseError("Error during db.Exec", errors.ToError(err), true),
		}
	}

	return result, nil
}

func ExecQuery(ctx context.Context, db DBorTx, query string, args ...any) (pgx.Rows, errors.Error) {
	logger.Debug(fmt.Sprintf("query: %s | params: %s \n", query, args))

	rows, err := db.Query(ctx, query, args...)

	if err != nil {
		logger.Error("Failed to execute query: ", errors.ToError(err))
		return nil, &QueryExecError{
			BaseError: errors.NewBaseError("Error during db.Query", errors.ToError(err), true),
		}
	}

	return rows, nil
}

func ExecQueryRow(ctx context.Context, db DBorTx, query string, args ...any) pgx.Row {
	logger.Debug(fmt.Sprintf("query: %s | params: %s \n", query, args))

	row := db.QueryRow(ctx, query, args...)

	return row
}

func BeginTransaction(ctx context.Context, db *pgxPool.Pool) (pgx.Tx, errors.Error) {
	transaction, err := db.Begin(ctx)

	if err != nil {
		logger.Error("error starting transaction", errors.ToError(err))
		return nil, &common.TransactionError{
			BaseError: errors.NewBaseError("error starting transaction", errors.ToError(err), true),
		}
	}

	return transaction, nil
}

func RollbackTransaction(ctx context.Context, transaction pgx.Tx) {
	transaction.Rollback(ctx)
}

func CommitTransaction(ctx context.Context, transaction pgx.Tx) errors.Error {
	err := transaction.Commit(ctx)

	if err != nil {
		logger.Error("error committing transaction", errors.ToError(err))
		return &common.TransactionError{
			BaseError: errors.NewBaseError("error committing transaction", errors.ToError(err), true),
		}
	}

	return nil
}
