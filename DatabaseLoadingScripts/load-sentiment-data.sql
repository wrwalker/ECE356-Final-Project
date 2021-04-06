SET foreign_key_checks = 0;
drop table if exists Sentiment;
SET foreign_key_checks = 1;
-- Sentiment ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create Sentiment' as '';

create table Sentiment
(
    tweetID        BIGINT  not null,
    trumpOrBiden   char(1) not null,
    sentimentScore bool,
    primary key (tweetID)
);

load data infile '/var/lib/mysql-files/datasets/new_sentiment.csv' ignore into table Sentiment
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    tweetID,
                    @trumpOrBiden,
                    sentimentScore)
    set trumpOrBiden = if(@trumpOrBiden = 'b', 'B', 'T');