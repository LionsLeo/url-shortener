# URL Shortener Application

This repository contains a **Full Stack** URL Shortener application built with:
- **Go (Golang)** as the backend server
- **Flutter** as the frontend client

The project provides a simple interface to shorten long URLs and redirect users seamlessly.

---

## Table of Contents

1. [Overview](#overview)  
2. [Features](#features)  
3. [Tech Stack](#tech-stack)  
4. [Project Structure](#project-structure)  
5. [Getting Started](#getting-started)  
   - [Prerequisites](#prerequisites)  
   - [Installation](#installation)  
   - [Running the Backend](#running-the-backend)  
   - [Running the Frontend](#running-the-frontend)  
6. [API Endpoints](#api-endpoints)  
7. [Contributing](#contributing)  
8. [License](#license)  

---

## Overview

This URL Shortener application allows users to:
- Enter a long URL.
- Generate a short URL based on a unique hash.
- Redirect to the original URL by visiting the short URL link.

It uses a **Go** server to handle API requests and manage redirections. The **Flutter** app communicates with the backend to display a friendly user interface for creating and managing shortened links.

---

## Features

- **Create Short URLs:** Generate shortened URLs quickly from any valid long URL.
- **Redirect:** Redirect instantly to the original URL when using the short URL.
- **Copy Short URL:** Conveniently copy the shortened URL to the clipboard.
- **Simple User Interface:** Flutter-based frontend providing an intuitive user experience.
- **Database/Storage Integration (optional):** Store your URL mappings in a database (e.g., PostgreSQL, MongoDB) or in-memory. (*The specific database usage can be adapted based on your needs.*)

---

## Tech Stack

- **Backend:** [Go (Golang)](https://go.dev/)
  - Web framework: Standard library `net/http` (or any preferred Go framework).
  - Database (optional): Choose a preferred driver or ORM if you want to persist data.
- **Frontend:** [Flutter](https://flutter.dev/)
  - Cross-platform mobile & web applications.
  - State management: Provider, Bloc, or any other approach you prefer.
  - API communication: HTTP requests to the Go server.

---

## Getting Started

### Prerequisites

- **Go**: v1.16 or newer  
  [Install Go](https://go.dev/doc/install)
- **Flutter**: Have the Flutter SDK set up  
  [Install Flutter](https://docs.flutter.dev/get-started/install)
- **Database** (Optional): If you choose to use one, ensure it’s installed and running.

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/LionsLeo/url-shortener.git
   cd url-shortener
