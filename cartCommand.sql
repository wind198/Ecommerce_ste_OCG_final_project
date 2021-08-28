create table cart(
userID int not null, 
productID int not null,
productNum int not null,
foreign key (productID) references product(productID),
foreign key (userID) references user(userID));

insert into cart(userID, productID,productNum)
values
(1,1,1),
(2,1,3),
(2,2,5);