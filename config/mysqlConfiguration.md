

Uninstall or Completely remove mysql from ubuntu 16-04
```sh
sudo apt-get remove --purge mysql*
sudo apt-get purge mysql*
sudo apt-get autoremove.
sudo apt-get autoclean.
sudo apt-get remove dbconfig-mysql.
sudo apt-get dist-upgrade.
sudo apt-get install mysql-server.
scp gfdbdump.sql gopherclub@gopherclub.jaynejacobs.com:~/.
 mysqladmin -p -u root version

systemctl enable mysql.service
  431  systemctl start mysql.service
  432  mysqladmin -p -u root version
  433  mysql -u root -p
```

DDL statements 	Database Definition Language

Create user:	CREATE USER 'username'@'%' IDENTIFIED BY 'password';
CREATE USER 'gopherclub'@'%' IDENTIFIED BY 'gopherclub';
Create DB: (backtics, not single quotes)	CREATE DATABASE IF NOT  EXISTS `gopherclubdb` DEFAULT CHARACTER SET 'utf8' COLLATE `utf8_unicode_ci`;
GRANT All PRIVILEGES ON gopherclubdb.* TO 'gopherclub'@'%';
FLUSH PRIVILEBES SO CHANGES TAKE AFFECT: 	FLUSH PRIVILEGES;

scp gfdbdump.sql gopherclub@gopherclub.jaynejacobs.com:~/.
mysql -u gopherclub -p gopherclubdb < /home/gopherclub/gfdbdump.sql 
mysql -u gopherclub -p gopherclubdb


```sql

mysql> show tables;

+------------------------+
| Tables_in_gopherclubdb |
+------------------------+
| friend_relation        |
| post                   |
| user                   |
| user_profile           |
+------------------------+

 select * from user;
```

