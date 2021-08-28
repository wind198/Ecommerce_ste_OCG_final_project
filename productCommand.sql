create table product(
productID int not null auto_increment, 
productName varchar(255) not null,
productPrice float(15,2) not null,
primary key(productID)
);

insert into product(productName,productPrice)
values
("ban",12.2),
("ghe",14.2);