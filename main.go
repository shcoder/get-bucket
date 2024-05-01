package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	ShardCount := 0
	SchemaCount := 0
	isPS := false
	isPDS := false
	flag.IntVar(&ShardCount, "shard", 0, "shard count")
	flag.IntVar(&SchemaCount, "schema", 0, "schema count")
	flag.BoolVar(&isPS, "ps", false, "posting-service")
	flag.BoolVar(&isPDS, "pds", false, "posting-data-service")
	flag.Parse()

	if isPS {
		ShardCount = 256
		SchemaCount = 32
	} else if isPDS {
		ShardCount = 512
		SchemaCount = 32
	}

	if ShardCount == 0 || SchemaCount == 0 {
		fmt.Printf("shard count and schema count must be greater than zero\n")
		return
	}

	companyIDStr := flag.Arg(0)
	if len(companyIDStr) == 0 {
		fmt.Printf("Not set company ID\n")
		return
	}
	companyId, err := strconv.Atoi(companyIDStr)
	if err != nil {
		fmt.Printf("Invalid company ID: %s\n", companyIDStr)
		return
	}

	fmt.Printf("Company ID: %d\n", companyId)
	fmt.Printf("Shard: %d\n", companyId%ShardCount/SchemaCount+1)
	fmt.Printf("Bucket: %d\n", companyId%ShardCount)
}
