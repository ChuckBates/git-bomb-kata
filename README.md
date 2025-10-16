# git-bomb-kata (Go)

A lightweight **Go CLI “git bomb”** plus a simple **Roman Numerals kata**.  
This setup is designed for code retreats and practice sessions focused on **tiny, atomic commits** and **fast feedback**.

The bomb periodically runs `git reset --hard` to erase any uncommitted work — forcing small, frequent commits.  
All configuration is built into the repo; you just clone, open in VS Code, and start coding.

## Prerequisites
- **macOS**
- **Go 1.24**
- **Visual Studio Code**
  - The first time you open this repo, VS Code will prompt to install the **Go extension** — click **Install**.

## 🚀 Quick Start

1. **Clone & open**
   ```bash
   git clone git@github.com:ChuckBates/git-bomb-kata.git
   cd git-bomb-kata
   code .
   ```

2. **Run a task**  
   Press **⇧⌘B** (or *Terminal → Run Build Task…*)  
   You’ll see:
   ```
   Select the task to run
   > Start Git Bomb
     Run kata tests (watch)
     Run kata tests
   ```

   These are handy tools for quick start up. 

3. **Start the bomb**  
   Choose **Start Git Bomb** → pick an interval such as `1m` or `2m`.  
   This opens a terminal running the bomb. Every cycle, it resets your repo to the last commit.

4. **Run tests (optional)**  
   Press **⇧⌘B** again → choose **Run kata tests (watch)** or **Run kata tests**.  
   This will either continuously run `go test` every 5 seconds so you see failures and passes instantly or one off run tests.


## The Kata: Roman Numerals
For more details on the Kata to solve, see the appropriate kata [README](katas/roman_numerals/README.md)

## Rules of Git Bomb
- TDD is your friend!
- A commit requires **both** a passing test and implementation
- If the bomb gets you, think of how you can break the work into smaller pieces

## Tips
- **Commit constantly.**  
  Anything uncommitted when the bomb fires will be lost — that’s the exercise!
- **Use small, clear commit messages.**  
  Think: *“add test: 4 → IV”*, *“green: 9 → IX”*, *“refactor: table-driven cases”*.
- **Pair or mob program.**  
  Switch drivers every reset.
- **Intervals:**  
  - Don't bother: `10m`
  - Warm-up: `8m`
  - Feeling stressed: `5m`  
  - Should I commit broken code?: `2m`  
  - Chaos: `30s`

## End of Session
When you’re done:
```bash
git log --oneline
```
to review your streak of tiny commits — that’s your progress snapshot!
