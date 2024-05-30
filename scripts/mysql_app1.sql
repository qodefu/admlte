CREATE database app1;

--@block
select * from mysql.user;
--@block
show databases


--@block
drop user 'app1'@'%'
--@block
CREATE USER 'app1'@'%' IDENTIFIED BY 'app1';

--@block
FLUSH PRIVILEGES;

--@block
GRANT ALL PRIVILEGES ON *.* TO 'app1'@'%' WITH GRANT OPTION;
--@block
SHOW GRANTS FOR 'app1'@'%'