---
description: Initialize React + Vite + TypeScript project with feature-based architecture
---

Create a new React project with Vite, TypeScript, and complete feature-based architecture setup.

**Discovery Process:**

First, understand what the user wants to build, then guide them through the setup process with intelligent defaults.

**Project Structure to Create:**

```
project-name/
├── src/
│   ├── features/
│   │   ├── [feature-name]/
│   │   │   ├── components/
│   │   │   ├── hooks/
│   │   │   ├── context/
│   │   │   ├── services/
│   │   │   ├── types/
│   │   │   ├── utils/
│   │   │   └── index.ts
│   │   └── shared/
│   │       ├── components/
│   │       ├── hooks/
│   │       ├── context/
│   │       ├── types/
│   │       └── utils/
│   ├── app/
│   │   ├── App.tsx
│   │   ├── App.test.tsx
│   │   ├── routes.tsx
│   │   └── providers.tsx
│   ├── assets/
│   ├── styles/
│   │   └── globals.css
│   └── main.tsx
├── tests/
│   ├── e2e/
│   ├── integration/
│   └── setup.ts
├── .storybook/
│   ├── main.ts
│   └── preview.ts
├── public/
├── vite.config.ts
├── vitest.config.ts
├── playwright.config.ts
├── tailwind.config.js
├── postcss.config.js
├── tsconfig.json
├── tsconfig.node.json
├── .eslintrc.json
├── .prettierrc
├── .gitignore
├── package.json
└── README.md
```

**Setup Steps:**

1. **Ask the user for project details using AskUserQuestion:**
   - **Project name** (text input, default: "my-react-app")
   - **Include example feature?** (yes/no, default: yes) - Creates a working "counter" feature to demonstrate structure
   - **Routing library** (select one):
     - React Router (v6) - Most popular
     - TanStack Router - Type-safe routing
     - None - For apps that don't need routing
   - **Additional features** (multiselect):
     - Auth scaffold - Adds AuthContext, login/register components
     - Form management - React Hook Form + Zod validation
     - Data fetching - TanStack Query for server state
     - State management - Zustand (alternative to Context API)

2. **Create project directory** and navigate into it

3. **Initialize package.json** with appropriate dependencies based on selections:

   **Core dependencies** (always included):
   ```json
   {
     "dependencies": {
       "react": "^18.3.1",
       "react-dom": "^18.3.1"
     }
   }
   ```

   **Routing dependencies** (based on selection):
   - React Router: `"react-router-dom": "^6.26.0"`
   - TanStack Router: `"@tanstack/react-router": "^1.58.0"`

   **Additional feature dependencies**:
   - Form management: `"react-hook-form": "^7.53.0"`, `"zod": "^3.23.0"`, `"@hookform/resolvers": "^3.9.0"`
   - Data fetching: `"@tanstack/react-query": "^5.56.0"`
   - State management: `"zustand": "^4.5.0"`

   **DevDependencies** (always included):
   ```json
   {
     "@types/react": "^18.3.0",
     "@types/react-dom": "^18.3.0",
     "@vitejs/plugin-react": "^4.3.0",
     "typescript": "^5.5.0",
     "vite": "^5.4.0",
     "vitest": "^2.1.0",
     "@testing-library/react": "^16.0.0",
     "@testing-library/user-event": "^14.5.0",
     "@testing-library/jest-dom": "^6.5.0",
     "@playwright/test": "^1.47.0",
     "tailwindcss": "^3.4.0",
     "postcss": "^8.4.0",
     "autoprefixer": "^10.4.0",
     "storybook": "^8.3.0",
     "@storybook/react-vite": "^8.3.0",
     "@storybook/addon-essentials": "^8.3.0",
     "@storybook/addon-interactions": "^8.3.0",
     "@storybook/addon-links": "^8.3.0",
     "@storybook/blocks": "^8.3.0",
     "eslint": "^9.11.0",
     "eslint-plugin-react": "^7.36.0",
     "eslint-plugin-react-hooks": "^5.1.0",
     "@typescript-eslint/eslint-plugin": "^8.6.0",
     "@typescript-eslint/parser": "^8.6.0",
     "prettier": "^3.3.0",
     "prettier-plugin-tailwindcss": "^0.6.6",
     "jest-axe": "^9.0.0",
     "@axe-core/react": "^4.10.0"
   }
   ```

   **Scripts** (add to package.json):
   ```json
   {
     "scripts": {
       "dev": "vite",
       "build": "tsc && vite build",
       "preview": "vite preview",
       "test": "vitest",
       "test:ui": "vitest --ui",
       "test:coverage": "vitest --coverage",
       "test:e2e": "playwright test",
       "test:e2e:ui": "playwright test --ui",
       "storybook": "storybook dev -p 6006",
       "build-storybook": "storybook build",
       "lint": "eslint . --ext ts,tsx",
       "lint:fix": "eslint . --ext ts,tsx --fix",
       "format": "prettier --write \"src/**/*.{ts,tsx,css}\"",
       "type-check": "tsc --noEmit"
     }
   }
   ```

