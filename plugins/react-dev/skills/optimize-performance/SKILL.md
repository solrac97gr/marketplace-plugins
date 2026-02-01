---
description: Analyze and optimize React application performance
---

Analyze and optimize React application performance with comprehensive checks for re-renders, memoization, bundle size, and code splitting.

**What to Analyze:**

1. **Ask the user:**
   - Analysis scope:
     - **Entire project**: Analyze all components and features
     - **Specific feature**: Analyze a specific feature directory
     - **Specific component**: Analyze a single component or component tree
   - Run bundle analysis? (yes/no)
     - If yes, suggest installing `vite-bundle-visualizer` if not present
   - Focus areas (select all that apply or "all"):
     - **re-renders**: Unnecessary re-render detection
     - **bundle size**: Bundle size analysis and optimization
     - **lazy loading**: Code splitting and lazy loading opportunities
     - **memoization**: useMemo/useCallback optimization
     - **all**: Comprehensive analysis of all areas

2. **Analysis Workflow:**

   **Phase 1: Component Analysis**
   - Scan components for performance anti-patterns
   - Identify components with expensive rendering logic
   - Check for proper React.memo usage
   - Analyze component prop dependencies

   **Phase 2: Hook Analysis**
   - Check useMemo usage for expensive computations
   - Check useCallback usage for callbacks passed to memoized children
   - Validate dependency arrays (missing or excessive dependencies)
   - Identify hooks that could benefit from memoization

   **Phase 3: Bundle Analysis** (if requested)
   - Run `npm run build` to generate production bundle
   - Run `npx vite-bundle-visualizer` to analyze bundle composition
   - Identify large dependencies that could be replaced
   - Check for duplicate dependencies
   - Analyze chunk sizes

   **Phase 4: Code Splitting Analysis**
   - Identify routes that should be lazy loaded
   - Check for heavy components that should use React.lazy
   - Verify Suspense boundaries are in place
   - Check for dynamic imports

   **Phase 5: List Rendering Analysis**
   - Find lists rendering >100 items without virtualization
   - Check for proper key usage in lists
   - Identify opportunities for virtual scrolling

   **Phase 6: Image Optimization**
   - Check for unoptimized images (missing width/height, no lazy loading)
   - Identify opportunities for next-gen formats (WebP, AVIF)
   - Check for missing loading="lazy" attributes

3. **Performance Checks:**

   **Re-render Issues:**
   - Components re-rendering unnecessarily
   - Missing React.memo on expensive components
   - Props changing when they shouldn't (object/array recreation)
   - Context value changes causing unnecessary re-renders
   - Inline function/object creation in render

   **Memoization Opportunities:**
   - Expensive computations without useMemo
   - Callbacks passed to children without useCallback
   - Over-memoization (memoizing cheap operations)
   - Missing memoization in custom hooks

   **Bundle Size Issues:**
   - Initial bundle >200KB gzipped
   - Route chunks >100KB gzipped
   - Vendor chunk >150KB gzipped
   - Large dependencies that could be replaced
   - Unused dependencies in package.json

   **Code Splitting Issues:**
   - Routes not lazy loaded
   - Heavy components loaded on initial render
   - Missing Suspense boundaries
   - Eager imports instead of dynamic imports

   **Virtual Scrolling:**
   - Lists with >100 items not using virtualization
   - Missing @tanstack/react-virtual for large datasets
   - Inefficient rendering of large tables/grids

   **Image Optimization:**
   - Images without width/height attributes (causing layout shift)
   - Missing loading="lazy" for below-fold images
   - Using PNG/JPG instead of WebP/AVIF
   - Large image file sizes

   **Dependency Arrays:**
   - Missing dependencies in useEffect/useMemo/useCallback
   - Excessive dependencies causing too many re-runs
   - Primitive vs reference type issues
   - Functions in dependency arrays without useCallback

