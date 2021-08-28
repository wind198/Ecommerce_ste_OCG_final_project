create table user(
userID int not null auto_increment, 
userEmail varchar(255) not null,
userPassword varchar(255) not null,
primary key(userID)
);

insert into user(userEmail,userPassword)
values
("user1@example.com","hello1234"),
("user2@example.com","loveyou123");