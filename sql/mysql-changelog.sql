CREATE TABLE schema_changelog
(
	id SERIAL PRIMARY KEY,
	sha256 TEXT NOT NULL,
	file TEXT NOT NULL,
	epoch INTEGER NOT NULL,
	author_name TEXT NOT NULL,
	author_email TEXT NOT NULL
);