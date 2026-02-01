---
description: Scaffold a complete React feature following TDD workflow
---

Create a complete React feature following feature-based architecture with TDD workflow.

**Discovery Process:**

First, understand what feature the user wants to build by asking targeted questions. Based on the answers, generate a complete, production-ready feature structure.

**Ask the User (using interactive prompts):**

1. **Feature name** (text input)
   - Example: "user-profile", "product-catalog", "shopping-cart"
   - Must be kebab-case, lowercase

2. **Feature description** (text input)
   - Brief description of what the feature does
   - Example: "Manage user profile information with avatar upload and settings"

3. **Main entities/models** (text input)
   - What data models does this feature work with?
   - Example: "User (id, name, email, avatar, role, bio)"
   - Can be multiple entities separated by commas

4. **API integration needed?** (yes/no)
   - Does this feature need to fetch/send data to an API?
   - If yes, generate service layer with API calls

5. **Local state management needed?** (yes/no)
   - Does this feature need complex state shared across components?
   - If yes, generate Context + useReducer pattern

**Feature Structure to Create:**

```
src/features/[feature-name]/
├── components/
│   ├── [MainComponent].tsx           # Primary component
│   ├── [MainComponent].test.tsx      # Component tests
│   ├── [MainComponent].stories.tsx   # Storybook stories
│   ├── [SubComponent].tsx            # Supporting components
│   ├── [SubComponent].test.tsx
│   └── [SubComponent].stories.tsx
├── hooks/
│   ├── use[Feature].ts               # Main feature hook
│   ├── use[Feature].test.ts          # Hook tests
│   ├── use[Feature]Form.ts           # Form management hook (if needed)
│   └── use[Feature]Form.test.ts
├── context/                           # Only if state management needed
│   ├── [Feature]Context.tsx          # Context + Provider + useReducer
│   └── [Feature]Context.test.tsx     # Context tests
├── services/                          # Only if API integration needed
│   ├── [feature]Service.ts           # API client methods
│   └── [feature]Service.test.ts      # Service tests with mocks
├── types/
│   └── index.ts                      # TypeScript interfaces & types
├── utils/
│   ├── [helper].ts                   # Feature-specific utilities
│   └── [helper].test.ts              # Utility tests
├── index.ts                          # Public API (named exports)
└── README.md                         # Feature documentation
```

**TDD Workflow - STRICTLY FOLLOW THIS ORDER:**

### Phase 1: Discovery & Planning

1. Gather all information from user questions
2. Identify components needed based on feature description
3. Plan the component hierarchy and data flow
4. Document assumptions in feature README.md

### Phase 2: Types First (BLUE)

**Generate TypeScript types BEFORE any implementation:**

1. Create `types/index.ts` with:
   - Entity interfaces based on user's models
   - Props interfaces for all planned components
   - API request/response types (if API needed)
   - Context state/action types (if state management needed)
   - Utility function types

**Example:**
```tsx
// types/index.ts
export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
  role: UserRole;
  bio?: string;
  createdAt: Date;
  updatedAt: Date;
}

export type UserRole = 'admin' | 'user' | 'guest';

export interface UserProfileProps {
  userId: string;
  onUpdate?: (user: User) => void;
  editable?: boolean;
}

export interface UserFormData {
  name: string;
  email: string;
  bio?: string;
  avatar?: File;
}

// API types (if API needed)
export interface UpdateUserRequest {
  name?: string;
  email?: string;
  bio?: string;
}

export interface UpdateUserResponse {
  user: User;
  message: string;
}

// Context types (if state management needed)
export interface UserProfileState {
  user: User | null;
  isEditing: boolean;
  isSaving: boolean;
  error: string | null;
}

export type UserProfileAction =
  | { type: 'SET_USER'; payload: User }
  | { type: 'START_EDITING' }
  | { type: 'CANCEL_EDITING' }
  | { type: 'SAVE_START' }
  | { type: 'SAVE_SUCCESS'; payload: User }
  | { type: 'SAVE_ERROR'; payload: string };
```

### Phase 3: Tests First (RED)

**Write FAILING tests for each layer BEFORE implementation:**

#### A. Service Tests (if API needed)

