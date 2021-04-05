SET foreign_key_checks = 0;
drop table if exists VotesByCounty;
SET foreign_key_checks = 1;
-- VotesByCounty ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create VotesByCounty' as '';

create table VotesByCounty
(
    state         char(20)     not null,
    county        char(30)     not null,
    totalVotes    int unsigned not null,
    percent       int unsigned,
    level         char,
    primary key (state, county, level)
);

-- Governor by County ------------------
load data infile '/var/lib/mysql-files/datasets/governors_county.csv' ignore into table VotesByCounty
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    @throwAway,
                    totalVotes,
                    percent,
                    @level)
    set level = 'G';

-- Presidential by County ------------------
load data infile '/var/lib/mysql-files/datasets/president_county.csv' ignore into table VotesByCounty
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    @throwAway,
                    totalVotes,
                    percent,
                    @level)
    set level = 'P';

-- Senate by County ------------------
load data infile '/var/lib/mysql-files/datasets/senate_county.csv' ignore into table VotesByCounty
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    @throwAway,
                    totalVotes,
                    percent,
                    @level)
    set level = 'S';