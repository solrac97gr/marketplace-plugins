---
description: Create a custom React hook with comprehensive tests
---

Create a new custom React hook following best practices with comprehensive tests for critical business logic.

**What to Create:**

1. **Ask the user:**
   - Hook name (must start with "use", e.g., "useUser", "useForm", "useDebounce")
   - Which feature does it belong to? (or "shared" for shared hooks)
   - What does the hook do? (description of purpose and behavior)
   - What parameters does it accept? (name, type, required/optional, default values)
   - What does it return? (return value structure and types)
   - Does it manage state? (yes/no) - if yes, what state?
   - Does it have side effects (useEffect)? (yes/no) - if yes, what side effects?
   - Does it depend on other hooks? (e.g., custom hooks, context)
   - Any specific edge cases or error handling requirements?

2. **Validate Hook Name:**
   - Ensure hook name starts with "use" (enforced by React hooks rules)
   - If invalid, prompt user to provide a valid hook name

3. **Generate Hook Implementation** in `src/features/[feature]/hooks/[hookName].ts` (or `src/features/shared/hooks/[hookName].ts` for shared):
   - Use functional component pattern
   - Follow **Rules of Hooks**:
     - No conditional hooks (hooks must be called in the same order every render)
     - No hooks in loops
     - No hooks in nested functions
     - Only call hooks at the top level
   - Proper **dependency arrays** in useEffect/useMemo/useCallback
   - TypeScript interfaces for:
     - Hook parameters (if any)
     - Hook return type (always explicitly type the return)
     - Internal state types
   - **Error handling**:
     - Validate input parameters
     - Handle async errors gracefully
     - Provide error state if needed
   - **Cleanup**:
     - Clean up effects (return cleanup function from useEffect)
     - Cancel pending requests on unmount
     - Remove event listeners
   - **Edge cases**:
     - Handle null/undefined inputs
     - Handle empty arrays/objects
     - Handle race conditions for async operations
   - Add JSDoc comments explaining:
     - Hook purpose
     - Parameters
     - Return value
     - Example usage

4. **Generate Tests** in `src/features/[feature]/hooks/[hookName].test.ts`:
   - Import `renderHook`, `waitFor`, `act` from `@testing-library/react`
   - **Initial state tests**: Hook returns correct initial values
   - **State update tests**: State updates work correctly
   - **Side effect tests**: Effects run when dependencies change
   - **Cleanup tests**: Cleanup functions are called
   - **Error handling tests**: Errors are caught and handled
   - **Edge case tests**: Null/undefined inputs, empty data, race conditions
   - **Dependency array tests**: Effects run only when dependencies change
   - **Re-render tests**: Hook behaves correctly on re-renders
   - **Async tests**: Handle async operations with `waitFor`
   - Aim for **100% coverage** (hooks contain critical business logic)
   - Use `describe` blocks to organize tests by concern
   - Mock external dependencies (services, context, timers)

**Hook Structure Pattern:**

```tsx
// useHookName.ts
import { useState, useEffect, useCallback } from 'react';

/**
 * Hook description explaining what it does and when to use it
 *
 * @param param1 - Description of parameter
 * @param param2 - Description of parameter
 * @returns Object containing state and helper functions
 *
 * @example
 * const { data, loading, error, refetch } = useHookName(userId);
 */

interface UseHookNameParams {
  param1: string;
  param2?: number;
}

interface UseHookNameReturn {
  data: DataType | null;
  loading: boolean;
  error: Error | null;
  refetch: () => void;
}

export function useHookName({
  param1,
  param2 = defaultValue
}: UseHookNameParams): UseHookNameReturn {
  // State management
  const [data, setData] = useState<DataType | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  // Memoized callbacks
  const refetch = useCallback(() => {
    setLoading(true);
    setError(null);
    // Fetch logic
  }, [param1, param2]);

  // Side effects
  useEffect(() => {
    let cancelled = false;

    async function fetchData() {
      try {
        setLoading(true);
        const result = await someAsyncOperation(param1);

        if (!cancelled) {
          setData(result);
          setError(null);
        }
      } catch (err) {
        if (!cancelled) {
          setError(err instanceof Error ? err : new Error('Unknown error'));
        }
      } finally {
        if (!cancelled) {
          setLoading(false);
        }
      }
    }

    fetchData();

    // Cleanup function
    return () => {
      cancelled = true;
      // Cancel any pending operations
      // Remove event listeners
    };
  }, [param1, param2]); // Proper dependency array

  return {
    data,
    loading,
    error,
    refetch
  };
}
```

