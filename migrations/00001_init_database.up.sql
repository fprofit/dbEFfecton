CREATE TABLE departments (
    departmentid serial PRIMARY KEY,
    departmentname varchar(100) UNIQUE
);

CREATE TABLE employees (
    employeeid serial PRIMARY KEY,
    firstname varchar(50),
    lastname varchar(50),
    departmentid int REFERENCES departments(departmentid)
);

INSERT INTO departments (departmentname) VALUES
    ('Отдел продаж'),
    ('Отдел маркетинга'),
    ('Отдел разработки'),
    ('Отдел качества'),
    ('Отдел поддержки'),
    ('Отдел финансов'),
    ('Отдел управления'),
    ('Отдел ресурсов');

INSERT INTO employees (firstname, lastname, departmentid) VALUES
    ('Иван', 'Иванов', 1),
    ('Петр', 'Петров', 1),
    ('Анна', 'Сидорова', 2),
    ('Мария', 'Иванова', 3),
    ('Алексей', 'Петров', 3),
    ('Елена', 'Смирнова', 4),
    ('Сергей', 'Козлов', 5),
    ('Ольга', 'Васильева', 6),
    ('Дмитрий', 'Попов', 7),
    ('Наталья', 'Соколова', 8),
    ('Андрей', 'Морозов', 1),
    ('Евгений', 'Николаев', 2),
    ('Татьяна', 'Захарова', 3),
    ('Александра', 'Павлова', 4),
    ('Михаил', 'Белов', 5);

