create table bill(
billID int not null auto_increment, 
userID int not null,
productID int not null,
productNum int not null,
primary key(billID),
foreign key (productID) references product(productID),
foreign key (userID) references user(userID));

insert into bill(userID, productID,productNum)
values
(1,1,3),
(1,2,2),
(2,1,12),
(2,2,23);