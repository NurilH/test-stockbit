-- Membuat database dengan nama db_test_stoctbit
CREATE DATABASE db_test_stockbit;

-- memilih untuk menggunakan database yang tadi dibuat  
USE db_test_stockbit;

-- membuat tabel users 
CREATE TABLE users(
ID INT PRIMARY KEY,
UserName VARCHAR(20),
Parent INT
);

-- mengisi data tabel users  
INSERT INTO users(ID,UserName,Parent)
VALUES (1,'Ali',2),(2,'Budi',0),(3,'Cecep',1);

-- menampilkan data tabel user sesuain soal 
SELECT a.ID, a.UserName, b.UserName AS ParentUserName
FROM users a
LEFT JOIN users b ON a.Parent = b.ID;

