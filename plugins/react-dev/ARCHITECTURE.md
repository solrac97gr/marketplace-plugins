# React-Dev Plugin Architecture

This document describes the feature-based architecture pattern used by the react-dev plugin.

## Philosophy

The react-dev plugin uses **feature-based architecture** (also called feature-driven or vertical slice architecture) instead of traditional technical layer separation.

### Traditional Layer-Based ❌

```
src/
├── components/          # All components together
├── hooks/               # All hooks together
├── contexts/            # All contexts together
└── services/            # All services together
```

**Problems:**
- Features scattered across multiple directories
- Hard to understand feature boundaries
- Difficult to reuse or remove features
- Encourages tight coupling

### Feature-Based Architecture ✅

```
src/
├── features/
│   ├── user/            # Everything user-related
│   ├── product/         # Everything product-related
│   └── shared/          # Shared utilities
```

**Benefits:**
- **Feature Isolation**: Each feature is self-contained
- **Clear Boundaries**: Easy to see what belongs to a feature
- **Scalability**: Add/remove features without affecting others
- **Team Collaboration**: Teams can own entire features
- **Code Reuse**: Features expose public APIs
- **Testing**: Easy to test features in isolation

---

## Project Structure

```
project-name/
├── src/
│   ├── features/                          # Feature-based organization
│   │   ├── [feature-name]/                # Each feature (user, product, etc.)
│   │   │   ├── components/                # Feature UI components
│   │   │   │   ├── Component.tsx
│   │   │   │   ├── Component.test.tsx
│   │   │   │   └── Component.stories.tsx
│   │   │   ├── hooks/                     # Feature custom hooks
│   │   │   │   ├── useFeature.ts
│   │   │   │   └── useFeature.test.ts
│   │   │   ├── context/                   # Feature state management
│   │   │   │   ├── FeatureContext.tsx
│   │   │   │   └── FeatureContext.test.tsx
│   │   │   ├── services/                  # API calls & business logic
│   │   │   │   ├── featureService.ts
│   │   │   │   └── featureService.test.ts
│   │   │   ├── types/                     # Feature TypeScript types
│   │   │   │   └── index.ts
│   │   │   ├── utils/                     # Feature-specific utilities
│   │   │   │   ├── helpers.ts
│   │   │   │   └── helpers.test.ts
│   │   │   └── index.ts                   # Public API (named exports)
│   │   │
│   │   └── shared/                        # Shared across features
│   │       ├── components/                # Reusable UI components
│   │       ├── hooks/                     # Shared hooks
│   │       ├── context/                   # Global state (auth, theme)
│   │       ├── types/                     # Global types
│   │       └── utils/                     # Shared utilities
│   │
│   ├── app/                               # App-level configuration
│   │   ├── App.tsx                        # Root component
│   │   ├── App.test.tsx
│   │   ├── routes.tsx                     # Route definitions
│   │   └── providers.tsx                  # Global providers
│   │
│   ├── assets/                            # Static assets
│   │   ├── images/
│   │   ├── icons/
│   │   └── fonts/
│   │
│   ├── styles/                            # Global styles
│   │   └── globals.css                    # Tailwind imports
│   │
│   └── main.tsx                           # Entry point
│
├── tests/
│   ├── e2e/                               # Playwright E2E tests
│   │   ├── auth/
│   │   ├── user-flows/
│   │   └── critical-paths/
│   │
│   ├── integration/                       # Integration tests
│   │   └── features/
│   │
│   └── setup.ts                           # Test configuration
│
├── .storybook/                            # Storybook configuration
│   ├── main.ts
│   └── preview.ts
│
├── public/                                # Public static files
│   ├── favicon.ico
│   └── robots.txt
│
├── vite.config.ts                         # Vite configuration
├── vitest.config.ts                       # Vitest configuration
├── playwright.config.ts                   # Playwright configuration
├── tailwind.config.js                     # Tailwind configuration
├── tsconfig.json                          # TypeScript configuration
├── .eslintrc.json                         # ESLint configuration
├── .prettierrc                            # Prettier configuration
├── package.json
└── README.md
```

---

## Feature Anatomy

Each feature follows a consistent internal structure:

### Components Layer

**Purpose**: UI components specific to the feature

```
features/user/components/
├── UserProfile.tsx                 # Main component
├── UserProfile.test.tsx            # Unit tests
├── UserProfile.stories.tsx         # Storybook stories
├── UserAvatar.tsx                  # Sub-component
├── UserAvatar.test.tsx
└── UserSettings.tsx
```

