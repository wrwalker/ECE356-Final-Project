SET foreign_key_checks = 0;
drop table if exists County;
SET foreign_key_checks = 1;
-- County ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create County' as '';

create table County
(
    state       char(30) not null,
    county      char(30) not null,
    annotations varchar(2048) not null,
    primary key (state, county)
);

-- Presidential by County ------------------
load data infile '/var/lib/mysql-files/datasets/president_county.csv' ignore into table County
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    state,
                    county,
                    @throwAway,
                    @throwAway,
                    @throwAway)
    set annotations = '';