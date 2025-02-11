#!/bin/bash

# This annoying fix is due to bugs in OpenAPI that don't make some fields public.
# Here we make edits the generated openapi rust code to make those fixes through regexes.
# This needs to run after the openapi generator and before the cargo build step.
# See /rover/roverd/Makefile

# Takes as arguments:
#   1. the filename of the file to make the changes in
#   2. the entire contents of the line which should be chaned
#   3. the entire contents of the replacement line
function edit_line_in_file() {
    echo "> editing $1"
    echo "     changing line: '$2'"
    echo "     to:            '$3'"

    sed -i "s/$2/$3/" $1
}

edit_line_in_file "openapi/src/models.rs" "pub struct DuplicateServiceError(String);" "pub struct DuplicateServiceError(pub String);"

edit_line_in_file "openapi/src/models.rs" "    Box<serde_json::value::RawValue>," "    pub Box<serde_json::value::RawValue>,"

edit_line_in_file "openapi/src/models.rs" "pub struct RoverdErrorErrorValue(Box<serde_json::value::RawValue>);" "pub struct RoverdErrorErrorValue(pub Box<serde_json::value::RawValue>);"