4. **Create TypeScript configuration files:**

   **tsconfig.json**:
   ```json
   {
     "compilerOptions": {
       "target": "ES2020",
       "useDefineForClassFields": true,
       "lib": ["ES2020", "DOM", "DOM.Iterable"],
       "module": "ESNext",
       "skipLibCheck": true,
       "moduleResolution": "bundler",
       "allowImportingTsExtensions": true,
       "resolveJsonModule": true,
       "isolatedModules": true,
       "noEmit": true,
       "jsx": "react-jsx",
       "strict": true,
       "noUnusedLocals": true,
       "noUnusedParameters": true,
       "noFallthroughCasesInSwitch": true,
       "baseUrl": ".",
       "paths": {
         "@/*": ["./src/*"],
         "@/features/*": ["./src/features/*"],
         "@/shared/*": ["./src/features/shared/*"],
         "@/app/*": ["./src/app/*"]
       }
     },
     "include": ["src"],
     "references": [{ "path": "./tsconfig.node.json" }]
   }
   ```

   **tsconfig.node.json**:
   ```json
   {
     "compilerOptions": {
       "composite": true,
       "skipLibCheck": true,
       "module": "ESNext",
       "moduleResolution": "bundler",
       "allowSyntheticDefaultImports": true
     },
     "include": ["vite.config.ts", "vitest.config.ts", "playwright.config.ts"]
   }
   ```

5. **Create Vite configuration (vite.config.ts)**:
   ```typescript
   import { defineConfig } from 'vite';
   import react from '@vitejs/plugin-react';
   import path from 'path';

   export default defineConfig({
     plugins: [react()],
     resolve: {
       alias: {
         '@': path.resolve(__dirname, './src'),
         '@/features': path.resolve(__dirname, './src/features'),
         '@/shared': path.resolve(__dirname, './src/features/shared'),
         '@/app': path.resolve(__dirname, './src/app'),
       },
     },
     server: {
       port: 3000,
     },
   });
   ```

6. **Create Vitest configuration (vitest.config.ts)**:
   ```typescript
   import { defineConfig } from 'vitest/config';
   import react from '@vitejs/plugin-react';
   import path from 'path';

   export default defineConfig({
     plugins: [react()],
     test: {
       globals: true,
       environment: 'jsdom',
       setupFiles: './tests/setup.ts',
       coverage: {
         provider: 'v8',
         reporter: ['text', 'json', 'html'],
         exclude: [
           'node_modules/',
           'tests/',
           '**/*.stories.tsx',
           '**/*.test.tsx',
           '.storybook/',
         ],
       },
     },
     resolve: {
       alias: {
         '@': path.resolve(__dirname, './src'),
         '@/features': path.resolve(__dirname, './src/features'),
         '@/shared': path.resolve(__dirname, './src/features/shared'),
         '@/app': path.resolve(__dirname, './src/app'),
       },
     },
   });
   ```

