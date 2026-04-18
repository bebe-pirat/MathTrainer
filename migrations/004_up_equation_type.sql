drop table equation_attempts;
DROP TABLE student_progress;
drop table level_equation_type;
drop table theory;
drop table levels;
drop table equations;

drop table equation_types;

create table equation_types (
	id SERIAL PRIMARY KEY, 
	class INT NOT NULL, 
	name VARCHAR(200) NOT NULL, 
	description TEXT NULL, 
	operations VARCHAR(50), 
	num_operands INT NOT NULL, 
	no_remainder BOOLEAN NULL DEFAULT 'true', 
	max_result INT NOT NULL
);

create table attempts (
	id SERIAL PRIMARY KEY,
	equation_type_id INT NOT NULL, 
	equation_text VARCHAR(200) NOT NULL, 
	correct_answer INT NOT NULL, 
	student_answer INT NOT NULL, 
	answered_at TIMESTAMP,
	FOREIGN KEY (equation_type_id) REFERENCES equation_types ON DELETE CASCADE ON UPDATE CASCADE
); 

create table operands (
	id SERIAL PRIMARY KEY, 
	equation_type_id INT NOT NULL, 
	operand_order INT NOT NULL, 
	min_value INT NOT NULL, 
	max_value INT NOT NULL
);

create table equation_hints (
	id SERIAL PRIMARY KEY, 
	equation_type_id INT NOT NULL, 
	step_order INT NOT NULL, 
	hint_text TEXT NOT NULL, 
	FOREIGN KEY (equation_type_id) REFERENCES equation_types
);	

CREATE TABLE sections (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(100) NOT NULL, 
	class INT NOT NULL, 
	section_order INT NOT NULL
);

CREATE TABLE section_equation_types (
	section_id INT,
	eqaution_type_id INT, 
	PRIMARY KEY (section_id, eqaution_type_id)
);


CREATE TABLE student_progress_level (
	id SERIAL PRIMARY KEY, 
	user_id INT NOT NULL, 
	section_id INT NOT NULL, 
	level_order INT NOT NULL, 
	count_stars INT NOT NULL, 
	finished_at TIMESTAMP NOT NULL, 
	FOREIGN KEY (user_id) REFERENCES users(id), 
	FOREIGN KEY (section_id) REFERENCES sections(id)
);

insert into sections (name, class, section_order)
values ('Сложение', 1, 1), 
('Вычитание', 1, 2), 
('Сложение', 2, 1), 
('Вычитание', 2, 2), 
('Учножение', 2, 3), 
('Деление', 2, 4), 
('Сложение', 3, 1), 
('Вычитание', 3, 2), 
('Учножение', 3, 3), 
('Деление', 3, 4), 
('Сложение', 4, 1), 
('Вычитание', 4, 2), 
('Учножение', 4, 3), 
('Деление', 4, 4);


