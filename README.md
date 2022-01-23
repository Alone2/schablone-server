# schablone-server

## Documentation API

The documentation of the API can be accessed right [here](https://alone2.github.io/schablone-server/).

## Installation Using Docker and Docker Compose
Prerequisites: You need to have Docker and Docker Compose installed.

```
# Clone Repo
git clone git@github.com:Alone2/schablone-server.git
cd schablone-server

# Set mysql root password
read password
echo "MARIADB_ROOT_PASSWORD: $password" > .mariadb.env

# Start server
docker-compose up -d
```

## Debugging

### Connect to Database (Docker)
```
docker-compose exec mariadb sh -c 'mysql -h 0.0.0.0 -u root -p '
```

## Populate Database With Quicktext Data
Prerequisites: You need to have python 3 installed.

```
./scripts/import_modules.py quicktext_module.xml
./scripts/import_templates.py quicktext_template.xml
```

## Development
### Generate OpenApi code
```
~/go/bin/oapi-codegen --generate types --package=api schablone-api.yaml > api/types.gen.go
~/go/bin/oapi-codegen --generate chi-server --package=api schablone-api.yaml > api/chi-server.gen.go
~/go/bin/oapi-codegen --generate spec --package=api schablone-api.yaml > api/spec.gen.go
```

### Install Go Dependencies (Local Install)
```
go get .
go test
```

### Python scripts
Install dependencies:
```
pip install openapi-python-client
```

Genereate python api to access server:
```
cd scripts
openapi-python-client update --path ../schablone-api.yaml
```