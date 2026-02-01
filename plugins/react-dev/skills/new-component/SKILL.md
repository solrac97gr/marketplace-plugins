---
description: Create a React component with tests and Storybook story
---

Create a new React component following best practices with comprehensive tests and Storybook stories.

**What to Create:**

1. **Ask the user:**
   - Component name (e.g., "Button", "UserCard", "LoginForm")
   - Which feature does it belong to? (or "shared" for shared components)
   - Component type:
     - **Presentational**: UI-only, no business logic (e.g., Button, Card)
     - **Container**: Connects to state/context (e.g., UserProfileContainer)
     - **Form**: Handles form logic (e.g., LoginForm, ContactForm)
   - Props needed (name, type, required/optional, default values)
   - Does it need variants? (e.g., primary/secondary/outline for Button)
   - Any specific accessibility requirements beyond WCAG 2.1 AA?

2. **Generate Component** in `src/features/[feature]/components/[ComponentName].tsx` (or `src/components/[ComponentName].tsx` for shared):
   - Use functional component
   - Export named interface `[ComponentName]Props`
   - Include comprehensive TypeScript typing
   - Use Tailwind CSS utility classes
   - Include proper semantic HTML
   - Add ARIA attributes for accessibility
   - Support keyboard navigation
   - Support responsive design (mobile-first)
   - Support dark mode with `dark:` classes
   - Add focus indicators
   - Ensure WCAG 2.1 AA compliance
   - For variants: use discriminated unions or simple string literals
   - Add JSDoc comments for complex props

3. **Generate Tests** in `src/features/[feature]/components/[ComponentName].test.tsx`:
   - Import from `@testing-library/react` and `@testing-library/user-event`
   - **Rendering tests**: Component renders correctly with different props
   - **User interaction tests**: Click, type, keyboard navigation (using userEvent)
   - **Accessibility tests**:
     - Use `jest-axe` for automated a11y testing
     - Test keyboard navigation (Tab, Enter, Space, Arrow keys)
     - Test ARIA attributes
     - Test focus management
   - **Prop validation tests**: Required props, optional props, defaults
   - **Variant tests**: All variants render correctly (if applicable)
   - Use `describe` blocks to organize tests
   - Use `getByRole`, `getByLabelText`, `getByText` (avoid `getByTestId`)
   - Follow user-centric testing (test behavior, not implementation)
   - Aim for 80%+ coverage

4. **Generate Storybook Story** in `src/features/[feature]/components/[ComponentName].stories.tsx`:
   - Import from `@storybook/react`
   - Export default meta object with title, component, tags
   - Add `argTypes` for interactive controls
   - Create stories using CSF3 format (Component Story Format 3)
   - **Default story**: Standard usage
   - **All variants**: One story per variant (if applicable)
   - **States**: Loading, disabled, error, etc.
   - **Interactive controls**: Enable all props in Storybook controls
   - **Responsive**: Show mobile/tablet/desktop if relevant
   - **Dark mode**: Include dark mode variant if applicable
   - Add decorators if needed (e.g., padding, background)

**Component Structure Pattern:**

```tsx
// ComponentName.tsx
interface ComponentNameProps {
  /** Description of prop */
  propName: string;
  onClick?: () => void;
  variant?: 'primary' | 'secondary' | 'outline';
  disabled?: boolean;
  children?: React.ReactNode;
}

export function ComponentName({
  propName,
  onClick,
  variant = 'primary',
  disabled = false,
  children
}: ComponentNameProps) {
  const baseClasses = 'base-utility-classes';
  const variantClasses = {
    primary: 'variant-specific-classes',
    secondary: 'variant-specific-classes',
    outline: 'variant-specific-classes'
  };

  return (
    <element
      className={`${baseClasses} ${variantClasses[variant]}`}
      onClick={onClick}
      disabled={disabled}
      aria-label="Descriptive label"
      // More ARIA and semantic attributes
    >
      {children}
    </element>
  );
}
```

