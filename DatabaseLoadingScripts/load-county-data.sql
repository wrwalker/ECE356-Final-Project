SET foreign_key_checks = 0;
drop table if exists VotesByCounty;
SET foreign_key_checks = 1;
-- VotesByCounty ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create VotesByCounty' as '';

create table VotesByCounty
(
    state      char(30)     not null,
    county     char(30)     not null,
    totalVotes int unsigned not null,
    percent    int unsigned,
    primary key (state, county),
    foreign key (state, county) references County(state, county)
);

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
                    percent);