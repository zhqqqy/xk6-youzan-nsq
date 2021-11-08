/*

This is a k6 test script that imports the xk6-kafka and
tests Kafka with a 200 JSON messages per iteration.

*/
import {check} from 'k6';
import {Counter, Rate, Trend} from 'k6/metrics';
import {close, consume, received,} from 'k6/x/nsq'; // import mqtt plugin

const lookupAddress = ["10.215.49.141:4161"];
const nsqTopic = "sinker_test";
const nsqChannel = "sinker";

const consumer = consume(lookupAddress, nsqTopic, nsqChannel);
let subscribe_trend = new Trend('subscribe_time', true);

let myCounter = new Counter('my_counter');
let myRate = new Rate("my_rate");
export default function () {
    let startTime = new Date().getTime();
    let messages = received(lookupAddress, consumer, 10, 3);
    let paased = check(messages, {
        "is content correct": msg => msg != null
    });
    myCounter.add(1)
    myRate.add(paased)
    subscribe_trend.add(new Date().getTime() - startTime);
}

export function teardown(data) {
    close(consumer);
}



