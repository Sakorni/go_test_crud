CREATE TABLE IF NOT EXISTS users
    (
        id serial PRIMARY KEY,
        firstname VARCHAR(50),
        lastname VARCHAR(50),
        email VARCHAR(100),
        age SMALLINT CHECK (age > 0),
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP

    );