insert into equation_types (class, name, description, operations, num_operands, no_remainder, max_result)
values (1, 'Сложение чисел до 20', 'Сложение двух чисел в пределах 20', '+', 2, true, 20),
(1, 'Вычитание чисел до 20', 'Вычитание чисел до 20', '-', 2, true, 20),
(2, 'Сложение чисел до 100', 'Сложение двух чисел в пределах 100', '+', 2, true, 100),
(2, 'Вычитание чисел до 100', 'Вычитание чисел до 100', '-', 2, true, 100),
(2, 'Табличное умножение', 'Табличное умножение в пределах 100', '+', 2, true, 100),
(2, 'Смешанные выражение в пределах 100 (сложение и вычитание)', 'Смешанные выражение в пределах 100 (сложение и вычитание)', '+-', 3, true, 1000),
(2, 'Смешанные выражение в пределах 100 (сложение, умножение и вычитание)', 'Смешанные выражение в пределах 100 (сложение, умножение и вычитание)', '+-*', 4, true, 1000),
(3, 'Сложение чисел до 1000', 'Сложение двух чисел в пределах 1000', '+', 2, true, 1000),
(3, 'Вычитание чисел до 1000', 'Вычитание чисел в пределах 1000', '-', 2, true, 1000),
(3, 'Табличное умножение', 'Табличное умножение в пределах 100', '*', 2, true, 100),
(3, 'Умножение двух чисел до 1000', 'Умножение двух чисел (каждое до 1000) с результатом до 1000', '*', 2, true, 1000),
(3, 'Деление чисел до 1000', 'Деление без остатка в пределах 1000', '/', 2, true, 1000),
(3, 'Смешанные выражения до 1000 (+ и -)', 'Сложение и вычитание трёх чисел в пределах 1000', '+-', 3, true, 1000),
(3, 'Смешанные выражения до 1000 (+ - *)', 'Сложение, вычитание и умножение (результат до 1000)', '+-*', 3, true, 1000),
(3, 'Смешанные выражения до 1000 (+ - /)', 'Сложение, вычитание и деление без остатка', '+-/', 3, true, 1000);
(4, 'Сложение чисел до 100000', 'Сложение двух чисел в пределах 100 000', '+', 2, true, 100000),
(4, 'Вычитание чисел до 100000', 'Вычитание чисел в пределах 100 000', '-', 2, true, 100000),
(4, 'Умножение на двузначное число', 'Умножение трёхзначного на двузначное (результат до 100 000)', '*', 2, true, 100000),
(4, 'Деление на двузначное число', 'Деление без остатка (результат и делимое до 100 000)', '/', 2, true, 100000),
(4, 'Смешанные выражения (+ -) до 100000', 'Сложение и вычитание 3-4 чисел в пределах 100 000', '+-', 4, true, 100000),
(4, 'Смешанные выражения (+ - *) до 100000', 'Сложение, вычитание, умножение (результат до 100 000)', '+-*', 4, true, 100000),
(4, 'Смешанные выражения (+ - /) до 100000', 'Сложение, вычитание, деление без остатка', '+-/', 4, true, 100000),
(4, 'Все действия (+, -, *, /) до 100000', 'Четыре арифметических действия без остатка', '+-*/', 4, true, 100000);

insert into operands (equation_type_id, operand_order, min_value, max_value) values
(1, 1, 0, 20),
(1, 2, 0, 20),
-- 2. Вычитание чисел до 20 (2 операнда)
(2, 1, 0, 20),
(2, 2, 0, 20),
-- 3. Сложение чисел до 100 (2 операнда)
(3, 1, 0, 100),
(3, 2, 0, 100),
-- 4. Вычитание чисел до 100 (2 операнда)
(4, 1, 0, 100),
(4, 2, 0, 100),
-- 5. Табличное умножение (2 операнда: множители от 1 до 9 или 2 до 9)
(5, 1, 1, 9),
(5, 2, 1, 9),
-- 6. Смешанные выражения (+ -) с 3 операндами
(6, 1, 0, 100),
(6, 2, 0, 100),
(6, 3, 0, 100),
-- 7. Смешанные выражения (+ * -) с 4 операндами
(7, 1, 0, 100),
(7, 2, 0, 100),
(7, 3, 0, 100),
(7, 4, 0, 100),
-- 8. Сложение чисел до 1000 (2 операнда)
(8, 1, 0, 1000),
(8, 2, 0, 1000),
-- 9. Вычитание чисел до 1000 (2 операнда)
(9, 1, 0, 1000),
(9, 2, 0, 1000),
-- 10. Табличное умножение (2 операнда)
(10, 1, 1, 9),
(10, 2, 1, 9),
-- 11. Умножение двух чисел до 1000 (2 операнда)
(11, 1, 1, 1000),
(11, 2, 1, 1000),
-- 12. Деление чисел до 1000 (2 операнда: делитель не 0)
(12, 1, 0, 1000),
(12, 2, 1, 1000),
-- 13. Смешанные (+ -) с 3 операндами
(13, 1, 0, 1000),
(13, 2, 0, 1000),
(13, 3, 0, 1000),
-- 14. Смешанные (+ - *) с 3 операндами
(14, 1, 0, 1000),
(14, 2, 0, 1000),
(14, 3, 0, 1000),
-- 15. Смешанные (+ - /) с 3 операндами
(15, 1, 0, 1000),
(15, 2, 0, 1000),
(15, 3, 1, 1000),
-- 16. Сложение чисел до 100000 (2 операнда)
(16, 1, 0, 100000),
(16, 2, 0, 100000),
-- 17. Вычитание чисел до 100000 (2 операнда)
(17, 1, 0, 100000),
(17, 2, 0, 100000),
-- 18. Умножение на двузначное число (2 операнда: трёхзначное * двузначное)
(18, 1, 100, 999),
(18, 2, 10, 99),
-- 19. Деление на двузначное число (2 операнда)
(19, 1, 0, 100000),
(19, 2, 10, 99),
-- 20. Смешанные (+ -) с 4 операндами
(20, 1, 0, 100000),
(20, 2, 0, 100000),
(20, 3, 0, 100000),
(20, 4, 0, 100000),
-- 21. Смешанные (+ - *) с 4 операндами
(21, 1, 0, 100000),
(21, 2, 0, 100000),
(21, 3, 0, 100000),
(21, 4, 0, 100000),
-- 22. Смешанные (+ - /) с 4 операндами
(22, 1, 0, 100000),
(22, 2, 0, 100000),
(22, 3, 1, 100000),
(22, 4, 1, 100000),
-- 23. Все действия (+ - * /) с 4 операндами
(23, 1, 0, 100000),
(23, 2, 0, 100000),
(23, 3, 1, 100000),
(23, 4, 1, 100000); 


