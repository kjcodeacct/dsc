# initialize repo
dsc init
> prompt for alias, db type etc, or cancel to create one manually

# clone repo
dsc clone 127.0.0.1:5432
> pulls change long, and passwordless config down into .dsc cfg file

# add file
dsc add test1.sql
> set epoch for sorting

dsc add test2.sql
> set epoch for sorting (is now second applied file in changeset)

dsc rm test1.sql
> removes tracking for test1.sql and sets test2.sql as first

dsc add test1.sql
> set epoch for sorting (is new second applied file in changeset)


dsc add test3.sql test4.sql test5.sql
> pop up termui for sorting in cli editor

dsc commit -m "adding test data for use case"
> creates comit message

dsc push production
> applies changes against production 'branch' AKA server (found in cfg after init)