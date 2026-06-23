# CWE-22: Path Traversal via File.read
def serve_file(filename)
  File.read("/var/www/uploads/" + filename)
end
