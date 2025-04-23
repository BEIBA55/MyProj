CREATE TABLE teachers (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          email VARCHAR NOT NULL,
                          subject_id INT REFERENCES subjects(id) ON DELETE CASCADE
);
