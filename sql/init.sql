DROP TABLE IF EXISTS session;
DROP TABLE IF EXISTS user;

/**
 * session table
 */
CREATE TABLE IF NOT EXISTS session (
	session_key    char(64) NOT NULL,
	session_data   blob,
	session_expiry INT(11)  unsigned  NOT NULL,

	PRIMARY KEY (session_key)
);

CREATE TABLE IF NOT EXISTS user(
	id 			INT 			NOT NULL AUTO_INCREMENT,
	username 	VARCHAR(32) 	NOT NULL,
	password 	VARCHAR(128) 	NOT NULL,
	email 		VARCHAR(32) 	NOT NULL,
    create_at  	DATE         	NOT NULL,
    is_active  	BOOLEAN      	DEFAULT FALSE,

	PRIMARY KEY (id)
);
