---
description: Perform WCAG 2.1 AA accessibility compliance review for React components
---

Perform a comprehensive WCAG 2.1 Level AA accessibility compliance review for React components and applications.

**Review Process:**

1. **Ask the user:**
   - Review scope:
     - **Entire project**: Review all components and pages
     - **Specific feature**: Review a bounded context (e.g., "authentication", "dashboard")
     - **Specific component**: Review a single component or set of components
   - Run automated tools? (yes/no) - Recommend running axe-core automated tests
   - Include code examples in report? (yes/no)
   - Target compliance level (default: WCAG 2.1 AA):
     - **Level A**: Basic accessibility
     - **Level AA**: Industry standard (recommended)
     - **Level AAA**: Enhanced accessibility

2. **Automated Testing** (if requested):
   - Check for existing jest-axe tests in component test files
   - If tests exist, run them: `npm test -- --testNamePattern="accessibility"`
   - Suggest adding jest-axe if not present:
     ```bash
     npm install --save-dev jest-axe
     ```
   - Review test results for automated violations

3. **Manual Review - WCAG 2.1 Principles:**

   **A. Perceivable** - Information must be presentable to users in ways they can perceive

   - **1.1 Text Alternatives**:
     - Images have descriptive alt text (not decorative "image" or filenames)
     - Icon buttons have aria-label or visible text
     - SVGs have title/desc elements or aria-label
     - Complex images (charts, diagrams) have detailed descriptions

   - **1.2 Time-based Media**:
     - Videos have captions/transcripts
     - Audio content has transcripts

   - **1.3 Adaptable**:
     - Semantic HTML elements (header, nav, main, article, section, footer)
     - Proper heading hierarchy (h1 -> h2 -> h3, no skipping levels)
     - Lists use ul/ol/li elements
     - Tables use proper structure (thead, tbody, th, td)
     - Form inputs have associated labels (htmlFor/id or aria-label)
     - Content order makes sense without CSS

   - **1.4 Distinguishable**:
     - Color contrast ratios:
       - Normal text (< 18pt): 4.5:1 minimum
       - Large text (≥ 18pt or 14pt bold): 3:1 minimum
       - UI components and graphics: 3:1 minimum
     - Color not used as only visual means of conveying information
     - Text can be resized up to 200% without loss of functionality
     - Images of text avoided (use actual text)
     - Focus indicators visible and sufficient contrast (3:1)

   **B. Operable** - User interface components must be operable

   - **2.1 Keyboard Accessible**:
     - All interactive elements accessible via keyboard (Tab, Enter, Space)
     - No keyboard traps (can navigate in and out)
     - Proper tabindex usage (avoid positive values)
     - Custom components handle keyboard events (Enter, Space, Arrows)
     - Skip links provided for navigation bypass

   - **2.2 Enough Time**:
     - Auto-updating content can be paused/stopped
     - Session timeouts warn users and allow extension
     - No time limits or can be extended

   - **2.3 Seizures and Physical Reactions**:
     - No content flashes more than 3 times per second
     - Motion can be disabled (prefers-reduced-motion)

   - **2.4 Navigable**:
     - Page has descriptive title
     - Focus order is logical and meaningful
     - Link purpose clear from link text or context
     - Multiple ways to find pages (navigation, search, sitemap)
     - Headings and labels descriptive
     - Focus visible on all interactive elements
     - Breadcrumbs or location indicators present

   **C. Understandable** - Information and operation must be understandable

   - **3.1 Readable**:
     - Page language declared (html lang attribute)
     - Language changes marked (lang attribute)
     - Content written in clear, simple language

   - **3.2 Predictable**:
     - Navigation consistent across pages
     - Components behave consistently
     - No automatic context changes on focus
     - No automatic form submission without warning

   - **3.3 Input Assistance**:
     - Form errors clearly identified
     - Labels and instructions provided
     - Error suggestions provided when possible
     - Error prevention for critical actions (confirmations)
     - Required fields clearly marked

   **D. Robust** - Content must be robust enough for assistive technologies

   - **4.1 Compatible**:
     - Valid HTML (no duplicate IDs, proper nesting)
     - ARIA usage follows WAI-ARIA specifications
     - Name, role, value provided for custom components
     - Status messages use appropriate ARIA live regions

