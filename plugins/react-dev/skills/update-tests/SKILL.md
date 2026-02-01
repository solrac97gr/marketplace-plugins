---
description: Update and improve test coverage for existing React code
---

Update test coverage for existing code by identifying gaps and generating missing tests.

**What to Do:**

1. **Ask the user:**
   - **Target scope**:
     - Entire project (all untested/undertested files)
     - Specific feature (e.g., "user", "auth", "dashboard")
     - Specific file path (e.g., "src/features/user/hooks/useUser.ts")
   - **Test types to add**:
     - Unit tests (components, hooks, utils, services)
     - Integration tests (components + context/providers)
     - E2E tests (critical user flows)
     - All of the above
   - **Run coverage report first?** (yes/no)
     - If yes: Run `npm run test:coverage` or `vitest run --coverage`

2. **Analysis Workflow:**

   a. **Scan target files** based on scope:
      - If entire project: Find all `.tsx`, `.ts` files in `src/`
      - If specific feature: Find all files in `src/features/[feature]/`
      - If specific file: Analyze that file only

   b. **Check for existing tests**:
      - For each source file (e.g., `Component.tsx`), look for `Component.test.tsx`
      - Note files with NO tests
      - Note files with INCOMPLETE tests (missing test cases)

   c. **Run coverage report** (if requested):
      - Execute coverage command
      - Parse output to identify:
        - Uncovered lines
        - Uncovered branches
        - Coverage percentages per file
      - Prioritize files below coverage goals:
        - Components: <80%
        - Hooks: <100%
        - Services: <90%
        - Utils: <100%

   d. **Identify untested code** (categorize by type):
      - **Components**: No tests or missing interaction/accessibility tests
      - **Hooks**: No tests or missing edge cases
      - **Services**: No tests or missing error handling tests
      - **Utils**: No tests or missing edge case tests
      - **Context providers**: No tests or missing integration tests

