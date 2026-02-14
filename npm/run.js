#!/usr/bin/env node

const { spawn, execSync } = require("child_process");
const fs = require("fs");
const https = require("https");
const os = require("os");
const path = require("path");

const REPO = "caiomoura1994/Extend-AI-MCP-Server";
const BIN_NAME = "extend-mcp-server";
const CACHE_DIR = path.join(os.homedir(), ".extend-mcp-server");

const PLATFORM_MAP = { darwin: "darwin", linux: "linux", win32: "windows" };
const ARCH_MAP = { x64: "amd64", arm64: "arm64" };

function getPlatformKey() {
  const plat = PLATFORM_MAP[os.platform()];
  const arch = ARCH_MAP[os.arch()];
  if (!plat || !arch) {
    throw new Error(
      `Unsupported platform: ${os.platform()}/${os.arch()}. ` +
        `Supported: darwin/linux/windows on amd64/arm64.`
    );
  }
  return { plat, arch };
}

function binaryPath(version) {
  const { plat } = getPlatformKey();
  const ext = plat === "windows" ? ".exe" : "";
  return path.join(CACHE_DIR, version, BIN_NAME + ext);
}

function follow(url, maxRedirects = 5) {
  return new Promise((resolve, reject) => {
    if (maxRedirects <= 0) return reject(new Error("Too many redirects"));

    const mod = url.startsWith("https") ? https : require("http");
    mod
      .get(url, { headers: { "User-Agent": "extend-ai-mcp-server-npm" } }, (res) => {
        if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
          return follow(res.headers.location, maxRedirects - 1).then(resolve, reject);
        }
        resolve(res);
      })
      .on("error", reject);
  });
}

async function getLatestVersion() {
  const res = await follow(`https://api.github.com/repos/${REPO}/releases/latest`);
  if (res.statusCode === 404) {
    throw new Error(
      `No releases found at github.com/${REPO}. Please check the repository or install manually.`
    );
  }
  if (res.statusCode !== 200) {
    throw new Error(`GitHub API returned HTTP ${res.statusCode}`);
  }
  const data = await readBody(res);
  const json = JSON.parse(data);
  if (!json.tag_name) {
    throw new Error("Could not determine latest version from GitHub releases.");
  }
  return json.tag_name;
}

function readBody(res) {
  return new Promise((resolve, reject) => {
    const chunks = [];
    res.on("data", (c) => chunks.push(c));
    res.on("end", () => resolve(Buffer.concat(chunks).toString()));
    res.on("error", reject);
  });
}

async function downloadToFile(url, dest) {
  const res = await follow(url);
  if (res.statusCode !== 200) {
    throw new Error(`Download failed (HTTP ${res.statusCode}): ${url}`);
  }
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    res.pipe(file);
    file.on("finish", () => {
      file.close();
      resolve();
    });
    file.on("error", reject);
  });
}

async function downloadBinary(version) {
  const { plat, arch } = getPlatformKey();
  const isWindows = plat === "windows";
  const archiveExt = isWindows ? "zip" : "tar.gz";
  const assetName = `${BIN_NAME}_${plat}_${arch}.${archiveExt}`;
  const url = `https://github.com/${REPO}/releases/download/${version}/${assetName}`;

  const versionDir = path.join(CACHE_DIR, version);
  fs.mkdirSync(versionDir, { recursive: true });

  const archivePath = path.join(versionDir, assetName);

  process.stderr.write(`[extend-ai-mcp-server] Downloading ${version} for ${plat}/${arch}...\n`);
  await downloadToFile(url, archivePath);

  // Extract
  if (isWindows) {
    execSync(
      `powershell -command "Expand-Archive -Force -Path '${archivePath}' -DestinationPath '${versionDir}'"`,
      { stdio: "ignore" }
    );
  } else {
    execSync(`tar xzf "${archivePath}" -C "${versionDir}"`, { stdio: "ignore" });
    // Make executable
    const bin = binaryPath(version);
    fs.chmodSync(bin, 0o755);
  }

  // Clean up archive
  fs.unlinkSync(archivePath);
  process.stderr.write(`[extend-ai-mcp-server] Ready.\n`);
}

async function main() {
  try {
    const version = await getLatestVersion();
    const bin = binaryPath(version);

    if (!fs.existsSync(bin)) {
      await downloadBinary(version);
    }

    if (!fs.existsSync(bin)) {
      throw new Error(`Binary not found after download: ${bin}`);
    }

    // Spawn the Go binary, passing through stdio for MCP
    const child = spawn(bin, process.argv.slice(2), {
      stdio: "inherit",
      env: process.env,
    });

    child.on("exit", (code) => process.exit(code ?? 1));
    child.on("error", (err) => {
      process.stderr.write(`[extend-ai-mcp-server] Failed to start: ${err.message}\n`);
      process.exit(1);
    });
  } catch (err) {
    process.stderr.write(`[extend-ai-mcp-server] Error: ${err.message}\n`);
    process.exit(1);
  }
}

main();
