package System

import "VGBackendFramework/Database"

//Model
type Model struct {
}

func (m Model) DB() Database.DB {
	postgresql := Database.NewPostgresqlDrive()
	return Database.NewDB(postgresql)
}
