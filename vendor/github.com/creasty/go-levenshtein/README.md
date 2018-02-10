go-levenshtein
==============

[![Build Status](https://travis-ci.org/creasty/go-levenshtein.svg?branch=master)](https://travis-ci.org/creasty/go-levenshtein)

Levenshtein algorithm in Golang (with C ext)


Usage
-----

```go
levenshtein.Distance("kitten", "sitting")
=> 3

// Multibyte string support
levenshtein.Distance("あいうえお", "aいうえo")
=> 2
```

```go
levenshtein.LcsLen("kitten", "sitting")
=> 4

// Multibyte string support
levenshtein.LcsLen("あいうえお", "aいうえo")
=> 3
```


License
-------

This project is copyright by [Creasty](http://creasty.com), released under the MIT license.  
See `LICENSE.txt` file for details.
