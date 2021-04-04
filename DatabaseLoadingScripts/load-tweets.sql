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
    trumpOrBiden       char(1)         not null,
    sentimentScore     decimal(5, 2),
    primary key (tweetID),
    foreign key (tweetID) references Location (tweetID)
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
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
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
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @TrumpOrBiden)
    set trumpOrBiden = 'B';
