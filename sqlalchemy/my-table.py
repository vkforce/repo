from sqlalchemy import create_engine

# engine = create_engine(
#     'redshift://insightsadmin:y5LCUMRTYgh5@dev-redshift-us-east-1-01.ccvenvsmaou8.us-east-1.redshift.amazonaws.com:5439/insightsdb')
# table_name = 'dev.rbi_youtube10'
# query = f"SELECT * FROM {table_name}"
# with engine.connect() as connection:
#     result = connection.execute(query)
#     rows = result.fetchall()
#
# for row in rows:
#     print(row)


class SiteVisit:
    

base_query = select(SiteVisit) \
    .where(and_(SiteVisit.tenantId == '002fb12e-f185-496a-8b49-636b877a1358',
                SiteVisit.eventTime.between('2024-02-01 23:19:00', '2024-02-21 23:19:00')))
tags_filter = base_query.filter(or_(SiteVisit.tags == 'movie', SiteVisit.tags.is_(None)))
country_filter = base_query.filter(or_(SiteVisit.country == 'Mexico', SiteVisit.country.is_(None)))