4. **Output Format:**

   ```
   ===== PERFORMANCE ANALYSIS REPORT =====

   Performance Score: X/100

   Scope: [entire project / feature: X / component: X]
   Date: [timestamp]

   ===== CRITICAL ISSUES (High Impact) =====

   1. [Component/File]: [Issue Description]
      Impact: [Performance impact explanation]
      Location: [file path:line number]

      Current Code:
      ```tsx
      [problematic code]
      ```

      Optimized Code:
      ```tsx
      [optimized code with explanation]
      ```

      Expected Improvement: [X% faster / Y KB smaller / etc.]

   ===== MEDIUM PRIORITY ISSUES =====

   [Same format as above]

   ===== LOW PRIORITY ISSUES =====

   [Same format as above]

   ===== BUNDLE SIZE ANALYSIS ===== (if run)

   Initial Bundle: XKB gzipped (Target: <200KB)
   Largest Chunks:
   - vendor.js: XKB (Target: <150KB)
   - main.js: XKB (Target: <100KB)
   - [route].js: XKB (Target: <100KB per route)

   Largest Dependencies:
   1. [dependency]: XKB - [suggestion: replace with Y or remove]
   2. [dependency]: XKB - [suggestion]

   ===== OPTIMIZATION RECOMMENDATIONS =====

   By Category:

   Re-renders:
   - [X components need React.memo]
   - [Y contexts could be split to prevent unnecessary re-renders]

   Memoization:
   - [X computations should use useMemo]
   - [Y callbacks should use useCallback]

   Code Splitting:
   - [X routes should be lazy loaded]
   - [Y components should use React.lazy]

   Virtual Scrolling:
   - [X lists need virtualization]

   Images:
   - [X images need optimization]

   ===== SUMMARY =====

   Total Issues Found: X (Y high, Z medium, W low)
   Estimated Performance Gain: X%
   Estimated Bundle Size Reduction: XKB
   Recommended Next Steps:
   1. [Priority action]
   2. [Priority action]
   3. [Priority action]
   ```

