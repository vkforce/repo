import logging
from logging.handlers import SysLogHandler
import json
from datetime import datetime, timezone


class QLogger:

    SYSLOG_LOG_FORMAT = '1 %(asctime)s forcepoint-one ForcepointONEApp PROCID - - [NOT:0000006000] %(message)s'
    SYSLOG_TIME_FORMAT = '%Y-%m-%dT%H:%M:%S%z'

    def __init__(self, logger_name):
        self.logger = logging.getLogger(logger_name)
        self.add_log_handler()


    def add_log_handler(self):
        syslog_handler = SysLogHandler(address=('127.0.0.1', 514))
        self.SYSLOG_LOG_FORMAT.replace('PROCID', '121313')
        formatter = logging.Formatter(self.SYSLOG_LOG_FORMAT, self.SYSLOG_TIME_FORMAT)
        syslog_handler.setFormatter(formatter)
        self.logger.addHandler(syslog_handler)

    def log(self, message, level='info'):
        msg = json.dumps(message)
        log_time = datetime.now(timezone.utc).strftime(self.SYSLOG_TIME_FORMAT)
        # Format the log message according to the specified format
        syslog_msg = self.SYSLOG_LOG_FORMAT % {'asctime': log_time, 'message': msg}
        if level == 'info':
            print(syslog_msg)
            self.logger.info(message)
        elif level == 'error':
            self.logger.error(message)
        elif level == 'warning':
            self.logger.warning(message)


logger = QLogger('com.qradar.forcepoint')
logger.log({"message": "This is an info message"}, level='info')