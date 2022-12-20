CREATE TABLE IF NOT EXISTS expense (id SERIAL  PRIMARY KEY , title TEXT , amount FLOAT , note TEXT);

Insert into expense (title, amount , note) values ('SomeTitle' , 20.0,'SomeNote')