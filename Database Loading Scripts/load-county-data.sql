drop table if exists VotesByCounty;
-- VotesByCounty ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create VotesByCounty' as '';

create table VotesByCounty
(
    state         char(20)     not null,
    county        char(30)     not null,
    currentVotes int unsigned not null,
    totalVotes    int unsigned not null,
    percent       int unsigned,
    level         char,
    primary key (state, county, level)
);

-- Governor by County ------------------
load data infile '/var/lib/mysql-files/ece356/us-election-2020/governors_county.csv' ignore into table VotesByCounty
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    currentVotes,
                    totalVotes,
                    percent,
                    @level)
    set level = 'G';

-- Presidential by County ------------------
load data infile '/var/lib/mysql-files/ece356/us-election-2020/president_county.csv' ignore into table VotesByCounty
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    currentVotes,
                    totalVotes,
                    percent,
                    @level)
    set level = 'P';

-- Senate by County ------------------
load data infile '/var/lib/mysql-files/ece356/us-election-2020/senate_county.csv' ignore into table VotesByCounty
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    currentVotes,
                    totalVotes,
                    percent,
                    @level)
    set level = 'S';