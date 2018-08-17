CREATE TABLE corpus (
    name VARCHAR(127),
    PRIMARY KEY(name)
);

CREATE TABLE token (
	id SERIAL,
	pa_word VARCHAR(127),
	pa_lemma VARCHAR(127),
	pa_tag VARCHAR(127),
	PRIMARY KEY(id)
);

CREATE TABLE struct (
	name VARCHAR(127),
	PRIMARY KEY(name)
);

CREATE TABLE structattr (
	name VARCHAR(127),
	struct_name VARCHAR(127),
	PRIMARY KEY(name, struct_name),
	FOREIGN KEY (struct_name) REFERENCES struct(name)
);

CREATE TABLE sattrvalue (
	id SERIAL,
	value TEXT,
    corpus_name VARCHAR(127),
	structattr_name VARCHAR(127),
	struct_name VARCHAR(127),
	PRIMARY KEY (id),
    FOREIGN KEY (corpus_name) REFERENCES corpus(name),
	FOREIGN KEY (struct_name, structattr_name) REFERENCES structattr(struct_name, name)
);

CREATE TABLE token_sattrvalue (
	token_id INTEGER,
	sattrvalue_id INTEGER,
	FOREIGN KEY (token_id) REFERENCES token(id),
	FOREIGN KEY (sattrvalue_id) REFERENCES sattrvalue(id)
);
