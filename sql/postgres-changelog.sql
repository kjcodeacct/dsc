CREATE TABLE dsc_changelog
(
	id SERIAL PRIMARY KEY,
	sha256 TEXT NOT NULL,
	-- created at = time commited
	created_at INTEGER NOT NULL,
	-- applied at = time sql is ran on remote 
	applied_at INTEGER,
	-- completed at = time sql is completed and transaction applied
	completed_at INTEGER,
	author_name TEXT NOT NULL,
	author_email TEXT NOT NULL
);

CREATE TABLE dsc_change_file
(
	id SERIAL PRIMARY KEY,
	file TEXT NOT NULL,
	created_at INTEGER NOT NULL
);

CREATE UNIQUE INDEX ON dsc_change_files(file);

CREATE TABLE dsc_reference_log
(
	id SERIAL PRIMARY KEY,
	dsc_change_file_id INTEGER NOT NULL REFERENCES dsc_change_file(id),
	updated_at INTEGER NOT NULL,
	change_type TEXT NOT NULL REFERENCES dsc_change_types(type),
	author_name TEXT NOT NULL,
	author_email TEXT NOT NULL
);

CREATE TABLE dsc_change_types
(
	id SERIAL PRIMARY KEY,
	type TEXT NOT NULL,
);

CREATE UNIQUE INDEX ON dsc_change_types(type);

CREATE TABLE dsc_config
(
	id SERIAL PRIMARY KEY,
	type TEXT NOT NULL,
	alias TEXT NOT NULL,
	dump_exec TEXT
);