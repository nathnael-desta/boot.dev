# boot.dev
a quick template for copying boot.dev code into editor to run it

# 🚀 Run Go Tests with a Shortcut in VS Code

This project is configured to run `go test -v` quickly using a custom task and keyboard shortcut in Visual Studio Code. This helps you run your tests instantly without needing to open the terminal or type commands.

---

## 🧪 What This Does

- Creates a reusable task to run `go test -v`
- Adds a keyboard shortcut (e.g. `Ctrl+Shift+T`) to run tests instantly
- Displays output in the terminal pane

---

## ⚙️ Step-by-Step Setup

### ✅ 1. Create the Test Task

1. Open your project in **Visual Studio Code**.
2. Press `Ctrl+Shift+P` to open the **Command Palette**.
3. Type and select: `Tasks: Configure Task`.
4. Choose: **Others**.
5. VS Code will generate a `.vscode/tasks.json` file.

Replace its content with the following:

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "Run Go Tests",
      "type": "shell",
      "command": "go test -v",
      "group": {
        "kind": "test",
        "isDefault": true
      },
      "problemMatcher": [],
      "detail": "Runs go test with verbose output"
    }
  ]
}

2. Add a Keyboard Shortcut to Run the Task
Press Ctrl+K, then Ctrl+S to open Keyboard Shortcuts.

In the search bar, type:

arduino
Copy
Edit
Run Test Task
Locate: Tasks: Run Test Task

Click the ✏️ pencil icon next to it.

Press your desired shortcut (e.g. Ctrl+Shift+T) and press Enter.
