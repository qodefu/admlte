
--@block
CREATE USER app1;

--@block
CREATE DATABASE app1;

--@block
GRANT ALL PRIVILEGES ON DATABASE app1 TO app1;

--@block
ALTER DATABASE app1 OWNER to app1;

--@block
ALTER USER app1 with password 'app1'
--@block
select * from user;