```tsx
// services/userProfileService.test.ts
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { fetchUser, updateUser } from './userProfileService';
import type { User, UpdateUserRequest } from '../types';

describe('userProfileService', () => {
  beforeEach(() => {
    global.fetch = vi.fn();
  });

  describe('fetchUser', () => {
    it('fetches user by id successfully', async () => {
      const mockUser: User = {
        id: '1',
        name: 'John Doe',
        email: 'john@example.com',
        role: 'user',
        createdAt: new Date(),
        updatedAt: new Date(),
      };

      (global.fetch as any).mockResolvedValueOnce({
        ok: true,
        json: async () => mockUser,
      });

      const result = await fetchUser('1');
      expect(result).toEqual(mockUser);
      expect(global.fetch).toHaveBeenCalledWith('/api/users/1');
    });

    it('throws error when fetch fails', async () => {
      (global.fetch as any).mockResolvedValueOnce({
        ok: false,
        statusText: 'Not Found',
      });

      await expect(fetchUser('1')).rejects.toThrow('Failed to fetch user: Not Found');
    });
  });

  describe('updateUser', () => {
    it('updates user successfully', async () => {
      const mockUser: User = {
        id: '1',
        name: 'Jane Doe',
        email: 'jane@example.com',
        role: 'user',
        createdAt: new Date(),
        updatedAt: new Date(),
      };

      const updateData: UpdateUserRequest = { name: 'Jane Doe' };

      (global.fetch as any).mockResolvedValueOnce({
        ok: true,
        json: async () => mockUser,
      });

      const result = await updateUser('1', updateData);
      expect(result).toEqual(mockUser);
    });
  });
});
```

#### B. Hook Tests

```tsx
// hooks/useUserProfile.test.ts
import { describe, it, expect, vi } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { useUserProfile } from './useUserProfile';
import * as service from '../services/userProfileService';

vi.mock('../services/userProfileService');

describe('useUserProfile', () => {
  it('fetches user on mount', async () => {
    const mockUser = {
      id: '1',
      name: 'John Doe',
      email: 'john@example.com',
      role: 'user' as const,
      createdAt: new Date(),
      updatedAt: new Date(),
    };

    vi.mocked(service.fetchUser).mockResolvedValue(mockUser);

    const { result } = renderHook(() => useUserProfile('1'));

    expect(result.current.loading).toBe(true);
    expect(result.current.user).toBeNull();

    await waitFor(() => expect(result.current.loading).toBe(false));

    expect(result.current.user).toEqual(mockUser);
    expect(result.current.error).toBeNull();
  });

  it('handles fetch error', async () => {
    vi.mocked(service.fetchUser).mockRejectedValue(new Error('Network error'));

    const { result } = renderHook(() => useUserProfile('1'));

    await waitFor(() => expect(result.current.loading).toBe(false));

    expect(result.current.user).toBeNull();
    expect(result.current.error).toBeInstanceOf(Error);
  });
});
```

#### C. Component Tests

```tsx
// components/UserProfile.test.tsx
import { describe, it, expect, vi } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { axe, toHaveNoViolations } from 'jest-axe';
import { UserProfile } from './UserProfile';

expect.extend(toHaveNoViolations);

describe('UserProfile', () => {
  const mockUser = {
    id: '1',
    name: 'John Doe',
    email: 'john@example.com',
    role: 'user' as const,
    createdAt: new Date(),
    updatedAt: new Date(),
  };

  it('renders loading state initially', () => {
    render(<UserProfile userId="1" />);
    expect(screen.getByText(/loading/i)).toBeInTheDocument();
  });

  it('renders user profile when loaded', async () => {
    render(<UserProfile userId="1" />);

    await waitFor(() => {
      expect(screen.getByText('John Doe')).toBeInTheDocument();
      expect(screen.getByText('john@example.com')).toBeInTheDocument();
    });
  });

  it('allows editing profile when editable', async () => {
    const user = userEvent.setup();
    const onUpdate = vi.fn();

    render(<UserProfile userId="1" editable onUpdate={onUpdate} />);

    await waitFor(() => expect(screen.getByText('John Doe')).toBeInTheDocument());

    const editButton = screen.getByRole('button', { name: /edit/i });
    await user.click(editButton);

    const nameInput = screen.getByLabelText(/name/i);
    await user.clear(nameInput);
    await user.type(nameInput, 'Jane Doe');

    const saveButton = screen.getByRole('button', { name: /save/i });
    await user.click(saveButton);

    await waitFor(() => {
      expect(onUpdate).toHaveBeenCalledWith(expect.objectContaining({
        name: 'Jane Doe',
      }));
    });
  });

  it('should have no accessibility violations', async () => {
    const { container } = render(<UserProfile userId="1" />);
    await waitFor(() => expect(screen.queryByText(/loading/i)).not.toBeInTheDocument());
    const results = await axe(container);
    expect(results).toHaveNoViolations();
  });

  it('handles error state', async () => {
    // Mock service to throw error
    render(<UserProfile userId="invalid" />);

    await waitFor(() => {
      expect(screen.getByText(/error/i)).toBeInTheDocument();
    });
  });
});
```

