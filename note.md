## section 9 
rules of data evolution

undating protocol rules
: 
1. don't change the numeric tags for any existing fields
2. you can add new fileds, and old code will just ignore them.

3. Likewise if the old/ new code read unknow data, the defalult will take place

4. OBSOLETE 

5. data type changes (don't change data type at all..!)


## section 10 

- 64 bits allow for a greater range
- 32 bit : 
    - int32/sint32 : -2,147,483,648 to 2,147,483,647
    - uint32:0 to 4,294,967,295

- 64 bit : 
    - int64/sint64 : -9,223,372,063 to 9,223,372,036,854,775,807
    - uint64 L 0 to 18,446,744,073,709,551,615

- unit32, uint64 not allow negative values
- int32, int64 do not encode negative values efficiently(negative values constantly use 10 bytes in space)

- sint32, sint64 encode negative values well 
(ZigZag)

- uint32, uint64, int32, int64, sint32, sint64 are variable encoding meaning that if they can use less space, they will.

- fixed32 use 4 bytes constantly.
- fixed64 use 8 bytes constantly.