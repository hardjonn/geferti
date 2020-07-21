CREATE TABLE geferti.nodes (
  id binary(16) NOT NULL,
	uuid CHAR(36) COLLATE utf8mb4_general_ci NOT NULL GENERATED ALWAYS AS (BIN_TO_UUID(id)) VIRTUAL,
	`key` CHAR(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
	CONSTRAINT servers_PK PRIMARY KEY (id),
	CONSTRAINT servers_UN UNIQUE KEY (`key`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
