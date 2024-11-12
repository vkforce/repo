from ingestion import Ingestion
import json
from collection import GetAllCollectionHandler

if __name__ == '__main__':
    data = {
        "name": "Youtube34",
        "payload": [
            {
                "ID": "1",
                "tenantId": "dd88422f-ee18-411b-8868-b0006a92f9c8",
                "eventTime": "2024-03-02 10:54:22",
                "name": "Video 1",
                "tags": ["motivational", "gym", "excersise"],
                "url": "youtube.com/video/1",
                "category": "fitness",
                "viewcount": 1500
            },
            {
                "ID": "2",
                "tenantId": "dd88422f-ee18-411b-8868-b0006a92f9c8",
                "eventTime": "2024-03-02 10:54:22",
                "name": "Video 2",
                "tags": ["comedy", "standup"],
                "url": "youtube.com/video/2",
                "category": "entertainement",
                "viewcount": 2000
            },
            {
                "ID": "3",
                "tenantId": "dd88422f-ee18-411b-8868-b0006a92f9c8",
                "eventTime": "2024-03-02 10:54:22",
                "name": "Video 3",
                "tags": ["movie", "comics", "robots", "animation", "comedy"],
                "url": "youtube.com/video/3",
                "category": "entertainement",
                "viewcount": 2500
            },
            {
                "ID": "4",
                "tenantId": "dd88422f-ee18-411b-8868-b0006a92f9c8",
                "eventTime": "2024-03-02 10:54:22",
                "name": "Video 4",
                "tags": ["spiritual", "god"],
                "url": "youtube.com/video/4",
                "category": "podcast",
                "viewcount": 300
            },
            {
                "ID": "5",
                "tenantId": "dd88422f-ee18-411b-8868-b0006a92f9c8",
                "eventTime": "2024-03-02 10:54:22",
                "name": "Video 5",
                "tags": [],
                "url": "youtube.com/video/5",
                "category": "fitness",
                "viewcount": 1500
            }
        ],
        "metadata": {
            "pointProductCode": "RBI"
        }
    }
    payload = json.dumps(data)
    gen = Ingestion(env='dev', payload=payload)
    result = gen.handler()
    print(result)
    # collections = GetAllCollectionHandler('dev')
    # result = collections.handler()
    # print(result)