7. **Create Playwright configuration (playwright.config.ts)**:
   ```typescript
   import { defineConfig, devices } from '@playwright/test';

   export default defineConfig({
     testDir: './tests/e2e',
     fullyParallel: true,
     forbidOnly: !!process.env.CI,
     retries: process.env.CI ? 2 : 0,
     workers: process.env.CI ? 1 : undefined,
     reporter: 'html',
     use: {
       baseURL: 'http://localhost:3000',
       trace: 'on-first-retry',
     },
     projects: [
       {
         name: 'chromium',
         use: { ...devices['Desktop Chrome'] },
       },
       {
         name: 'firefox',
         use: { ...devices['Desktop Firefox'] },
       },
       {
         name: 'webkit',
         use: { ...devices['Desktop Safari'] },
       },
     ],
     webServer: {
       command: 'npm run dev',
       url: 'http://localhost:3000',
       reuseExistingServer: !process.env.CI,
     },
   });
   ```

8. **Create Tailwind configuration (tailwind.config.js)**:
   ```javascript
   /** @type {import('tailwindcss').Config} */
   export default {
     content: [
       "./index.html",
       "./src/**/*.{js,ts,jsx,tsx}",
     ],
     theme: {
       extend: {
         colors: {
           brand: {
             50: '#f0f9ff',
             100: '#e0f2fe',
             200: '#bae6fd',
             300: '#7dd3fc',
             400: '#38bdf8',
             500: '#0ea5e9',
             600: '#0284c7',
             700: '#0369a1',
             800: '#075985',
             900: '#0c4a6e',
           },
         },
       },
     },
     plugins: [],
   }
   ```

9. **Create PostCSS configuration (postcss.config.js)**:
   ```javascript
   export default {
     plugins: {
       tailwindcss: {},
       autoprefixer: {},
     },
   }
   ```

10. **Create ESLint configuration (.eslintrc.json)**:
    ```json
    {
      "extends": [
        "eslint:recommended",
        "plugin:@typescript-eslint/recommended",
        "plugin:react/recommended",
        "plugin:react-hooks/recommended"
      ],
      "parser": "@typescript-eslint/parser",
      "parserOptions": {
        "ecmaVersion": "latest",
        "sourceType": "module",
        "ecmaFeatures": {
          "jsx": true
        }
      },
      "plugins": ["@typescript-eslint", "react", "react-hooks"],
      "rules": {
        "react/react-in-jsx-scope": "off",
        "react/prop-types": "off",
        "@typescript-eslint/no-explicit-any": "error",
        "@typescript-eslint/explicit-module-boundary-types": "off",
        "react-hooks/rules-of-hooks": "error",
        "react-hooks/exhaustive-deps": "warn"
      },
      "settings": {
        "react": {
          "version": "detect"
        }
      }
    }
    ```

11. **Create Prettier configuration (.prettierrc)**:
    ```json
    {
      "semi": true,
      "trailingComma": "es5",
      "singleQuote": true,
      "printWidth": 100,
      "tabWidth": 2,
      "useTabs": false,
      "arrowParens": "always",
      "plugins": ["prettier-plugin-tailwindcss"]
    }
    ```

12. **Create Storybook configuration:**

    **.storybook/main.ts**:
    ```typescript
    import type { StorybookConfig } from '@storybook/react-vite';

    const config: StorybookConfig = {
      stories: ['../src/**/*.stories.@(js|jsx|ts|tsx)'],
      addons: [
        '@storybook/addon-links',
        '@storybook/addon-essentials',
        '@storybook/addon-interactions',
      ],
      framework: {
        name: '@storybook/react-vite',
        options: {},
      },
      docs: {
        autodocs: 'tag',
      },
    };

    export default config;
    ```

    **.storybook/preview.ts**:
    ```typescript
    import type { Preview } from '@storybook/react';
    import '../src/styles/globals.css';

    const preview: Preview = {
      parameters: {
        actions: { argTypesRegex: '^on[A-Z].*' },
        controls: {
          matchers: {
            color: /(background|color)$/i,
            date: /Date$/,
          },
        },
      },
    };

    export default preview;
    ```

