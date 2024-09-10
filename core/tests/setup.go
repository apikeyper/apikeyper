package tests

import "apikeyper/internal/database"

// CleanupDb cleans up the database by dropping all tables
func CleanupDb() {
	dbConn := database.GetGormDb()
	dbConn.Exec(`DROP TABLE IF EXISTS "api_keys";
DROP TABLE IF EXISTS "apis";
DROP TABLE IF EXISTS "root_keys";
DROP TABLE IF EXISTS "workspaces";
DROP TABLE IF EXISTS "api_key_usages";
`)
	dbConn.Commit()
}
