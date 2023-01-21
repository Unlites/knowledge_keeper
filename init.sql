SELECT 'CREATE DATABASE knowledge_keeper_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'knowledge_keeper_db')\gexec;