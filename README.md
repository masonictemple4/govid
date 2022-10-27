# govid
An opencv experiment in golang. The program will open your webcam and detect faces if it
detects less than 1 face 10 consecutive times it will play an alert and exit.

## Setup
Requires that you install opencv and pkgconfig.

`brew install opencv`

then run 

`brew install pkgconfig`


## Run
You can run this project a few different ways.

`go run main.go` from the source directory.
Or you can install it and run it using the following steps.
`go install` then `govid`
