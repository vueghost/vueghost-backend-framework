package Database

import (
	System "Framework"
	"testing"
)

var tUnit = System.TestUnit{}

func TestConstant(t *testing.T) {
	if tUnit.IsEmpty(dbHost) ||
		tUnit.IsEmpty(dbUsername) ||
		tUnit.IsEmpty(dbPassword) ||
		tUnit.IsEmpty(dbPORT) ||
		tUnit.IsEmpty(dbName) ||
		tUnit.IsEmpty(dbSchema) {
		t.Fatal()
	}
}

func TestNewPostgresqlDrive(t *testing.T) {
	_ = NewPostgresqlDrive()
}

func TestPostgresqlDrive_dbDevelopment(t *testing.T) {
	db := PostgresqlDrive{}
	dbDev := db.dbDevelopment()
	if tUnit.IsEmpty(dbDev) {
		t.Fatal()
	}
	if !tUnit.Contains(dbDev, "host=") {
		t.Fatal()
	}
}

func TestPostgresqlDrive_dbProduction(t *testing.T) {
	db := PostgresqlDrive{}
	dbPro := db.dbProduction()
	if tUnit.IsEmpty(dbPro) {
		t.Fatal()
	}
	if !tUnit.Contains(dbPro, "host=") {
		t.Fatal()
	}
}

func TestPostgresqlDrive_Connect(t *testing.T) {
	db := PostgresqlDrive{}
	db.Connect()
}

func TestPostgresqlDrive_Close(t *testing.T) {
	t.Run("When there is a open connection", func(t *testing.T) {
		db := PostgresqlDrive{}
		db.Connect()
		t.Log(db.Close())
	})
}
func TestPostgresqlDrive_Reconnect(t *testing.T) {
	db := NewPostgresqlDrive()
	t.Run("When not connected", func(t *testing.T) {
		db.Reconnect()
	})

	t.Run("when connected", func(t *testing.T) {
		db.Connect()
		db.Reconnect()
	})
}

func TestPostgresqlDrive_Version(t *testing.T) {
	db := NewPostgresqlDrive()
	db.Connect()
	t.Log(db.Version())
}