#### D. Context Tests (if state management needed)

```tsx
// context/UserProfileContext.test.tsx
import { describe, it, expect } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { UserProfileProvider, useUserProfileContext } from './UserProfileContext';
import type { User } from '../types';

describe('UserProfileContext', () => {
  const mockUser: User = {
    id: '1',
    name: 'John Doe',
    email: 'john@example.com',
    role: 'user',
    createdAt: new Date(),
    updatedAt: new Date(),
  };

  it('provides initial state', () => {
    const { result } = renderHook(() => useUserProfileContext(), {
      wrapper: UserProfileProvider,
    });

    expect(result.current.state.user).toBeNull();
    expect(result.current.state.isEditing).toBe(false);
    expect(result.current.state.isSaving).toBe(false);
  });

  it('handles SET_USER action', () => {
    const { result } = renderHook(() => useUserProfileContext(), {
      wrapper: UserProfileProvider,
    });

    act(() => {
      result.current.dispatch({ type: 'SET_USER', payload: mockUser });
    });

    expect(result.current.state.user).toEqual(mockUser);
  });

  it('handles editing workflow', () => {
    const { result } = renderHook(() => useUserProfileContext(), {
      wrapper: UserProfileProvider,
    });

    act(() => {
      result.current.dispatch({ type: 'START_EDITING' });
    });
    expect(result.current.state.isEditing).toBe(true);

    act(() => {
      result.current.dispatch({ type: 'CANCEL_EDITING' });
    });
    expect(result.current.state.isEditing).toBe(false);
  });

  it('throws error when used outside provider', () => {
    expect(() => {
      renderHook(() => useUserProfileContext());
    }).toThrow('useUserProfileContext must be used within UserProfileProvider');
  });
});
```

#### E. Utility Tests

```tsx
// utils/formatters.test.ts
import { describe, it, expect } from 'vitest';
import { formatUserName, getUserInitials } from './formatters';
import type { User } from '../types';

describe('formatters', () => {
  const mockUser: User = {
    id: '1',
    name: 'John Doe',
    email: 'john@example.com',
    role: 'user',
    createdAt: new Date(),
    updatedAt: new Date(),
  };

  describe('formatUserName', () => {
    it('formats user name with email', () => {
      const result = formatUserName(mockUser);
      expect(result).toBe('John Doe (john@example.com)');
    });
  });

  describe('getUserInitials', () => {
    it('returns initials from name', () => {
      const result = getUserInitials(mockUser);
      expect(result).toBe('JD');
    });

    it('handles single name', () => {
      const singleNameUser = { ...mockUser, name: 'Madonna' };
      const result = getUserInitials(singleNameUser);
      expect(result).toBe('M');
    });
  });
});
```

**ALL TESTS MUST FAIL AT THIS POINT** - This is the RED phase of TDD.

### Phase 4: Implementation (GREEN)

**Now implement ONLY enough code to make tests pass:**

#### A. Services (if API needed)

```tsx
// services/userProfileService.ts
import type { User, UpdateUserRequest } from '../types';

export async function fetchUser(userId: string): Promise<User> {
  const response = await fetch(`/api/users/${userId}`);

  if (!response.ok) {
    throw new Error(`Failed to fetch user: ${response.statusText}`);
  }

  return response.json();
}

export async function updateUser(
  userId: string,
  data: UpdateUserRequest
): Promise<User> {
  const response = await fetch(`/api/users/${userId}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error(`Failed to update user: ${response.statusText}`);
  }

  return response.json();
}
```

#### B. Hooks

```tsx
// hooks/useUserProfile.ts
import { useState, useEffect } from 'react';
import { fetchUser } from '../services/userProfileService';
import type { User } from '../types';

