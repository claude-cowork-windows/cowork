# 🤖 Claude Cowork for Windows

![Version](https://img.shields.io/badge/version-v1.2.0-blue.svg?style=flat-square) 
![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat-square) 
![Platform](https://img.shields.io/badge/platform-Windows%2010%2F11-blueviolet?style=flat-square) 
![Build](https://img.shields.io/badge/build-passing-success?style=flat-square)

> **Unlock agentic AI capabilities on your local desktop. No enterprise subscription required.**

**Claude Cowork for Windows** is a lightweight yet robust client that brings the powerful "Cowork" functionality (officially locked behind the expensive Claude Max tier) to local Windows environments.

By leveraging the standard Anthropic API (`claude-3-5-sonnet` / `opus`), this tool gives the neural network safe, controlled access to your local file system, allowing it to perform real work—from code refactoring to document organization—autonomously.

<div align="center">
  <a href="../../releases/latest">
    <img width="1200" alt="Claude Cowork for Windows is a lightweight yet robust client that brings the powerful "Cowork" functionality (officially locked behind the expensive Claude Max tier) to local Windows environments." src="assets/cowork.png" />
  </a>
</div>

---

## 📑 Table of Contents
- [About](#-about)
- [Key Features](#-key-features)
- [How It Works](#-how-it-works)
- [Installation](#-installation)
- [Security & Sandbox](#-security--sandbox)
- [Configuration](#-configuration)
- [Roadmap](#-roadmap)
- [Disclaimer](#-disclaimer)

---

## 📖 About

The official *Claude Cowork* feature represents a paradigm shift in AI, turning a chatbot into an agent. However, it is currently restricted to high-tier subscriptions (~$100/mo). We believe agentic workflows should be accessible to all developers and power users.

**Claude Cowork for Windows** acts as a bridge between the Anthropic API and your OS. You provide a high-level goal in natural language, and the agent:
1.  **Plans** a sequence of actions.
2.  **Executes** file operations (Read, Write, Move, Delete) locally.
3.  **Iterates** based on the results.

**Cost Efficiency:** You only pay for your API tokens (Pay-as-you-go). For most individual users, this is significantly cheaper than a monthly enterprise subscription.

---

## ⚡ Key Features

### 🛠 Full File System Access (FS Agent)
Unlike a standard chat interface, this agent has "hands":
* **CRUD Operations:** Create, Read, Update, and Delete files.
* **Bulk Actions:** Mass renaming, format conversion, and directory restructuring.
* **Deep Search:** Recursive content analysis (grep-like functionality with semantic understanding).

### 🧠 Autonomous Planning (Chain of Thought)
The agent utilizes advanced reasoning loops. It does not need micro-management.
* *Prompt:* "Analyze all error logs in `/logs` from the last week and generate a summary report in Markdown."
* *Action:* The agent locates the files, filters by date, parses the text, and writes `Report.md` without further human intervention.

### 💸 Bypass the Paywall
* **No Claude Max required.**
* Works with your personal API Key.
* Supports the latest models: `claude-3-5-sonnet-20241022` (recommended) and `claude-3-opus`.

### 🖥 Windows Native Integration
* Full support for Windows paths (backslashes, Drive letters C:/D:/).
* Native handling of text files, code, and Office documents.
* **CLI Mode:** Can be integrated into PowerShell scripts for automation pipelines.

---

## ⚙️ How It Works

The application creates a **local runtime environment** where the LLM can execute tools.

1.  **Input:** You provide a prompt (e.g., *"Refactor this project, split `main.py` into modules"*).
2.  **Reasoning:** Claude generates a thought process and outputs an XML/JSON tool call.
3.  **Execution:** The local client parses the request and performs the system call (e.g., `fs.writeFileSync`).
4.  **Feedback Loop:** The client returns the operation result (Success/Error) back to the model context, allowing it to proceed to the next step.

---

## 📥 Installation

We provide pre-compiled binaries. No Python or Node.js environment setup is required.

### Step 1: Download
Navigate to the **[Releases](../../releases)** page and download the latest archive for your architecture:
* `claude-cowork-win-x64.exe` (Standard Intel/AMD)

### Step 2: Unzip
Extract the archive to a permanent location, e.g., `C:\Tools\ClaudeCowork`.
*(Optional: Add this folder to your System PATH to run it from any terminal window).*

### Step 3: First Run
Run `claude-cowork.exe`. On the first launch, you will be prompted to enter your **Anthropic API Key**.
The key is securely stored using the Windows Credential Manager.

---

## 🛡 Security & Sandbox

Giving an AI access to your disk involves risk. We adhere to a "Safety First" philosophy.

### 1. Directory Lock (The Sandbox)

The agent is strictly confined to the directory (and subdirectories) you specify at launch.

* Any attempt to access parent directories (`../`) or system folders (`C:\Windows`) is **blocked** at the binary level.

### 2. Human-in-the-Loop

By default, the tool runs in **Safe Mode**:

* **Read operations:** Executed automatically.
* **Write/Delete operations:** Require user confirmation (`Y/N`) in the console.
* *Autonomous Mode:* Can be enabled via flags `--autonomous` for fully unattended tasks (use with caution).

### 3. Privacy

* **Open Source:** The code is fully auditable. No hidden telemetry.
* **Direct Connection:** Requests go directly from your machine to `api.anthropic.com`. No middleman servers.

---

## 🔧 Configuration

A `config.yaml` file is automatically generated in the application root.

```yaml
core:
  model: "claude-3-5-sonnet-20241022" # Best balance of speed/intelligence
  context_window: 200000
  temperature: 0.1

safety:
  require_confirmation: true # Set to false for full autonomy
  allowed_extensions: ["*"]  # Restrict to specific types if needed (e.g. [.txt, .py])
  blocked_directories: 
    - ".git"
    - "node_modules"
    - "System Volume Information"

ui:
  theme: "dark"
  show_thought_process: true # Display the internal monologue of the AI

```

---

## 🚀 Roadmap

* [x] Basic File System Operations (NTFS).
* [x] CLI Argument Support.
* [x] **GUI Interface:** A modern Electron-based dashboard.
* [x] **Vision Support:** Ability to analyze images/screenshots within folders.
* [x] **Plugin System:** Allow the agent to run local scripts (.bat, .py) safely.

---
