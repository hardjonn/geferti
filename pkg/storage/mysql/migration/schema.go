package migration

var schemaMigrations = `
	CREATE TABLE IF NOT EXISTS migrations (
		name VARCHAR(255) NOT NULL,
		migrated_at DATETIME NULL,
		CONSTRAINT migrations_PK PRIMARY KEY (name)
	)
	ENGINE=InnoDB
	DEFAULT CHARSET=utf8mb4
	COLLATE=utf8mb4_general_ci;
`

var schemaNodes = `
	CREATE TABLE nodes (
		id binary(16) NOT NULL,
		uuid CHAR(36) COLLATE utf8mb4_general_ci GENERATED ALWAYS AS (BIN_TO_UUID(id)) VIRTUAL NOT NULL,
		host_key CHAR(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
		deleted_at DATETIME NULL DEFAULT NULL,
		created_at DATETIME NULL DEFAULT NULL,
		updated_at DATETIME NULL DEFAULT NULL,
		CONSTRAINT nodes_PK PRIMARY KEY (id),
		CONSTRAINT nodes_UN UNIQUE KEY (host_key)
	)
	ENGINE=InnoDB
	DEFAULT CHARSET=utf8mb4
	COLLATE=utf8mb4_general_ci;
`
