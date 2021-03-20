-- this solution is written using MySQL syntax

-- create USER table
CREATE TABLE USER(ID INTEGER, UserName VARCHAR(25), Parent INTEGER);

-- insert data to USER table
INSERT INTO USER(ID, UserName, Parent) 
    VALUES (1, "Ali", 2), (2, "Budi", 0), (3, "Cecep", 1);

-- query to get all user data with their respctive "Creator"
SELECT U.ID, U.UserName,
    (
        SELECT UserName
        FROM USER
        WHERE ID = U.Parent
    ) AS ParentUserName
FROM USER AS U