CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          email VARCHAR NOT NULL,
                          teacher_id INT REFERENCES teachers(id) ON DELETE SET NULL
);
