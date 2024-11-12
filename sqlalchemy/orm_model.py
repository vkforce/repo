from sqlalchemy import Column, String, Boolean, Integer, BigInteger, TIMESTAMP
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

Base = declarative_base()


class SiteVisit(Base):
    __tablename__ = 'idm_sitevisitarray_2024_february_1'

    id = Column(String(256), primary_key=True)
    tenantid = Column(String(256))
    tenantname = Column(String(256))
    eventtime = Column(TIMESTAMP)
    userid = Column(String(256))
    userfirstname = Column(String(256))
    userlastname = Column(String(256))
    username = Column(String(256))
    renderingaction = Column(String(256))
    url = Column(String(65535))
    urlcategories = Column(String(256))
    malicious = Column(Boolean)
    rbipolicyrulename = Column(String(256))
    applianceid = Column(String(256))
    nodeid = Column(String(256))
    title = Column(String(65535))
    host = Column(String(65535))
    tabid = Column(String(256))
    deviceid = Column(String(256))
    sessionid = Column(String(256))
    ipaddress = Column(String(256))
    clienttype = Column(String(256))
    clientversion = Column(String(256))
    os = Column(String(256))
    osversion = Column(String(256))
    browsertype = Column(String(256))
    tags = Column(String)
    country = Column(String)
    thumbnail = Column(String)
    threatscore = Column(BigInteger)
    year = Column(Integer)
    month = Column(String(256))
    day = Column(Integer)
    hour = Column(Integer)
    eventreceivedtime = Column(TIMESTAMP)
    eventprocessingstarttime = Column(TIMESTAMP)
    eventprocessingendtime = Column(TIMESTAMP)
    eventinsertedtime = Column(TIMESTAMP)
    eventid = Column(String(256), unique=True)


'''
select * from perf.idm_sitevisitarray_2024_february_1 where tenantid='0000000000000000000' limit 10;
'''

# Create the engine to connect to the Redshift database
engine = create_engine('redshift+psycopg2://username:password@hostname:port/database')

# Create a session factory
Session = sessionmaker(bind=engine)

# Create a session
session = Session()

# Perform the query using the ORM model and filter
results = session.query(SiteVisit).filter(SiteVisit.tenantid == '0000000000000000000').limit(10).all()

# Print the results
for result in results:
    print(result)
