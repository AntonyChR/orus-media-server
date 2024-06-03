#!/bin/bash

#Create fake data

mkdir -p media
mkdir -p subtitles

cd media

# create "movies"
touch "The matrix(1999).mp4"
touch "Godzilla(2014).mp4"
touch "The lord of the rings(2001).mp4"
touhc "avatar(2009).mp4"
touch "The dark knight(2008).mp4"
touch "Dune(2021).mp4"
touch "The shawshank redemption(1994).mp4"
touch "Pacific rim(2013).mp4"
touch "The avengers(2012).mp4"
touch "The hunger games(2012).mp4"

# create series
mkdir -p "Friends (1994)"
mkdir -p "Breaking bad (2008)"
mkdir -p "Game of thrones (2011)"
mkdir -p "The walking dead (2010)"
mkdir -p "The big bang theory (2007)"


# create episodes, format sXeX.mp4

cd "Game of thrones (2011)"

for ((i=1; i<=3; i++));do
    for ((j=1; j<=5; j++));do
        touch "s${i}e${j}.mp4"
    done
done



