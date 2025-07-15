CREATE TABLE IF NOT EXISTS vouchers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    crew_name TEXT,
    crew_id TEXT,
    flight_number TEXT,
    flight_date TEXT,
    aircraft_type TEXT,
    seat1 TEXT,
    seat2 TEXT,
    seat3 TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_flight_date ON vouchers(flight_number, flight_date);