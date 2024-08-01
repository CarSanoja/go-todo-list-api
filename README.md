# To-Do List API

This project represents the backend of a to-do list application. It allows users to manage their to-do items via HTTP requests, supporting various CRUD operations. The server is built using the Gin web framework and uses a CSV file for data storage.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installing Dependencies](#installing-dependencies)
  - [Configuration](#configuration)
- [Running the Server](#running-the-server)
- [API Endpoints](#api-endpoints)
  - [Get All To-Do Items](#1-get-all-to-do-items)
  - [Get a Single To-Do Item](#2-get-a-single-to-do-item)
  - [Add a New To-Do Item](#3-add-a-new-to-do-item)
  - [Update a To-Do Item](#4-update-a-to-do-item)
  - [Delete a To-Do Item](#5-delete-a-to-do-item)
- [License](#license)

## Getting Started

To get started with this project, you need to install the necessary dependencies and set up the configuration.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)

### Installing Dependencies

A script is provided to install all necessary dependencies. Run the following command:

```
./scripts/install_lib.sh
```