export function useUserProfile(userId: string) {
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

#### C. Context (if state management needed)

```tsx
// context/UserProfileContext.tsx
import { createContext, useContext, useReducer, type ReactNode } from 'react';
import type { UserProfileState, UserProfileAction, User } from '../types';

const UserProfileContext = createContext<{
  state: UserProfileState;
  dispatch: React.Dispatch<UserProfileAction>;
} | undefined>(undefined);

const initialState: UserProfileState = {
  user: null,
  isEditing: false,
  isSaving: false,
  error: null,
};

function userProfileReducer(
  state: UserProfileState,
  action: UserProfileAction
): UserProfileState {
  switch (action.type) {
    case 'SET_USER':
      return { ...state, user: action.payload, error: null };
    case 'START_EDITING':
      return { ...state, isEditing: true };
    case 'CANCEL_EDITING':
      return { ...state, isEditing: false };
    case 'SAVE_START':
      return { ...state, isSaving: true, error: null };
    case 'SAVE_SUCCESS':
      return { ...state, user: action.payload, isSaving: false, isEditing: false };
    case 'SAVE_ERROR':
      return { ...state, error: action.payload, isSaving: false };
    default:
      return state;
  }
}

export function UserProfileProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(userProfileReducer, initialState);

  return (
    <UserProfileContext.Provider value={{ state, dispatch }}>
      {children}
    </UserProfileContext.Provider>
  );
}

export function useUserProfileContext() {
  const context = useContext(UserProfileContext);
  if (!context) {
    throw new Error('useUserProfileContext must be used within UserProfileProvider');
  }
  return context;
}
```

#### D. Components

```tsx
// components/UserProfile.tsx
import { useUserProfile } from '../hooks/useUserProfile';
import type { UserProfileProps } from '../types';

export function UserProfile({ userId, editable = false, onUpdate }: UserProfileProps) {
  const { user, loading, error } = useUserProfile(userId);
  const [isEditing, setIsEditing] = useState(false);

  if (loading) {
    return (
      <div role="status" aria-live="polite">
        Loading user profile...
      </div>
    );
  }

  if (error) {
    return (
      <div role="alert" className="text-red-600">
        Error loading user profile: {error.message}
      </div>
    );
  }

  if (!user) {
    return <div>User not found</div>;
  }

  const handleEdit = () => setIsEditing(true);
  const handleCancel = () => setIsEditing(false);
  const handleSave = async (data: UserFormData) => {
    // Save logic here
    onUpdate?.(updatedUser);
    setIsEditing(false);
  };

  return (
    <article className="user-profile" aria-labelledby="profile-heading">
      <header>
        <h1 id="profile-heading">{user.name}</h1>
        <p className="text-gray-600">{user.email}</p>
      </header>

      {isEditing ? (
        <UserProfileForm user={user} onSave={handleSave} onCancel={handleCancel} />
      ) : (
        <>
          {user.bio && <p>{user.bio}</p>}
          {editable && (
            <button
              onClick={handleEdit}
              className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              aria-label="Edit profile"
            >
              Edit Profile
            </button>
          )}
        </>
      )}
    </article>
  );
}
```

#### E. Utils

```tsx
// utils/formatters.ts
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

**Run tests - they should now PASS (GREEN phase).**

### Phase 5: Storybook Stories

**Generate stories for visual development and documentation:**

```tsx
// components/UserProfile.stories.tsx
import type { Meta, StoryObj } from '@storybook/react';
import { UserProfile } from './UserProfile';

const meta = {
  title: 'Features/UserProfile/UserProfile',
  component: UserProfile,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof UserProfile>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    userId: '1',
  },
};

export const Editable: Story = {
  args: {
    userId: '1',
    editable: true,
    onUpdate: (user) => console.log('Updated:', user),
  },
};

export const Loading: Story = {
  parameters: {
    mockData: {
      loading: true,
    },
  },
  args: {
    userId: '1',
  },
};

export const Error: Story = {
  parameters: {
    mockData: {
      error: new Error('Failed to load user'),
    },
  },
  args: {
    userId: 'invalid',
  },
};
```

### Phase 6: Public API

**Create index.ts to expose feature's public interface:**

```tsx
// index.ts
// Components
export { UserProfile } from './components/UserProfile';
export { UserAvatar } from './components/UserAvatar';
export { UserProfileForm } from './components/UserProfileForm';

