# ECE356-Final-Project
The final project of our ECE 356 course. In this project, we build a MySQL based database system to predict the results of the 2020 US Election using sentiment analysis on Twitter data for #trump and #joebiden.

# Docker set up
## Prerequisites
We use docker to containerize our db. If you don't already have it installed, run this from the repo root
```
$ make presSetup
```
You will also need to add the needed csv files to the `datasets` dir. Due to github's file limits, you will need to add the
`hastag_donaldtrump.csv` and `hashtag_joebiden.csv` files found [here](https://www.kaggle.com/manchunhui/us-election-2020-tweets).

## Load up the DB and connect to mysql cli
To create, build and start the container, load the csvs into the tables and connect to the mysql cli run this from the repo root
```
$ make startContainerAndLoad
```

## Only connect to mysql cli
If you already have the container running:
```
$ make connectToContainer
```