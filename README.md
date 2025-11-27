# Nexus SHAi-Hulud Malware NPM-Repo Scanner 

A CLI utility written on GO for interacting with a Nexus Repository to list and scan packages for SHAi-Hulud malware packages

Malware list is pulling from live-update `.CSV`: https://www.koi.ai/incident/live-updates-sha1-hulud

# Usage
```bash
./shaihd-nx-scan.exe [flags]
```

## Required flags
- `-uri`: Nexus instance base URI (e.g., `nexus.example.com` or `https://nexus.example.com`)
- `-repo`: Target Nexus repository name with NPM packages (e.g. `npm-hosted`, `npm`).

## Optional flags
By default utility work in "anonymous mode" but if you need special permission to pull assets via API you can specify your credentials:
- `-u`: Nexus username
- `-p`: Nexus Password

## Extra flags
- `-no-scan`: Skip malware scanning. Only fetch and parse the package list from the specified repository (`-repo`)
- `-help`: CLI-tool command help

# Results
At `shaihd-compromised-checklist.csv` is located compromised pkgs by "Chai-Hulud" which scrapped from [koi.ai](https://koi.ai) `.CSV`

## Repository's assets
Repository's packages which parsed from Nexus API is located at utility execution path at folder `.nx-shaihd-scan` with name-fomat `nx-<your-repo>-npm-pkgs.csv`

## Malware scan results
At this time, results of the malware scanning is located at the CLI-output with message:
```
> Malware pkg list asset found: <infected-package-name>
```

If your repository is clean you will see: 
```
> OK! Malware pkg assets not found at your repository.
```
