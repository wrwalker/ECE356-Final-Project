-- VotesByState ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create VotesByState' as '';

create table VotesByState
(
    state      char(20)     not null,
    totalVotes int unsigned not null,
    level      char,
    primary key (state, level)
);

-- Governor by State ------------------
load data infile '/var/lib/mysql-files/datasets/governors_state.csv' ignore into table VotesByState
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    totalVotes,
                    @level)
    set level = 'G';

-- Presidential by State ------------------
load data infile '/var/lib/mysql-files/datasets/president_state.csv' ignore into table VotesByState
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    totalVotes,
                    @level)
    set level = 'P';

-- Senate by State ------------------
load data infile '/var/lib/mysql-files/datasets/senate_state.csv' ignore into table VotesByState
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    totalVotes,
                    @level)
    set level = 'S';