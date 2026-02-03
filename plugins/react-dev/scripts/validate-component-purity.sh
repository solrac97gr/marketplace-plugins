#!/bin/bash

# Validate React Component Purity
# This script checks React components follow best practices and architectural rules

set -e

FILE_PATH="$1"

# Check if file is a component or hook
if [[ ! "$FILE_PATH" =~ \.(tsx|ts|jsx|js)$ ]]; then
    exit 0  # Not a TypeScript/JavaScript file, skip
fi

# Skip test files
if [[ "$FILE_PATH" =~ \.(test|spec)\.(tsx|ts|jsx|js)$ ]]; then
    exit 0
fi

echo "🔍 Validating component purity: $FILE_PATH"

VIOLATIONS=0

# Read file content
CONTENT=$(cat "$FILE_PATH")

# Check 1: No direct fetch/axios calls in component files (should be in hooks or services)
if [[ "$FILE_PATH" =~ src/.*components/.* ]]; then
    if echo "$CONTENT" | grep -qE "(fetch\(|axios\.|\.get\(|\.post\(|\.put\(|\.delete\()" && \
       ! echo "$CONTENT" | grep -qE "// @allow-api-call"; then
        echo "❌ VIOLATION: Direct API calls found in component"
        echo "   Move API calls to custom hooks (src/features/*/hooks/) or services"
        VIOLATIONS=$((VIOLATIONS + 1))
    fi
fi

# Check 2: Hooks must follow naming convention (use*)
if [[ "$FILE_PATH" =~ src/.*/hooks/.*\.(tsx|ts)$ ]]; then
    HOOK_FUNCTIONS=$(echo "$CONTENT" | grep -oE "export (function|const) [a-zA-Z]+" | awk '{print $3}')
    while IFS= read -r func_name; do
        if [ -z "$func_name" ]; then
            continue
        fi
        if [[ ! "$func_name" =~ ^use[A-Z] ]] && [[ ! "$func_name" =~ ^(create|get|make) ]]; then
            echo "❌ VIOLATION: Hook function '$func_name' must start with 'use' (e.g., useMyHook)"
            VIOLATIONS=$((VIOLATIONS + 1))
        fi
    done <<< "$HOOK_FUNCTIONS"
fi

# Check 3: Components should export types/interfaces for props
if [[ "$FILE_PATH" =~ src/.*components/.*\.tsx$ ]]; then
    COMPONENT_NAME=$(basename "$FILE_PATH" .tsx)

    # Check if component exports props interface
    if ! echo "$CONTENT" | grep -qE "(interface|type) ${COMPONENT_NAME}Props"; then
        # Check if it's a simple component without props
        if ! echo "$CONTENT" | grep -qE "export (function|const) ${COMPONENT_NAME}\s*\(\s*\)"; then
            echo "⚠️  WARNING: Component should export '${COMPONENT_NAME}Props' interface/type"
            # Don't increment violations for this - it's a warning
        fi
    fi
fi

# Check 4: No console.log in production code (allow in development utilities)
if [[ "$FILE_PATH" =~ src/(features|components)/.* ]]; then
    if echo "$CONTENT" | grep -qE "console\.(log|debug|info)" && \
       ! echo "$CONTENT" | grep -qE "// @dev-only|process\.env\.NODE_ENV"; then
        echo "⚠️  WARNING: Found console.log - consider using proper logging or remove before production"
        # Don't increment violations - it's a warning
    fi
fi

# Check 5: Ensure React is imported (or using new JSX transform)
if [[ "$FILE_PATH" =~ \.tsx$ ]]; then
    if echo "$CONTENT" | grep -qE "(<[A-Z]|<>)" && \
       ! echo "$CONTENT" | grep -qE "import.*React" && \
       ! echo "$CONTENT" | grep -qE "\/\/ @jsx-runtime"; then
        # This is OK with React 17+ new JSX transform, so just a note
        : # No warning needed
    fi
fi

# Check 6: No business logic in presentational components
if [[ "$FILE_PATH" =~ src/.*/components/.*\.tsx$ ]]; then
    # Check for complex business logic patterns
    if echo "$CONTENT" | grep -qE "(useEffect.*fetch|useEffect.*axios|localStorage\.setItem.*useEffect)"; then
        echo "⚠️  WARNING: Presentational component may contain business logic"
        echo "   Consider moving data fetching to custom hooks or container components"
        # Don't increment violations - it's a warning
    fi
fi

# Check 7: Validate proper TypeScript usage
if [[ "$FILE_PATH" =~ \.tsx?$ ]]; then
    # Check for 'any' type usage
    ANY_COUNT=$(echo "$CONTENT" | grep -oE ": any[ ,;\)]" | wc -l | tr -d ' ')
    if [ "$ANY_COUNT" -gt 0 ]; then
        echo "⚠️  WARNING: Found $ANY_COUNT usage(s) of 'any' type - consider using specific types"
        # Don't increment violations - it's a warning
    fi
fi

# Check 8: Ensure accessibility attributes in interactive elements
if [[ "$FILE_PATH" =~ src/.*components/.*\.tsx$ ]]; then
    if echo "$CONTENT" | grep -qE "<button" && \
       ! echo "$CONTENT" | grep -qE "(aria-label|aria-describedby)"; then
        # Only warn for buttons without obvious text content
        if ! echo "$CONTENT" | grep -qE "<button[^>]*>[^<{]"; then
            echo "⚠️  WARNING: Interactive elements should have ARIA attributes for accessibility"
        fi
    fi
fi

# Final result
if [ $VIOLATIONS -gt 0 ]; then
    echo ""
    echo "⚠️  Component validation failed with $VIOLATIONS violation(s)"
    echo ""
    echo "React Component Best Practices:"
    echo "  - Move API calls to hooks (src/features/*/hooks/) or services"
    echo "  - Hook functions must start with 'use' (useMyHook)"
    echo "  - Export Props interfaces for components"
    echo "  - Keep presentational components pure (no business logic)"
    echo "  - Use TypeScript strict types (avoid 'any')"
    echo "  - Add ARIA attributes for accessibility"
    echo ""
    exit 1
fi

echo "✅ Component purity validated"
exit 0
