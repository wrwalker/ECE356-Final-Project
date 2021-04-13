SET foreign_key_checks = 0;
drop table if exists VotesByState;
SET foreign_key_checks = 1;
-- VotesByState ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create VotesByState' as '';

create table VotesByState
(
    state      char(20)     not null,
    totalVotes int unsigned not null,
    primary key (state)
);

-- Presidential by State ------------------
load data infile '/var/lib/mysql-files/datasets/president_state.csv' ignore into table VotesByState
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    totalVotes);