5. **Code Examples for Optimizations:**

   **React.memo with Custom Comparison:**
   ```tsx
   // Before: Component re-renders on every parent render
   export function UserCard({ user, onUpdate }: UserCardProps) {
     return <div>{user.name}</div>;
   }

   // After: Only re-renders when user changes
   export const UserCard = React.memo(
     function UserCard({ user, onUpdate }: UserCardProps) {
       return <div>{user.name}</div>;
     },
     (prevProps, nextProps) => {
       // Custom comparison: only re-render if user ID changes
       return prevProps.user.id === nextProps.user.id;
     }
   );
   ```

   **useMemo for Expensive Computations:**
   ```tsx
   // Before: Sorts on every render
   function DataTable({ data }: DataTableProps) {
     const sortedData = [...data].sort((a, b) =>
       a.name.localeCompare(b.name)
     );
     return <table>{/* render sortedData */}</table>;
   }

   // After: Only sorts when data changes
   function DataTable({ data }: DataTableProps) {
     const sortedData = useMemo(() => {
       return [...data].sort((a, b) => a.name.localeCompare(b.name));
     }, [data]);

     return <table>{/* render sortedData */}</table>;
   }
   ```

   **useCallback for Callbacks:**
   ```tsx
   // Before: New function on every render
   function Parent() {
     const [count, setCount] = useState(0);

     const handleClick = () => {
       setCount(c => c + 1);
     };

     return <MemoizedChild onClick={handleClick} />;
   }

   // After: Stable function reference
   function Parent() {
     const [count, setCount] = useState(0);

     const handleClick = useCallback(() => {
       setCount(c => c + 1);
     }, []); // Empty deps: function never changes

     return <MemoizedChild onClick={handleClick} />;
   }
   ```

   **React.lazy and Suspense:**
   ```tsx
   // Before: Eager import increases initial bundle
   import Dashboard from './features/dashboard/Dashboard';
   import Settings from './features/settings/Settings';

   function App() {
     return (
       <Routes>
         <Route path="/dashboard" element={<Dashboard />} />
         <Route path="/settings" element={<Settings />} />
       </Routes>
     );
   }

   // After: Lazy loading reduces initial bundle
   import { lazy, Suspense } from 'react';

   const Dashboard = lazy(() => import('./features/dashboard/Dashboard'));
   const Settings = lazy(() => import('./features/settings/Settings'));

   function App() {
     return (
       <Suspense fallback={<LoadingSpinner />}>
         <Routes>
           <Route path="/dashboard" element={<Dashboard />} />
           <Route path="/settings" element={<Settings />} />
         </Routes>
       </Suspense>
     );
   }
   ```

   **Virtual Scrolling Setup:**
   ```tsx
   // Before: Renders all 10,000 items (slow)
   function UserList({ users }: UserListProps) {
     return (
       <div>
         {users.map(user => (
           <UserCard key={user.id} user={user} />
         ))}
       </div>
     );
   }

   // After: Only renders visible items (fast)
   import { useVirtualizer } from '@tanstack/react-virtual';

   function UserList({ users }: UserListProps) {
     const parentRef = useRef<HTMLDivElement>(null);

     const virtualizer = useVirtualizer({
       count: users.length,
       getScrollElement: () => parentRef.current,
       estimateSize: () => 80, // Height of each item in px
       overscan: 5, // Render 5 extra items above/below viewport
     });

     return (
       <div
         ref={parentRef}
         style={{ height: '600px', overflow: 'auto' }}
       >
         <div
           style={{
             height: `${virtualizer.getTotalSize()}px`,
             position: 'relative'
           }}
         >
           {virtualizer.getVirtualItems().map((virtualItem) => (
             <div
               key={virtualItem.key}
               data-index={virtualItem.index}
               style={{
                 position: 'absolute',
                 top: 0,
                 left: 0,
                 width: '100%',
                 height: `${virtualItem.size}px`,
                 transform: `translateY(${virtualItem.start}px)`
               }}
             >
               <UserCard user={users[virtualItem.index]} />
             </div>
           ))}
         </div>
       </div>
     );
   }
   ```

   **Image Optimization:**
   ```tsx
   // Before: Unoptimized image
   <img src="/large-hero.jpg" alt="Hero" />

   // After: Optimized with lazy loading and dimensions
   <img
     src="/large-hero.webp"
     alt="Hero"
     width={1200}
     height={600}
     loading="lazy"
     decoding="async"
   />
   ```

   **Dependency Array Fixes:**
   ```tsx
   // Before: Missing dependency
   function useUserData(userId: string) {
     const [data, setData] = useState(null);

     useEffect(() => {
       fetchUser(userId).then(setData);
     }, []); // Missing userId - stale closure!

     return data;
   }

   // After: Correct dependencies
   function useUserData(userId: string) {
     const [data, setData] = useState(null);

     useEffect(() => {
       fetchUser(userId).then(setData);
     }, [userId]); // Correct

     return data;
   }

   // Before: Excessive dependencies
   function useSearch(query: string) {
     const config = { caseSensitive: true }; // Created on every render

     useEffect(() => {
       search(query, config);
     }, [query, config]); // config changes every render!
   }

   // After: Stable config
   function useSearch(query: string) {
     const config = useMemo(() => ({
       caseSensitive: true
     }), []); // Stable reference

     useEffect(() => {
       search(query, config);
     }, [query, config]); // Now config is stable
   }
   ```

**Reference Standards:**

All optimizations must follow CODE_STANDARDS.md "Performance Standards" section:

- React.memo for expensive components
- useMemo for expensive computations (not trivial ones)
- useCallback for callbacks passed to memoized children
- Code splitting for routes and heavy components
- Bundle size targets: <200KB initial, <100KB per route, <150KB vendor
- Virtual scrolling for lists >100 items
- Proper dependency arrays in hooks
- Image optimization with lazy loading

**Analysis Tools:**

```bash
# Bundle analysis
npm run build
npx vite-bundle-visualizer

# Bundle size check
npm run build -- --mode production
ls -lh dist/assets/*.js

# Type checking
npm run type-check

# Lint performance issues
npm run lint
```

**After Analysis:**

1. Present findings in the structured report format
2. Prioritize issues by impact (high/medium/low)
3. Provide specific file locations and line numbers
4. Show before/after code examples
5. Estimate performance improvements
6. Give actionable next steps

Be thorough and data-driven. Focus on measurable improvements. Provide specific, actionable recommendations with code examples.
