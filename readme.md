go get -u github.com/go-sql-driver/mysql
go get github.com/julienschmidt/httprouter
go get github.com/go-playground/validator/v10

create database go_customer_api;

create table category
(
    id integer primary key auto_increment,
    name varchar(200) not null
    
) engine = INNODB;