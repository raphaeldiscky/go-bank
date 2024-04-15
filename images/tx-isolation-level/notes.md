# Notes

MySQL -> default: repeatable read
PostgreSQL -> default: read committed

MySQL -> set the whole session isolation level before starting the transactions

PostgreSQL -> set the isolation level within the transatctions and only have effects on that 1 specific transaction

[Postgres Transaction Isolation](https://www.postgresql.org/docs/current/transaction-iso.html)

[MySQL Transaction Isolation](https://dev.mysql.com/doc/refman/8.0/en/innodb-transaction-isolation-levels.html)

PostgreSQL -> detect potential read phenomena and stop them by throwing out an error

MySQL -> use locking mechanism to achieve similar result
