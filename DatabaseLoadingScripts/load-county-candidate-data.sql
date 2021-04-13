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
    votes     int unsigned not null,
    won       bool         not null,
    primary key (state, county, candidate),
    foreign key (candidate) references Candidate(candidate)
);

-- Presidential by County ------------------
load data infile '/var/lib/mysql-files/datasets/president_county_candidate.csv' ignore into table VotesByCountyCandidate
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    candidate,
                    @throwAway,
                    votes,
                    @won)
    set won = if(@won = 'True', true, false);