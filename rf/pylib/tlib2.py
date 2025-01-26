from robot.api import logger
class SubLibrary2:
    def __init__(self, host, port, table='test'):
        self.host = host
        self.port = port
        self.table = table

    def printaddr2(self):
        logger.console('host:%s,port:%s,table:%s' % (self.host, self.port, self.table))

class SubLibrary:
    def __init__(self):
        pass

    def returnint(self):
        return 2020

    def _returnint2(self):
        return 4