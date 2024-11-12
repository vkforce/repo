package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {
	namespace := "CWAgent"
	metricName := "ingestion/commandservice/totalRequest_count"
	dimensions := []*cloudwatch.Dimension{
		{Name: aws.String("path"), Value: aws.String("/api/data:")},
		{Name: aws.String("metric_type"), Value: aws.String("counter")},
	}
	targetValue := float64(100)
	// collectionIntervalSeconds := int64(30)

	for {
		// endTime := time.Now().UTC()
		// startTime := endTime.Add(-time.Duration(collectionIntervalSeconds) * time.Second)
		var endTime, startTime time.Time
		layout := "2006-01-02T15:04:05Z"
		endTime, err := time.Parse(layout, "2024-06-05T05:52:33Z")
		startTime, err = time.Parse(layout, "2024-06-05T05:51:33Z")
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		params := &cloudwatch.GetMetricDataInput{
			EndTime:   &endTime,
			StartTime: &startTime,
			MetricDataQueries: []*cloudwatch.MetricDataQuery{
				{
					Id: aws.String("m1"),
					MetricStat: &cloudwatch.MetricStat{
						Metric: &cloudwatch.Metric{
							Namespace:  aws.String(namespace),
							MetricName: aws.String(metricName),
							Dimensions: dimensions,
						},
						Period: aws.Int64(10),
						Stat:   aws.String("Sum"),
					},
				},
			},
		}

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		})
		if err != nil {
			fmt.Println("Error creating session:", err)
			return
		}

		svc := cloudwatch.New(sess)
		result, err := svc.GetMetricData(params)
		if err != nil {
			fmt.Println("Error getting metric data:", err)
			continue
		}

		var aggregatedValue float64
		if len(result.MetricDataResults) > 0 && result.MetricDataResults[0].Values != nil {
			aggregatedValue = sumFloat64Slice(result.MetricDataResults[0].Values)
		}
		fmt.Printf("Start Time: %s\n", startTime.Format(time.RFC3339))
		fmt.Printf("End Time: %s\n", endTime.Format(time.RFC3339))
		// for _, value := range result.MetricDataResults[0].Values {
		// 	fmt.Println(*value)
		// }
		if aggregatedValue >= targetValue {
			fmt.Printf("Target metric value met or exceeded: %.2f\n", aggregatedValue)
		} else {
			fmt.Printf("Current aggregated value: %.2f, below target.\n", aggregatedValue)
		}

		time.Sleep(10 * time.Second)
	}
}

func sumFloat64Slice(slice []*float64) float64 {
	sum := float64(0)
	for _, v := range slice {
		sum += *v
	}
	return sum
}
