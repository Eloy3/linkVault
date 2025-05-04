# LinkVault 🔖

LinkVault is a personal bookmark manager with a RESTful API backend and a terminal-based user interface.  
It lets you organize, store, and open links easily from your terminal.

---

## 📦 Features

- 📋 Store bookmarks with title, URL, description, and timestamp
- 🔍 TUI built with [`tview`](https://github.com/rivo/tview) for fast, keyboard-driven navigation
- 🌐 REST API for fetching and managing bookmarks
- 🧱 PostgreSQL database backend

---

## 🧰 Tech Stack

- **Backend**: Go (`net/http`), PostgreSQL, `database/sql`
- **TUI**: Go + `tview`
- **Dev Tools**: VS Code, Delve (for debugging), `curl` or Postman (for API testing)

---

### ⚙️ How It Works

1. **PostgreSQL Initialization**  
   The backend expects an existing PostgreSQL database.  
   Connection details are provided via environment variables in a `.env` file:

   ```env
   DB_HOST=localhost
   DB_USER=youruser
   DB_NAME=linkvault
   ```

2. **Starting the Server**
    The API server runs on localhost:8080. When you send a GET request to:
    ``` http://localhost:8080/bookmarks ``` 
    it returns all saved bookmarks as JSON.

3. **Running the TUI**
    The terminal-based UI is launched separately. When started, it automatically fetches bookmarks from the API and displays them using tview.
    You can select a bookmark to open it in your browser directly from the terminal.
