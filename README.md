# schablone-server

* [Documentation API](#documentation)  
* [Installation Using Docker](#docker)  
* [Usage](#usage)  
* [Debugging](#debugging)  
    * [Access Database](#debugging_db)  
    * [Debugging](#debugging_populate)  
    * [DB Schema](#schema)  
* [Developemet](#debugging)  
    * [Go Dependencies](#code_generation)  
    * [Go Dependencies](#go_dependencies)  
    * [Python Scripts](#python)  

## Documentation API
<a name="documentation"/>

The documentation of the API can be accessed right [here](https://alone2.github.io/schablone-server/).

Screenshot:
![text](/docs/screenshot.png)

## Installation Using Docker
<a name="docker"/>
Prerequisites: You need to have Docker and Docker Compose installed.

```
# Clone Repo
git clone git@github.com:Alone2/schablone-server.git
cd schablone-server

# Set mysql root password
read password
echo "MARIADB_ROOT_PASSWORD=$password
MARIADB_USER_PASSWORD=$password
MARIADB_USER_NAME=root
MARIADB_HOST=mariadb" > .mariadb.env

# Start server
docker-compose up -d
```

## Usage
<a name="usage"/>

Sign in with the default password:
```
curl "http://0.0.0.0:8080/user/login?username=admin&password=12345"
```

## Debugging
<a name="debugging"/>

### Connect to Database (Docker)
<a name="debugging_db"/>

```
docker-compose exec mariadb sh -c 'mysql -h 0.0.0.0 -u root -p '
```

### Populate Database With Quicktext Data
<a name="debugging_populate"/>
Prerequisites: You need to have python 3 installed and already have followed the following instructions:

* [Python Scripts](#python)

```
# GET API token
curl "http://0.0.0.0:8080/user/login?username=admin&password=12345"

# Set it
API_TOKEN="INSERT_TOKEN_HERE"

# Import scripts
./scripts/schablone-server-client/import_macros.py http://localhost:8080 ${API_TOKEN} quicktext_module.xml
./scripts/schablone-server-client/import_templates.py http://localhost:8080 ${API_TOKEN} quicktext_template.xml
```

### DB Schema
<a name="debugging_populate"/>
The database SQL CREATE file (as well as the MySQL Workbench file) is in the ./db folder:

![](/docs/db_schema.svg)

## Development
<a name="developement"/>

### Generate OpenApi Code
<a name="code_generation"/>

```
~/go/bin/oapi-codegen --generate types --package=api schablone-api.yaml > api/types.gen.go
~/go/bin/oapi-codegen --generate chi-server --package=api schablone-api.yaml > api/chi-server.gen.go
~/go/bin/oapi-codegen --generate spec --package=api schablone-api.yaml > api/spec.gen.go
```

### Install Go Dependencies (Local Install)
<a name="go_dependencies"/>

```
go get .
go test
```

### Python scripts
<a name="python"/>
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
