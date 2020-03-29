# expoliring-go-concurrency

Concurrently incrementing an integer by one a thousand times using different techinques of syncronisation

* base - case with race condition 
* atomic - same as previous but with atomic addition
* lock - using mutex explicitly to protect data
* channels - using channel to protect state implicitly