**Rules:**
- Components are presentational or container components
- Use TypeScript interfaces for props (named exports)
- Include accessibility attributes
- Test user interactions with React Testing Library

**Example:**
```tsx
// UserProfile.tsx
import { useUser } from '../hooks/useUser';
import { UserAvatar } from './UserAvatar';
import type { UserProfileProps } from '../types';

export function UserProfile({ userId }: UserProfileProps) {
  const { user, loading } = useUser(userId);

  if (loading) return <div>Loading...</div>;
  if (!user) return <div>User not found</div>;

  return (
    <div className="user-profile">
      <UserAvatar src={user.avatar} alt={user.name} />
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </div>
  );
}
```

### Hooks Layer

**Purpose**: Custom hooks for feature logic and state

```
features/user/hooks/
├── useUser.ts                      # Fetch user data
├── useUser.test.ts
├── useUserForm.ts                  # Form management
└── useUserPermissions.ts           # Permission checks
```

**Rules:**
- Prefix with `use` (enforced by hooks rules)
- 100% test coverage (critical business logic)
- Use `renderHook` from React Testing Library

**Example:**
```tsx
// useUser.ts
import { useState, useEffect } from 'react';
import { fetchUser } from '../services/userService';
import type { User } from '../types';

export function useUser(userId: string) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    fetchUser(userId)
      .then(setUser)
      .catch(setError)
      .finally(() => setLoading(false));
  }, [userId]);

  return { user, loading, error };
}
```

### Context Layer

**Purpose**: Feature state management using React Context + useReducer

```
features/user/context/
├── UserContext.tsx                 # Context + Provider
└── UserContext.test.tsx
```

**Rules:**
- Use `useReducer` for complex state
- Separate context from provider
- Provide custom hook for consuming context
- Type state and actions

**Example:**
```tsx
// UserContext.tsx
import { createContext, useContext, useReducer } from 'react';
import type { User } from '../types';

interface UserState {
  currentUser: User | null;
  isAuthenticated: boolean;
}

type UserAction =
  | { type: 'LOGIN'; payload: User }
  | { type: 'LOGOUT' }
  | { type: 'UPDATE_USER'; payload: Partial<User> };

const UserContext = createContext<{
  state: UserState;
  dispatch: React.Dispatch<UserAction>;
} | undefined>(undefined);

function userReducer(state: UserState, action: UserAction): UserState {
  switch (action.type) {
    case 'LOGIN':
      return { currentUser: action.payload, isAuthenticated: true };
    case 'LOGOUT':
      return { currentUser: null, isAuthenticated: false };
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

export function UserProvider({ children }: { children: React.ReactNode }) {
  const [state, dispatch] = useReducer(userReducer, {
    currentUser: null,
    isAuthenticated: false,
  });

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

### Services Layer

**Purpose**: API calls and business logic

```
features/user/services/
├── userService.ts                  # API client
├── userService.test.ts
└── userValidation.ts               # Business logic
```

**Rules:**
- Handle API calls (fetch, axios)
- Transform API responses to domain types
- Error handling and retry logic
- 90%+ test coverage

**Example:**
```tsx
// userService.ts
import type { User } from '../types';

export async function fetchUser(userId: string): Promise<User> {
  const response = await fetch(`/api/users/${userId}`);
  if (!response.ok) {
    throw new Error(`Failed to fetch user: ${response.statusText}`);
  }
  return response.json();
}

