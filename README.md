# Avialeta

## Development

Create port tunnel to work with API from local machine.

*Linux*

```
$ ssh dev@petruchik.com -L 8080:localhost:8080
```

## API

### Search Location

Search location by name.

*Request*

```
GET /locations?search=<name>
```

```
$ curl http://localhost:8080/locations/?search=<name>
```

*Response*

```json
[
    {
        "Code": "COD",
        "Name": "Match Location Name",
        "CityCode": ""
    },
    {
        "Code": "COD",
        "Name": "Match Location Name",
        "CityCode": ""
    },

    ...

    {
        "Code": "COD",
        "Name": "Match Location Name",
        "CityCode": ""
    }
]
```

### Search Flights