13. **Create test setup file (tests/setup.ts)**:
    ```typescript
    import '@testing-library/jest-dom';
    import { expect, afterEach } from 'vitest';
    import { cleanup } from '@testing-library/react';
    import * as matchers from '@testing-library/jest-dom/matchers';

    expect.extend(matchers);

    afterEach(() => {
      cleanup();
    });
    ```

14. **Create directory structure:**
    ```
    src/features/
    src/features/shared/components/
    src/features/shared/hooks/
    src/features/shared/context/
    src/features/shared/types/
    src/features/shared/utils/
    src/app/
    src/assets/
    src/styles/
    tests/e2e/
    tests/integration/
    public/
    ```

15. **Create global styles (src/styles/globals.css)**:
    ```css
    @tailwind base;
    @tailwind components;
    @tailwind utilities;

    @layer base {
      * {
        @apply border-border;
      }
      body {
        @apply bg-background text-foreground;
      }
    }
    ```

16. **Create main entry point (src/main.tsx)**:
    ```tsx
    import React from 'react';
    import ReactDOM from 'react-dom/client';
    import { App } from './app/App';
    import './styles/globals.css';

    ReactDOM.createRoot(document.getElementById('root')!).render(
      <React.StrictMode>
        <App />
      </React.StrictMode>
    );
    ```

17. **Create index.html**:
    ```html
    <!doctype html>
    <html lang="en">
      <head>
        <meta charset="UTF-8" />
        <link rel="icon" type="image/svg+xml" href="/vite.svg" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>{Project Name}</title>
      </head>
      <body>
        <div id="root"></div>
        <script type="module" src="/src/main.tsx"></script>
      </body>
    </html>
    ```

18. **Create App component and providers:**

    **src/app/App.tsx** (basic version without routing):
    ```tsx
    import { Providers } from './providers';

    export function App() {
      return (
        <Providers>
          <div className="min-h-screen bg-gray-50">
            <header className="bg-white shadow">
              <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
                <h1 className="text-3xl font-bold tracking-tight text-gray-900">
                  {Project Name}
                </h1>
              </div>
            </header>
            <main className="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
              {/* Content will go here */}
            </main>
          </div>
        </Providers>
      );
    }
    ```

    **src/app/App.test.tsx**:
    ```tsx
    import { render, screen } from '@testing-library/react';
    import { App } from './App';

    describe('App', () => {
      it('renders the app title', () => {
        render(<App />);
        expect(screen.getByText('{Project Name}')).toBeInTheDocument();
      });
    });
    ```

    **src/app/providers.tsx**:
    ```tsx
    import { ReactNode } from 'react';

    interface ProvidersProps {
      children: ReactNode;
    }

    export function Providers({ children }: ProvidersProps) {
      return <>{children}</>;
    }
    ```

19. **If routing is selected, create routes:**

    **For React Router:**
    - Create `src/app/routes.tsx` with BrowserRouter setup
    - Update App.tsx to use Routes
    - Create example route components

    **For TanStack Router:**
    - Create `src/app/routes.tsx` with Router and Route definitions
    - Update App.tsx to use RouterProvider
    - Create example route components

