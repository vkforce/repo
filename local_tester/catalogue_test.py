from insights.config import Config
from insights import client
from insights.catalogue import ImportCatalogue
config = Config(
    insights_host="portal-new.insights.dev.forcepointone.com",
    platform_host="portal.dev.forcepointone.com",
    client_id="01ec7292-f7fb-4f6f-94eb-89f54a5d7184",
    client_secret="snJQuiCeiVXIPVZhEBgg_ojAXr3zhcbLxjEaYi1HyPE=",
    tenant_id="c69d6257-9dc9-4dde-9bff-f3a1d7d81317",
    # proxy_host="54.242.73.41",
    # proxy_port=3128,
    # proxy_scheme="http",
    # proxy_username="morpheus",
    # proxy_password="squid123"
)


ie_catalogue_client = client("Catalogues", config)
export_body = {
   "catalogueIds": [
      "9756b411-20f5-47e5-9a82-2b8b8c2d5fb6"
   ]
}
import_body = {
    "catalogue": [
        {
            "catalogueType": "widget",
            "name": "NewWidgetForTestImport-1",
            "type": "Custom",
            "description": "",
            "path": "Global/user2",
            "config": {
                "showDescInExportedPdf": "disabled",
                "showTooltipDescInDashboard": "disabled",
                "visualization": True
            },
            "content": {
                "collectionName": "FileTransfer",
                "fields": [
                    {
                        "aggregation": "",
                        "name": "avsPerformed",
                        "primaryAggregationField": False,
                        "sortOrder": 0,
                        "topXCount": 5,
                        "text": "Avscan Performed"
                    },
                    {
                        "aggregation": "count",
                        "name": "processedFilesize",
                        "primaryAggregationField": True,
                        "sortOrder": 0,
                        "topXCount": 0,
                        "text": "Processed Filesize Total Count"
                    }
                ],
                "filters": {
                    "FileTransfer__eventTime": {
                        "between": {
                            "mode": "inclusive",
                            "values": [
                                "2024-06-24 00:00:00",
                                "2024-07-24 05:40:16"
                            ]
                        }
                    }
                },
                "ppcode": "RBI"
            },
            "tags": [
                "automation"
            ]
        }
    ]
}
catalogues = [
    {
            "catalogueType": "widget",
            "name": "NewWidgetForTestImport-1",
            "type": "Custom",
            "description": "",
            "path": "Global/user2",
            "config": {
                "showDescInExportedPdf": "disabled",
                "showTooltipDescInDashboard": "disabled",
                "visualization": True
            },
            "content": {
                "collectionName": "FileTransfer",
                "fields": [
                    {
                        "aggregation": "",
                        "name": "avsPerformed",
                        "primaryAggregationField": False,
                        "sortOrder": 0,
                        "topXCount": 5,
                        "text": "Avscan Performed"
                    },
                    {
                        "aggregation": "count",
                        "name": "processedFilesize",
                        "primaryAggregationField": True,
                        "sortOrder": 0,
                        "topXCount": 0,
                        "text": "Processed Filesize Total Count"
                    }
                ],
                "filters": {
                    "FileTransfer__eventTime": {
                        "between": {
                            "mode": "inclusive",
                            "values": [
                                "2024-06-24 00:00:00",
                                "2024-07-24 05:40:16"
                            ]
                        }
                    }
                },
                "ppcode": "RBI"
            },
            "tags": [
                "automation"
            ]
        }
]
import_object = ImportCatalogue(
    catalogue=catalogues,
)



# result = ie_catalogue_client.export_catalogues_data(body=export_body)
# # old_name = result.result['catalogue'][0]["name"]
# print(result.result)
# print(result.error)

# catalogue_body = result.result['catalogue'][0]["name"] = old_name + "-recreated-widget"
# import_result = ie_catalogue_client.import_catalogues_data(body=import_body)
# print(import_result.result)
# print(import_result.error)

summary = ie_catalogue_client.retrieve_catalogues_summary(catalogue_object=import_object)
print(summary.result)
print(summary.error)

