CREATE TABLE Equation_attempts (
    Id SERIAL PRIMARY KEY,
    Student_Id INT NOT NULL,
    Equation_Id INT NOT NULL,
    Given_answer INT,
    Correct BOOLEAN,
    Attempted_at TIMESTAMP,
    FOREIGN KEY(Student_Id) REFERENCES Users(Id),
    FOREIGN KEY(Equation_Id) REFERENCES Equations(Id)
);