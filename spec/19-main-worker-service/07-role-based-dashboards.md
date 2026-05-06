# 07 — Role-Based Dashboards

**Spec:** `19-main-worker-service`
**Version:** 2.0.0

> **v2.0.0 rename (Phase 1):** all references to `AccessItem` are now `AccessItem`. Access checks resolve by `Code` or by `PageUrlSuffix` matcher. See `03-main-db-schema.md` §2.6.1 and `14-rbac-and-status-seed.md`.

Roles, dashboards, and the **non-negotiable** access-check pattern.

---

## 1. The Rule (READ FIRST)

Per verbatim §Roles.3:

> ❌ **NEVER** check `if (user.role === 'PowerAdmin')`.
> ✅ **ALWAYS** check `User has access to {AccessItem}`.

Why: roles change. Capabilities don't. Tomorrow a new role (`SupportAgent`) needs to see the billing page — with the role-based check you change every site of the check; with the page-based check you grant `SupportAgent → BillingPage` once.

---

## 2. Built-in Roles

| Role | RoleCode | Notes |
|------|----------|-------|
| Power Admin | `PowerAdmin` | Application owners (Riseup Asia LLC). System-wide settings. |
| Admin User | `AdminUser` | Paying customer admin. Scoped to their `Company`. |
| Member | `Member` | Default for users without elevated access. |

Roles are extensible — new ones added to `Role` table via Seedable-Config seed file. No code change required to introduce a role.

Role assignment happens from the Power Admin dashboard (verbatim §Roles.4).

---

## 3. `AccessItem` Capability Catalog

```php
enum AccessItem: string
{
    case PowerAdminPage     = 'PowerAdminPage';
    case AdminPage          = 'AdminPage';
    case BillingPage        = 'BillingPage';
    case CompanySettingsPage = 'CompanySettingsPage';
    case UserManagementPage = 'UserManagementPage';
    case WorkerRegistryPage = 'WorkerRegistryPage';
    case PushUpdatePage     = 'PushUpdatePage';
    case AuditLogPage       = 'AuditLogPage';
    case DashboardPage      = 'DashboardPage';
}
```

Extensible. Adding a page = add an enum case + a row in `RoleAccessItem` join table seed.

---

## 4. Schema Additions for Access (Main DB)

### 4.1 `RoleAccessItem` (join, exempt from Description rule)

| Column | Type | Null |
|--------|------|------|
| `RoleAccessItemId` | INTEGER | NO (PK) |
| `RoleId` | INTEGER | NO (FK) |
| `Code` | TEXT | NO (matches `AccessItem` value) |

Unique: `(RoleId, Code)`.

### 4.2 Default seed (via Seedable-Config)

| Role | Pages |
|------|-------|
| `PowerAdmin` | All `AccessItem` values |
| `AdminUser` | `AdminPage`, `BillingPage`, `CompanySettingsPage`, `UserManagementPage`, `DashboardPage` |
| `Member` | `DashboardPage` |

---

## 5. Access Check Implementation

CODE RED compliant (≤8 lines, positive guard, max 2 operands):

```php
public function userHasAccessToItem(int $userId, AccessItem $accessItem): bool
{
    $roleIds = $this->userRoleRepo->roleIdsFor($userId);
    return $this->rolePageAccessRepo->anyRoleGrants($roleIds, $page->value);
}
```

Used in middleware:

```php
public function handle(Request $request, Closure $next, string $pageCode): Response
{
    $page = AccessItem::from($accessItemCode);
    $this->guardUserHasAccess($request->user()->id, $page);
    return $next($request);
}
```

`guardUserHasAccess` throws `AccessDenied` (per `08-error-contract.md`) when access is missing.

Route declaration (resolves F-A-34 — stack-agnostic contract first, Laravel example second):

**Stack-agnostic contract.** Every route that mutates or reads a governed page MUST be wrapped by an access guard that:
1. Resolves the route to its `Code` (PascalCase, from `AccessItem` ref table per `03-` §2.6.1).
2. Calls `guardUserHasAccess(userId, pageCode)` BEFORE the controller body.
3. On denial, returns the `AccessDenied` envelope per `08-error-contract.md` §3.5 with HTTP 403.

The wiring mechanism is implementation-defined: Laravel middleware, Express middleware, Go HTTP middleware chain, ASP.NET filter, etc. The contract is the same.

Laravel example (one valid binding of the contract above):

```php
Route::post('/API/V1/Workers/All/Update', UpdateAllController::class)
    ->middleware('access:PushUpdatePage');
```

Express example (equivalent contract):

```ts
app.post('/API/V1/Workers/All/Update',
    requireAccess('PushUpdatePage'),
    updateAllController);
```

---

## 6. The Three Dashboards (default)

### 6.1 Power Admin Dashboard
- Worker registry view (status, version, assigned-company count)
- Push-update controls (one / all)
- Endpoint auth settings (OQ-1)
- Update schedule settings
- Audit log viewer
- Role/access matrix editor

Required pages: `PowerAdmin`, `WorkerRegistryPage`, `PushUpdatePage`, `AuditLogPage`.

### 6.2 Admin User Dashboard
- Company profile editor (calls Worker)
- User management for their company
- Billing
- Their company's analytics

Required pages: `AdminPage`, `BillingPage`, `CompanySettingsPage`, `UserManagementPage`.

### 6.3 Member Dashboard
- The actual product surface (graphs, business data — all from Worker).

Required pages: `DashboardPage`.

---

## 7. Frontend Gating

React components use a `<RequiresAccess page={AccessItem.PushUpdate}>` wrapper:

```tsx
<RequiresAccess page={AccessItem.PushUpdate} fallback={<Hidden/>}>
  <PushUpdatePanel />
</RequiresAccess>
```

The wrapper reads access flags from the worker JWT's `roles` claim resolved against the local `RoleAccessItem` cache (refreshed when token refreshes). NEVER hardcode role names in JSX.

Per `.lovable/coding-guidelines/coding-guidelines.md`: React components < 100 lines, small and reusable.

---

## 8. Audit

Every access denial writes to an `AccessDenialEvent` table (transactional, includes `Notes` + `Comments`). Per `spec/03-error-manage/`: log, don't swallow.

---

*Role-based dashboards v1.0.0 — 2026-05-04*
