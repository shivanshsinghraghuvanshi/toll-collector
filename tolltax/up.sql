Create table if not exists owner(
    ownerid INT GENERATED ALWAYS AS IDENTITY,
    accountnumber bigserial,
    name varchar(40) not null,
    primary key(ownerid)
);
create table if not exists car(
    carid INT GENERATED ALWAYS AS IDENTITY,
    make varchar(30),
    cartype varchar(10) not null,
    carnumber varchar(12) not null,
    primary key(carid)
);

create table if not exists netc(
    netcid INT GENERATED ALWAYS AS IDENTITY,
    ownerid bigserial,
    carid bigserial,
    rfid varchar(50) not null,
    primary key(netcid),
    foreign key(ownerid) references owner(ownerid),
    foreign key(carid) references car(carid)
);

create table if not exists tollbooth(
    tollboothid INT GENERATED ALWAYS as IDENTITY,
    name varchar(40),
    accountnumber bigserial,
    primary key(tollboothid)
);


create table if not exists deductible(
    id INT GENERATED ALWAYS as IDENTITY,
    cartype varchar(10) not null,
    amount INT not null
);

INSERT INTO owner(accountnumber, name) values (1234567891, 'Shivansh Singh Raghuvanshi');
INSERT INTO owner(accountnumber, name) values (1624725182, 'Devansh Raghuvanshi');
INSERT INTO car(make, cartype, carnumber) values ('Maruti Suzuki Vitara Brezza', 'LMV', 'KA01MP0226');
INSERT INTO car(make, cartype, carnumber) values ('Renault Kwid', 'LMV', 'KA09MP9978');
INSERT INTO tollbooth(accountnumber, name) values (9876543219, 'NICE Road Toll Booth');
INSERT INTO tollbooth(accountnumber, name) values (6379352314, 'Hosur Toll Booth');
INSERT INTO deductible(cartype, amount) values ('LMV', 100);

