import boto3
from datetime import datetime, timedelta
import time
import pytz


cloudwatch = boto3.client('cloudwatch')

def get_cloudwatch_metrics(namespace, metric_name, dimensions, start_time, end_time):
    params = {
        'StartTime': start_time,
        'EndTime': end_time,
        'MetricDataQueries': [
            {
                'Id': 'm1',
                'MetricStat': {
                    'Metric': {
                        'Namespace': namespace,
                        'MetricName': metric_name,
                        'Dimensions': dimensions
                    },
                    'Period': 10,
                    'Stat': 'Sum'
                }
            }
        ]
    }

    try:
        response = cloudwatch.get_metric_data(**params)
        return response['MetricDataResults']
    except Exception as e:
        print(f"Error fetching metrics: {e}")
        return []

def main():
    namespace = "CWAgent"
    metric_name = "ingestion/commandservice/totalRequest_count"
    dimensions = [{'Name': 'path', 'Value': '/api/data:'}, {'Name':'metric_type', 'Value': 'counter'}]
    target_value = 500
    
    collection_interval_seconds = 30
    
    # utc_now = datetime.now(pytz.UTC)
    
    while True:
        end_time = datetime.now(pytz.UTC) -  timedelta(seconds=10)
        start_time = end_time - timedelta(seconds=collection_interval_seconds)
        
        start_time_str = start_time.strftime('%Y-%m-%dT%H:%M:%S')
        end_time_str = end_time.strftime('%Y-%m-%dT%H:%M:%S')

        metrics_data = get_cloudwatch_metrics(namespace, metric_name, dimensions, start_time_str, end_time_str)
        print("Start Time: ", start_time_str)
        print("End Time: ", end_time_str)
        print(metrics_data)
        if metrics_data:
            aggregated_value = sum(metrics_data[0]['Values']) if metrics_data[0]['Values'] else None
            
            if aggregated_value and aggregated_value >= target_value:
                print(f"Target metric value met or exceeded: {aggregated_value}")
            else:
                print(f"Current aggregated value: {aggregated_value}, below target.")
        else:
            print("No metrics data received.")

        print(f"Sleeping for 10 seconds...")
        time.sleep(10)
        utc_now = datetime.now(pytz.UTC)

if __name__ == "__main__":
    main()
