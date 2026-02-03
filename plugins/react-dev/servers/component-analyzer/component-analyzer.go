package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ComponentAnalyzer struct {
	projectRoot string
	mcpServer   *server.MCPServer
}

func NewComponentAnalyzer(projectRoot string) *ComponentAnalyzer {
	if projectRoot == "" {
		projectRoot, _ = os.Getwd()
	}

	s := &ComponentAnalyzer{
		projectRoot: projectRoot,
	}

	mcpServer := server.NewMCPServer(
		"component-analyzer",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	s.mcpServer = mcpServer
	s.setupHandlers()

	return s
}

func (s *ComponentAnalyzer) setupHandlers() {
	// Register tools
	s.mcpServer.AddTool(
		mcp.NewTool("analyze_component_tree",
			mcp.WithDescription("Analyze component hierarchy and nesting depth"),
			mcp.WithString("componentPath",
				mcp.Required(),
				mcp.Description("Path to the component file to analyze"),
			),
		),
		s.analyzeComponentTree,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("detect_prop_drilling",
			mcp.WithDescription("Detect prop drilling patterns in component tree (depth threshold: 3)"),
			mcp.WithString("featurePath",
				mcp.Required(),
				mcp.Description("Path to feature directory (e.g., 'src/features/user')"),
			),
		),
		s.detectPropDrilling,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("check_hook_dependencies",
			mcp.WithDescription("Analyze hook dependency arrays for issues"),
			mcp.WithString("filePath",
				mcp.Required(),
				mcp.Description("Path to file with hooks to analyze"),
			),
		),
		s.checkHookDependencies,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("analyze_component_complexity",
			mcp.WithDescription("Analyze component complexity (lines, branches, hooks count)"),
			mcp.WithString("componentPath",
				mcp.Required(),
				mcp.Description("Path to component to analyze"),
			),
		),
		s.analyzeComponentComplexity,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("find_unused_props",
			mcp.WithDescription("Find props that are defined but not used in component"),
			mcp.WithString("componentPath",
				mcp.Required(),
				mcp.Description("Path to component file"),
			),
		),
		s.findUnusedProps,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("check_accessibility",
			mcp.WithDescription("Check component for basic accessibility issues"),
			mcp.WithString("componentPath",
				mcp.Required(),
				mcp.Description("Path to component file"),
			),
		),
		s.checkAccessibility,
	)
}

func (s *ComponentAnalyzer) analyzeComponentTree(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	componentPath, err := request.RequireString("componentPath")
	if err != nil {
		return mcp.NewToolResultError("componentPath parameter is required"), nil
	}

	fullPath := filepath.Join(s.projectRoot, componentPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error reading file: %v", err)), nil
	}

	analysis := analyzeTreeStructure(string(content))

	message := fmt.Sprintf(`## Component Tree Analysis: %s

**Nesting Depth**: %d levels
**Child Components**: %d
**Custom Hooks Used**: %d
**Conditional Rendering**: %d instances

%s`,
		filepath.Base(componentPath),
		analysis.nestingDepth,
		analysis.childComponents,
		analysis.hooksCount,
		analysis.conditionalCount,
		analysis.recommendations,
	)

	return mcp.NewToolResultText(message), nil
}

func (s *ComponentAnalyzer) detectPropDrilling(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	featurePath, err := request.RequireString("featurePath")
	if err != nil {
		return mcp.NewToolResultError("featurePath parameter is required"), nil
	}

	// Use default depth of 3 for now
	minDepth := 3

	fullPath := filepath.Join(s.projectRoot, featurePath)
	issues := detectPropDrillingIssues(fullPath, minDepth)

	if len(issues) == 0 {
		return mcp.NewToolResultText(fmt.Sprintf("✅ No prop drilling issues found (depth >= %d)", minDepth)), nil
	}

	message := fmt.Sprintf("❌ Prop drilling detected in %d location(s):\n\n", len(issues))
	for _, issue := range issues {
		message += fmt.Sprintf("- %s\n", issue)
	}
	message += "\n💡 Consider using Context API or state management library"

	return mcp.NewToolResultText(message), nil
}

func (s *ComponentAnalyzer) checkHookDependencies(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	filePath, err := request.RequireString("filePath")
	if err != nil {
		return mcp.NewToolResultError("filePath parameter is required"), nil
	}

	fullPath := filepath.Join(s.projectRoot, filePath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error reading file: %v", err)), nil
	}

	issues := analyzeHookDependencies(string(content))

	if len(issues) == 0 {
		return mcp.NewToolResultText("✅ No hook dependency issues found"), nil
	}

	message := fmt.Sprintf("⚠️  Hook dependency issues found:\n\n")
	for _, issue := range issues {
		message += fmt.Sprintf("- %s\n", issue)
	}

	return mcp.NewToolResultText(message), nil
}

func (s *ComponentAnalyzer) analyzeComponentComplexity(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	componentPath, err := request.RequireString("componentPath")
	if err != nil {
		return mcp.NewToolResultError("componentPath parameter is required"), nil
	}

	fullPath := filepath.Join(s.projectRoot, componentPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error reading file: %v", err)), nil
	}

	complexity := calculateComplexity(string(content))

	status := "✅"
	if complexity.score > 20 {
		status = "⚠️"
	}
	if complexity.score > 40 {
		status = "❌"
	}

	message := fmt.Sprintf(`%s Component Complexity Analysis

**Complexity Score**: %d/100
**Lines of Code**: %d
**Hooks Used**: %d
**Conditional Branches**: %d
**State Variables**: %d

%s`,
		status,
		complexity.score,
		complexity.lines,
		complexity.hooks,
		complexity.branches,
		complexity.stateVars,
		complexity.recommendation,
	)

	return mcp.NewToolResultText(message), nil
}

func (s *ComponentAnalyzer) findUnusedProps(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	componentPath, err := request.RequireString("componentPath")
	if err != nil {
		return mcp.NewToolResultError("componentPath parameter is required"), nil
	}

	fullPath := filepath.Join(s.projectRoot, componentPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error reading file: %v", err)), nil
	}

	unusedProps := findUnusedPropsList(string(content))

	if len(unusedProps) == 0 {
		return mcp.NewToolResultText("✅ No unused props found"), nil
	}

	message := fmt.Sprintf("❌ Found %d unused prop(s):\n\n", len(unusedProps))
	for _, prop := range unusedProps {
		message += fmt.Sprintf("- %s\n", prop)
	}

	return mcp.NewToolResultText(message), nil
}

func (s *ComponentAnalyzer) checkAccessibility(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	componentPath, err := request.RequireString("componentPath")
	if err != nil {
		return mcp.NewToolResultError("componentPath parameter is required"), nil
	}

	fullPath := filepath.Join(s.projectRoot, componentPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error reading file: %v", err)), nil
	}

	issues := checkA11yIssues(string(content))

	if len(issues) == 0 {
		return mcp.NewToolResultText("✅ No obvious accessibility issues found"), nil
	}

	message := fmt.Sprintf("⚠️  Accessibility issues found:\n\n")
	for _, issue := range issues {
		message += fmt.Sprintf("- %s\n", issue)
	}
	message += "\n💡 Run full accessibility tests with jest-axe for comprehensive validation"

	return mcp.NewToolResultText(message), nil
}

// Analysis helper functions

type treeAnalysis struct {
	nestingDepth      int
	childComponents   int
	hooksCount        int
	conditionalCount  int
	recommendations   string
}

func analyzeTreeStructure(content string) treeAnalysis {
	analysis := treeAnalysis{}

	// Count hooks
	hookPattern := regexp.MustCompile(`use[A-Z]\w+\(`)
	analysis.hooksCount = len(hookPattern.FindAllString(content, -1))

	// Count child components (JSX tags starting with capital letter)
	componentPattern := regexp.MustCompile(`<[A-Z]\w+`)
	analysis.childComponents = len(componentPattern.FindAllString(content, -1))

	// Count conditional rendering
	conditionalPattern := regexp.MustCompile(`\{.*\?.*:|\{.*&&`)
	analysis.conditionalCount = len(conditionalPattern.FindAllString(content, -1))

	// Estimate nesting depth by counting nested divs/elements
	maxNesting := 0
	currentNesting := 0
	for _, char := range content {
		if char == '<' {
			currentNesting++
			if currentNesting > maxNesting {
				maxNesting = currentNesting
			}
		} else if char == '/' && currentNesting > 0 {
			currentNesting--
		}
	}
	analysis.nestingDepth = maxNesting / 10 // Rough estimate

	// Recommendations
	if analysis.hooksCount > 8 {
		analysis.recommendations = "⚠️  High hook usage - consider extracting logic to custom hooks"
	} else if analysis.childComponents > 15 {
		analysis.recommendations = "⚠️  High component count - consider breaking down into smaller components"
	} else if analysis.nestingDepth > 6 {
		analysis.recommendations = "⚠️  Deep nesting detected - consider flattening component structure"
	} else {
		analysis.recommendations = "✅ Component structure looks good"
	}

	return analysis
}

func detectPropDrillingIssues(featurePath string, minDepth int) []string {
	issues := []string{}

	// This is a simplified version - in production, you'd walk the file tree
	// and analyze prop passing patterns

	filepath.Walk(featurePath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".tsx") || strings.HasSuffix(path, ".jsx") {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil
			}

			// Simple heuristic: look for prop destructuring patterns
			propPattern := regexp.MustCompile(`\{([^}]+)\}.*=.*props`)
			matches := propPattern.FindAllStringSubmatch(string(content), -1)

			if len(matches) > 0 {
				for _, match := range matches {
					props := strings.Split(match[1], ",")
					if len(props) >= minDepth {
						relPath := strings.TrimPrefix(path, featurePath+"/")
						issues = append(issues, fmt.Sprintf("%s: %d props passed through", relPath, len(props)))
					}
				}
			}
		}

		return nil
	})

	return issues
}

