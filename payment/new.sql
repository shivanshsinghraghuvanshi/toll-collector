CREATE TABLE if not exists accountdetails
(
    accountid     INT GENERATED ALWAYS AS IDENTITY,
    accountnumber bigserial   not null,
    name          varchar(40) not null,
    balance       numeric     not null,
    lastUpdated   timestamp
);

CREATE TABLE if not exists transactionDetails
(
    id                  INT GENERATED ALWAYS AS IDENTITY,
    timestamp           timestamp,
    debitaccountnumber  bigserial not null,
    creditaccountnumber bigserial not null,
    amount              numeric   not null,
    remarks             varchar(100)
);

INSERT INTO accountdetails(accountnumber, name, balance, lastUpdated)
values (1234567891, 'Shivansh Singh Raghuvanshi', 10000.00, '2016-06-22 19:10:25-07');

INSERT INTO accountdetails(accountnumber, name, balance, lastUpdated)
values (1624725182, 'Devansh Raghuvanshi', 10000.00, '2016-06-22 19:10:25-07');

INSERT INTO accountdetails(accountnumber, name, balance, lastUpdated)
values (9876543219, 'NICE Road Toll Booth', 10000.00, '2016-06-22 19:10:25-07');

INSERT INTO accountdetails(accountnumber, name, balance, lastUpdated)
values (6379352314, 'Hosur Toll Booth', 10000.00, '2016-06-22 19:10:25-07');
