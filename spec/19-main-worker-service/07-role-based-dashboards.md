# 07 — Role-Based Dashboards

**Spec:** `19-main-worker-service`
**Version:** 1.0.0

Roles, dashboards, and the **non-negotiable** access-check pattern.

---

## 1. The Rule (READ FIRST)

Per verbatim §Roles.3:

> ❌ **NEVER** check `if (user.role === 'PowerAdmin')`.
> ✅ **ALWAYS** check `User has access to {EnumPage}`.

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

## 3. `EnumPage` Capability Catalog

```php
enum EnumPage: string
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

Extensible. Adding a page = add an enum case + a row in `RolePageAccess` join table seed.

---

## 4. Schema Additions for Access (Main DB)

### 4.1 `RolePageAccess` (join, exempt from Description rule)

| Column | Type | Null |
|--------|------|------|
| `RolePageAccessId` | INTEGER | NO (PK) |
| `RoleId` | INTEGER | NO (FK) |
| `EnumPageCode` | TEXT | NO (matches `EnumPage` value) |

Unique: `(RoleId, EnumPageCode)`.

### 4.2 Default seed (via Seedable-Config)

| Role | Pages |
|------|-------|
| `PowerAdmin` | All `EnumPage` values |
| `AdminUser` | `AdminPage`, `BillingPage`, `CompanySettingsPage`, `UserManagementPage`, `DashboardPage` |
| `Member` | `DashboardPage` |

---

## 5. Access Check Implementation

CODE RED compliant (≤8 lines, positive guard, max 2 operands):

```php
public function userHasAccessToPage(int $userId, EnumPage $page): bool
{
    $roleIds = $this->userRoleRepo->roleIdsFor($userId);
    return $this->rolePageAccessRepo->anyRoleGrants($roleIds, $page->value);
}
```

Used in middleware:

```php
public function handle(Request $request, Closure $next, string $pageCode): Response
{
    $page = EnumPage::from($pageCode);
    $this->guardUserHasAccess($request->user()->id, $page);
    return $next($request);
}
```

`guardUserHasAccess` throws `AccessDenied` (per `08-error-contract.md`) when access is missing.

Route declaration:

```php
Route::post('/API/V1/Workers/All/Update', UpdateAllController::class)
    ->middleware('access:PushUpdatePage');
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

Required pages: `PowerAdminPage`, `WorkerRegistryPage`, `PushUpdatePage`, `AuditLogPage`.

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

React components use a `<RequiresAccess page={EnumPage.PushUpdatePage}>` wrapper:

```tsx
<RequiresAccess page={EnumPage.PushUpdatePage} fallback={<Hidden/>}>
  <PushUpdatePanel />
</RequiresAccess>
```

The wrapper reads access flags from the worker JWT's `roles` claim resolved against the local `RolePageAccess` cache (refreshed when token refreshes). NEVER hardcode role names in JSX.

Per `.lovable/coding-guidelines/coding-guidelines.md`: React components < 100 lines, small and reusable.

---

## 8. Audit

Every access denial writes to an `AccessDenialEvent` table (transactional, includes `Notes` + `Comments`). Per `spec/03-error-manage/`: log, don't swallow.

---

*Role-based dashboards v1.0.0 — 2026-05-04*
