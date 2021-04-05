drop table if exists User;
-- User ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create User' as '';

create table User
(
    userID             BIGINT unsigned not null,
    userFollowersCount int unsigned,
    primary key (userID)
);

load data infile '/var/lib/mysql-files/datasets/hashtag_donaldtrump.csv' ignore into table User
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    userID,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    userFollowersCount,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway);

load data infile '/var/lib/mysql-files/datasets/hashtag_joebiden.csv' ignore into table User
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    userID,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    userFollowersCount,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway);
