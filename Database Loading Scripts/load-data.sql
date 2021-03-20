use Election;
drop table if exists ElectionTweets;
SHOW VARIABLES LIKE "secure_file_priv";
-- Game ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create ElectionTweets' as '';



-- create table ElectionTweets (state char(14),
-- 							county char(25),
--                             candidate char(40),
--                             party char(4),
--                             votes, 
--                             won

create table ElectionTweets (tweetID BIGINT not null,
							tweet varchar(1960) not null,
                            likes int,
                            retweetCount int,
                            userID BIGINT not null,
                            userFollowersCount int,
                            city varchar(255),
                            country char(30) not null,
                            stateCode char(4) not null,
                            trumpOrBiden char(1) not null,
                            sentimentScore decimal(5, 2),
                            primary key (tweetID));

load data infile '/var/lib/mysql-files/ece356/election-tweets/hashtag_donaldtrump.csv' ignore into table ElectionTweets
     fields terminated by ','
     enclosed by '"'
     lines terminated by '\n'
     ignore 1 lines(
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
     city,
     country,
     @throwAway,
     @throwAway,
     stateCode,
     @throwAway,
     @TrumpOrBiden)
     set trumpOrBiden = 'T';
     
load data infile '/var/lib/mysql-files/ece356/election-tweets/hashtag_joebiden.csv' ignore into table ElectionTweets
     fields terminated by ','
     enclosed by '"'
     lines terminated by '\n'
     ignore 1 lines(
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
     city,
     country,
     @throwAway,
     @throwAway,
     stateCode,
     @throwAway,
     @TrumpOrBiden)
     set trumpOrBiden = 'B';