**Test Structure Pattern:**

```tsx
// useHookName.test.ts
import { renderHook, waitFor, act } from '@testing-library/react';
import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { useHookName } from './useHookName';

// Mock dependencies
vi.mock('../services/someService', () => ({
  fetchData: vi.fn()
}));

describe('useHookName', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  describe('Initial State', () => {
    it('returns correct initial values', () => {
      const { result } = renderHook(() => useHookName({ param1: 'test' }));

      expect(result.current.data).toBeNull();
      expect(result.current.loading).toBe(true);
      expect(result.current.error).toBeNull();
    });

    it('handles default parameters', () => {
      const { result } = renderHook(() => useHookName({ param1: 'test' }));

      // Assert default parameter behavior
      expect(result.current).toBeDefined();
    });
  });

  describe('State Updates', () => {
    it('updates state when data is fetched successfully', async () => {
      const mockData = { id: '1', name: 'Test' };
      vi.mocked(fetchData).mockResolvedValue(mockData);

      const { result } = renderHook(() => useHookName({ param1: 'test' }));

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      expect(result.current.data).toEqual(mockData);
      expect(result.current.error).toBeNull();
    });

    it('updates error state when fetch fails', async () => {
      const mockError = new Error('Fetch failed');
      vi.mocked(fetchData).mockRejectedValue(mockError);

      const { result } = renderHook(() => useHookName({ param1: 'test' }));

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      expect(result.current.data).toBeNull();
      expect(result.current.error).toEqual(mockError);
    });
  });

  describe('Side Effects', () => {
    it('refetches data when dependencies change', async () => {
      const { result, rerender } = renderHook(
        ({ param1 }) => useHookName({ param1 }),
        { initialProps: { param1: 'initial' } }
      );

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      // Change dependency
      rerender({ param1: 'updated' });

      expect(result.current.loading).toBe(true);
      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      expect(fetchData).toHaveBeenCalledTimes(2);
    });

    it('does not refetch when non-dependencies change', async () => {
      const { result, rerender } = renderHook(
        ({ param1, other }) => useHookName({ param1 }),
        { initialProps: { param1: 'test', other: 'value1' } }
      );

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      const callCount = vi.mocked(fetchData).mock.calls.length;

      // Change non-dependency
      rerender({ param1: 'test', other: 'value2' });

      // Should not trigger refetch
      expect(vi.mocked(fetchData).mock.calls.length).toBe(callCount);
    });
  });

  describe('Cleanup', () => {
    it('cancels pending requests on unmount', async () => {
      const { unmount } = renderHook(() => useHookName({ param1: 'test' }));

      unmount();

      // Assert cleanup was called
      // e.g., abort controller signal was triggered
    });

    it('removes event listeners on unmount', () => {
      const removeEventListenerSpy = vi.spyOn(window, 'removeEventListener');

      const { unmount } = renderHook(() => useHookName({ param1: 'test' }));

      unmount();

      expect(removeEventListenerSpy).toHaveBeenCalled();
    });
  });

  describe('Error Handling', () => {
    it('handles null/undefined parameters gracefully', () => {
      const { result } = renderHook(() => useHookName({ param1: null as any }));

      expect(result.current.error).toBeDefined();
    });

    it('handles empty data', async () => {
      vi.mocked(fetchData).mockResolvedValue(null);

      const { result } = renderHook(() => useHookName({ param1: 'test' }));

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      expect(result.current.data).toBeNull();
    });

    it('handles race conditions', async () => {
      let resolveFirst: (value: any) => void;
      let resolveSecond: (value: any) => void;

      const firstPromise = new Promise(resolve => { resolveFirst = resolve; });
      const secondPromise = new Promise(resolve => { resolveSecond = resolve; });

      vi.mocked(fetchData)
        .mockReturnValueOnce(firstPromise as any)
        .mockReturnValueOnce(secondPromise as any);

      const { result, rerender } = renderHook(
        ({ param1 }) => useHookName({ param1 }),
        { initialProps: { param1: 'first' } }
      );

      // Trigger second request
      rerender({ param1: 'second' });

      // Resolve first request (should be ignored)
      act(() => {
        resolveFirst!({ id: '1', name: 'First' });
      });

      // Resolve second request
      act(() => {
        resolveSecond!({ id: '2', name: 'Second' });
      });

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      // Should only use second result
      expect(result.current.data).toEqual({ id: '2', name: 'Second' });
    });
  });

  describe('Refetch Functionality', () => {
    it('refetches data when refetch is called', async () => {
      const { result } = renderHook(() => useHookName({ param1: 'test' }));

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      const callCount = vi.mocked(fetchData).mock.calls.length;

      act(() => {
        result.current.refetch();
      });

      expect(result.current.loading).toBe(true);
      expect(vi.mocked(fetchData).mock.calls.length).toBe(callCount + 1);
    });
  });

  describe('Edge Cases', () => {
    it('handles rapid re-renders', async () => {
      const { result, rerender } = renderHook(
        ({ param1 }) => useHookName({ param1 }),
        { initialProps: { param1: 'test1' } }
      );

      // Rapid re-renders
      rerender({ param1: 'test2' });
      rerender({ param1: 'test3' });
      rerender({ param1: 'test4' });

      await waitFor(() => {
        expect(result.current.loading).toBe(false);
      });

      // Should handle gracefully
      expect(result.current.error).toBeNull();
    });

    it('handles concurrent renders in React 18', async () => {
      // Test React 18 concurrent features if applicable
    });
  });
});
```

