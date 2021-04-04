drop table if exists ElectionTweets;
-- ElectionTweets ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create ElectionTweets' as '';

create table ElectionTweets
(
    tweetID            BIGINT          not null,
    tweet              varchar(1960)   not null CHECK (tweet <> ''),
    likes              int unsigned,
    retweetCount       int unsigned,
    userID             BIGINT unsigned not null,
    userFollowersCount int unsigned,
    latitude           decimal(25, 20) not null CHECK (latitude <> ''),
    longitude          decimal(25, 20) not null CHECK (longitude <> ''),
    city               varchar(255),
    country            char(30),
    stateCode          char(4),
    trumpOrBiden       char(1)         not null,
    sentimentScore     decimal(5, 2),
    primary key (tweetID)
);

load data infile '/var/lib/mysql-files/datasets/hashtag_donaldtrump.csv' ignore into table ElectionTweets
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    @throwAway,
                    tweetID,
                    tweet,
                    likes,
                    retweetCount,
                    @throwAway,
                    userID,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    userFollowersCount,
                    @throwAway,
                    latitude,
                    longitude,
                    city,
                    country,
                    @throwAway,
                    @throwAway,
                    stateCode,
                    @throwAway,
                    @TrumpOrBiden)
    set trumpOrBiden = 'T';

load data infile '/var/lib/mysql-files/datasets/hashtag_joebiden.csv' ignore into table ElectionTweets
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    @throwAway,
                    tweetID,
                    tweet,
                    likes,
                    retweetCount,
                    @throwAway,
                    userID,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    userFollowersCount,
                    @throwAway,
                    latitude,
                    longitude,
                    city,
                    country,
                    @throwAway,
                    @throwAway,
                    stateCode,
                    @throwAway,
                    @TrumpOrBiden)
    set trumpOrBiden = 'B';
