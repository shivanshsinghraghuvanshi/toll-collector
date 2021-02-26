Create table if not exists owner{
    ownerid INT GENERATED ALWAYS AS IDENTITY,
    accountnumber bigserial,
    name varchar(40) not null ,
    primary key(ownerid)
    }

create table if not exists car{
    carid INT GENERATED ALWAYS AS IDENTITY,
    make varchar (30),
    cartype varchar (10) not null,
    carnumber varchar (12) not null,
    primary key(carid)
    }

create table if not exists netc{
    netcid INT GENERATED ALWAYS AS IDENTITY,
    primary key(netcid)
    rfid varchar (50) not null ,
    constraint fkownerid foreign key(ownerid) references owner(ownerid)
    constraint fkcarid foreign key(carid) references car(carid)
    }

create table if not exists tollbooth{
    tollboothid INT GENERATED ALWAYS as IDENTITY,
    name varchar(40),
    accountnumber bigserial,
    primary key(tollboothid)
    }


create table if not exists deductible{
    id INT GENERATED ALWAYS as IDENTITY,
    cartype varchar (10) not null,
    amount INT not null
    }
