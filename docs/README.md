# PicoOraClaw Documentation

Welcome to the PicoOraClaw documentation. This directory contains detailed guides for understanding and working with the system.

## 📚 Documentation Index

### Core Documentation

- **[PostgreSQL Support](POSTGRESQL_SUPPORT.md)** ⭐ NEW
  - PostgreSQL backend configuration and setup
  - Environment variables and config.json format
  - Database schema and feature overview
  - Troubleshooting and performance optimization

- **[Storage Architecture](STORAGE_ARCHITECTURE.md)** ⭐ NEW
  - Multi-backend storage design using factory pattern
  - Core interfaces and how to extend with new backends
  - Configuration-driven storage selection
  - Testing strategy and deployment considerations

### Getting Started

For initial setup and quick start:
- See [README.md](../README.md) - Main project documentation
- See [README.zh.md](../README.zh.md) - Chinese documentation

## 🗂️ Directory Structure

```
docs/
├── README.md                      # This file
├── POSTGRESQL_SUPPORT.md          # PostgreSQL setup and configuration
├── STORAGE_ARCHITECTURE.md        # Storage layer design
└── slides/                        # Presentation slides
    └── *.jpg                      # Slide images
```

## 🔄 Storage Backend Decision Tree

Choose your storage backend based on your needs:

```
Are you existing Oracle user?
│
├─ YES → Use Oracle AI Database
│   ├─ Built-in ONNX embeddings
│   ├─ No external API required
│   └─ See: main README.md
│
└─ NO → Choose based on your setup:
    │
    ├─ Want lightweight self-hosted DB?
    │   └─ Use PostgreSQL + pgvector
    │       └─ See: POSTGRESQL_SUPPORT.md
    │
    ├─ Want cloud-managed database?
    │   └─ Use Cloud SQL, RDS, or similar
    │       └─ Adapt PostgreSQL setup
    │
    └─ Want embedded/minimal setup?
        └─ SQLite support (coming soon)
```

## 🚀 Quick Setup by Backend

### Oracle AI Database

```bash
# 1. Setup
./build/picooraclaw setup-database

# 2. Run
./build/picooraclaw agent
```

### PostgreSQL

```bash
# 1. Set storage type
export PICO_STORAGE_TYPE=postgres

# 2. Configure connection
export PICO_POSTGRES_HOST=localhost
export PICO_POSTGRES_PORT=5432
export PICO_POSTGRES_DATABASE=picooraclaw
export PICO_POSTGRES_USER=postgres
export PICO_POSTGRES_PASSWORD=password

# 3. Setup
./build/picooraclaw setup-database

# 4. Run
./build/picooraclaw agent
```

## 📖 How to Read This Documentation

1. **First time?** Start with the main [README.md](../README.md)
2. **Using PostgreSQL?** Read [POSTGRESQL_SUPPORT.md](POSTGRESQL_SUPPORT.md)
3. **Extending the system?** Read [STORAGE_ARCHITECTURE.md](STORAGE_ARCHITECTURE.md)
4. **Troubleshooting?** Check the appropriate guide's troubleshooting section

## 🔧 Configuration Reference

| Aspect | Details |
|--------|---------|
| **Config Location** | `~/.picooraclaw/config.json` |
| **Environment Prefix** | `PICO_` |
| **Storage Type** | `PICO_STORAGE_TYPE` (oracle/postgres) |
| **Commands** | See main README or run `./build/picooraclaw --help` |

## 📋 Common Tasks

### Task: Switch from Oracle to PostgreSQL

1. Export current data from Oracle (if any)
2. Update `config.json`: set `"storageType": "postgres"`
3. Configure PostgreSQL connection parameters
4. Run: `./build/picooraclaw setup-database`
5. Import data if needed

→ See [POSTGRESQL_SUPPORT.md - Migration Guide](POSTGRESQL_SUPPORT.md#migration-guide-oracle--postgresql)

### Task: Add Support for New Database Backend

1. Create `pkg/yourdb/` package
2. Implement interfaces from `pkg/storage/types.go`
3. Add config struct to `pkg/config/config.go`
4. Add factory case in `pkg/storage/factory.go`
5. Create tests in `pkg/yourdb/`

→ See [STORAGE_ARCHITECTURE.md - Adding New Backend](STORAGE_ARCHITECTURE.md#adding-a-new-storage-backend-mysql)

### Task: Optimize Vector Search Performance

→ See [POSTGRESQL_SUPPORT.md - Performance](POSTGRESQL_SUPPORT.md#performance-considerations) or main README for Oracle

### Task: Deploy to Production

1. Choose storage backend (Oracle or PostgreSQL)
2. Set up database instance
3. Configure connection in `config.json` or env vars
4. Run: `./build/picooraclaw setup-database`
5. Deploy application

→ See [STORAGE_ARCHITECTURE.md - Deployment](STORAGE_ARCHITECTURE.md#deployment-considerations)

## 🐛 Troubleshooting

### Database Connection Issues

- **Check logs**: Look for connection error messages
- **Verify connectivity**: Test database connection manually
- **Check credentials**: Ensure username/password are correct
- **Firewall**: Ensure database port is accessible

→ See relevant backend guide for detailed troubleshooting

### Vector Search Not Working

- **Check embeddings**: Verify embedding API is running
- **Check indexes**: Database indexes may need to be created
- **Performance**: Large databases may need index tuning

→ See [POSTGRESQL_SUPPORT.md - Troubleshooting](POSTGRESQL_SUPPORT.md#troubleshooting)

### Configuration Issues

- **Config not loading**: Check file permissions and format
- **Env vars not working**: Ensure `PICO_` prefix is correct
- **Defaults**: Falls back to Oracle if storage type not specified

## 📚 Architecture Diagrams

### Storage Layer

```
Application
    ↓
Factory Pattern (pkg/storage/)
    ↓
┌───────────────────┐
│ Oracle | PostSQL  │
└───────────────────┘
    ↓
┌───────────────────┐
│ Database Instance │
└───────────────────┘
```

### Database Selection

```
Config/Env Vars
    ↓
Storage Type Selection
    ├─ oracle    → Oracle AI Database
    └─ postgres  → PostgreSQL + pgvector
```

## 🔗 External Resources

### PostgreSQL
- [PostgreSQL Official Docs](https://www.postgresql.org/docs/)
- [pgvector GitHub](https://github.com/pgvector/pgvector)
- [pgvector Python](https://github.com/pgvector/pgvector-python)

### Oracle
- [Oracle AI Database 26ai](https://www.oracle.com/database/free/)
- [Oracle Database Docs](https://docs.oracle.com/database/)

### Go Database
- [sql/database Package](https://pkg.go.dev/database/sql)
- [lib/pq Driver](https://github.com/lib/pq)

## 📝 Contributing Documentation

When adding new features:

1. Create/update relevant documentation
2. Include setup instructions
3. Add troubleshooting section
4. Include configuration examples
5. Link from this README

## 📞 Support

### For PostgreSQL Issues
- Check [POSTGRESQL_SUPPORT.md](POSTGRESQL_SUPPORT.md)
- Verify pgvector installation
- Check embedding API connectivity

### For Architecture Questions
- Check [STORAGE_ARCHITECTURE.md](STORAGE_ARCHITECTURE.md)
- Review source code in `pkg/storage/`

### For General Issues
- Check main [README.md](../README.md)
- Run `./build/picooraclaw --help`

---

**Last Updated**: April 17, 2026
**Latest PostgreSQL Support**: feat commit 5d84ce6
