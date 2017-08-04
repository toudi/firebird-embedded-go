These are my experiments with golang (go) and firebird SQL.

The reason behind this is that SQLite (which is de-facto standard of server-less db) doesn't support full ALTER TABLE
as well as it doesn't check variable types. As a result, the code can sometimes work on dev db but fail to execute on production.

Firebird supports embedded server, which technically means that you can achieve the same result as with SQLite3 - which is a
SQL database in a single file. Of course, it's not as lightweight as SQLite, but on the other hand, you don't require a fully
fledged SQL server, so this would make it ideal for development purposes, where performance is not important at all, but
SQL / ACID is.

this is my first project with CGO

I have no idea whether I can include libfbembed.so file in the repo, so in case you're wondering - you have to download firebird 
(I used version 2.5.x), then look into provided 'buildroot.tar.gz' in which you will find '/opt/firebird/lib/libfbembed.so'

 