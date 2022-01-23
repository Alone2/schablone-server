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
echo "MARIADB_ROOT_PASSWORD: $password
MARIADB_USER_PASSWORD: $password
MARIADB_USER_NAME: root
MARIADB_HOST=mariadb" > .mariadb.env

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
./scripts/schablone-server-client/import_modules.py localhost:8080 ${API_TOKEN} quicktext_module.xml
./scripts/schablone-server-client/import_templates.py localhost:8080 ${API_TOKEN} quicktext_template.xml
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
Setup virtual environement and install dependencies:
```
cd scripts
python -m venv schablone-python
source schablone-python/bin/activate
pip install -r requirements.txt
```

Update python client api:
```
cd scripts
source schablone-python/bin/activate
openapi-python-client update --path ../schablone-api.yaml
```