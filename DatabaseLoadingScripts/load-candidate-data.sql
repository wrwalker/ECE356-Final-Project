SET foreign_key_checks = 0;
drop table if exists Candidate;
SET foreign_key_checks = 1;
-- Candidate ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create Candidate' as '';

create table Candidate
(
    candidate char(30) not null,
    party     char(3)  not null,
    primary key (candidate)
);

-- Presidential by County ------------------
load data infile '/var/lib/mysql-files/datasets/president_county_candidate.csv' ignore into table Candidate
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    @throwAway,
                    @throwAway,
                    candidate,
                    party,
                    @throwAway,
                    @throwAway);