insert into section_equation_types (section_id, eqaution_type_id) values
(1, 1),  -- Сложение чисел до 20
(2, 2),  -- Вычитание чисел до 20
(3, 3),  -- Сложение чисел до 100
(4, 4),  -- Вычитание чисел до 100
(5, 5),  -- Табличное умножение
(6, 6),  -- Смешанные выражение в пределах 100 (+ -)
(6, 7),  -- Смешанные выражение в пределах 100 (+ * -)
(7, 8),  -- Сложение чисел до 1000
(8, 9),  -- Вычитание чисел до 1000
(9, 10), -- Табличное умножение
(9, 11), -- Умножение двух чисел до 1000
(10, 12), -- Деление чисел до 1000
(11, 13), -- Смешанные выражения до 1000 (+ и -)
(11, 14), -- Смешанные выражения до 1000 (+ - *)
(11, 15), -- Смешанные выражения до 1000 (+ - /)
(12, 16), -- Сложение чисел до 100000
(13, 17), -- Вычитание чисел до 100000
(14, 18), -- Умножение на двузначное число
(15, 19), -- Деление на двузначное число
(16, 20), -- Смешанные выражения (+ -) до 100000
(16, 21), -- Смешанные выражения (+ - *) до 100000
(16, 22), -- Смешанные выражения (+ - /) до 100000
(16, 23); -- Все действия (+, -, *, /) до 100000


alter table attempts
add student_id int not null;

alter table attempts 
add constraint fk_student_id_attempts foreign key (student_id) references users(id);

CREATE VIEW student_stats AS
SELECT 
    student_id,
    equation_type_id,
    COUNT(*) as total_attempts,
    SUM(CASE WHEN correct_answer = student_answer THEN 1 ELSE 0 END) as correct_attempts,
    AVG(CASE WHEN correct_answer = student_answer THEN 1.0 ELSE 0 END) as success_rate,
    MAX(answered_at) as last_attempt
FROM attempts
GROUP BY student_id, equation_type_id;

CREATE VIEW student_section_stats AS
SELECT 
    student_id,
	section_id,
    equation_type_id,
    COUNT(*) as total_attempts,
    SUM(CASE WHEN correct_answer = student_answer THEN 1 ELSE 0 END) as correct_attempts,
    AVG(CASE WHEN correct_answer = student_answer THEN 1.0 ELSE 0 END) as success_rate,
    MAX(answered_at) as last_attempt
FROM attempts
JOIN section_equation_types ON section_equation_types.eqaution_type_id = attempts.equation_type_id
GROUP BY student_id, section_id, equation_type_id;
