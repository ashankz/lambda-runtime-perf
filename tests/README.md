# Running performance test

## Steps to run jmeter load tests locally: 

_If jmeter is not installed:_
1. Download jmeter from [this location][https://jmeter.apache.org/download_jmeter.cgi]
2. Unzip the download package to a particular location. 
3. Add the jmeter bin folder to your PATH:
   On  MAC, run this in terminal (Double check the location and jemter version): export PATH=$PATH:~/Downloads/apache-jmeter-5.4.3/bin
   On Windows, follow this [link][https://stackoverflow.com/a/44272417] to add to PATH]


_If jmeter is installed:_

## Test Script

Run the following command after replacing the domain name with appropriate API gateway domain.
This test script executes 1000 requests against the api. Results are in the runs folder.
Check the correct location for the jmeter binary.

```

~/Downloads/apache-jmeter-5.4.3/bin/jmeter -n -t jmeter/setup/createapi-test.jmx  -Jdomain=4elyfdf53bn4.execute-api.us-west-2.amazonaws.com -Jthreads=50 -Jrampup=4 -Jduration=100 -Jiterations=20 -e -l jmeter/runs/logs_$(date '+%Y%m%d%H%M%S').jtl -o jmeter/runs/run_$(date '+%Y%m%d%H%M%S')

```


## Postman collection
Import the collection and set the variables to appropriate API gateway domains

domain-node
domain-go
domain-python

 