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
    packageRules: [
        {
            matchManagers: ["github-actions"],
            enabled: true,
        },
    ],
    "github-actions": {
        fileMatch: [".github/*/*.ya?ml$"],
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
        {
            description: "Look for npm install",
            customType: "regex",
            fileMatch: [".*"],
            matchStrings: [
                "npm\\s+install\\s+(?:--global|-g)\\s+(?<depName>[@\\/_\\-a-zA-Z]+)@(?<currentValue>[\\w.]+)",
            ],
            datasourceTemplate: "npm",
        },
        {
            description: "Look for docker://",
            customType: "regex",
            fileMatch: [".*"],
            matchStrings: [
                "uses\\s*:\\s*['\"]?docker:\\\/\\\/(?<depName>.*?):(?<currentValue>.*)(@(?<currentDigest>.*?))?['\"]?\\s*"
            ],
            datasourceTemplate: "docker",
        },
        {
            description: "Comment",
            customType: "regex",
            fileMatch: [".*"],
            matchStrings: [
                "#\\s*renovate:\\s*datasource=(?<datasource>.*?)\\s+depName=(?<depName>.*?)\\s+(versioning=(?<versioning>.*?))?\\s*(ARG)?.*_(ver|VER)(=|:\\s+)['\"]?(?<currentValue>[\\w+\\.\\-]*)['\"]?",
            ],
            versioningTemplate:
                "{{#if versioning}}{{{versioning}}}{{else}}ver{{/if}}",
        },
    ],
};
