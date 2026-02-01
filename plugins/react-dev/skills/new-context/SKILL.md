---
description: Create Context + Provider + useReducer for state management with tests
---

Create a new React Context with Provider and useReducer for state management following best practices.

**What to Create:**

1. **Ask the user:**
   - Context name (e.g., "User", "Theme", "Cart")
   - Which feature does it belong to? (or "shared" for global context)
   - What state does it manage? (list of state fields with types)
     - Example: "currentUser (User | null), isAuthenticated (boolean), loading (boolean)"
   - What actions are needed? (list of actions with payload types)
     - Example: "LOGIN (User), LOGOUT (void), UPDATE_USER (Partial<User>), SET_LOADING (boolean)"

2. **Generate Context** in `src/features/[feature]/context/[ContextName]Context.tsx` (or `src/features/shared/context/[ContextName]Context.tsx` for global):
   - State interface with TypeScript types
   - Action discriminated union types
   - Initial state constant
   - Reducer function with switch statement
   - Context with createContext
   - Provider component wrapping children
   - Custom useContext hook with error checking (throws if used outside provider)
   - All exports are named (no default exports)
   - Proper TypeScript typing (no `any`)

3. **Generate Tests** in `src/features/[feature]/context/[ContextName]Context.test.tsx`:
   - Import from `@testing-library/react`
   - **Reducer tests**:
     - Test each action updates state correctly
     - Test initial state
     - Test invalid actions return unchanged state
   - **Provider tests**:
     - Test provider renders children
     - Test initial state is accessible
   - **Hook tests** (using renderHook):
     - Test hook provides state and dispatch
     - Test dispatching actions updates state
     - Test error when hook used outside provider
   - Use `describe` blocks to organize tests
   - Aim for 100% coverage (critical state management)

**Context Structure Pattern:**

```tsx
// [ContextName]Context.tsx
import { createContext, useContext, useReducer, type ReactNode } from 'react';

// State interface
export interface [ContextName]State {
  field1: Type1;
  field2: Type2;
  // ... more fields
}

// Action types (discriminated union)
export type [ContextName]Action =
  | { type: 'ACTION_1'; payload: PayloadType1 }
  | { type: 'ACTION_2'; payload: PayloadType2 }
  | { type: 'ACTION_3' } // No payload
  | { type: 'ACTION_4'; payload: PayloadType4 };

// Context type
interface [ContextName]ContextType {
  state: [ContextName]State;
  dispatch: React.Dispatch<[ContextName]Action>;
}

// Create context
const [ContextName]Context = createContext<[ContextName]ContextType | undefined>(
  undefined
);

// Initial state
const initialState: [ContextName]State = {
  field1: defaultValue1,
  field2: defaultValue2,
  // ... more fields
};

// Reducer function
function [contextName]Reducer(
  state: [ContextName]State,
  action: [ContextName]Action
): [ContextName]State {
  switch (action.type) {
    case 'ACTION_1':
      return {
        ...state,
        field1: action.payload,
        // Update relevant fields
      };
    case 'ACTION_2':
      return {
        ...state,
        field2: action.payload,
      };
    case 'ACTION_3':
      return {
        ...state,
        // Update without payload
      };
    case 'ACTION_4':
      return {
        ...state,
        field3: action.payload,
      };
    default:
      return state;
  }
}

// Provider component
export function [ContextName]Provider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer([contextName]Reducer, initialState);

  return (
    <[ContextName]Context.Provider value={{ state, dispatch }}>
      {children}
    </[ContextName]Context.Provider>
  );
}

// Custom hook
export function use[ContextName]Context() {
  const context = useContext([ContextName]Context);
  if (!context) {
    throw new Error('use[ContextName]Context must be used within [ContextName]Provider');
  }
  return context;
}
```

**Test Structure Pattern:**