3. **Generate Missing Tests:**

   a. **Component Tests** (`ComponentName.test.tsx`):
      ```tsx
      import { render, screen } from '@testing-library/react';
      import userEvent from '@testing-library/user-event';
      import { axe, toHaveNoViolations } from 'jest-axe';
      import { ComponentName } from './ComponentName';

      expect.extend(toHaveNoViolations);

      describe('ComponentName', () => {
        describe('Rendering', () => {
          it('renders with required props', () => {
            render(<ComponentName prop="value" />);
            expect(screen.getByRole('...')).toBeInTheDocument();
          });

          it('renders all variants/states', () => {
            // Test each variant
          });
        });

        describe('User Interactions', () => {
          it('handles user actions', async () => {
            const user = userEvent.setup();
            const onAction = vi.fn();

            render(<ComponentName onAction={onAction} />);
            await user.click(screen.getByRole('button'));

            expect(onAction).toHaveBeenCalled();
          });

          it('supports keyboard navigation', async () => {
            const user = userEvent.setup();
            render(<ComponentName />);

            await user.tab();
            expect(screen.getByRole('button')).toHaveFocus();

            await user.keyboard('{Enter}');
            // Assert expected behavior
          });
        });

        describe('Accessibility', () => {
          it('has no accessibility violations', async () => {
            const { container } = render(<ComponentName prop="value" />);
            const results = await axe(container);
            expect(results).toHaveNoViolations();
          });

          it('has proper ARIA attributes', () => {
            render(<ComponentName />);
            expect(screen.getByRole('...')).toHaveAttribute('aria-label');
          });
        });

        describe('Edge Cases', () => {
          it('handles empty/null values gracefully', () => {
            // Test edge cases
          });

          it('handles error states', () => {
            // Test error handling
          });
        });
      });
      ```

   b. **Hook Tests** (`useHookName.test.ts`):
      ```tsx
      import { renderHook, waitFor } from '@testing-library/react';
      import { useHookName } from './useHookName';

      describe('useHookName', () => {
        it('returns initial state', () => {
          const { result } = renderHook(() => useHookName());

          expect(result.current.data).toBeNull();
          expect(result.current.loading).toBe(false);
        });

        it('handles state updates', async () => {
          const { result } = renderHook(() => useHookName());

          act(() => {
            result.current.updateData('new value');
          });

          expect(result.current.data).toBe('new value');
        });

        it('handles async operations', async () => {
          const { result } = renderHook(() => useHookName());

          act(() => {
            result.current.fetchData();
          });

          expect(result.current.loading).toBe(true);

          await waitFor(() => {
            expect(result.current.loading).toBe(false);
          });

          expect(result.current.data).toBeDefined();
        });

        it('handles errors', async () => {
          const { result } = renderHook(() => useHookName());

          act(() => {
            result.current.fetchData('invalid-id');
          });

          await waitFor(() => {
            expect(result.current.error).toBeDefined();
          });
        });

        it('cleans up on unmount', () => {
          const { unmount } = renderHook(() => useHookName());

          unmount();

          // Assert cleanup behavior
        });
      });
      ```

   c. **Service Tests** (`serviceName.test.ts`):
      ```tsx
      import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
      import { serviceName } from './serviceName';

      describe('serviceName', () => {
        beforeEach(() => {
          global.fetch = vi.fn();
        });

        afterEach(() => {
          vi.restoreAllMocks();
        });

        it('fetches data successfully', async () => {
          const mockData = { id: '1', name: 'Test' };

          (global.fetch as any).mockResolvedValueOnce({
            ok: true,
            json: async () => mockData,
          });

          const result = await serviceName.getData('1');

          expect(result).toEqual(mockData);
          expect(global.fetch).toHaveBeenCalledWith('/api/data/1');
        });

        it('handles network errors', async () => {
          (global.fetch as any).mockRejectedValueOnce(new Error('Network error'));

          await expect(serviceName.getData('1')).rejects.toThrow('Network error');
        });

        it('handles HTTP errors', async () => {
          (global.fetch as any).mockResolvedValueOnce({
            ok: false,
            status: 404,
            statusText: 'Not Found',
          });

          await expect(serviceName.getData('1')).rejects.toThrow();
        });

        it('sends correct request headers', async () => {
          (global.fetch as any).mockResolvedValueOnce({
            ok: true,
            json: async () => ({}),
          });

          await serviceName.createData({ name: 'Test' });

          expect(global.fetch).toHaveBeenCalledWith(
            '/api/data',
            expect.objectContaining({
              method: 'POST',
              headers: expect.objectContaining({
                'Content-Type': 'application/json',
              }),
            })
          );
        });
      });
      ```

   d. **Util/Helper Tests** (`helpers.test.ts`):
      ```tsx
      import { describe, it, expect } from 'vitest';
      import { utilFunction } from './helpers';

      describe('utilFunction', () => {
        it('handles normal cases', () => {
          expect(utilFunction('input')).toBe('expected output');
        });

        it('handles edge cases', () => {
          expect(utilFunction('')).toBe('');
          expect(utilFunction(null)).toBe(null);
          expect(utilFunction(undefined)).toBe(undefined);
        });

        it('handles special characters', () => {
          expect(utilFunction('!@#$%')).toBe('...');
        });

        it('handles large inputs', () => {
          const largeInput = 'x'.repeat(10000);
          expect(utilFunction(largeInput)).toBeDefined();
        });
      });
      ```

   e. **Integration Tests** (components + context):
      ```tsx
      // tests/integration/[feature]/FeatureFlow.test.tsx
      import { render, screen } from '@testing-library/react';
      import userEvent from '@testing-library/user-event';
      import { FeatureProvider } from '@/features/[feature]/context/FeatureContext';
      import { FeatureComponent } from '@/features/[feature]/components/FeatureComponent';

      describe('Feature Integration', () => {
        it('integrates component with context', async () => {
          const user = userEvent.setup();

          render(
            <FeatureProvider>
              <FeatureComponent />
            </FeatureProvider>
          );

          // Test component interacting with context
          await user.click(screen.getByRole('button', { name: /action/i }));

          expect(await screen.findByText(/success/i)).toBeInTheDocument();
        });

        it('handles context state updates', async () => {
          // Test state propagation through context
        });
      });
      ```

   f. **E2E Tests** (critical flows):
      ```ts
      // tests/e2e/[feature]/critical-flow.spec.ts
      import { test, expect } from '@playwright/test';

      test('user can complete critical flow', async ({ page }) => {
        await page.goto('/feature');

        // Fill form
        await page.fill('input[name="field"]', 'value');
        await page.click('button[type="submit"]');

        // Verify success
        await expect(page.locator('text=Success')).toBeVisible();
        await expect(page).toHaveURL('/success');
      });

      test('handles errors gracefully', async ({ page }) => {
        await page.goto('/feature');

        // Trigger error
        await page.fill('input[name="field"]', 'invalid');
        await page.click('button[type="submit"]');

        // Verify error message
        await expect(page.locator('text=Error')).toBeVisible();
      });
      ```

