#!/bin/bash

# Validate Domain Layer Purity
# This script checks that domain layer files only import from stdlib and shared package

set -e

FILE_PATH="$1"

# Check if file is in domain layer
if [[ ! "$FILE_PATH" =~ internal/.*/domain/.*\.go$ ]]; then
    exit 0  # Not a domain file, skip
fi

echo "🔍 Validating domain layer purity: $FILE_PATH"

# List of allowed import prefixes (Go stdlib only)
STDLIB_IMPORTS=(
    "context"
    "encoding"
    "errors"
    "fmt"
    "io"
    "log"
    "math"
    "net"
    "os"
    "path"
    "reflect"
    "regexp"
    "runtime"
    "sort"
    "strconv"
    "strings"
    "sync"
    "syscall"
    "testing"
    "time"
    "unicode"
)

# Extract imports from the file
IMPORTS=$(grep -E '^\s*"[^"]+"' "$FILE_PATH" | sed 's/^\s*"\(.*\)".*/\1/' || true)

# Check each import
VIOLATIONS=0
while IFS= read -r import; do
    if [ -z "$import" ]; then
        continue
    fi

    # Allow internal/shared imports
    if [[ "$import" =~ /internal/shared/ ]]; then
        continue
    fi

    # Check if it's a stdlib import
    IS_STDLIB=0
    for prefix in "${STDLIB_IMPORTS[@]}"; do
        if [[ "$import" == "$prefix" ]] || [[ "$import" == "$prefix"/* ]]; then
            IS_STDLIB=1
            break
        fi
    fi

    # If not stdlib and not shared, it's a violation
    if [ $IS_STDLIB -eq 0 ]; then
        echo "❌ VIOLATION: Domain layer importing non-stdlib package: $import"
        VIOLATIONS=$((VIOLATIONS + 1))
    fi
done <<< "$IMPORTS"

if [ $VIOLATIONS -gt 0 ]; then
    echo ""
    echo "⚠️  Domain layer should only import:"
    echo "   - Go standard library"
    echo "   - internal/shared/* (shared kernel)"
    echo ""
    echo "To fix: Move infrastructure concerns to infrastructure layer"
    exit 1
fi

echo "✅ Domain layer purity validated"
exit 0
