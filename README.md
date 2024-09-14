# Process Log Records


### Problem Statement

Write a program that can parse a file containing flow log data and maps each row to a tag based on a lookup table. The lookup table is defined as a csv file, and it has 3 columns, dstport,protocol,tag. The dstport and protocol combination decide what tag can be applied.  

Input:

- Log file with records, Size - 10 MB
- Lookup Table CSV file - dstport,protocol,tag fields


Intuition:

- Each log record is mapped to a tag based on the lookup table.
- The tag is unique based on the protocol and dstport combination.
- We want to know `tag` occurrences and `port, protocol` occurrences. Since, the lookup table has information about the tag, port and protocol, it can be used to derive all our results.
- We will be reading `lookup.txt` and defining a map to store `tag` and `port, protocol` occurrences.

Assumptions:

- `lookup` table is relevant in this case.
- when `tag` field is empty, we assign `untagged` as its tag.
- input is structured as `port,protocol,tag`.


### Run the program

To compile the program, clone the github repository:

```
git clone https://github.com/salonich/flow-log-records.git
```

```
cd flow-log-records
```

```
go build ./
go run flow-log-records <lookup.txt>
```

### Tests

Test is located in file `log-records-test.go`. To run unit tests:

```
go test -v
```

