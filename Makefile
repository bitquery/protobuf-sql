compile:

	go install github.com/bitquery/protobuf-sql/cmd/protoc-gen-sql && \
	rm -rf ./examples/sql && \
	mkdir -p ./examples/sql && \
	protoc \
	-I=./examples/proto \
	--proto_path=./examples/proto \
	--sql_out=./examples/sql \
	--sql_opt=paths=source_relative \
	--sql_opt=template_path=examples/templates/create_table.sql \
	--sql_opt=message_suffix=Message \
	--sql_opt=db='$${CLICKHOUSE_RT_DATABASE}' \
	tables.proto
