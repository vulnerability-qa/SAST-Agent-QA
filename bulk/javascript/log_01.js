// CWE-117: Log Injection
const winston = require('winston');
const logger = winston.createLogger({ transports: [new winston.transports.Console()] });
function logLogin(username) {
  logger.info('Login: ' + username); // newline in username injects fake log entry
}
