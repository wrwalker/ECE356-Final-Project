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
							Tweet varchar(300) not null,
                            Likes int,
                            RetweetCount int,
                            UserID BIGINT not null,
                            UserFollowersCount int,
                            City char(20),
                            Country char(20) not null,
                            StateCode char(4) not null,
                            primary key (TweetID));
                            

-- gameID decimal(10),
--        	     	   season decimal(8) not null,
-- 				   type char(1) not null,
-- 				   dateTimeGMT datetime not null,
-- 				   awayTeamID int not null,
-- 				   homeTeamID int not null,
-- 				   awayGoals int not null check(awayGoals >= 0),
-- 				   homeGoals int not null check(homeGoals >= 0),
-- 				   outcome char(12) not null,
-- 				   homeRinkSideStart char(50) not null,
-- 				   venue varchar(64) not null,
-- 				   venueTimeZoneID char(20) not null,
-- 				   venueTimeZoneOffset char(50) not null,
-- 				   venueTimeZoneTZ char(3) not null,
-- -- Key Constraints
-- 				   primary key (gameID),
-- 				   foreign key (awayTeamID) references TeamInfo(teamID),
-- 				   foreign key (homeTeamID) references TeamInfo(teamID)
-- -- Integrity Constraints
-- -- Set above showing inline format
-- 		  );

load data infile '/var/lib/mysql-files/ece356/election-tweets/hashtag_donaldtrump.csv' ignore into table ElectionTweets
     fields terminated by ','
     enclosed by '"'
     lines terminated by '\n'
     ignore 1 lines;
