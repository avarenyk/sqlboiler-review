### Background

As an DEVELOPER I want to write less code and boost my productivity  
I don't want to create models for DB entities  
I want to have an ability to load multiple relationships

### Requirements
I had my own requirements on ORM

1. Models generation
2. Exact representation of DB models
3. Load related entities
4. Load related entities with conditions
5. Eager loading. N+1 problem
6. Save entity with related ones in one action
7. Update only selected columns
8. Complex queries or/and where conditions
9. Locks: select for update. select for read

### Issues
1. Could not parse MySQL enums. todo: create a ticket and try fix it by myself
2. Methods Set() and Add() are performing sql query immediate, but I would like to call it when everything set. Hacked it with Executor() that do nothing.
3. Does not have "lock in share mode" for mysql 5.7


### Conclusion

Overall very good ORM. 
To learn it please read docs:
https://github.com/volatiletech/sqlboiler
and watch these videos:
https://www.youtube.com/watch?v=y5utRS9axfg
https://www.youtube.com/watch?v=iiJuM9NR8No
https://www.youtube.com/watch?v=-B-OPsYRZJA

This repo provides a good DB to play with and explore more about SqlBoiler.

