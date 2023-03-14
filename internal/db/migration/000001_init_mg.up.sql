CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          full_name VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL,
                          year INT NOT NULL,
                          department VARCHAR(255) NOT NULL,
                          created_at TIMESTAMP DEFAULT NOW(),
                          updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX students_id_index ON students (id);
CREATE INDEX students_email_index ON students (email);

CREATE TABLE classes (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         professor VARCHAR(255) NOT NULL,
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX classes_id_index ON classes (id);
CREATE INDEX classes_name_index ON classes (name);

CREATE TABLE student_classes (
                                 id SERIAL PRIMARY KEY,
                                 student_id INT NOT NULL REFERENCES students (id) ON DELETE CASCADE,
                                 class_id INT NOT NULL REFERENCES classes (id) ON DELETE CASCADE,
                                 created_at TIMESTAMP DEFAULT NOW(),
                                 updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX student_classes_student_id_index ON student_classes (student_id);
CREATE INDEX student_classes_class_id_index ON student_classes (class_id);