**Common Hook Patterns:**

**Data Fetching Hook:**
```tsx
export function useUser(userId: string) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    let cancelled = false;

    fetchUser(userId)
      .then(data => !cancelled && setUser(data))
      .catch(err => !cancelled && setError(err))
      .finally(() => !cancelled && setLoading(false));

    return () => { cancelled = true; };
  }, [userId]);

  return { user, loading, error };
}
```

**Form Hook:**
```tsx
export function useForm<T>(initialValues: T) {
  const [values, setValues] = useState<T>(initialValues);
  const [errors, setErrors] = useState<Partial<Record<keyof T, string>>>({});
  const [touched, setTouched] = useState<Partial<Record<keyof T, boolean>>>({});

  const handleChange = useCallback((field: keyof T, value: any) => {
    setValues(prev => ({ ...prev, [field]: value }));
  }, []);

  const handleBlur = useCallback((field: keyof T) => {
    setTouched(prev => ({ ...prev, [field]: true }));
  }, []);

  const validate = useCallback(() => {
    // Validation logic
  }, [values]);

  return { values, errors, touched, handleChange, handleBlur, validate };
}
```

**Debounce Hook:**
```tsx
export function useDebounce<T>(value: T, delay: number): T {
  const [debouncedValue, setDebouncedValue] = useState<T>(value);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value);
    }, delay);

    return () => {
      clearTimeout(handler);
    };
  }, [value, delay]);

  return debouncedValue;
}
```

**Event Listener Hook:**
```tsx
export function useEventListener<K extends keyof WindowEventMap>(
  eventName: K,
  handler: (event: WindowEventMap[K]) => void,
  element: HTMLElement | Window = window
) {
  const savedHandler = useRef(handler);

  useEffect(() => {
    savedHandler.current = handler;
  }, [handler]);

  useEffect(() => {
    const eventListener = (event: WindowEventMap[K]) => savedHandler.current(event);
    element.addEventListener(eventName, eventListener as any);

    return () => {
      element.removeEventListener(eventName, eventListener as any);
    };
  }, [eventName, element]);
}
```

**Code Standards Compliance:**

Follow all standards from CODE_STANDARDS.md:

- **Hooks Rules**: Follow all React hooks rules (no conditional hooks, proper dependencies)
- **TypeScript Standards**: Strict mode, explicit return types, no `any`
- **Testing Standards**: 100% coverage for hooks (critical business logic)
- **Naming Conventions**: camelCase with `use` prefix
- **Error Handling**: Validate inputs, handle async errors, provide error state
- **Cleanup**: Always clean up effects (event listeners, timers, subscriptions)
- **Performance**: Use useCallback/useMemo when returning functions/objects

**File Organization:**

For feature hooks:
```
src/features/[feature]/hooks/
├── useHookName.ts
└── useHookName.test.ts
```

For shared hooks:
```
src/features/shared/hooks/
├── useHookName.ts
└── useHookName.test.ts
```

**After Creation:**

1. Verify files are created in correct locations
2. Run tests: `npm test useHookName.test.ts`
3. Check coverage: `npm run test:coverage -- useHookName`
4. Verify 100% coverage is achieved
5. Check TypeScript compilation: `npm run type-check`
6. Test hook in a real component if needed

**Reference:**

See ARCHITECTURE.md "Hooks Layer" section (lines 174-212) for architectural context and examples.

Be thorough and ask clarifying questions if requirements are unclear. Hooks are critical business logic and require comprehensive testing.
