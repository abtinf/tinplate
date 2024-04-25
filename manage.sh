#!/bin/bash

Help() {
	echo "Go Service Template - Manage script"
	echo "Usage: manage.sh [options]"
	echo "Options:"
	echo "  -i  Init"
	echo "  -o  Obliterate"
	echo "  -u  Up"
	echo "  -d  Down"
	echo "  -h  Help"
	echo ""
	echo "  -a  Generate all"
	echo "  -p  Generate Protobufs"
	echo "  -s  Generate SQL"
}

Init() {
	echo "Init"
}

Obliterate() {
	echo "Obliterate"
}

Up() {
	echo "Up"
}

Down() {
	echo "Down"
}

Genall() {
	echo "Generating all"
	Protogen
	Sqlgen
}

Protogen() {
	echo "Generating Protobuf"
	protoc -I ./proto \
		--go_out=./proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./proto \
		--grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=./proto \
		--grpc-gateway-ts_out=./proto \
		--grpc-gateway-ts_opt=paths=source_relative \
		./proto/api.proto
	mv ./proto/api.swagger.json ./static/http/swagger-ui/swagger.json
}

Sqlgen() {
	echo "Generating SQL"
	pushd db
	cat `ls sql/migrations/*.sql` > gen/schema.sql
	cat `ls sql/queries/*.sql` > gen/queries.sql
	sqlc generate
	popd
}


while getopts ":ioudhasp" option; do
	case $option in
		i)
			Init
			exit;;
		o)
			Obliterate
			exit;;
		u)
			Up
			exit;;
		d)
			Down
			exit;;
		h)
			Help
			exit;;
		a)
			Genall
			exit;;
		p)
			Protogen
			exit;;
		s)
			Sqlgen
			exit;;
		\?)
			Help
			exit;;
	esac
done
Help