export async function updateUser(userId: string, data: Partial<User>): Promise<User> {
  const response = await fetch(`/api/users/${userId}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    throw new Error(`Failed to update user: ${response.statusText}`);
  }
  return response.json();
}
```

### Types Layer

**Purpose**: TypeScript type definitions

```
features/user/types/
└── index.ts                        # All feature types
```

**Rules:**
- Named exports for all types
- Use interfaces for object shapes
- Use type aliases for unions, intersections
- Document complex types

**Example:**
```tsx
// types/index.ts
export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
  role: UserRole;
  createdAt: Date;
  updatedAt: Date;
}

export type UserRole = 'admin' | 'user' | 'guest';

export interface UserProfileProps {
  userId: string;
  onUpdate?: (user: User) => void;
}

export interface UserFormData {
  name: string;
  email: string;
  avatar?: File;
}
```

### Utils Layer

**Purpose**: Feature-specific utility functions

```
features/user/utils/
├── formatUserName.ts
├── formatUserName.test.ts
├── validateEmail.ts
└── validateEmail.test.ts
```

**Rules:**
- Pure functions (no side effects)
- 100% test coverage
- Single responsibility

**Example:**
```tsx
// formatUserName.ts
import type { User } from '../types';

export function formatUserName(user: User): string {
  return `${user.name} (${user.email})`;
}

export function getUserInitials(user: User): string {
  return user.name
    .split(' ')
    .map(n => n[0])
    .join('')
    .toUpperCase();
}
```

### Public API (index.ts)

**Purpose**: Expose feature's public interface

```tsx
// features/user/index.ts
export { UserProfile } from './components/UserProfile';
export { UserAvatar } from './components/UserAvatar';
export { useUser } from './hooks/useUser';
export { UserProvider, useUserContext } from './context/UserContext';
export { fetchUser, updateUser } from './services/userService';
export type { User, UserRole, UserProfileProps } from './types';
```

**Rules:**
- Named exports only (no default exports)
- Expose only what's needed by other features
- Keep internal implementation details private

---

## State Management Strategy

### Local Component State

**Use `useState` when:**
- State is only used in a single component
- State is UI-specific (e.g., form inputs, toggles)

```tsx
function SearchInput() {
  const [query, setQuery] = useState('');
  return <input value={query} onChange={e => setQuery(e.target.value)} />;
}
```

### Feature Context

**Use Context + useReducer when:**
- State is shared across multiple components in a feature
- State changes are complex (multiple actions)
- State needs to be testable in isolation

```tsx
// Feature-level context
<UserProvider>
  <UserProfile />
  <UserSettings />
</UserProvider>
```

### Global Context

**Use global context (in `shared/context`) when:**
- State is truly global (auth, theme, i18n)
- State is needed by multiple features

```tsx
// src/shared/context/AuthContext.tsx
export function AuthProvider({ children }) {
  // Global auth state
}

// src/app/providers.tsx
export function Providers({ children }) {
  return (
    <AuthProvider>
      <ThemeProvider>
        {children}
      </ThemeProvider>
    </AuthProvider>
  );
}
```

### Server State

**Use data fetching libraries for:**
- Remote data (API calls)
- Caching and synchronization

```tsx
// Consider react-query or SWR for complex server state
import { useQuery } from '@tanstack/react-query';

function useUser(userId: string) {
  return useQuery({
    queryKey: ['user', userId],
    queryFn: () => fetchUser(userId),
  });
}
```

---

## Testing Strategy

### Testing Pyramid

```
      E2E (Playwright)
     /                \
    /   Integration    \
   /     (RTL + CTX)    \
  /_______________________\
         Unit Tests
    (RTL + renderHook)
```

### Unit Tests (80% of tests)

**What to test:**
- Component rendering
- User interactions
- Hook behavior
- Service functions
- Utility functions

**Tools:** Vitest + React Testing Library

```tsx
// Component unit test
describe('UserProfile', () => {
  it('renders user information', () => {
    const user = { id: '1', name: 'John', email: 'john@example.com' };
    render(<UserProfile user={user} />);
    expect(screen.getByText('John')).toBeInTheDocument();
  });
});

// Hook unit test
import { renderHook, waitFor } from '@testing-library/react';

describe('useUser', () => {
  it('fetches user data', async () => {
    const { result } = renderHook(() => useUser('1'));
    expect(result.current.loading).toBe(true);
    await waitFor(() => expect(result.current.user).toBeDefined());
  });
});
```

### Integration Tests (15% of tests)

**What to test:**
- Feature workflows (multiple components + context)
- Component interaction with services
- Context provider behavior

**Tools:** Vitest + React Testing Library + Context

```tsx
// Integration test with context
describe('User Management Flow', () => {
  it('allows updating user profile', async () => {
    const user = userEvent.setup();
    render(
      <UserProvider>
        <UserProfile userId="1" />
      </UserProvider>
    );

    await user.click(screen.getByRole('button', { name: /edit/i }));
    await user.type(screen.getByLabelText(/name/i), 'New Name');
    await user.click(screen.getByRole('button', { name: /save/i }));

    expect(await screen.findByText('Profile updated')).toBeInTheDocument();
  });
});
```

### E2E Tests (5% of tests)

**What to test:**
- Critical user flows
- Multi-page workflows
- Authentication flows
- Payment flows

**Tools:** Playwright

```ts
// E2E test
test('complete user registration flow', async ({ page }) => {
  await page.goto('/register');
  await page.fill('input[name="email"]', 'user@example.com');
  await page.fill('input[name="password"]', 'password123');
  await page.click('button[type="submit"]');
  await expect(page).toHaveURL('/dashboard');
  await expect(page.getByText('Welcome')).toBeVisible();
});
```

---

## Code Organization Patterns

### Dependency Flow

```
Components → Hooks → Services → API
     ↓         ↓
  Context    Utils
```

**Rules:**
- Components depend on hooks and context
- Hooks depend on services
- Services are independent (no dependencies on hooks/components)
- Utils are pure functions (no dependencies)

### Import Rules

**✅ DO:**
```tsx
// Import from feature's public API
import { UserProfile, useUser } from '@/features/user';

// Import from shared
import { Button } from '@/features/shared/components/Button';

// Import types
import type { User } from '@/features/user';
```

**❌ DON'T:**
```tsx
// Don't import from internal paths
import { UserProfile } from '@/features/user/components/UserProfile';

// Don't import from other feature internals
import { userReducer } from '@/features/user/context/UserContext';
```

### File Naming

- **Components**: `PascalCase.tsx`
- **Hooks**: `camelCase.ts` with `use` prefix
- **Services**: `camelCase.ts` with `Service` suffix
- **Types**: `index.ts` in types directory
- **Tests**: Same as source + `.test.tsx`
- **Stories**: Same as component + `.stories.tsx`

---

## Routing

### Feature Routes

```tsx
// src/app/routes.tsx
import { lazy } from 'react';

// Lazy load feature routes
const UserRoutes = lazy(() => import('@/features/user/routes'));
const ProductRoutes = lazy(() => import('@/features/product/routes'));

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/users/*" element={<UserRoutes />} />
      <Route path="/products/*" element={<ProductRoutes />} />
    </Routes>
  );
}

// src/features/user/routes.tsx
export default function UserRoutes() {
  return (
    <Routes>
      <Route path="/" element={<UserList />} />
      <Route path="/:userId" element={<UserProfile />} />
      <Route path="/:userId/edit" element={<UserEdit />} />
    </Routes>
  );
}
```

---

## Shared Code

### When to Share

**Move to `shared/` when:**
- Used by 3+ features
- Truly generic (no feature-specific logic)
- Stable API (unlikely to change)

**Keep in feature when:**
- Used by 1-2 features
- Contains feature-specific logic
- API is still evolving

### Shared Components

```
features/shared/components/
├── Button/
│   ├── Button.tsx
│   ├── Button.test.tsx
│   └── Button.stories.tsx
├── Input/
├── Modal/
└── index.ts                    # Export all shared components
```

### Shared Hooks

```
features/shared/hooks/
├── useLocalStorage.ts
├── useDebounce.ts
├── useMediaQuery.ts
└── index.ts
```

### Shared Context

```
features/shared/context/
├── AuthContext.tsx             # Global authentication
├── ThemeContext.tsx            # Global theme
└── index.ts
```

---

## Migration Path

### From Layer-Based to Feature-Based

**Step 1**: Identify features
```
Current:
- UserProfile, UserList, UserSettings components
- useUser, useUserForm hooks

Target:
- features/user/ (all user-related code)
```

**Step 2**: Create feature directory
```bash
mkdir -p src/features/user/{components,hooks,context,services,types,utils}
```

**Step 3**: Move files
```bash
# Move components
mv src/components/UserProfile.tsx src/features/user/components/
mv src/components/UserList.tsx src/features/user/components/

# Move hooks
mv src/hooks/useUser.ts src/features/user/hooks/
```

**Step 4**: Create public API
```tsx
// src/features/user/index.ts
export { UserProfile } from './components/UserProfile';
export { useUser } from './hooks/useUser';
```

**Step 5**: Update imports
```tsx
// Before
import { UserProfile } from '@/components/UserProfile';

// After
import { UserProfile } from '@/features/user';
```

---

## Best Practices Summary

1. **Feature Isolation**: Each feature is self-contained
2. **Public API**: Expose only what's needed via index.ts
3. **Named Exports**: Use named exports everywhere
4. **Dependency Flow**: Components → Hooks → Services
5. **Test Coverage**: 80%+ components, 100% hooks
6. **Type Safety**: TypeScript strict mode, no any
7. **Accessibility**: WCAG 2.1 AA compliance
8. **Performance**: Lazy load features, memoize expensive operations
9. **Consistency**: Follow naming conventions and file structure
10. **Documentation**: Document complex features in README.md

---

## Example: Complete Feature

See the `/new-feature` skill for a complete walkthrough of creating a feature from scratch with TDD workflow.
