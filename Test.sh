#!/bin/bash
function timediff() {

# time format:date +"%s.%N", such as 1502758855.907197692
    start_time=$1
    end_time=$2
    
    start_s=${start_time%.*}
    start_nanos=${start_time#*.}
    end_s=${end_time%.*}
    end_nanos=${end_time#*.}
    
    # end_nanos > start_nanos? 
    # Another way, the time part may start with 0, which means
    # it will be regarded as oct format, use "10#" to ensure
    # calculateing with decimal
    if [ "$end_nanos" -lt "$start_nanos" ];then
        end_s=$(( 10#$end_s - 1 ))
        end_nanos=$(( 10#$end_nanos + 10**9 ))
    fi
    
# get timediff
    time=$(( 10#$end_s - 10#$start_s )).`printf "%03d\n" $(( (10#$end_nanos - 10#$start_nanos)/10**6 ))`
    
    echo $time
}

# ego-go build -o sgx_new_abs sgx_abs.go lagRange.go define.go
# ego sign sgx_new_abs
# OE_SIMULATION=1 ego run sgx_new_abs -t 7 -n 10

# echo "t = 300, n = 450"
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 300 -n 450
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 300 -n 450
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# echo "t = 2, n = 3"
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 2 -n 3
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# echo "t = 2, n = 3"
# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 2 -n 3
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 2 -n 3
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# echo "t = 7, n = 10"
# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 7 -n 10
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 7 -n 10
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# echo "t = 20, n = 30"
# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 20 -n 30
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 20 -n 30
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 40 -n 60
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 40 -n 60
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 60 -n 90
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 60 -n 90
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 80 -n 120
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 80 -n 120
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

# startTime=$(date +"%s.%N")
# OE_SIMULATION=1 ego run sgx_new_abs -t 100 -n 150
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime
# startTime=$(date +"%s.%N")
# ./nonsgx_abs -t 100 -n 150
# endTime=$(date +"%s.%N")
# timediff $startTime $endTime

startTime=$(date +"%s.%N")
OE_SIMULATION=1 ego run sgx_new_abs -t 10 -n 15
endTime=$(date +"%s.%N")
timediff $startTime $endTime
startTime=$(date +"%s.%N")
./nonsgx_abs -t 10 -n 15
endTime=$(date +"%s.%N")
timediff $startTime $endTime

startTime=$(date +"%s.%N")
OE_SIMULATION=1 ego run sgx_new_abs -t 30 -n 45
endTime=$(date +"%s.%N")
timediff $startTime $endTime
startTime=$(date +"%s.%N")
./nonsgx_abs -t 30 -n 45
endTime=$(date +"%s.%N")
timediff $startTime $endTime

startTime=$(date +"%s.%N")
OE_SIMULATION=1 ego run sgx_new_abs -t 50 -n 75
endTime=$(date +"%s.%N")
timediff $startTime $endTime
startTime=$(date +"%s.%N")
./nonsgx_abs -t 50 -n 75
endTime=$(date +"%s.%N")
timediff $startTime $endTime

startTime=$(date +"%s.%N")
OE_SIMULATION=1 ego run sgx_new_abs -t 70 -n 105
endTime=$(date +"%s.%N")
timediff $startTime $endTime
startTime=$(date +"%s.%N")
./nonsgx_abs -t 70 -n 105
endTime=$(date +"%s.%N")
timediff $startTime $endTime
