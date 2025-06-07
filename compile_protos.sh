#!/bin/bash

PROTO_DIR="./protos"

OUT_BASE="./go/gen"

mkdir -p "$OUT_BASE"

find "$PROTO_DIR" -name "*.proto" | while read -r proto_file; do
  file_name=$(basename "$proto_file" .proto)

  OUT_DIR="$OUT_BASE/$file_name"
  mkdir -p "$OUT_DIR"

  echo "Compiling $proto_file -> $OUT_DIR"

  protoc \
    --proto_path="$PROTO_DIR" \
    --go_out=paths=source_relative:"$OUT_DIR" \
    --go-grpc_out=paths=source_relative:"$OUT_DIR" \
    "$proto_file"
done

echo "✅ Compilation finished."
