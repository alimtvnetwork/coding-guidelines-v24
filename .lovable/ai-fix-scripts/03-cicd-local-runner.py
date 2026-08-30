import sys
import subprocess
from concurrent.futures import ThreadPoolExecutor

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

JOBS = {
    "Relative Path Check": [sys.executable, "linter-scripts/check-relative-paths.py"],
    "Prompts Loaded Check": [sys.executable, "linter-scripts/check-prompts-loaded.py"],
    "Readme Install Section Check": [sys.executable, "linter-scripts/check-readme-install-section.py"],
    "Forbidden Strings Check": [sys.executable, "linter-scripts/check-forbidden-strings.py"],
    "Newline Styling Check": [sys.executable, "linter-scripts/check-newline-styling.py"],
}

def run_job(job_name, cmd):
    print(f"🔄 Starting: {job_name}...")
    try:
        result = subprocess.run(
            cmd,
            capture_output=True,
            text=True,
            encoding='utf-8',
            errors='replace'
        )
        if result.returncode == 0:
            return (job_name, True, result.stdout)
        else:
            return (job_name, False, result.stdout + "\n" + result.stderr)
    except Exception as e:
        return (job_name, False, str(e))

def main():
    print("🚀 Running Local CI/CD Pipeline via ThreadPoolExecutor...")
    print(f"📋 Enqueued Jobs: {', '.join(JOBS.keys())}\n")
    
    results = []
    with ThreadPoolExecutor(max_workers=3) as executor:
        futures = {executor.submit(run_job, name, cmd): name for name, cmd in JOBS.items()}
        for future in futures:
            results.append(future.result())

    failed = False
    print("\n================ FINAL SUMMARY ================")
    for name, success, output in results:
        if success:
            print(f"✅ {name}: PASSED")
        else:
            print(f"❌ {name}: FAILED")
            failed = True
            print("--- LOG ---")
            print(output.strip())
            print("-----------\n")

    if failed:
        sys.exit(1)
    else:
        print("🎉 All jobs passed successfully!")
        sys.exit(0)

if __name__ == "__main__":
    main()
