create table phone_numbers(
    id      serial primary key, 
    phone   text                not null
);

insert into phone_numbers (phone) values ('1234567890');
insert into phone_numbers (phone) values ('123 456 7891');
insert into phone_numbers (phone) values ('(123) 456 7892');
insert into phone_numbers (phone) values ('(123) 456-7893');
insert into phone_numbers (phone) values ('123-456-7894');
insert into phone_numbers (phone) values ('123-456-7890');
insert into phone_numbers (phone) values ('1234567892');
insert into phone_numbers (phone) values ('(123)456-7892');