4. **Update Existing Tests:**

   a. **Modernize to React Testing Library best practices**:
      - Replace `getByTestId` with `getByRole`, `getByLabelText`, `getByText`
      - Replace Enzyme shallow/mount with RTL `render`
      - Replace wrapper.find() with screen queries
      - Use `userEvent` instead of `fireEvent`

   b. **Add missing accessibility tests**:
      ```tsx
      it('has no accessibility violations', async () => {
        const { container } = render(<Component />);
        const results = await axe(container);
        expect(results).toHaveNoViolations();
      });
      ```

   c. **Improve query priorities** (prefer accessible queries):
      ```tsx
      // Before (not ideal)
      screen.getByTestId('submit-button')

      // After (better)
      screen.getByRole('button', { name: /submit/i })
      ```

   d. **Add user interaction tests**:
      ```tsx
      it('handles user interactions', async () => {
        const user = userEvent.setup();
        const onClick = vi.fn();

        render(<Button onClick={onClick}>Click me</Button>);

        await user.click(screen.getByRole('button'));
        expect(onClick).toHaveBeenCalled();
      });
      ```

   e. **Add edge case tests**:
      - Null/undefined values
      - Empty arrays/strings
      - Error states
      - Loading states
      - Boundary conditions

5. **Coverage Goals:**

   Ensure each file type meets minimum coverage:
   - **Components**: 80%+ coverage
     - All variants rendered
     - All user interactions tested
     - Accessibility validated
     - Edge cases covered

   - **Hooks**: 100% coverage
     - All state transitions
     - All side effects
     - Cleanup functions
     - Error handling

   - **Services**: 90%+ coverage
     - All API calls
     - Success responses
     - Error responses
     - Network failures

   - **Utils**: 100% coverage
     - All code paths
     - All edge cases
     - All error conditions

6. **Output Format:**

   After updating tests, provide:

   ```
   ## Test Coverage Update Summary

   ### Coverage Before:
   - Components: X% (Y files below 80%)
   - Hooks: X% (Y files below 100%)
   - Services: X% (Y files below 90%)
   - Utils: X% (Y files below 100%)

   ### Coverage After:
   - Components: X% (↑ +N%)
   - Hooks: X% (↑ +N%)
   - Services: X% (↑ +N%)
   - Utils: X% (↑ +N%)

   ### Files Updated:
   1. ✅ src/features/user/components/UserCard.tsx
      - Added: Rendering tests, interaction tests, accessibility tests
      - Coverage: 55% → 85%

   2. ✅ src/features/user/hooks/useUser.ts
      - Added: Edge case tests, error handling tests
      - Coverage: 70% → 100%

   3. ✅ src/features/user/services/userService.ts
      - Added: HTTP error tests, network failure tests
      - Coverage: 65% → 92%

   ### Tests Added:
   - Unit tests: N new test files, M new test cases
   - Integration tests: N new test files
   - E2E tests: N new test files
   - Accessibility tests: N components now tested

   ### Remaining Gaps:
   - File: src/features/dashboard/components/Chart.tsx
     - Missing: E2E tests for chart interactions
     - Coverage: 78% (below 80% goal)

   - File: src/features/auth/hooks/useAuth.ts
     - Missing: Token refresh flow tests
     - Coverage: 85% (below 100% goal)

   ### Next Steps:
   1. Run: `npm run test:coverage` to verify coverage
   2. Review: Check if new tests pass with `npm test`
   3. Fix: Address any remaining gaps listed above
   4. Commit: Create commit with test updates
   ```

**Code Standards Compliance:**

Follow all standards from CODE_STANDARDS.md:

- **Testing Standards**: User-centric tests, accessibility tests, proper coverage
- **Query Priorities**: getByRole > getByLabelText > getByText > getByTestId
- **User Interactions**: Use `userEvent` over `fireEvent`
- **Accessibility**: Include `jest-axe` tests for all components
- **TypeScript**: Properly type test mocks and fixtures
- **Coverage Goals**: Meet minimum coverage percentages per file type

**Commands to Run:**

```bash
# Generate coverage report
npm run test:coverage
# or
vitest run --coverage

# Run specific test file
npm test -- path/to/file.test.tsx

# Run tests in watch mode
npm test
# or
vitest

# Run E2E tests
npm run test:e2e
# or
npx playwright test
```

**Best Practices:**

1. **Prioritize user-visible behavior** over implementation details
2. **Use accessible queries** (getByRole, getByLabelText) to ensure components are accessible
3. **Test edge cases** including null, undefined, empty, and error states
4. **Add accessibility tests** for all interactive components
5. **Mock external dependencies** (APIs, context, localStorage) properly
6. **Test async behavior** with waitFor, findBy queries
7. **Organize tests** with describe blocks (Rendering, Interactions, Accessibility, Edge Cases)
8. **Follow AAA pattern**: Arrange, Act, Assert

Be thorough and ensure all untested code paths are covered. Ask clarifying questions if scope is unclear.
