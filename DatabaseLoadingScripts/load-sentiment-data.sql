SET foreign_key_checks = 0;
drop table if exists Sentiment;
SET foreign_key_checks = 1;
-- Sentiment ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create Sentiment' as '';

create table Sentiment
(
    tweetID        BIGINT          not null,
    tweet          varchar(1960)   not null CHECK (tweet <> ''),
    likes          int unsigned,
    retweetCount   int unsigned,
    userID         BIGINT unsigned not null,
    trumpOrBiden   char(1)         not null,
    sentimentScore bool,
    primary key (tweetID),
    foreign key (tweetID) references Location (tweetID),
    foreign key (userID) references User (userID)
);

load data infile '/var/lib/mysql-files/datasets/hashtag_donaldtrump.csv' ignore into table Sentiment
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
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @throwAway,
                    @TrumpOrBiden);