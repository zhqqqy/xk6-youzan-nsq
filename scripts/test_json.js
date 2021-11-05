/*

This is a k6 test script that imports the xk6-kafka and
tests Kafka with a 200 JSON messages per iteration.

*/

import {consume, reader} from "k6/x/nsq"; // import kafka extension

const lookupaddress = ["10.215.49.141:4161"];
const nsqTopic = "sinker_test";
const nsqChannel = "sinker";

const consumer = reader(nsqTopic, nsqChannel);


export default function () {
    // for (let index = 0; index < 100; index++) {
    // Read 10 messages only
    consume(lookupaddress, consumer, 10,3);
    // check(messages, {
    //     "10 messages returned": (msgs) => msgs.length == 10,
    // });
    // }
}

export function teardown(data) {

}



