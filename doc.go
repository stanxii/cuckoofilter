// Package cuckoofilter is an implementation of the Cuckoo filter, a Bloom filter
// replacement for approximated set-membership queries. Cuckoo filters support
// adding and removing items dynamically while achieving even higher performance
// than Bloom filters.
//
// As documented in the original implementation:
//      Cuckoo filters provide the flexibility to add and remove items dynamically. A
//      cuckoo filter is based on cuckoo hashing (and therefore named as cuckoo
//      filter). It is essentially a cuckoo hash table storing each key's fingerprint.
//      Cuckoo hash tables can be highly compact, thus a cuckoo filter could use less
//      space than conventional Bloom filters, for applications that require low false
//      positive rates (< 3%).
//
// For details about the algorithm and citations please refer to the original
// research paper, "Cuckoo Filter: Better Than Bloom" by Bin Fan, Dave Andersen
// and Michael Kaminsky (https://www.cs.cmu.edu/~dga/papers/cuckoo-conext2014.pdf).
package cuckoofilter