20. **If "Include example feature" is selected, create a counter feature:**

    Create complete counter feature with:
    - `src/features/counter/components/Counter.tsx` - Main component
    - `src/features/counter/components/Counter.test.tsx` - Unit tests
    - `src/features/counter/components/Counter.stories.tsx` - Storybook stories
    - `src/features/counter/hooks/useCounter.ts` - Custom hook with logic
    - `src/features/counter/hooks/useCounter.test.ts` - Hook tests
    - `src/features/counter/types/index.ts` - TypeScript types
    - `src/features/counter/index.ts` - Public API exports

    The counter should demonstrate:
    - Feature-based architecture
    - TypeScript strict mode
    - Custom hooks
    - Component testing
    - Storybook integration
    - Tailwind styling
    - Accessibility (proper buttons, ARIA labels)

21. **If additional features are selected:**

    **Auth scaffold:**
    - Create `src/features/shared/context/AuthContext.tsx`
    - Create `src/features/auth/` with login/register components
    - Add auth types and hooks
    - Update providers.tsx to include AuthProvider

    **Form management:**
    - Create example form component using react-hook-form + zod
    - Add form utilities in shared

    **Data fetching:**
    - Add QueryClient setup in providers.tsx
    - Create example hook using useQuery
    - Add query utilities in shared

    **State management:**
    - Create example Zustand store
    - Add store utilities in shared

22. **Create .gitignore**:
    ```
    # Dependencies
    node_modules/

    # Production
    dist/
    build/

    # Testing
    coverage/
    playwright-report/
    test-results/

    # Storybook
    storybook-static/

    # Environment
    .env
    .env.local
    .env.*.local

    # Editor
    .vscode/
    .idea/

    # OS
    .DS_Store
    Thumbs.db

    # Logs
    *.log
    npm-debug.log*
    ```

23. **Create comprehensive README.md** with:
    - Project title and description
    - Architecture overview (feature-based)
    - Tech stack details
    - Getting started instructions
    - Available scripts explanation
    - Project structure overview
    - Testing strategy
    - Development workflow
    - Deployment notes
    - Link to ARCHITECTURE.md and CODE_STANDARDS.md (from plugin docs)

24. **Initialize git repository:**
    ```bash
    git init
    git add .
    git commit -m "Initial commit: React + Vite + TypeScript project with feature-based architecture"
    ```

25. **Install dependencies:**
    ```bash
    npm install
    ```

26. **Run initial build and tests to verify setup:**
    ```bash
    npm run type-check
    npm run lint
    npm run test
    ```

27. **Provide next steps to the user:**
    ```
    Project created successfully!

    Next steps:
    1. cd {project-name}
    2. npm run dev - Start development server
    3. npm run test - Run unit tests
    4. npm run storybook - Open Storybook
    5. npm run test:e2e - Run E2E tests

    To add a new feature:
    Use the /new-feature skill to create a complete feature with TDD workflow

    To review code quality:
    Use the /review-code skill

    To run architecture validation:
    Use the /review-arch skill

    Project structure follows feature-based architecture.
    See ARCHITECTURE.md in the react-dev plugin for detailed guidelines.
    ```

**Important Notes:**

- All components use functional components (no class components)
- TypeScript strict mode enabled
- Named exports only (no default exports)
- Accessibility built-in (WCAG 2.1 AA compliance)
- Mobile-first responsive design with Tailwind
- Test coverage: 80%+ components, 100% hooks
- All generated code follows CODE_STANDARDS.md from react-dev plugin
- Feature isolation - each feature has its own directory with clear boundaries
- Public API pattern - features expose only what's needed via index.ts

**Architecture Principles:**

- Feature-based organization (not layer-based)
- Each feature is self-contained with domain/application/infrastructure concerns
- Shared code only for truly generic utilities (3+ feature usage)
- Dependency flow: Components → Hooks → Services
- Type safety: No `any` types, strict TypeScript
- Testing: User-centric tests with React Testing Library
- Performance: Lazy loading, memoization where appropriate
- Accessibility: Semantic HTML, ARIA, keyboard navigation

Be thorough and create a production-ready structure that follows all best practices from ARCHITECTURE.md and CODE_STANDARDS.md.
