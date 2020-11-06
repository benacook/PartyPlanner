## Summary

Hopefully this tech task allows you to strut your stuff as much as you decide to!

We'd like to implement a guestlist service for the GetGround year end party!
We haven't decide on the venue yet so the number of tables and the capacity are subject to change.

Before the party, we should be able to add and remove guests with their entourages from our guest list and generate a nice looking invitation for the invited guests. 

When the party begins, guests will arrive with an entourage. This party may not be the size indicated on the guest list. 
However, if it is expected that the guest's table can accommodate the extra people, then the whole party should be let in. Otherwise, they will be turned away.
Guests will also leave throughout the course of the party. Note that when a guest leaves, their accompanying guests will leave with them.

At any point in the party, we should be able to know:
- Our guests at the party
- How many empty seats there are

## Sample API guide

This is a directional API guide.

### Add a guest to the guestlist

If there is insufficient space at the specified table, then an error should be thrown.

```
POST /guest_list/name
body: 
{
    "table": int,
    "accompanying_guests": int
}
response: 
{
    "name": "string"
}
```

### Remove a guest from the guest list

```
DELETE /guest_list/name
```

### Get the guest list

```
GET /guest_list
response: 
{
    "guests": [
        {
            "name": "string",
            "table": int,
            "accompanying_guests": int
        }, ...
    ]
}
```

### Generate an invitation

Output an invitation in Markdown or HTML that contains the guest's name and allocated table.

```
GET /invitation/name
response: File download
```

### Guest Arrives

A guest may arrive with an entourage that is not the size indicated at the guest list.
If the table is expected to have space for the extras, allow them to come. Otherwise, this method should throw an error.

```
PUT /guests/name
body:
{
    "accompanying_guests": int
}
response:
{
    "name": "string"
}
```

### Guest Leaves

When a guest leaves, all their accompanying guests leave as well.

```
DELETE /guests/name
```

### Get arrived guests

```
GET /guests
response: 
{
    "guests": [
        {
            "name": "string",
            "accompanying_guests": int,
            "time_arrived": "string"
        }
    ]
}
```

### Count number of empty seats

```
GET /seats_empty
response:
{
    "seats_empty": int
}
```


## Requirements

- Please use MySQL version 5.7 as a database
- We are expecting well tested and structured code
- Good documentation to get us started and understand your implementation

Have fun!
