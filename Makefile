compile:
	go install github.com/bitquery/protobuf-sql/cmd/protoc-gen-sql@test/test && \
	rm -rf ./examples/sql && \
	mkdir -p ./examples/sql && \
	protoc \
	-I=./examples/proto \
	--proto_path=./examples/proto \
	--sql_out=./examples/sql \
	--sql_opt=paths=source_relative \
	--sql_opt=template_path=examples/templates/create_table.sql \
	--sql_opt=message_suffix=Message \
	--sql_opt=ignore_fields='Indexing_SegmentHash;Indexing_LogicalTime;Indexing_Time;Indexing_OnTree;Indexing_OnTrunk' \
	tables.proto
