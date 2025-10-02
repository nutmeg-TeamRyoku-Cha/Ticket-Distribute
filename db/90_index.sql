CREATE INDEX IF NOT EXISTS idx_projects_building_id ON projects(building_id);
CREATE INDEX IF NOT EXISTS idx_tickets_visitor_id  ON tickets(visitor_id);
CREATE INDEX IF NOT EXISTS idx_tickets_project_id  ON tickets(project_id);
CREATE INDEX IF NOT EXISTS idx_sessions_visitor_id ON login_sessions(visitor_id);