func analyzeHookDependencies(content string) []string {
	issues := []string{}

	// Find useEffect/useCallback/useMemo hooks
	hookPattern := regexp.MustCompile(`use(Effect|Callback|Memo)\s*\([^,]+,\s*\[([^\]]*)\]`)
	matches := hookPattern.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		hookType := match[1]
		deps := strings.TrimSpace(match[2])

		// Check for empty dependency array with useEffect
		if hookType == "Effect" && deps == "" {
			issues = append(issues, "useEffect with empty deps [] - ensure this is intentional (runs once)")
		}

		// Check for missing deps
		if deps == "" && hookType != "Effect" {
			issues = append(issues, fmt.Sprintf("use%s with empty deps [] - may cause stale closures", hookType))
		}
	}

	// Check for useEffect without deps array
	noDepPattern := regexp.MustCompile(`useEffect\s*\([^,]+\s*\)`)
	if noDepPattern.MatchString(content) {
		issues = append(issues, "useEffect without dependency array - runs on every render")
	}

	return issues
}

type complexityResult struct {
	score          int
	lines          int
	hooks          int
	branches       int
	stateVars      int
	recommendation string
}

func calculateComplexity(content string) complexityResult {
	result := complexityResult{}

	lines := strings.Split(content, "\n")
	result.lines = len(lines)

	// Count hooks
	hookPattern := regexp.MustCompile(`use[A-Z]\w+\(`)
	result.hooks = len(hookPattern.FindAllString(content, -1))

	// Count state variables
	statePattern := regexp.MustCompile(`useState\(`)
	result.stateVars = len(statePattern.FindAllString(content, -1))

	// Count branches (if/else/ternary/switch)
	branchPattern := regexp.MustCompile(`\b(if|else|switch|\?)\b`)
	result.branches = len(branchPattern.FindAllString(content, -1))

	// Calculate score
	result.score = (result.lines / 10) + (result.hooks * 2) + result.branches + (result.stateVars * 3)

	// Recommendation
	if result.score > 40 {
		result.recommendation = "❌ High complexity - strongly consider refactoring into smaller components"
	} else if result.score > 20 {
		result.recommendation = "⚠️  Moderate complexity - consider splitting responsibilities"
	} else {
		result.recommendation = "✅ Complexity is manageable"
	}

	return result
}

