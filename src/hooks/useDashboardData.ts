import { useQuery } from "@tanstack/react-query";
import type {
  CiRunsPayload,
  DashboardSnapshot,
  HealthScore,
  VersionInfo,
} from "@/types/dashboard";
import versionJson from "../../version.json";

const HEALTH_URL = "/health-score.json";
const CI_URL = "/ci-runs.json";

const isOkResponse = (response: Response): boolean => response.ok === true;

async function fetchJson<T>(url: string, label: string): Promise<T> {
  const response = await fetch(url, { cache: "no-store" });

  if (isOkResponse(response) === false) {
    throw new Error(`Failed to load ${label} (${response.status})`);
  }

  return (await response.json()) as T;
}

async function loadSnapshot(): Promise<DashboardSnapshot> {
  const [health, ci] = await Promise.all([
    fetchJson<HealthScore>(HEALTH_URL, "health score"),
    fetchJson<CiRunsPayload>(CI_URL, "CI runs"),
  ]);

  return { health, version: versionJson as VersionInfo, ci };
}

export function useDashboardData() {
  return useQuery({
    queryKey: ["dashboard-snapshot"],
    queryFn: loadSnapshot,
    staleTime: 60_000,
  });
}
