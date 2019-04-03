create database songslist;
use songslist
drop table songs;
create table songs (id serial, title varchar(32), author varchar(32), musicurl varchar(32));


insert into songs (title, author, musicurl) values ('One', 'U2', 'http://guitargods.com/U2/One');
select * from songs
