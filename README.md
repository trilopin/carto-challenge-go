
## Carto Challenge
Build the following and make it run as fast as you possibly can using Python 3 (vanilla). The faster it runs, the more you will impress us!

Your code should:

Download this 2.2GB file: https://s3.amazonaws.com/carto-1000x/data/yellow_tripdata_2016-01.csv
Count the lines in the file
Calculate the average value of the tip_amount field.
All of that in the most efficient way you can come up with.

So, we'd do:

$ time python script.py https://s3.amazonaws.com/carto-1000x/data/yellow_tripdata_2016-01.csv
AVG_VALUE
x.00s user x.00s system 81% cpu x.000 total
That's it. Make it fly!


## Solution

No language wars is intended, I am only learning while playing.

Get it
```
$ go get github.com/trilopin/carto-challenge-go
```

Run tests

```
$ cd GOPATH/src/github.com/trilopin/carto-challenge-go
$ dep ensure && go test
```

### Simple stream

It takes *12 minutes* (macbook pro (i5) and a 50MB Fiber network).
```
$ carto-challenge-go stream_simple https://s3.amazonaws.com/carto-1000x/data/yellow_tripdata_2016-01.csv

Total lines: 10906858
Average Tip amount 1.7507./carto-challenge stream_simple
34.77s user 23.82s system 8% cpu 12:00.71 total
```

