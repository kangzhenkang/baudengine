# Roadmap of Baud Development

## First, Storage

A scalable document storage, just store, no search, no graph

Every json document has a builtin uint32 field, _slotid_, which indicates the partition (slot range) it will be located at. When it is inserted, the accepting PartitionServer (PS) raft group will allocate a UID for the document. 


## Indexing

After implementing the document CRUD functions, then the Indexing capability inside partitions. 

numberic, boolean, string match, fulltext, geo... 

We build a unified index tech for the above data types. 

both data and index are converted into key-value log-structured store. 

## Search & Ranking

## SQL

tables sharding the same partition key belong to one space. 

MyRouter will seperated from JSONRouter

## Graph

edge partitions are co-located with the source partitions. 


