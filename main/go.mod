module main

go 1.16

replace example.com/hello_prev => ../hello-prev

require example.com/hello_prev v0.0.0-00010101000000-000000000000