4. **Specific Component Checks:**

   - **Buttons**:
     - Use button element (not div/span with onClick)
     - Descriptive text or aria-label
     - Disabled state properly indicated
     - Focus visible
     - Keyboard accessible (Enter/Space)

   - **Forms**:
     - Labels associated with inputs (htmlFor/id)
     - Required fields marked (aria-required or required)
     - Error messages linked (aria-describedby)
     - Fieldsets/legends for grouped inputs
     - Autocomplete attributes where appropriate
     - Validation messages clear and helpful

   - **Navigation**:
     - nav element wraps navigation
     - Current page indicated (aria-current="page")
     - Skip links present
     - Logical tab order

   - **Modals/Dialogs**:
     - role="dialog" or role="alertdialog"
     - aria-modal="true"
     - aria-labelledby points to title
     - Focus trapped within modal
     - Focus returns to trigger on close
     - Escape key closes modal
     - Background inert (aria-hidden or inert)

   - **Dropdowns/Menus**:
     - ARIA roles (menu, menuitem, menubar)
     - aria-expanded state
     - aria-haspopup attribute
     - Keyboard navigation (Arrow keys, Enter, Escape)

   - **Tabs**:
     - role="tablist", role="tab", role="tabpanel"
     - aria-selected on active tab
     - aria-controls linking tabs to panels
     - Arrow key navigation between tabs

   - **Images**:
     - Informative images have descriptive alt
     - Decorative images have alt=""
     - Complex images have long descriptions

   - **Tables**:
     - Data tables have th elements
     - Headers have scope attribute
     - Complex tables use headers/id association

   - **Live Regions**:
     - Status updates use aria-live
     - Proper politeness level (polite/assertive)
     - Role="status" or role="alert" where appropriate

5. **React-Specific Accessibility Patterns:**

   - **Fragment usage**: Proper use to avoid unnecessary divs
   - **Focus management**: useRef and useEffect for focus control
   - **Dynamic content**: ARIA live regions for updates
   - **Conditional rendering**: Proper cleanup and focus handling
   - **Event handlers**: Both mouse and keyboard events
   - **Custom hooks**: Accessibility logic encapsulated (useFocusTrap, useAnnounce)

6. **Code Pattern Review:**

   Check for common anti-patterns:
   ```tsx
   // ❌ BAD
   <div onClick={handleClick}>Click me</div>
   <img src="photo.jpg" />
   <input type="text" />
   <div className="button">Submit</div>

   // ✅ GOOD
   <button onClick={handleClick}>Click me</button>
   <img src="photo.jpg" alt="Team photo at 2024 conference" />
   <label htmlFor="email">Email</label>
   <input id="email" type="text" />
   <button type="submit">Submit</button>
   ```

7. **Automated Tool Results** (if run):
   - Parse jest-axe test results
   - Categorize violations by WCAG criterion
   - Show violation count by impact level (critical, serious, moderate, minor)
   - Include code snippets showing violations

**Output Format:**

Generate a comprehensive accessibility report:

```markdown
# Accessibility Review Report

**Scope**: [Component/Feature/Project name]
**Date**: [Current date]
**Compliance Target**: WCAG 2.1 Level AA
**Automated Tools**: [Yes/No - jest-axe results if run]

---

## Executive Summary

- **Overall Compliance**: [Pass/Fail/Partial]
- **Total Issues Found**: [Number]
- **Critical Issues**: [Number] (Level A failures)
- **Important Issues**: [Number] (Level AA failures)
- **Best Practice Issues**: [Number] (AAA or enhancements)

---

## WCAG 2.1 Compliance by Principle

### 1. Perceivable
**Status**: [Pass/Fail/Partial]

#### 1.1 Text Alternatives
- [✓/✗] Images have descriptive alt text
- [✓/✗] Icon buttons labeled
- Issues found: [Number]

#### 1.3 Adaptable
- [✓/✗] Semantic HTML structure
- [✓/✗] Proper heading hierarchy
- [✓/✗] Form labels associated
- Issues found: [Number]

#### 1.4 Distinguishable
- [✓/✗] Color contrast (4.5:1 for text)
- [✓/✗] Focus indicators visible
- [✓/✗] Color not sole indicator
- Issues found: [Number]

### 2. Operable
**Status**: [Pass/Fail/Partial]

#### 2.1 Keyboard Accessible
- [✓/✗] All features keyboard accessible
- [✓/✗] No keyboard traps
- [✓/✗] Logical tab order
- Issues found: [Number]

#### 2.4 Navigable
- [✓/✗] Skip links present
- [✓/✗] Focus order logical
- [✓/✗] Link purpose clear
- Issues found: [Number]

### 3. Understandable
**Status**: [Pass/Fail/Partial]

#### 3.1 Readable
- [✓/✗] Language declared
- Issues found: [Number]

#### 3.2 Predictable
- [✓/✗] Consistent navigation
- [✓/✗] No unexpected context changes
- Issues found: [Number]

#### 3.3 Input Assistance
- [✓/✗] Error identification
- [✓/✗] Labels provided
- [✓/✗] Error suggestions
- Issues found: [Number]

### 4. Robust
**Status**: [Pass/Fail/Partial]

#### 4.1 Compatible
- [✓/✗] Valid HTML
- [✓/✗] ARIA properly used
- [✓/✗] Name, role, value present
- Issues found: [Number]

---

## Issues by Severity

### Critical (Level A) - Must Fix
[Number] issues found

**Issue 1: [WCAG Criterion] - [Component/File]**
- **Description**: [What's wrong]
- **Impact**: [How it affects users]
- **Location**: `src/features/[feature]/components/[Component].tsx:123`
- **Current Code**:
  ```tsx
  [Code snippet showing the issue]
  ```
- **Recommended Fix**:
  ```tsx
  [Code snippet showing the solution]
  ```
- **WCAG Reference**: [Criterion number and name]

### Important (Level AA) - Should Fix
[Number] issues found

[Same format as Critical]

### Best Practice (Enhancements) - Nice to Have
[Number] issues found

[Same format as Critical]

---

## Components with Issues

### [Component Name] (`path/to/Component.tsx`)
- [Number] issues found
- Severity: [Critical/Important/Minor]
- Issues:
  1. [Issue description]
  2. [Issue description]

---

## Automated Test Results

[If jest-axe was run]

**Total Violations**: [Number]
**Impact Breakdown**:
- Critical: [Number]
- Serious: [Number]
- Moderate: [Number]
- Minor: [Number]

**Sample Violations**:
```
[Paste relevant jest-axe output]
```

---

## Actionable Fixes

### Quick Wins (Easy fixes with high impact)
1. [Fix description] - `Component.tsx:123`
2. [Fix description] - `Component.tsx:456`

### Structural Changes (Require refactoring)
1. [Fix description] - `Component.tsx`
2. [Fix description] - `Component.tsx`

### Recommended Testing
1. Add jest-axe tests to all components
2. Test with screen readers (NVDA, JAWS, VoiceOver)
3. Test keyboard-only navigation
4. Test with browser zoom at 200%
5. Test in high contrast mode

---

## Resources

- [WCAG 2.1 Guidelines](https://www.w3.org/WAI/WCAG21/quickref/)
- [MDN Accessibility](https://developer.mozilla.org/en-US/docs/Web/Accessibility)
- [React Accessibility Docs](https://react.dev/learn/accessibility)
- [WAI-ARIA Authoring Practices](https://www.w3.org/WAI/ARIA/apg/)
- [WebAIM Contrast Checker](https://webaim.org/resources/contrastchecker/)

---

## Next Steps

1. Fix all Critical (Level A) issues immediately
2. Address Important (Level AA) issues within [timeframe]
3. Consider Best Practice enhancements for future iterations
4. Implement automated accessibility testing in CI/CD
5. Schedule regular accessibility audits
6. Consider user testing with people with disabilities
```

**Code Standards Compliance:**

Follow all standards from CODE_STANDARDS.md, specifically:

- **Accessibility Standards** (Section 4):
  - WCAG 2.1 Level AA compliance
  - Semantic HTML elements
  - ARIA only when HTML insufficient
  - Keyboard navigation
  - Focus indicators
  - Color contrast (4.5:1 for text, 3:1 for UI)
  - Form labels and instructions
  - Screen reader compatibility

- **Testing Standards** (Section 3):
  - Use jest-axe for automated testing
  - Test with accessible queries (getByRole, getByLabelText)

**Review Methodology:**

1. Start broad: Check overall structure and patterns
2. Get specific: Review individual components
3. Test practically: Try keyboard navigation, screen reader if possible
4. Document clearly: Provide actionable fixes with code examples
5. Prioritize: Critical issues first, then important, then nice-to-have

**After Review:**

1. Generate the accessibility report
2. Offer to create GitHub issues for critical violations
3. Suggest implementing automated accessibility testing if not present
4. Recommend accessibility testing tools and practices
5. Provide guidance on manual testing with assistive technologies

Be thorough but practical. Focus on violations that impact real users. Provide clear, actionable fixes with code examples.
