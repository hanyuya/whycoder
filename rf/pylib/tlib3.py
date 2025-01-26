from robot.api import logger

class tlib3:
    def __init__(self, host, port):
        self.host = host
        self.port = port

    def printaddr(self):
        logger.console('host:%s,port:%s' % (self.host, self.port))