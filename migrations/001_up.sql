CREATE TABLE Schools (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Address VARCHAR(200) NOT NULL,
	Created_at TIMESTAMP NOT NULL
);

CREATE TABLE Classes (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Grade INT CHECK (Grade > 0 AND Grade < 5) NOT NULL,
	School_Id INT NOT NULL,
	Created_at TIMESTAMP NOT NULL,
	CONSTRAINT FK_Classes_School_id  FOREIGN KEY(School_Id) REFERENCES Schools(Id) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE Roles (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(100) UNIQUE
);

CREATE TABLE Users (
	Id SERIAL PRIMARY KEY,
	Email VARCHAR(100) NULL UNIQUE,
	Login VARCHAR(100) NOT NULL UNIQUE,
	Password_hash VARCHAR(100) NOT NULL, 
	Role_Id INT NULL,
	Blocked BOOLEAN NOT NULL,
	FullName VARCHAR(200) NOT NULL,
	Class_Id INT NULL,
	Created_at TIMESTAMP NOT NULL,
	Last_login TIMESTAMP NULL,
	CONSTRAINT FK_Users_Class_Id FOREIGN KEY(Class_Id) REFERENCES Classes(Id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT FK_Users_Role_Id FOREIGN KEY(Role_Id) REFERENCES Roles(Id) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE Achievements (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Description TEXT NULL
);

CREATE TABLE Achievements_of_students (
	Student_Id INT NOT NULL,
	Achievement_Id INT NOT NULL,
	Got_at TIMESTAMP NOT NULL,
	CONSTRAINT PK_Achievements_of_students PRIMARY KEY(Student_Id, Achievement_Id),
	CONSTRAINT FK_Achievements_of_students_Student_Id  FOREIGN KEY(Student_Id) REFERENCES Users(Id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT FK_Achievements_of_students_Achievement_Id  FOREIGN KEY(Achievement_Id) REFERENCES Achievements(Id) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE Equation_types (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Description TEXT NULL
);

CREATE TABLE Levels (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Test_level BOOLEAN NOT NULL,
	Difficulty INTEGER NOT NULL
);

CREATE TABLE Equations (
	Id SERIAL PRIMARY KEY,
	Expression VARCHAR(100) NOT NULL,
	Correct_answer INT NOT NULL,
	Equaition_type_Id INT NOT NULL,
	Difficulty INTEGER NOT NULL,
	CONSTRAINT FK_Equations_Equaition_type_Id  FOREIGN KEY(Equaition_type_Id) REFERENCES Equation_types(Id) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE Student_progress (
	Id SERIAL PRIMARY KEY,
	Student_Id INT NOT NULL,
	Level_id INT NOT NULL,
	Count_stars INT,
	Finished_at TIMESTAMP NULL,
	CONSTRAINT FK_Student_progress_Student_Id  FOREIGN KEY(Student_Id) REFERENCES Users(Id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT FK_Student_progress_Level_Id  FOREIGN KEY(Level_Id) REFERENCES Levels(Id) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE Level_equation_type (
	Level_Id INT NOT NULL,
	Equation_type_Id INT NOT NULL,
	CONSTRAINT PK_Level_equation_type PRIMARY KEY(Level_Id, Equation_type_Id),
	CONSTRAINT FK_Level_equation_type_Equation_type_Id  FOREIGN KEY(Equation_type_Id) REFERENCES Equation_types(Id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT FK_Level_equation_type_Level_Id  FOREIGN KEY(Level_Id) REFERENCES Levels(Id) ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE Theory (
	Id SERIAL PRIMARY KEY,
	Equation_type_Id INT NOT NULL,
	Name VARCHAR(100),
	Content TEXT,
	CONSTRAINT FK_Theory_Equation_type_Id  FOREIGN KEY(Equation_type_Id) REFERENCES Equation_types(Id) ON UPDATE CASCADE ON DELETE NO ACTION
);

INSERT INTO Schools(Name, Address, Created_at)
VALUES ('Общеобразовательная школа «Средняя общеобразовательная школа № 322»', 'Санкт-Петербург, улица Олеко Дундича, 38, корп. 3, метро «Дунайская»', CURRENT_TIMESTAMP),
('ГБОУ СОШ № 555 с углубленным изучением английского языка «Белогорье»','Санкт-Петербург, Комендантский проспект, 17, корп. 3, метро «Комендантский проспект»', CURRENT_TIMESTAMP),
('ГБОУ Средняя общеобразовательная школа № 111 с углубленным изучением немецкого языка', 'Санкт-Петербург, улица Фаворского, дом 16, литер А', CURRENT_TIMESTAMP);


INSERT INTO Classes	(Name, Grade, School_Id, Created_at)
VALUES ('3А', 3, 1, CURRENT_TIMESTAMP),
('2А', 2, 1, CURRENT_TIMESTAMP),
('2Б', 2, 1, CURRENT_TIMESTAMP);

INSERT INTO Roles(Name)
VALUES ('Student'),
('Administrator'),
('Teacher'),
('Head');

INSERT INTO Users (Email, Login, Password_hash, Role_Id, Blocked, FullName, Class_Id, Created_at)
VALUES ('student1@gmail.com', 'student1', '$2a$10$Qtc61.Gn4rXMbajz/iBYV.kV4Tv4jQIB50sBcfWX6s6CtHbXsuGS2', 1, FALSE, 'Семенов Семен Семенович', 3, CURRENT_TIMESTAMP),
('student2@gmail.com', 'student2', '$2a$10$hTAYihicrLbDYqqPoIXFBus5EqXbE/mwm5tXQlljp3ECREHVAeECG', 1, FALSE, 'Николаев Николай Николаевич', 3, CURRENT_TIMESTAMP),
('admin@gmail.com', 'admin', '$2a$10$ToSI7ka8vIXEMwPhobZZkuAqu3NTlYFQj3.eE7oXnqxZtl7twak4C', 2, FALSE, 'Дарьина Дарья Дарьевна', NULL, CURRENT_TIMESTAMP),
('teacher1@gmail.com', 'teacher1', '$2a$10$alzWzsvxHWvdsihcWfu2rezwplu/hKrVGLKhyVhiKjaHA4zGs44OO', 3, FALSE, 'Марьина Марья Маровна', 3, CURRENT_TIMESTAMP);

INSERT INTO Achievements (Name, Description)
VALUES ('Бронзовый кубок', 'Мог бы и лучше сделать)'),
('Серебряный кубок', 'А почему не первое место?'),
('Золотой кубок', 'Ты чемпион, выше всех, знай это!');

INSERT INTO Achievements_of_students (Student_Id, Achievement_Id, Got_at)
VALUES (1, 1, CURRENT_TIMESTAMP),
(2, 1, CURRENT_TIMESTAMP),
(1, 2, CURRENT_TIMESTAMP);

INSERT INTO Equation_types(Name, Description)
VALUES ('Простейшие уравнения на сложение', 'Нахождение неизвестного слагаемого'),
('Простейшие уравнения на вычитание', 'Нахождение вычитаемого и уменьшаемого'),
('Простейшие уравнения на умножение', 'Нахождение множителя');
INSERT INTO Equation_types(Name, Description)
VALUES 
('Простейшие уравнения на деление', 'Нахождение делителя'),
('Составные уравнения', 'Уравнения в несколько действий');
INSERT INTO Levels (Name, Test_level, Difficulty)
VALUES
('Уровень 1', FALSE, 1),
('Уровень 2', FALSE, 2),
('Итоговый тест', TRUE, 3);
INSERT INTO Equations (Expression, Correct_answer, Equaition_type_Id, Difficulty)
VALUES
('x + 3 = 7', 4, 1, 1),
('x - 5 = 2', 7, 2, 1),
('3 * x = 12', 4, 3, 2),
('x / 2 = 5', 10, 4, 2),
('2 * x + 3 = 11', 4, 5, 3);
INSERT INTO Student_progress (Student_Id, Level_Id, Count_stars, Finished_at)
VALUES
(1, 1, 3, CURRENT_TIMESTAMP),
(1, 2, 2, CURRENT_TIMESTAMP),
(2, 1, 1, CURRENT_TIMESTAMP);
INSERT INTO Level_equation_type (Level_Id, Equation_type_Id)
VALUES
(1, 1),
(1, 2),
(2, 3),
(3, 5);
INSERT INTO Theory (Equation_type_Id, Name, Content)
VALUES
(1, 'Как находить неизвестное слагаемое', 
 'Чтобы найти неизвестное слагаемое, нужно из суммы вычесть известное слагаемое.'),

(2, 'Как находить уменьшаемое и вычитаемое', 
 'Чтобы найти уменьшаемое, нужно к разности прибавить вычитаемое.'),

(3, 'Как находить неизвестный множитель', 
 'Чтобы найти множитель, нужно произведение разделить на известный множитель.');
