# Implementation of a Bloom Filter

<span>
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
</span>

When a simple login process is initiated, the database is scanned through to get data (say - username). That username, if the database is to run the standard methods of search, would be found very late. That makes for some inefficiency. <br />
A solution is to query the database in such a way that if the item does exist, only then will it go on to query it fully. In any other case, it would spit out the fact that the item does not exist. Such implementations are done using **Bloom Filters**.


## Concept

A Bloom Filter is a probabilistic system that *may have false positives, but never false negatives.* It accepts a key, hashes it and checks if the hash was previously used in the system. If it was, there can be two cases:
- Either the item actually exists in the system.
- Or there is a collision. So the item did not actually exist, but the system would raise a *false positive* saying that it does.

If there is no hash, as obvious, the system never had that key. This is why false negatives are never the case here.

It is important to note that Bloom Filters are not responsible for the storage of actual data in any way.


## Implementation

The program is responsible for:
- The complete implementation of an optimized Bloom Filter.
- The analysis of the statistical data as to the increase of the false positives with increase in the size of the number of keys actually existing in the system (with constant filter size).


## Design Decisions for Optimization

- Strict typing with the use of `byte` over `bool` to implement flags. This ends up boosting performance.
- Use of Murmur3 hash function over the likes of SHA for faster computation (compromised with loose security).
