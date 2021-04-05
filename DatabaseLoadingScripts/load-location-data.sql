-- Location ------------------------------------------------------------------------
select '-----------------------------------------------------------------' as '';
select 'Create Location' as '';

create table Location
(
    tweetID     BIGINT          not null,
    latitude    decimal(25, 20) not null CHECK (latitude <> ''),
    longitude   decimal(25, 20) not null CHECK (longitude <> ''),
    county_name char(20)        not null,
    state_name  char(20)        not null,
    primary key (tweetID)
);

-- Governor by State ------------------
load data infile '/var/lib/mysql-files/datasets/lat_lon.csv' ignore into table Location
    fields terminated by ','
    enclosed by '"'
    lines terminated by '\n'
    ignore 1 lines (
                    @lat,
                    @lon,
                    @tweet_id,
                    @county_nm,
                    @state_nm)
    set tweetID = @tweet_id,
        latitude = @lat,
        longitude = @lon,
        county_name = @county_nm,
        state_name = @state_nm;