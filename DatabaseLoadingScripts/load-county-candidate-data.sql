SET foreign_key_checks = 0;
drop table if exists VotesByCountyCandidate;
SET foreign_key_checks = 1;
-- VotesByCountyCandidate ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create VotesByCountyCandidate' as '';

create table VotesByCountyCandidate
(
    state     char(20)     not null,
    county    char(30)     not null,
    candidate char(30)     not null,
    party     char(3)      not null,
    votes     int unsigned not null,
    won       bool         not null,
    level     char         not null,
    primary key (state, county, candidate, level)
);

-- Governor by County ------------------
load data infile '/var/lib/mysql-files/datasets/governors_county_candidate.csv' ignore into table VotesByCountyCandidate
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    candidate,
                    party,
                    votes,
                    @won,
                    @level)
    set level = 'G',
        won = if(@won = 'True', true, false);

-- Presidential by County ------------------
load data infile '/var/lib/mysql-files/datasets/president_county_candidate.csv' ignore into table VotesByCountyCandidate
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    candidate,
                    party,
                    votes,
                    @won,
                    @level)
    set level = 'P',
        won = if(@won = 'True', true, false);

-- Senate by County ------------------
load data infile '/var/lib/mysql-files/datasets/senate_county_candidate.csv' ignore into table VotesByCountyCandidate
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    candidate,
                    party,
                    votes,
                    @won,
                    @level)
    set level = 'S',
        won = if(@won = 'True', true, false);