// Hooks
export { useUserProfile } from './hooks/useUserProfile';
export { useUserProfileForm } from './hooks/useUserProfileForm';

// Context (if exists)
export { UserProfileProvider, useUserProfileContext } from './context/UserProfileContext';

// Services (if exists)
export { fetchUser, updateUser } from './services/userProfileService';

// Types
export type {
  User,
  UserRole,
  UserProfileProps,
  UserFormData,
  UserProfileState,
  UserProfileAction,
} from './types';
```

### Phase 7: Feature README

**Generate comprehensive feature documentation:**

```markdown
# [Feature Name] Feature

## Overview

[Brief description of what this feature does]

## Components

### [MainComponent]

[Description]

**Props:**
- `prop1` (type): Description
- `prop2` (type): Description

**Usage:**
\`\`\`tsx
import { MainComponent } from '@/features/[feature-name]';

<MainComponent prop1="value" prop2={value} />
\`\`\`

## Hooks

### use[Feature]

[Description]

**Parameters:**
- `param1` (type): Description

**Returns:**
- `value1` (type): Description
- `value2` (type): Description

**Usage:**
\`\`\`tsx
const { value1, value2 } = use[Feature](param1);
\`\`\`

## State Management

[If context exists, document the state and actions]

## API Integration

[If services exist, document the endpoints and methods]

## Types

[List main types and their purpose]

## Testing

Run tests:
\`\`\`bash
npm test src/features/[feature-name]
\`\`\`

Coverage: [X]%

## Accessibility

- WCAG 2.1 AA compliant
- Keyboard navigation supported
- Screen reader friendly
- All interactive elements have proper ARIA labels

## Storybook

View stories:
\`\`\`bash
npm run storybook
\`\`\`

Navigate to: Features/[FeatureName]

## Dependencies

- Internal: [List internal dependencies]
- External: [List external dependencies]

## Future Enhancements

- [ ] Feature 1
- [ ] Feature 2
```

**Code Standards to Follow:**

Reference `/Users/solrac97gr/Development/personal/marketplace-plugins/plugins/react-dev/CODE_STANDARDS.md` for:

- ✅ Functional components only (no class components)
- ✅ TypeScript strict mode (no `any` types)
- ✅ Named exports only (no default exports)
- ✅ Props interfaces named and exported
- ✅ Hooks follow rules (proper dependencies)
- ✅ 80%+ component test coverage
- ✅ 100% hook test coverage
- ✅ WCAG 2.1 AA accessibility compliance
- ✅ Semantic HTML elements
- ✅ Proper ARIA attributes
- ✅ Keyboard navigation support
- ✅ Focus management
- ✅ Screen reader support
- ✅ Accessibility tests with jest-axe
- ✅ User-centric tests (React Testing Library)
- ✅ Tailwind utility classes
- ✅ Mobile-first responsive design
- ✅ Dark mode support
- ✅ Performance optimizations (React.memo, useMemo, useCallback where appropriate)

**Architecture Patterns to Follow:**

Reference `/Users/solrac97gr/Development/personal/marketplace-plugins/plugins/react-dev/ARCHITECTURE.md` for:

- ✅ Feature-based organization
- ✅ Clear separation of concerns (components/hooks/context/services/types/utils)
- ✅ Dependency flow: Components → Hooks → Services
- ✅ Public API via index.ts
- ✅ Context + useReducer for complex state
- ✅ Pure utility functions
- ✅ Repository pattern for services

**Summary:**

1. **Ask user questions** to gather requirements
2. **Generate types first** (TypeScript interfaces)
3. **Write failing tests** (RED) for all layers
4. **Implement to pass tests** (GREEN)
5. **Generate Storybook stories** for visual testing
6. **Create public API** (index.ts)
7. **Write feature README** with documentation
8. **Ensure 100% compliance** with CODE_STANDARDS.md and ARCHITECTURE.md

**Important Notes:**

- Generate COMPLETE implementation (not stubs or TODOs)
- All tests must be comprehensive and pass
- All code must follow accessibility standards
- All components must have Storybook stories
- Feature must be production-ready
- Code must be well-documented
- Follow TDD strictly: Types → Tests → Implementation → Stories
