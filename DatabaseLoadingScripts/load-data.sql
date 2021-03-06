create database If NOT Exists Election;
use Election;

source /var/lib/mysql-files/DatabaseLoadingScripts/load-county-annotations-data.sql
source /var/lib/mysql-files/DatabaseLoadingScripts/load-candidate-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-sentiment-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-user-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-location-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-tweets.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-state-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-county-data.sql;
source /var/lib/mysql-files/DatabaseLoadingScripts/load-county-candidate-data.sql;
