module.exports = {
    extends: [
        "config:recommended",
        "docker:enableMajor",
        ":disableRateLimiting",
        "helpers:pinGitHubActionDigests"
    ],
    autodiscover: false,
    repositories: [
        "yellowhat/hosts"
    ],
    dependencyDashboard: true,
    baseBranches: ["main"],
    labels: ["renovate"],
    separateMajorMinor: true,
    lockFileMaintenance: {
        enabled: true,
    },
    customManagers: [
        {
            description: "Look for pip install",
            customType: "regex",
            fileMatch: [".*"],
            matchStrings: [
                "(pip|pip3)\\s+install\\s+(?<depName>[_\\-a-zA-Z0-9]+)(\\[.*\\])?==(?<currentValue>[\\w.]+)",
            ],
            datasourceTemplate: "pypi",
        },
    ],
};
