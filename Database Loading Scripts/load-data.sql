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

create table ElectionTweets (TweetID BIGINT not null,
							Tweet varchar(1960) not null,
                            Likes int,
                            RetweetCount int,
                            UserID BIGINT not null,
                            UserFollowersCount int,
                            City varchar(255),
                            Country char(30) not null,
                            StateCode char(4) not null,
                            primary key (TweetID));

load data infile '/var/lib/mysql-files/ece356/election-tweets/hashtag_donaldtrump.csv' ignore into table ElectionTweets
     fields terminated by ','
     enclosed by '"'
     lines terminated by '\n'
     ignore 1 lines(
     @throwAway,
     TweetID,
     Tweet,
     Likes,
     RetweetCount,
     @throwAway,
     UserID,
     @throwAway,
     @throwAway,
     @throwAway,
     @throwAway,
     UserFollowersCount,
     @throwAway,
     @throwAway,
     @throwAway,
     City,
     Country,
     @throwAway,
     @throwAway,
     StateCode,
     @throwAway);