```tsx
// [ContextName]Context.test.tsx
import { describe, it, expect } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { render, screen } from '@testing-library/react';
import {
  [ContextName]Provider,
  use[ContextName]Context,
  type [ContextName]State,
  type [ContextName]Action,
} from './[ContextName]Context';

describe('[ContextName]Context', () => {
  describe('Reducer', () => {
    it('returns initial state', () => {
      const { result } = renderHook(() => use[ContextName]Context(), {
        wrapper: [ContextName]Provider,
      });

      expect(result.current.state).toEqual({
        field1: defaultValue1,
        field2: defaultValue2,
        // ... expected initial state
      });
    });

    it('handles ACTION_1', () => {
      const { result } = renderHook(() => use[ContextName]Context(), {
        wrapper: [ContextName]Provider,
      });

      act(() => {
        result.current.dispatch({ type: 'ACTION_1', payload: testValue });
      });

      expect(result.current.state.field1).toBe(testValue);
    });

    it('handles ACTION_2', () => {
      const { result } = renderHook(() => use[ContextName]Context(), {
        wrapper: [ContextName]Provider,
      });

      act(() => {
        result.current.dispatch({ type: 'ACTION_2', payload: testValue });
      });

      expect(result.current.state.field2).toBe(testValue);
    });

    // Add more action tests...
  });

  describe('Provider', () => {
    it('renders children', () => {
      render(
        <[ContextName]Provider>
          <div>Test Child</div>
        </[ContextName]Provider>
      );

      expect(screen.getByText('Test Child')).toBeInTheDocument();
    });

    it('provides state and dispatch', () => {
      const { result } = renderHook(() => use[ContextName]Context(), {
        wrapper: [ContextName]Provider,
      });

      expect(result.current.state).toBeDefined();
      expect(result.current.dispatch).toBeDefined();
      expect(typeof result.current.dispatch).toBe('function');
    });
  });

  describe('use[ContextName]Context hook', () => {
    it('throws error when used outside provider', () => {
      // Suppress console.error for this test
      const spy = vi.spyOn(console, 'error').mockImplementation(() => {});

      expect(() => {
        renderHook(() => use[ContextName]Context());
      }).toThrow('use[ContextName]Context must be used within [ContextName]Provider');

      spy.mockRestore();
    });

    it('provides context value when used inside provider', () => {
      const { result } = renderHook(() => use[ContextName]Context(), {
        wrapper: [ContextName]Provider,
      });

      expect(result.current).toBeDefined();
      expect(result.current.state).toBeDefined();
      expect(result.current.dispatch).toBeDefined();
    });

    it('allows multiple actions in sequence', () => {
      const { result } = renderHook(() => use[ContextName]Context(), {
        wrapper: [ContextName]Provider,
      });

      act(() => {
        result.current.dispatch({ type: 'ACTION_1', payload: value1 });
        result.current.dispatch({ type: 'ACTION_2', payload: value2 });
      });

      expect(result.current.state.field1).toBe(value1);
      expect(result.current.state.field2).toBe(value2);
    });
  });
});
```

**Real-World Example (UserContext):**

Based on the user asking for "User" context managing "currentUser (User | null), isAuthenticated (boolean)" with actions "LOGIN (User), LOGOUT (void), UPDATE_USER (Partial<User>)":

```tsx
// UserContext.tsx
import { createContext, useContext, useReducer, type ReactNode } from 'react';

export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
}

export interface UserState {
  currentUser: User | null;
  isAuthenticated: boolean;
}

export type UserAction =
  | { type: 'LOGIN'; payload: User }
  | { type: 'LOGOUT' }
  | { type: 'UPDATE_USER'; payload: Partial<User> };

interface UserContextType {
  state: UserState;
  dispatch: React.Dispatch<UserAction>;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

const initialState: UserState = {
  currentUser: null,
  isAuthenticated: false,
};

function userReducer(state: UserState, action: UserAction): UserState {
  switch (action.type) {
    case 'LOGIN':
      return {
        ...state,
        currentUser: action.payload,
        isAuthenticated: true,
      };
    case 'LOGOUT':
      return {
        ...state,
        currentUser: null,
        isAuthenticated: false,
      };
    case 'UPDATE_USER':
      return {
        ...state,
        currentUser: state.currentUser
          ? { ...state.currentUser, ...action.payload }
          : null,
      };
    default:
      return state;
  }
}

export function UserProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(userReducer, initialState);

  return (
    <UserContext.Provider value={{ state, dispatch }}>
      {children}
    </UserContext.Provider>
  );
}

export function useUserContext() {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error('useUserContext must be used within UserProvider');
  }
  return context;
}
```

**Code Standards Compliance:**

Follow all standards from CODE_STANDARDS.md:

- **React Standards**: Functional components, named exports, hooks rules
- **TypeScript Standards**: Strict mode, no `any`, proper typing, discriminated unions
- **Testing Standards**: 100% coverage for state management (critical logic)
- **Naming Conventions**:
  - PascalCase for Context, Provider, State interface
  - camelCase for reducer function
  - UPPER_SNAKE_CASE for action types
  - use prefix for custom hook

**Architecture Patterns:**

Follow the Context Layer pattern from ARCHITECTURE.md:

- Use `useReducer` for complex state (multiple related state values)
- Separate context creation from provider
- Provide custom hook for consuming context
- Type state and actions properly
- Export state, action types, provider, and hook

**File Organization:**

For feature-specific context:
```
src/features/[feature]/context/
├── [ContextName]Context.tsx
└── [ContextName]Context.test.tsx
```

For global/shared context:
```
src/features/shared/context/
├── [ContextName]Context.tsx
└── [ContextName]Context.test.tsx
```

**After Creation:**

1. Verify both files are created in correct locations
2. Run tests: `npm test [ContextName]Context.test.tsx`
3. Verify 100% test coverage
4. Check TypeScript compilation: `npm run type-check`
5. Update feature's `index.ts` to export:
   ```tsx
   export { [ContextName]Provider, use[ContextName]Context } from './context/[ContextName]Context';
   export type { [ContextName]State, [ContextName]Action } from './context/[ContextName]Context';
   ```

**When to Use Context:**

- **Feature Context**: State shared across multiple components within a feature
- **Global Context**: State needed by multiple features (auth, theme, i18n)
- **Use Local State**: If state is only used in one component, use `useState` instead

**Best Practices:**

1. Keep reducer functions pure (no side effects)
2. Use discriminated unions for type-safe actions
3. Always provide error message in custom hook
4. Test all actions thoroughly
5. Document complex state transitions
6. Consider using action creators for complex payloads
7. Keep context focused (single responsibility)

Be concise and follow React Context best practices. Ask clarifying questions if state structure or actions are unclear.
