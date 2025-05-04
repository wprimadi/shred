# SHRED (Secure High-level Removal & Erasure Drivers)

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/wprimadi/shred)](https://goreportcard.com/report/github.com/wprimadi/shred)
[![Go Reference](https://pkg.go.dev/badge/github.com/wprimadi/shred.svg)](https://pkg.go.dev/github.com/wprimadi/shred)

**SHRED** is a Go library that provides secure file deletion using multiple industry-standard erasure algorithms. It supports a wide range of secure data wiping techniques to ensure that sensitive files are unrecoverable after deletion.

---

## 🧠 Background

This project was born out of both my personal and professional commitment to data security and privacy. I created SHRED (Secure High-level Removal & Erasure Drivers) to provide a reliable and standards-based approach to secure file deletion.

On a personal note, I even named my third son `Gutmann Alvarendra Primadi`, inspired by the `Gutmann method`, a highly respected and comprehensive algorithm in the world of data sanitization. This isn’t just a technical topic for me; it reflects how deeply I value strong, irreversible data protection and secure digital practices.

---

## ✨ Features

- ✅ Multiple secure deletion algorithms
- ✅ Simple and consistent API
- ✅ Follows Go best practices and SonarQube compliance
- ✅ Suitable for integration into CLI tools, backend services, or DevSecOps workflows

---

## 🔒 Supported Algorithms

| Algorithm                | Description                                                                 |
|--------------------------|-----------------------------------------------------------------------------|
| **Gutmann**              | 35-pass algorithm that writes specific patterns designed for older disk tech |
| **DoD 5220.22-M**        | US Department of Defense 3-pass standard                                     |
| **DoD 5220.22-M (ECE)**  | Extended 7-pass version of DoD method                                       |
| **NIST 800-88**          | 3-pass wipe based on NIST SP 800-88 guidelines (pattern + random)           |
| **Random Overwrite**     | Overwrites the file with random data for 3 passes                           |
| **Zero Fill**            | Overwrites the file with `0x00` for 3 passes                                |
| **One Fill**             | Overwrites the file with `0xFF` for 3 passes                                |
| **Cryptographic Erase**  | Simulated encryption and deletion of key (CE principle)                     |

---

## 🚀 Getting Started

### 📦 Installation

Use `go get` to include SHRED in your project:

```bash
go get github.com/wprimadi/shred
```

### 🧩 Usage

```go
package main

import (
    "log"
    "github.com/wprimadi/shred"
)

func main() {
    err := shred.SecureDelete("/path/to/sensitive-file.txt", "gutmann")
    if err != nil {
        log.Fatalf("Secure deletion failed: %v", err)
    }

    log.Println("File successfully wiped using Gutmann method.")
}
```

### 🛠️ Available Methods

- `gutmann`
- `dod`
- `dod-ece`
- `nist`
- `random`
- `zero-fill`
- `one-fill`
- `cryptographic-erase`

---

## 🧑‍💻 Author

**Wahyu Primadi**  
📧 [saya@wahyuprimadi.com](mailto:saya@wahyuprimadi.com)  
🌐 [https://wahyuprimadi.com](https://wahyuprimadi.com)

---

## 📋 License

MIT License. See LICENSE for details.

---

## 🙏 Acknowledgements

- Inspired by secure deletion standards including [NIST 800-88](https://csrc.nist.gov/publications/detail/sp/800-88/rev-1/final) and [DoD 5220.22-M](https://www.dss.mil/).
- Gutmann algorithm reference: Peter Gutmann’s paper “[Secure Deletion of Data from Magnetic and Solid-State Memory](https://www.cs.auckland.ac.nz/~pgut001/pubs/secure_del.html)”.