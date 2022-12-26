CREATE TABLE IF NOT EXISTS expense (id SERIAL  PRIMARY KEY , title TEXT , amount FLOAT , note TEXT ,tags TEXT[]);
Insert into expense (title, amount , note,tags) values ('SeedData1' , 20.0,'SomeNote',ARRAY['tags1']);
Insert into expense (title, amount , note,tags) values ('SeedData2' , 20.0,'SomeNote',ARRAY['tags1']);