**Test Structure Pattern:**

```tsx
// ComponentName.test.tsx
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { axe, toHaveNoViolations } from 'jest-axe';
import { ComponentName } from './ComponentName';

expect.extend(toHaveNoViolations);

describe('ComponentName', () => {
  describe('Rendering', () => {
    it('renders with required props', () => {
      render(<ComponentName propName="value" />);
      expect(screen.getByRole('...')).toBeInTheDocument();
    });

    it('renders all variants', () => {
      // Test each variant
    });
  });

  describe('User Interactions', () => {
    it('handles click events', async () => {
      const user = userEvent.setup();
      const onClick = vi.fn();

      render(<ComponentName propName="value" onClick={onClick} />);
      await user.click(screen.getByRole('button'));

      expect(onClick).toHaveBeenCalled();
    });

    it('supports keyboard navigation', async () => {
      const user = userEvent.setup();
      // Test Tab, Enter, Space, Arrow keys
    });
  });

  describe('Accessibility', () => {
    it('has no accessibility violations', async () => {
      const { container } = render(<ComponentName propName="value" />);
      const results = await axe(container);
      expect(results).toHaveNoViolations();
    });

    it('has proper ARIA attributes', () => {
      // Test specific ARIA attributes
    });
  });

  describe('Props', () => {
    it('applies default values', () => {
      // Test defaults
    });

    it('handles optional props', () => {
      // Test optional props
    });
  });
});
```

**Storybook Story Structure Pattern:**

```tsx
// ComponentName.stories.tsx
import type { Meta, StoryObj } from '@storybook/react';
import { ComponentName } from './ComponentName';

const meta: Meta<typeof ComponentName> = {
  title: 'Features/[Feature]/ComponentName',
  component: ComponentName,
  tags: ['autodocs'],
  argTypes: {
    variant: {
      control: 'select',
      options: ['primary', 'secondary', 'outline'],
    },
    disabled: {
      control: 'boolean',
    },
  },
};

export default meta;
type Story = StoryObj<typeof ComponentName>;

export const Default: Story = {
  args: {
    propName: 'Default value',
    variant: 'primary',
  },
};

export const Primary: Story = {
  args: {
    propName: 'Primary',
    variant: 'primary',
  },
};

export const Secondary: Story = {
  args: {
    propName: 'Secondary',
    variant: 'secondary',
  },
};

export const Disabled: Story = {
  args: {
    propName: 'Disabled',
    disabled: true,
  },
};

export const DarkMode: Story = {
  args: {
    propName: 'Dark Mode',
  },
  parameters: {
    backgrounds: { default: 'dark' },
  },
};
```

**Code Standards Compliance:**

Follow all standards from CODE_STANDARDS.md:

- **React Standards**: Functional components, named exports, hooks rules
- **TypeScript Standards**: Strict mode, no `any`, proper typing
- **Testing Standards**: User-centric tests, accessibility tests, 80%+ coverage
- **Accessibility Standards**: WCAG 2.1 AA, semantic HTML, ARIA, keyboard navigation
- **Tailwind Standards**: Utility-first, mobile-first, dark mode support
- **Naming Conventions**: PascalCase for components, camelCase for handlers

**File Organization:**

For feature components:
```
src/features/[feature]/components/
├── ComponentName.tsx
├── ComponentName.test.tsx
└── ComponentName.stories.tsx
```

For shared components:
```
src/components/
├── ComponentName.tsx
├── ComponentName.test.tsx
└── ComponentName.stories.tsx
```

**After Creation:**

1. Verify all files are created in correct locations
2. Run tests: `npm test ComponentName.test.tsx`
3. Run Storybook: `npm run storybook` (if available)
4. Verify accessibility with `jest-axe`
5. Check TypeScript compilation: `npm run type-check`

Be concise and follow React best practices. Ask clarifying questions if requirements are unclear.
