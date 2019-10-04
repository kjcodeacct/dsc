CREATE TABLE dsc_changelog
(
	id SERIAL PRIMARY KEY,
	sha256 TEXT NOT NULL,
	epoch INTEGER NOT NULL,
	author_name TEXT NOT NULL,
	author_email TEXT NOT NULL
);

CREATE TABLE dsc_change_files
(
	id SERIAL PRIMARY KEY,
	changelog_id INTEGER NOT NULL REFERENCES dsc_changelog(id) ON DELETE CASCADE,
	file TEXT NOT NULL
);

CREATE TABLE dsc_config
(
	id SERIAL PRIMARY KEY,
	type TEXT NOT NULL,
	alias TEXT NOT NULL,
	dump_exec TEXT
);