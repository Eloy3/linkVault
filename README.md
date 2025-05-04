# LinkVault

LinkVault is a personal bookmark manager with a RESTful API backend and a terminal-based user interface.  
It lets you organize, store, and open links easily from your terminal.

---

## 📦 Features

- Store bookmarks with title, URL, description, and timestamp
- TUI built with [`tview`](https://github.com/rivo/tview) for fast, keyboard-driven navigation
- REST API for fetching and managing bookmarks
- PostgreSQL database backend

---

## 🧰 Tech Stack

- **Backend**: Go (`net/http`), PostgreSQL, `database/sql`
- **TUI**: Go + `tview`

---

## ⚙️ How It Works

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

---

## 🔧 TODO / Roadmap

- [ ] **Add full CRUD API**  
  - [ ] POST endpoint to add new bookmarks  
  - [ ] PUT/PATCH endpoint to update existing bookmarks  
  - [ ] DELETE endpoint to remove bookmarks  
- [ ] **Implement user authentication**  
  - [ ] Token-based login (JWT)  
  - [ ] Authenticated access to bookmark endpoints  
- [ ] **Enhance bookmark structure**  
  - [ ] Support for tags or categories  
  - [ ] Add search/filter API (by tag, title, etc.)  
- [ ] **Improve TUI**  
  - [ ] Add filtering and search in the UI  
  - [ ] Display tags or categories alongside bookmarks  
  - [ ] Support keyboard shortcuts for quick navigation  
- [ ] **Settings & Config**  
  - [ ] Config file for default API URL, theme, etc.  
  - [ ] TUI theming (light/dark modes)  

