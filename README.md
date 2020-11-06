# GetGround-technical-task
This API gives the user the ability to plan a party!  
Guests cannot be added to the guest list until a venue has been added, as this is used
to calculate if the guests and their entourage can fit on a table and in the venue
. You can change your mind later and delete and add a new venue. Doing this will
 recalculate the used capacity of the venue, but will not check that the guests will
  still fit on their specified tables if the table size has changed, so it is
   recommended that you keep the table size the same.
   
To add a guest, the user specifies a preferred table number, but if this table does
 not have the capacity to accommodate them, the system will try to find a free table
  to put them on. If there is no table they can fit on, an error will be returned. Try
   splitting them into two separate guests.
##API Documentation
### Add a guest to the guestlist

If there is insufficient space at the specified table, then it will attempt to find a
 table which can accommodate the guest and accompanying guests. If no table is
  available, or the venue  is full, or there is no venue an error will be thrown.

```
POST /guest_list/name
body: 
{
    "table": int,
    "accompanying_guests": int
}
```
Example response for successfully adding a guest to the list:
```
status: 202
response: 
{
    "Id":78,
    "Name":"ben231","
    accompanying_guests":5,
    "TableNumber":2,
    "Arrived":false,
    "ArrivalTime":""
}
```
Example error response:
```
status: 400
response: 
{
    "no table avaliable for number of guests"
}
```

### Remove a guest from the guest list
If the user does not exist, an error will be thrown
```
DELETE /guest_list/name
```
Successful response:
```
status: 202
response: 
{
    "Guest removed"
}
```
Example error response:
```
status: 400
response: 
{
    "error: no guest by that name"
}
```

### Get the guest list
Will return an empty array if no guests on the list. 
```
GET /guest_list
```
Example response:
```
status: 200
response: 
{
    "Guests":[
        {
            "Id":75,
            "Name":"ben",
            "accompanying_guests":8,
            "TableNumber":8,
            "Arrived":false,
            "ArrivalTime":""},
        }, ...
    ]
}
```
Example error response:
```
status: 500
response: 
{
    "error: could not get guest list"
}
```
### Generate an invitation

Output an invitation in Markdown or HTML that contains the guest's name and allocated table.

```
GET /invitation/name
```
Successful response:
```
status: 200
response: File download
```

Example error response:
```
status: 500
response:
{
    "no guest by that name"
}
```

### Guest Arrives

A guest may arrive with an entourage that is not the size indicated at the guest list.
If the table is expected to have space for the extras, allow them to come.
Otherwise, this method should throw an error.

```
PUT /guests/name
body:
{
    "accompanying_guests": int
}
```
Example success response:
```
status: 202
response:
{
    "Id":75,
    "Name":"ben",
    "accompanying_guests":8,
    "TableNumber":8,
    "Arrived":true
    "ArrivalTime":"2020-11-06 12:35:14"
}
```
Example error response:
```
status: 400
response:
{
    "error: no guest by that name"
}
```

### Guest Leaves

When a guest leaves, all their accompanying guests leave as well.

```
DELETE /guests/name
```
Success response:
```
status 202
response:
{
    "Guest left the party"
}
```
Example error response:
```
status:400
response:
{
    "error: could not get guest by that name"
}
```

### Get arrived guests
returns a list of arrived guests. If there are no guests an empty list is returned.
```
GET /guests
```
Example response:
```
status: 200
response: 
{
    "guests": [
        {
            "Id":75,
            "Name":"ben",
            "accompanying_guests":8,
            "TableNumber":8,
            "Arrived":true,
            "ArrivalTime":"2020-11-06 12:35:14"
        }
    ]
}
```

### Count number of empty seats

```
GET /seats_empty
```
Example response:
```
status: 200
response:
{
    "seats_empty": 150
}
```

### Count number of remaining bookable seats
The number of seats remaining in the venue that can be used to invite guests into.
```
GET /seats_bookable
```
Example response:
```
response:
{
    "seats_empty": 5
}
```

### Add the venue

```
POST /venue
response:
{
    "name": "Hilton London Bankside",
    "numberoftables": 21,
    "capacity": 210
}
```
Example response:
```
status: 200
response:
{
    "Id": 1
    "Name":"Hilton London Bankside",
    "Capacity":210,
    "NumberOfTables":21,
    "TableSize":10,
    "NextFreeTable":1,
    "UsedCapacity":12
}
```
Example error response:
```
status: 500
response:
{
    "invalid venue"
}
```

### Get the venue
Will only return the first venue in the database. Teh field "NextFreeTable" is not
 currently implemented and is reserved for future versions.
```
GET /venue
```
Example response:
```
status: 200
response:
{
    "Id":30,
    "Name":"Hilton London Bankside",
    "Capacity":210,
    "NumberOfTables":21,
    "TableSize":10,
    "NextFreeTable":1,
    "UsedCapacity":12
}
```
Example error response:
```
status: 500
response:
{
    "error: could not get venue"
}
```

### Remove the venue

```
DELETE /venue
```
Example response:
```
status: 200
```