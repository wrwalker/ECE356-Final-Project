create database If NOT Exists Election;
use Election;

source /var/lib/mysql-files/DatabaseLoadingScripts/load-tweets.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-state-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-county-data.sql;
