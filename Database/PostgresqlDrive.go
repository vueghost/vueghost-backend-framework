package Database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

//Postgresql DB drive constant.
const (
	dbHost     string = "localhost"
	dbUsername string = "postgres"
	dbPassword string = "root"
	dbPORT     int    = 5433
	dbName     string = "vueghost_local"
	dbSchema   string = "public"

	rdsARN                  string = "arn:aws:rds:us-east-2:743010515815:cluster:vueghost-database"
	rdsEndPoint             string = "vueghost-database.cluster-cqviybilwaj4.us-east-2.rds.amazonaws.com"
	rdsPort                 int    = 5432
	rdsMasterUsername       string = "vueghost__db"
	rdsPassword             string = "HWry6fgvZXAu52zvAJwQ"
	rdsDB                   string = "postgres"
	rdsSchema               string = "public"
	ErrorCodeDuplication    int    = 502
	ErrorCodeExecution      int    = 500
	ErrorCodeRecordNotExist int    = 503
)

//PostgresqlDrive
type PostgresqlDrive struct {
	db               *sqlx.DB
	Jsonb            Jsonb
	isOpenConnection bool
}

//NewPostgresqlDrive Create new
func NewPostgresqlDrive() *PostgresqlDrive {
	return &PostgresqlDrive{}
}

//dbDevelopment database development.
func (d *PostgresqlDrive) dbDevelopment() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=disable search_path=%s",
		dbHost,
		dbPORT,
		dbName,
		dbUsername,
		dbPassword,
		dbSchema)
}

//dbProduction database production.
func (d *PostgresqlDrive) dbProduction() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=disable search_path=%s",
		rdsEndPoint,
		rdsPort,
		rdsDB,
		rdsMasterUsername,
		rdsPassword,
		rdsSchema)
}

//Connect open a connection to database.
func (d *PostgresqlDrive) Connect() {
	dbConnectionSource := d.dbDevelopment()
	db, err := sqlx.Open("postgres", dbConnectionSource)
	defer d.throwError(err)
	if err != nil {
		d.isOpenConnection = false
		return
	}

	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(3 * time.Second)
	d.db = db
	d.isOpenConnection = true
	fmt.Println("@Database 'PostgresqlDrive' 🔌  connected successfully ✅")
}

//Reconnect reconnect to database connection.
func (d *PostgresqlDrive) Reconnect() {
	if d.isOpenConnection {
		d.Close()
		d.Connect()
	} else {
		d.Connect()
	}
}

//Close closing the database connection.
func (d *PostgresqlDrive) Close() (success bool) {
	err := d.db.Close()
	if err != nil {
		d.throwError(err)
		return false
	}

	fmt.Println("@Database:", d.db.Stats().OpenConnections)
	return true
}

//Version current database drive version.
func (d *PostgresqlDrive) Version() (version interface{}, error error) {
	var result interface{}
	error = d.db.QueryRow("select version()").Scan(&version)
	return result, error
}

//GetDbInstance
func (d *PostgresqlDrive) GetDbInstance() *sqlx.DB {
	return d.db
}

//TransactionBegin
func (d *PostgresqlDrive) TransactionBegin() *sqlx.Tx {
	return d.db.MustBegin()
}

//Get
func (d *PostgresqlDrive) Get(dest interface{}, query string, args ...interface{}) error {
	return d.db.Get(dest, query, args...)
}

//Execute
func (d *PostgresqlDrive) Execute(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

//MapExecute
func (d *PostgresqlDrive) MapExecute(query string, arg interface{}) (sql.Result, error) {
	return d.db.NamedExec(query, arg)
}

//throwError database error exception error.
func (d *PostgresqlDrive) throwError(err error) {
	if err != nil {
		if d.isOpenConnection == true {
			d.Close()

		}
		panic("⭕ @Database 'PostgresqlDrive' Error:  " + err.Error())
	}
}
