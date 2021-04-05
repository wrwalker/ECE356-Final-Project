create database If NOT Exists Election;
use Election;
SET foreign_key_checks = 0;
drop table if exists VotesByCounty;
drop table if exists VotesByCountyCandidate;
drop table if exists Location;
drop table if exists VotesByState;
drop table if exists ElectionTweets;
drop table if exists User;
SET foreign_key_checks = 1;

source /var/lib/mysql-files/DatabaseLoadingScripts/load-user-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-location-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-tweets.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-state-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-county-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-county-candidate-data.sql;