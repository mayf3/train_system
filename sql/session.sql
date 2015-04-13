
DROP TABLE IF EXISTS session;

/**
 * session table
 */
CREATE TABLE IF NOT EXISTS session (
    session_key    char(64) NOT NULL,
    session_data   blob,
    session_expiry int(11)  unsigned  NOT NULL,
    PRIMARY KEY (session_key)
);
