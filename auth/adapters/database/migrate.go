package database

// Migrate migrate the models
func Migrate(db *Database) (err error) {
	modelList := []interface{}{}

	for _, model := range modelList {
		err = db.AutoMigrate(model)
		if err != nil {
			return
		}
	}

	return
}