func findUnusedPropsList(content string) []string {
	unused := []string{}

	// Extract prop names from interface/type definition
	propsPattern := regexp.MustCompile(`interface\s+\w+Props\s*\{([^}]+)\}`)
	matches := propsPattern.FindStringSubmatch(content)

	if len(matches) < 2 {
		return unused
	}

	propsBlock := matches[1]
	propLinePattern := regexp.MustCompile(`(\w+)[\?]?:\s*`)
	propMatches := propLinePattern.FindAllStringSubmatch(propsBlock, -1)

	for _, match := range propMatches {
		propName := match[1]

		// Check if prop is used in component body (simple check)
		propUsagePattern := regexp.MustCompile(fmt.Sprintf(`\b%s\b`, propName))
		usageCount := len(propUsagePattern.FindAllString(content, -1))

		// If only appears once (in definition), it's unused
		if usageCount <= 1 {
			unused = append(unused, propName)
		}
	}

	return unused
}

func checkA11yIssues(content string) []string {
	issues := []string{}

	// Check for img without alt
	if strings.Contains(content, "<img") && !regexp.MustCompile(`<img[^>]+alt=`).MatchString(content) {
		issues = append(issues, "Image(s) missing alt attribute")
	}

	// Check for button without aria-label or text
	buttonPattern := regexp.MustCompile(`<button[^>]*>`)
	buttons := buttonPattern.FindAllString(content, -1)
	for _, button := range buttons {
		if !strings.Contains(button, "aria-label") && !strings.Contains(button, ">") {
			issues = append(issues, "Button without aria-label or text content")
		}
	}

	// Check for input without label or aria-label
	inputPattern := regexp.MustCompile(`<input[^>]*>`)
	inputs := inputPattern.FindAllString(content, -1)
	for _, input := range inputs {
		if !strings.Contains(input, "aria-label") && !strings.Contains(input, "id=") {
			issues = append(issues, "Input field without aria-label or associated label")
		}
	}

	// Check for onClick on non-interactive elements
	if regexp.MustCompile(`<div[^>]*onClick`).MatchString(content) {
		issues = append(issues, "onClick on div - consider using button for keyboard accessibility")
	}

	return issues
}

func main() {
	projectRoot := ""
	if len(os.Args) > 1 {
		projectRoot = os.Args[1]
	}

	s := NewComponentAnalyzer(projectRoot)

	fmt.Fprintln(os.Stderr, "Component Analyzer MCP server running")

	if err := server.ServeStdio(s.mcpServer); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
