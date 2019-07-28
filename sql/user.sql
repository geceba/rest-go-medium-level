CREATE TABLE users
(
  id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username varchar(30) NOT NULL,
  password varchar(64) NOT NULL,
  email varchar(40),
  created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, first_name, last_name) values ('geceba', 'Gerardo', 'Cetzal Balam');
