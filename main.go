package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

//// lists all sessions
//func main(){
//	sess := session.Must(session.NewSessionWithOptions(session.Options{
//		SharedConfigState: session.SharedConfigEnable,
//	}))
//
//	svc := sns.New(sess)
//
//	result, err := svc.ListTopics(nil)
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//
//	for _, t := range result.Topics {
//		fmt.Println(*t.TopicArn)
//	}
//
//}

////create new topic
//func main() {
//	if len(os.Args) < 2 {
//		fmt.Println("You must supply a topic name")
//		fmt.Println("Usage: go run SnsCreateTopic.go TOPIC")
//		os.Exit(1)
//	}
//
//	// Initialize a session that the SDK will use to load
//	// credentials from the shared credentials file. (~/.aws/credentials).
//	sess := session.Must(session.NewSessionWithOptions(session.Options{
//		SharedConfigState: session.SharedConfigEnable,
//	}))
//
//	svc := sns.New(sess)
//
//	result, err := svc.CreateTopic(&sns.CreateTopicInput{
//		Name: aws.String(os.Args[1]),
//	})
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//
//	fmt.Println("Topic Created",*result.TopicArn)
//}


// creates new subscription or gives ID in case it is already created.

func main() {
	emailPtr := flag.String("e", "", "The email address of the user subscribing to the topic")
	topicPtr := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *emailPtr == "" || *topicPtr == "" {
		fmt.Println("You must supply an email address and topic ARN")
		fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
		os.Exit(1)
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              emailPtr,
		Protocol:              aws.String("https"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.SubscriptionArn)
}

// sends message to all subscribers of a given topic.
//func main() {
//	// Initialize a session in us-west-2 that the SDK will use to load
//	// credentials from the shared credentials file ~/.aws/credentials.
//	sess := session.Must(session.NewSessionWithOptions(session.Options{
//				SharedConfigState: session.SharedConfigEnable,
//			}))
//
//	client := sns.New(sess)
//
//	input := &sns.PublishInput{
//		Message:  aws.String(`{"name":"bhavya","message":"aa gaya"}`),
//		TopicArn: aws.String("----------------"),
//	}
//
//	result, err := client.Publish(input)
//	if err != nil {
//		fmt.Println("Publish error:", err)
//		return
//	}
//
//	fmt.Println(result)
//}